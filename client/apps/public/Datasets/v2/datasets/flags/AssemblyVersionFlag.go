package flags

import (
	"strings"

	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type GenomeAssemblyVersion enumflag.Flag

// Define the enumeration values for GenomeAssemblyVersion
const (
	Current GenomeAssemblyVersion = iota
	All_assemblies
	Default
)

// Note that all_assemblies matches protobuf value but we want more intuitive 'all' for the cli parameter
var GenomeAssemblyVersionIds = map[GenomeAssemblyVersion][]string{
	Current:        {"latest", "current"},
	All_assemblies: {"all"},
	Default:        {""},
}

var GenomeAssemblyVersionOpenapi = map[GenomeAssemblyVersion]openapi.V2AssemblyDatasetDescriptorsFilterAssemblyVersion{
	Current:        openapi.V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYVERSION_CURRENT,
	All_assemblies: openapi.V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYVERSION_ALL_ASSEMBLIES,
}

func RemoveDuplicateStrings(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type AssemblyVersionFlag struct {
	FlagInterface
	AssemblyVersion *GenomeAssemblyVersion
}

func NewAssemblyVersionFlag(defaultAssemblyVersionFlag GenomeAssemblyVersion) *AssemblyVersionFlag {
	avf := &AssemblyVersionFlag{
		AssemblyVersion: &defaultAssemblyVersionFlag,
	}
	return avf
}

func (avf *AssemblyVersionFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(avf.AssemblyVersion, "string", GenomeAssemblyVersionIds, enumflag.EnumCaseInsensitive),
		"assembly-version",
		`Limit to 'latest' assembly accession version or include 'all' (latest + previous versions)`)
}

func (avf *AssemblyVersionFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (avf *AssemblyVersionFlag) GetAssemblyVersion() openapi.V2AssemblyDatasetDescriptorsFilterAssemblyVersion {
	// Default to current if the command is not an accessions command and user did not explicitly set the value
	if *avf.AssemblyVersion == Default {
		return openapi.V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYVERSION_CURRENT
	}
	return GenomeAssemblyVersionOpenapi[*avf.AssemblyVersion]
}

func isExplicitVersionRequested(accessions []string) bool {
	for _, acc := range accessions {
		if strings.Index(acc, ".") != -1 {
			return true
		}
	}
	return false
}

func stripVersion(accessions []string) []string {
	var accessions_no_ver []string
	for _, acc := range accessions {
		dot_found := strings.Index(acc, ".")
		if dot_found != -1 {
			accessions_no_ver = append(accessions_no_ver, acc[:dot_found])
		} else {
			accessions_no_ver = append(accessions_no_ver, acc)
		}
	}
	return accessions_no_ver
}

func (avf *AssemblyVersionFlag) UpdateForInputAccessions(accessions []string) []string {
	new_accs := accessions

	if *avf.AssemblyVersion == All_assemblies {
		// If user specified All_assemblies: Strip ".versions" to return all versions of requested assemblies
		new_accs = stripVersion(new_accs)
	} else {
		if isExplicitVersionRequested(new_accs) {
			if *avf.AssemblyVersion == Current {
				// If accessions have '.version', strip it if user explicitly request current accessions
				new_accs = stripVersion(new_accs)
			} else {
				// If accessions have a .version and AssemblyVersion==Default, use All_assemblies
				// to return the specifically requested version
				*avf.AssemblyVersion = All_assemblies
			}
		} else {
			// If accessions do not have '.version' and user did not request 'All_assemblies', return current version
			// (regardless of whether user picked 'Current' or 'Default')
			*avf.AssemblyVersion = Current
		}
	}

	return RemoveDuplicateStrings(new_accs)
}
