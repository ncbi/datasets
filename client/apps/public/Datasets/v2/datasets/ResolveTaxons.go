package datasets

import (
	"context"
	openapi "datasets/openapi/v2"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var (
	// Message terms to use with taxon filtering
	lineageTerms = map[int32]string{
		10239: "virus",
	}
)

const maxAlternateSuggestionNames = 10

// Ranks that we know to be at-or-below species
var belowSpeciesRanks = []openapi.V2reportsRankType{
	openapi.V2REPORTSRANKTYPE_SPECIES,
	openapi.V2REPORTSRANKTYPE_SUBSPECIES,
	openapi.V2REPORTSRANKTYPE_TRIBE,
	openapi.V2REPORTSRANKTYPE_SUBTRIBE,
	openapi.V2REPORTSRANKTYPE_STRAIN,
	openapi.V2REPORTSRANKTYPE_SUBVARIETY,
	openapi.V2REPORTSRANKTYPE_SEROTYPE,
	openapi.V2REPORTSRANKTYPE_ISOLATE,
}

// Ranks that we know to be above species
var aboveSpeciesRanks = []openapi.V2reportsRankType{
	openapi.V2REPORTSRANKTYPE_SUPERKINGDOM,
	openapi.V2REPORTSRANKTYPE_KINGDOM,
	openapi.V2REPORTSRANKTYPE_SUBKINGDOM,
	openapi.V2REPORTSRANKTYPE_SUPERPHYLUM,
	openapi.V2REPORTSRANKTYPE_SUBPHYLUM,
	openapi.V2REPORTSRANKTYPE_PHYLUM,
	openapi.V2REPORTSRANKTYPE_CLADE,
	openapi.V2REPORTSRANKTYPE_SUPERCLASS,
	openapi.V2REPORTSRANKTYPE_CLASS,
	openapi.V2REPORTSRANKTYPE_SUBCLASS,
	openapi.V2REPORTSRANKTYPE_INFRACLASS,
	openapi.V2REPORTSRANKTYPE_SUPERORDER,
	openapi.V2REPORTSRANKTYPE_ORDER,
	openapi.V2REPORTSRANKTYPE_SUBORDER,
	openapi.V2REPORTSRANKTYPE_INFRAORDER,
	openapi.V2REPORTSRANKTYPE_PARVORDER,
	openapi.V2REPORTSRANKTYPE_SUPERFAMILY,
	openapi.V2REPORTSRANKTYPE_FAMILY,
	openapi.V2REPORTSRANKTYPE_SUBFAMILY,
	openapi.V2REPORTSRANKTYPE_GENUS,
	openapi.V2REPORTSRANKTYPE_SUBGENUS,
}

var resourceCountType = map[openapi.V2OrganismQueryRequestTaxonResourceFilter]openapi.V2reportsCountType{
	openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENOME: openapi.V2REPORTSCOUNTTYPE_ASSEMBLY,
	openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE:   openapi.V2REPORTSCOUNTTYPE_GENE,
}

type taxonAutosuggestApi struct {
	TaxonApi *openapi.TaxonomyAPIService
}

func RetrieveTaxIdForTaxon(
	taxName string,
	aboveSpecies bool,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
	cmd string,
	args ...int32) (string, error) {

	cli, cliErr := createOAClient()
	if cliErr != nil {
		return "", cliErr
	}
	api := taxonAutosuggestApi{TaxonApi: cli.TaxonomyAPI}
	return api.GetOrganisms(taxName, aboveSpecies, resource, cmd, maxAlternateSuggestionNames, args...)
}

func (apiService *taxonAutosuggestApi) GetMetadata(taxName string, returnedContent openapi.V2TaxonomyMetadataRequestContentType) (*openapi.V2TaxonomyMetadataResponse, bool, error) {
	result, _, err := apiService.TaxonApi.TaxonomyMetadataPost(context.TODO()).V2TaxonomyMetadataRequest(
		openapi.V2TaxonomyMetadataRequest{
			Taxons:          []string{taxName},
			ReturnedContent: &returnedContent,
		},
	).Execute()

	hasResults := false
	if (result.TaxonomyNodes != nil) && (len(result.TaxonomyNodes) == 1) && ((result.TaxonomyNodes)[0].Taxonomy != nil) {
		var namesMap map[string]bool = make(map[string]bool)
		namesMap[strings.ToUpper(*(result.TaxonomyNodes)[0].Taxonomy.OrganismName)] = true
		namesMap[strconv.Itoa(int(*(result.TaxonomyNodes)[0].Taxonomy.TaxId))] = true
		if (result.TaxonomyNodes)[0].Taxonomy.CommonName != nil {
			namesMap[strings.ToUpper(*(result.TaxonomyNodes)[0].Taxonomy.CommonName)] = true
		}
		if (result.TaxonomyNodes)[0].Taxonomy.GenbankCommonName != nil {
			namesMap[strings.ToUpper(*(result.TaxonomyNodes)[0].Taxonomy.GenbankCommonName)] = true
		}
		if (result.TaxonomyNodes)[0].Taxonomy.GenbankAcronym != nil {
			namesMap[strings.ToUpper(*(result.TaxonomyNodes)[0].Taxonomy.GenbankAcronym)] = true
		}
		if namesMap[strings.ToUpper(taxName)] {
			hasResults = true
		}
	}

	return result, hasResults, err
}

func (apiService *taxonAutosuggestApi) CheckLineage(taxId string, lineageFilter int32) (bool, error) {
	metadata, hasResults, err := apiService.GetMetadata(taxId, openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_METADATA)
	if err != nil {
		return false, err
	}

	if !hasResults {
		err = fmt.Errorf("Taxonomy metadata not found for tax ID %s", taxId)
		return false, err
	}

	if *(metadata.TaxonomyNodes)[0].Taxonomy.TaxId == lineageFilter {
		return true, nil
	}
	for _, lineageTaxId := range metadata.TaxonomyNodes[0].Taxonomy.Lineage {
		if lineageTaxId == lineageFilter {
			return true, nil
		}
	}
	return false, nil
}

// returns atOrBelowSpecies, aboveSpecies (both false if we are not sure)
func rankStatus(rank *openapi.V2reportsRankType) (bool, bool) {
	if rank == nil {
		return false, false
	}

	if slices.Contains(belowSpeciesRanks, *rank) {
		return true, false
	}

	if slices.Contains(aboveSpeciesRanks, *rank) {
		return false, true
	}

	return false, false
}

// Return true if the rank of the given taxid is at or below species. Use lineage
// for verification, if needed.
func (apiService *taxonAutosuggestApi) IsAtOrBelowSpecies(taxId string, resource openapi.V2OrganismQueryRequestTaxonResourceFilter) (bool, int32, error) {
	metadata, hasResults, err := apiService.GetMetadata(taxId, openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_COMPLETE)
	var resourceCount int32 = 0
	if err != nil {
		return false, resourceCount, err
	}
	if !hasResults {
		err = fmt.Errorf("Taxonomy metadata not found for tax ID %s", taxId)
		return false, resourceCount, err
	}

	if countType, hasCounts := resourceCountType[resource]; hasCounts {
		for _, count := range (metadata.TaxonomyNodes)[0].Taxonomy.Counts {
			if countType == count.GetType() {
				resourceCount = count.GetCount()
			}
		}
	}

	atOrBelowSpecies, aboveSpcies := rankStatus((metadata.TaxonomyNodes)[0].Taxonomy.Rank)
	if atOrBelowSpecies || aboveSpcies {
		return atOrBelowSpecies, resourceCount, nil
	}

	// To determine if we are at species or below, iterate over lineage in reverse order until
	// all taxons (except root (1)) are checked and if any are of rank species, return true.
	lineageLen := len(metadata.TaxonomyNodes[0].Taxonomy.Lineage) - 1
	for idx := range metadata.TaxonomyNodes[0].Taxonomy.Lineage {
		lineageTaxId := metadata.TaxonomyNodes[0].Taxonomy.Lineage[lineageLen-idx]
		if lineageTaxId == 1 {
			break
		}
		linMetadata, linHasResults, linErr := apiService.GetMetadata(strconv.Itoa(int(lineageTaxId)), openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_METADATA)
		if linErr != nil {
			return false, resourceCount, linErr
		}
		if !linHasResults {
			err = fmt.Errorf("Taxonomy metadata not found for tax ID %d", lineageTaxId)
			return false, resourceCount, err
		}

		atOrBelowSpecies, aboveSpcies = rankStatus((linMetadata.TaxonomyNodes)[0].Taxonomy.Rank)
		if atOrBelowSpecies || aboveSpcies {
			return atOrBelowSpecies, resourceCount, nil
		}
	}
	return false, resourceCount, nil
}

func nameMatches(taxName string, name openapi.V2SciNameAndIdsSciNameAndId) bool {
	// seems to be that the returned values do not always have the matching name in the MatchedTerm field, so
	// check them all.
	taxNameUpper := strings.ToUpper(taxName)
	if (taxNameUpper == strings.ToUpper(*name.MatchedTerm)) ||
		(taxNameUpper == strings.ToUpper(*name.SciName)) ||
		((name.CommonName != nil) && (taxNameUpper == strings.ToUpper(*name.CommonName))) {
		return true
	}
	return false
}

func (apiService *taxonAutosuggestApi) GetMatchingOrganisms(taxName string, aboveSpecies bool, exactMatch bool, resource openapi.V2OrganismQueryRequestTaxonResourceFilter) (*openapi.V2SciNameAndIds, bool, error) {
	taxRank := openapi.V2ORGANISMQUERYREQUESTTAXRANKFILTER_SPECIES
	if aboveSpecies {
		taxRank = openapi.V2ORGANISMQUERYREQUESTTAXRANKFILTER_HIGHER_TAXON
	}
	result, _, err := apiService.TaxonApi.TaxNameQueryByPost(context.TODO()).
		V2OrganismQueryRequest(
			openapi.V2OrganismQueryRequest{
				TaxonQuery:          &taxName,
				TaxRankFilter:       &taxRank,
				TaxonResourceFilter: &resource,
				ExactMatch:          &exactMatch,
			}).Execute()

	if result == nil {
		return nil, false, fmt.Errorf("%s appears to be an invalid tax name.  This could be due to unexpected special characters.", taxName)
	}

	if err != nil {
		return result, false, err
	}

	hasResults := (result.SciNameAndIds != nil) && (len(result.SciNameAndIds) != 0) && ((result.SciNameAndIds)[0].SciName != nil)
	return result, hasResults, nil
}

func (apiService *taxonAutosuggestApi) resultMatches(
	names []openapi.V2SciNameAndIdsSciNameAndId,
	taxName string,
	taxIdFilter int32,
	maxAlternateNames int) ([]string, string, int) {

	var taxIdMatch = []string{}
	taxNameAlternates := ""
	numAlternateNames := 0

	for _, name := range names {
		validLineage := true
		if taxIdFilter > 0 {
			// Allows us to restrict suggestions to virus-only (for virus commands)
			validLineage, _ = apiService.CheckLineage(*name.TaxId, taxIdFilter)
		}
		if validLineage {
			// exact matches - may be more than 1 so we still collect alternates
			exactMatch := nameMatches(taxName, name)
			if exactMatch {
				taxIdMatch = append(taxIdMatch, *name.TaxId)
			}

			if !exactMatch && numAlternateNames >= maxAlternateNames {
				break
			}
			numAlternateNames += 1

			if numAlternateNames != 1 {
				taxNameAlternates += "\n"
			}
			var rankStr = "no-rank"
			if name.Rank != nil {
				rankStr = strings.ToLower(string(*name.Rank))
			}
			taxNameAlternates += *name.SciName + " (" + rankStr + ", taxid: " + *name.TaxId

			// Add matched term to the alternates if it is not the scientific name, and the
			// common name otherwise (adding both when both available seems unwieldy)
			if name.MatchedTerm != nil && *name.MatchedTerm != *name.SciName {
				taxNameAlternates += ", " + *name.MatchedTerm
			} else if name.CommonName != nil {
				taxNameAlternates += ", " + *name.CommonName
			}
			taxNameAlternates += ")"
		}
	}
	return taxIdMatch, taxNameAlternates, numAlternateNames
}

func (apiService *taxonAutosuggestApi) handleExactMatch(
	name openapi.V2SciNameAndIdsSciNameAndId,
	cmd string,
	taxQuery string,
	isTaxIdQuery bool,
	aboveSpecies bool,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
) error {
	sciName := *name.SciName
	taxId := *name.TaxId

	msgText := ""
	if isTaxIdQuery {
		msgText = fmt.Sprintf("The taxonomy ID '%s' is valid for '%s'", taxId, sciName)
	} else {
		msgText = fmt.Sprintf("The taxonomy name '%s' (taxid: '%s') is valid", taxQuery, taxId)
	}

	atOrBelowSpecies, resourceCount, absErr := apiService.IsAtOrBelowSpecies(taxId, resource)
	if !aboveSpecies {
		if absErr != nil {
			return absErr
		}
		if !atOrBelowSpecies {
			// Please use the command `datasets summary taxonomy taxon` to explore this taxonomic name.
			if isTaxIdQuery {
				return fmt.Errorf(msgText+", but %s requires an at-or-below-species-level taxon.", cmd)
			}
			return fmt.Errorf(msgText+", but %s requires an at-or-below-species-level taxon.\nPlease use the command `datasets summary taxonomy taxon` to explore this taxonomic name", cmd)
		}
	}

	// Check if the result is valid but has no entries (calls to summary would also know this, but not calls to download)
	if resource != openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL {
		_, hasResults, err := apiService.GetMatchingOrganisms(sciName, aboveSpecies, true, resource)
		if err != nil {
			return err
		}

		if !hasResults {
			if resourceCount > 0 {
				var rank = strings.ToLower(string(name.GetRank()))
				err = fmt.Errorf("There are no annotations matching your query at the %s level for '%s' (taxid: '%s') please use a subspecies tax name or ID.", rank, sciName, taxId)
			} else {
				err = fmt.Errorf(msgText+", but no %s data is currently available for this taxon.", cmd)
			}
			return err
		}
	}
	return nil
}

func (apiService *taxonAutosuggestApi) GetOrganisms(
	taxName string,
	aboveSpecies bool,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
	cmd string,
	maxAlternateNames int,
	args ...int32) (string, error) {

	// Trim blanks
	taxName = strings.TrimSpace(taxName)
	// Consider adding a blank here.  Good: it increases odds of seeing suggestions where the word
	// is a complete word. Bad: if the user typed a partial word on purpose, they don't see anything useful.
	autosuggestTaxName := taxName

	// Do initial query for all taxons so that we catch exact matches (of id or name) even if they are not otherwise usable for the query
	// (e.g. no genome data or not below species).
	result, hasResults, err := apiService.GetMatchingOrganisms(autosuggestTaxName, true, true, openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL)
	if err != nil {
		return "", err
	}
	if result == nil {
		err = fmt.Errorf("There is an unknown error for tax name '%s'", taxName)
		return "", err
	}

	var taxIdFilter int32 = 0
	var lineageFilterName string = ""
	if len(args) > 0 {
		taxIdFilter = args[0]
		lineageFilterName = lineageTerms[taxIdFilter]
	}

	// If taxon was numeric (a taxid) validate it and check if at-or below species if needed
	if _, err := strconv.Atoi(taxName); err == nil {
		// We don't return any autosuggest results if the query term was only a number.
		if !hasResults {
			err = fmt.Errorf("The taxonomy ID '%s' does not match any existing taxids for '%s'", taxName, cmd)
			return "", err
		}

		if taxIdFilter != 0 {
			validLineage, linErr := apiService.CheckLineage(taxName, taxIdFilter)
			if linErr != nil {
				return "", linErr
			}
			if !validLineage {
				return "", fmt.Errorf("The taxonomy ID '%s' is valid for '%s', but it is not a %s taxon", taxName, *(result.SciNameAndIds)[0].SciName, lineageFilterName)
			}
		}

		err = apiService.handleExactMatch(
			(result.SciNameAndIds)[0],
			cmd,
			taxName,
			true,
			aboveSpecies,
			resource,
		)
		if err != nil {
			return "", err
		}

		return taxName, nil
	}

	// Not a numeric taxon - check for exact match in current (non-restricted) results
	exactMatches, alternateNames, _ := apiService.resultMatches(result.SciNameAndIds, taxName, taxIdFilter, maxAlternateNames)
	if len(exactMatches) == 1 {
		err = apiService.handleExactMatch(
			(result.SciNameAndIds)[0],
			cmd,
			taxName,
			false,
			aboveSpecies,
			resource,
		)
		if err != nil {
			return "", err
		}

		return exactMatches[0], nil
	}

	// Some scientfic and common names are duplicated, e.g. mus, drosophila, coffee, so have user choose in those cases. Note that
	// in cases where the user may be querying for genome/gene specific data, this can also include above-species results
	if len(exactMatches) > 1 {
		msg := fmt.Sprintf("The taxonomy name '%s' is an exact match for more than one taxid. Please select one of these taxids:\n%s", taxName, alternateNames)
		return "", fmt.Errorf("%s", msg)
	}

	// No exact matches in autosuggest search, now search specified resource and above/below species, and
	// return suggestions as appropriate.
	result, _, err = apiService.GetMatchingOrganisms(autosuggestTaxName, aboveSpecies, false, resource)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", err
	}

	var numAlternateNames int
	exactMatches, alternateNames, numAlternateNames = apiService.resultMatches(result.SciNameAndIds, taxName, taxIdFilter, maxAlternateNames)
	if len(exactMatches) != 0 {
		// This should not happen because we already checked for exact matches above.
		return exactMatches[0], nil
	}

	msg := fmt.Sprintf("The taxonomy name '%s'", taxName)
	if alternateNames != "" {
		var phrase string = "is not exact"
		if lineageFilterName != "" {
			phrase = fmt.Sprintf(" is not a %s taxon", lineageFilterName)
		}
		if numAlternateNames == 1 {
			msg += fmt.Sprintf(" %s. Try using the taxid:\n%s", phrase, alternateNames)
		} else {
			msg += fmt.Sprintf(" %s. Try using one of the suggested taxids:\n%s", phrase, alternateNames)
		}
	} else {
		if lineageFilterName == "" {
			msg += " is not recognized."
		} else {
			msg += fmt.Sprintf(" is not a recognized %s taxon.", lineageFilterName)
		}
	}
	return "", fmt.Errorf("%s", msg)
}
