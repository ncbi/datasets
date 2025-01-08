package flags

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const dateFormat = "YYYY-MM-DD"
const dateDescr = " a specified date (free format, ISO 8601 " + dateFormat + " recommended)"

func CheckDateTime(dt string, flag string) (time.Time, string, error) {
	if dt == "" {
		return time.Time{}, "", nil
	}
	t, dateErr := dateparse.ParseAny(dt)
	if dateErr != nil {
		return time.Time{}, "", fmt.Errorf("--%s requires a valid date, ISO 8601 standard "+dateFormat+" is recommended", flag)
	}
	// Golang format accepts this special 'layout', equivalent to MM/DD/YYYY
	return t, t.Format("01/02/2006"), nil
}

const (
	VirusProteinReleasedAfterDesc string = "Limit to coronavirus genomes released on or after" + dateDescr
	VirusProteinUpdatedAfterDesc  string = "Limit to coronavirus genomes updated on or after" + dateDescr
	VirusGenomeUpdatedAfterDesc   string = "Limit to genomes updated on or after" + dateDescr
	GenomeReleasedBeforeDesc      string = "Limit to genomes released on or before" + dateDescr
	GenomeReleasedAfterDesc       string = "Limit to genomes released on or after" + dateDescr
)

// released-before
type ReleasedBeforeFlag struct {
	FlagInterface
	beforeFlagDesc              string
	releasedBeforeFlag          string
	releasedBeforeSanitized     string
	releasedBeforeDateSanitized time.Time
}

func NewReleasedBeforeFlag(beforeDesc string) *ReleasedBeforeFlag {
	rbf := &ReleasedBeforeFlag{
		beforeFlagDesc: beforeDesc,
	}
	return rbf
}

func (rbf *ReleasedBeforeFlag) ReleasedBeforeFlagDate() string {
	return rbf.releasedBeforeSanitized
}

func (uaf *ReleasedBeforeFlag) ReleasedBeforeFlagDateAsTime() time.Time {
	return uaf.releasedBeforeDateSanitized
}

func (rbf *ReleasedBeforeFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&rbf.releasedBeforeFlag, "released-before", "", rbf.beforeFlagDesc)
}

func (rbf *ReleasedBeforeFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	rbf.releasedBeforeDateSanitized, rbf.releasedBeforeSanitized, err = CheckDateTime(rbf.releasedBeforeFlag, "released-before")
	return
}

// released-after
type ReleasedAfterFlag struct {
	FlagInterface
	afterFlagDesc              string
	releasedAfterFlag          string
	releasedAfterSanitized     string
	releasedAfterDateSanitized time.Time
}

func NewReleasedAfterFlag(afterDesc string) *ReleasedAfterFlag {
	raf := &ReleasedAfterFlag{
		afterFlagDesc: afterDesc,
	}
	return raf
}

func (raf *ReleasedAfterFlag) ReleasedAfterFlagDate() string {
	return raf.releasedAfterSanitized
}

func (uaf *ReleasedAfterFlag) ReleasedAfterFlagDateAsTime() time.Time {
	return uaf.releasedAfterDateSanitized
}

func (raf *ReleasedAfterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&raf.releasedAfterFlag, "released-after", "", raf.afterFlagDesc)
}

func (raf *ReleasedAfterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	raf.releasedAfterDateSanitized, raf.releasedAfterSanitized, err = CheckDateTime(raf.releasedAfterFlag, "released-after")
	return
}

// updated-after
type UpdatedAfterFlag struct {
	FlagInterface
	updatedAfterFlagDesc      string
	updatedAfterFlag          string
	updatedAfterSanitized     string
	updatedAfterDateSanitized time.Time
}

func NewUpdatedAfterFlag(updatedAfterDesc string) *UpdatedAfterFlag {
	uaf := &UpdatedAfterFlag{
		updatedAfterFlagDesc: updatedAfterDesc,
	}
	return uaf
}

func (uaf *UpdatedAfterFlag) UpdatedAfterFlagDate() string {
	return uaf.updatedAfterSanitized
}

func (uaf *UpdatedAfterFlag) UpdatedAfterFlagDateAsTime() time.Time {
	return uaf.updatedAfterDateSanitized
}

func (uaf *UpdatedAfterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&uaf.updatedAfterFlag, "updated-after", "", uaf.updatedAfterFlagDesc)
}

func (uaf *UpdatedAfterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	uaf.updatedAfterDateSanitized, uaf.updatedAfterSanitized, err = CheckDateTime(uaf.updatedAfterFlag, "updated-after")
	return
}
