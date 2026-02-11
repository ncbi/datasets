package datasets

import (
	"context"
	openapi "datasets/openapi/v2"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
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
	openapi.V2REPORTSRANKTYPE_STRAIN,
	openapi.V2REPORTSRANKTYPE_SUBVARIETY,
	openapi.V2REPORTSRANKTYPE_SEROTYPE,
	openapi.V2REPORTSRANKTYPE_ISOLATE,
}

// Ranks that we know to be above species
var aboveSpeciesRanks = []openapi.V2reportsRankType{
	openapi.V2REPORTSRANKTYPE_ACELLULAR_ROOT,
	openapi.V2REPORTSRANKTYPE_CELLULAR_ROOT,
	openapi.V2REPORTSRANKTYPE_REALM,
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

type AutoSuggestError struct {
	Msg string
}

func (e *AutoSuggestError) Error() string {
	return e.Msg
}

type taxonAutosuggestApi struct {
	TaxonApi *openapi.TaxonomyAPIService
}

func RetrieveTaxIdForTaxon(
	taxName string,
	aboveSpecies bool,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
	resourceType string,
	args ...int32) (string, error) {

	cli, cliErr := createOAClient()
	if cliErr != nil {
		return "", cliErr
	}
	api := taxonAutosuggestApi{TaxonApi: cli.TaxonomyAPI}
	return api.GetOrganisms(taxName, aboveSpecies, resource, resourceType, maxAlternateSuggestionNames, args...)
}

// Returns valid input taxids mapped to matching taxnames, and writes messages to stderr for any invalid choices
func RetrieveTaxIdsForTaxons(
	cmd *cobra.Command,
	taxNames []string,
	aboveSpecies bool,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
	resourceType string,
	args ...int32) (map[string][]string, error) {

	var taxIdsMap map[string][]string = make(map[string][]string)
	cli, cliErr := createOAClient()
	if cliErr != nil {
		return taxIdsMap, cliErr
	}
	api := taxonAutosuggestApi{TaxonApi: cli.TaxonomyAPI}

	var asgErr *AutoSuggestError
	for _, taxName := range taxNames {
		taxId, taxError := api.GetOrganisms(taxName, aboveSpecies, resource, resourceType, maxAlternateSuggestionNames, args...)
		if taxError != nil {
			// Return connection or other undefined errors, or if there was only 1 taxon to check, return its error
			if !errors.As(taxError, &asgErr) || len(taxNames) == 1 {
				return taxIdsMap, taxError
			}
			cmd.PrintErr(fmt.Errorf("%w\n\n", taxError))
			continue
		}

		if _, ok := taxIdsMap[taxId]; !ok {
			taxIdsMap[taxId] = []string{taxName}
		} else {
			taxIdsMap[taxId] = append(taxIdsMap[taxId], taxName)
		}
	}
	if len(taxIdsMap) == 0 {
		return taxIdsMap, fmt.Errorf("No valid taxons provided")
	}

	return taxIdsMap, nil
}

func (apiService *taxonAutosuggestApi) GetMetadata(taxId string, returnedContent openapi.V2TaxonomyMetadataRequestContentType) (*openapi.V2reportsTaxonomyDataReportPage, bool, error) {
	result, _, err := apiService.TaxonApi.TaxonomyDataReportPost(context.TODO()).V2TaxonomyMetadataRequest(
		openapi.V2TaxonomyMetadataRequest{
			Taxons:          []string{taxId},
			ReturnedContent: &returnedContent,
		},
	).Execute()

	hasResults := (result.Reports != nil) && (len(result.Reports) == 1) && ((result.Reports)[0].Taxonomy != nil)
	return result, hasResults, err
}

func (apiService *taxonAutosuggestApi) CheckLineage(taxId string, lineageFilter int32) (bool, error) {
	metadata, hasResults, err := apiService.GetMetadata(taxId, openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_METADATA)
	if err != nil {
		return false, err
	}

	if !hasResults {
		return false, &AutoSuggestError{Msg: fmt.Sprintf("Taxonomy metadata not found for tax ID %s", taxId)}
	}

	if *(metadata.Reports)[0].Taxonomy.TaxId == lineageFilter {
		return true, nil
	}
	for _, lineageTaxId := range metadata.Reports[0].Taxonomy.Parents {
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
func (apiService *taxonAutosuggestApi) IsAtOrBelowSpecies(
	taxId string,
	rank *openapi.V2reportsRankType,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
) (bool, error) {
	atOrBelowSpecies, aboveSpcies := rankStatus(rank)
	if atOrBelowSpecies || aboveSpcies {
		return atOrBelowSpecies, nil
	}

	metadata, hasResults, err := apiService.GetMetadata(taxId, openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_METADATA)
	if err != nil {
		return false, err
	}
	if !hasResults {
		err = &AutoSuggestError{Msg: fmt.Sprintf("Taxonomy metadata not found for tax ID %s", taxId)}
		return false, err
	}

	// To determine if we are at species or below, iterate over lineage in reverse order until
	// all taxons (except root (1)) are checked and if any are of rank species, return true.
	lineageLen := len(metadata.Reports[0].Taxonomy.Parents) - 1
	for idx := range metadata.Reports[0].Taxonomy.Parents {
		lineageTaxId := metadata.Reports[0].Taxonomy.Parents[lineageLen-idx]
		if lineageTaxId == 1 {
			break
		}
		linMetadata, linHasResults, linErr := apiService.GetMetadata(strconv.Itoa(int(lineageTaxId)), openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_METADATA)
		if linErr != nil {
			return false, linErr
		}
		if !linHasResults {
			err = &AutoSuggestError{Msg: fmt.Sprintf("Taxonomy metadata not found for tax ID %d", lineageTaxId)}
			return false, err
		}

		atOrBelowSpecies, aboveSpcies = rankStatus((linMetadata.Reports)[0].Taxonomy.Rank)
		if atOrBelowSpecies || aboveSpcies {
			return atOrBelowSpecies, nil
		}
	}
	return false, nil
}

func (apiService *taxonAutosuggestApi) GetOrganismResourceCount(
	taxId string,
	resource openapi.V2OrganismQueryRequestTaxonResourceFilter,
) (int32, error) {
	metadata, hasResults, err := apiService.GetMetadata(taxId, openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_COMPLETE)
	var resourceCount int32 = 0
	if err != nil {
		return resourceCount, err
	}
	if !hasResults {
		err = &AutoSuggestError{Msg: fmt.Sprintf("Taxonomy metadata not found for tax ID %s", taxId)}
		return resourceCount, err
	}

	if countType, hasCounts := resourceCountType[resource]; hasCounts {
		for _, count := range (metadata.Reports)[0].Taxonomy.Counts {
			if countType == count.GetType() {
				resourceCount = count.GetCount()
			}
		}
	}

	return resourceCount, err
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

	if err != nil {
		return result, false, err
	}

	if result == nil {
		return nil, false, &AutoSuggestError{Msg: fmt.Sprintf("%s appears to be an invalid tax name.  This could be due to unexpected special characters.", taxName)}
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
			// common or group name otherwise (adding both when both available seems unwieldy)
			if name.MatchedTerm != nil && *name.MatchedTerm != *name.SciName {
				taxNameAlternates += ", " + *name.MatchedTerm
			} else if name.CommonName != nil {
				taxNameAlternates += ", " + *name.CommonName
			} else if name.GroupName != nil {
				taxNameAlternates += ", " + *name.GroupName
			}
			taxNameAlternates += ")"
		}
	}
	return taxIdMatch, taxNameAlternates, numAlternateNames
}

func (apiService *taxonAutosuggestApi) handleExactMatch(
	name openapi.V2SciNameAndIdsSciNameAndId,
	resourceType string,
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

	if !aboveSpecies {
		atOrBelowSpecies, absErr := apiService.IsAtOrBelowSpecies(taxId, name.Rank, resource)
		if absErr != nil {
			return absErr
		}
		if !atOrBelowSpecies {
			// Please use the command `datasets summary taxonomy taxon` to explore this taxonomic name.
			if isTaxIdQuery {
				return &AutoSuggestError{Msg: fmt.Sprintf(msgText+", but %s requires an at-or-below-species-level taxon.", resourceType)}
			}
			return &AutoSuggestError{Msg: fmt.Sprintf(msgText+", but %s requires an at-or-below-species-level taxon.\nPlease use the command `datasets summary taxonomy taxon` to explore this taxonomic name", resourceType)}
		}
	}

	// Check if the result is valid but has no entries (calls to summary would also know this, but not calls to download)
	if resource != openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL {
		_, hasResults, err := apiService.GetMatchingOrganisms(sciName, aboveSpecies, true, resource)
		if err != nil {
			return err
		}

		if !hasResults {
			resourceCount, err := apiService.GetOrganismResourceCount(taxId, resource)
			if err != nil {
				return err
			}
			if resourceCount > 0 {
				var rank = strings.ToLower(string(name.GetRank()))
				err = &AutoSuggestError{Msg: fmt.Sprintf("There are no annotations matching your query at the %s level for '%s' (taxid: '%s') please use a subspecies tax name or ID.", rank, sciName, taxId)}
			} else {
				err = &AutoSuggestError{Msg: fmt.Sprintf(msgText+", but no %s data is currently available for this taxon.", resourceType)}
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
	resourceType string,
	maxAlternateNames int,
	args ...int32) (string, error) {

	// Trim blanks
	taxName = strings.TrimSpace(taxName)
	autosuggestTaxName := taxName

	var taxIdFilter int32 = 0
	var lineageFilterName string = ""
	if len(args) > 0 {
		taxIdFilter = args[0]
		lineageFilterName = lineageTerms[taxIdFilter]
	}

	// Check for an exact match with the incoming taxon first and just return that if found. If not found,
	// do a more general search to get suggestions, taking into account above/below species, lineage, and
	// resource type requirements.
	exactMatch, hasExactMatch, err := apiService.GetMatchingOrganisms(autosuggestTaxName, aboveSpecies, true, resource)
	if err != nil {
		return "", err
	}
	if hasExactMatch && len(exactMatch.SciNameAndIds) == 1 {
		validLineage := true
		if taxIdFilter != 0 {
			validLineage, err = apiService.CheckLineage(autosuggestTaxName, taxIdFilter)
			if err != nil {
				return "", err
			}
		}
		if validLineage {
			return *exactMatch.SciNameAndIds[0].TaxId, nil
		}
	}

	// Do query for all taxons so that we catch exact matches (of id or name) even if they are not otherwise usable for the query
	// (e.g. no genome data or not below species, or does not have data for the requested resource).
	result, hasResults, err := apiService.GetMatchingOrganisms(autosuggestTaxName, true, true, openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", &AutoSuggestError{Msg: fmt.Sprintf("There is an unknown error for tax name '%s'", taxName)}
	}

	// If taxon was numeric (a taxid) validate it and check if at-or below species if needed
	if _, err := strconv.Atoi(taxName); err == nil {
		// We don't return any autosuggest results if the query term was only a number.
		if !hasResults {
			return "", &AutoSuggestError{Msg: fmt.Sprintf("The taxonomy ID '%s' does not match any existing taxids for '%s'", taxName, resourceType)}
		}

		if taxIdFilter != 0 {
			validLineage, linErr := apiService.CheckLineage(taxName, taxIdFilter)
			if linErr != nil {
				return "", linErr
			}
			if !validLineage {
				return "", &AutoSuggestError{Msg: fmt.Sprintf("The taxonomy ID '%s' is valid for '%s', but it is not a %s taxon", taxName, *(result.SciNameAndIds)[0].SciName, lineageFilterName)}
			}
		}

		err = apiService.handleExactMatch(
			(result.SciNameAndIds)[0],
			resourceType,
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
	// Already checked for valid lineage here
	if len(exactMatches) == 1 {
		err = apiService.handleExactMatch(
			(result.SciNameAndIds)[0],
			resourceType,
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
		return "", &AutoSuggestError{Msg: msg}
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
	return "", &AutoSuggestError{Msg: msg}
}
