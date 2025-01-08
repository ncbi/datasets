package flags

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type LimitFlag struct {
	FlagInterface
	limitType    string // genome, gene, or virus
	limitRaw     string // all, none or number
	maxRetrieval int
	countOnly    bool
}

func NewLimitFlag(limitType string) *LimitFlag {
	lf := &LimitFlag{
		limitType: limitType,
	}
	return lf
}

// Unused AFAICT
// func NewLimitFlagFor(limitType string, limitArgument string) *LimitFlag {
// 	lf := &LimitFlag{
// 		limitType: limitType,
// 		limitRaw:  limitArgument,
// 	}
// 	return lf
// }

func (lf *LimitFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&lf.limitRaw, "limit", "all", strings.Replace(`Limit the number of TYPE summaries returned
  * all:      returns all matching TYPE summaries
  * a number: returns the specified number of matching TYPE summaries
    `, "TYPE", lf.limitType, -1))
}

func (lf *LimitFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return lf.SetRetrievalAndCount()
}

func (lf *LimitFlag) RetrievalCount() int {
	return lf.maxRetrieval
}

func (lf *LimitFlag) CountOnly() bool {
	return lf.countOnly
}

func (lf *LimitFlag) LimitRaw() string {
	return lf.limitRaw
}

// Helper function
func (lf *LimitFlag) SetRetrievalAndCount() (err error) {
	lf.maxRetrieval = math.MaxInt
	lf.countOnly = false
	if lf.limitRaw == "" {
		return
	}
	lowerCaseLimit := strings.ToLower(lf.limitRaw)
	if lowerCaseLimit == "none" || lowerCaseLimit == "0" {
		lf.maxRetrieval = 1
		lf.countOnly = true
		return
	}

	if lowerCaseLimit != "all" {
		lf.maxRetrieval, err = strconv.Atoi(lowerCaseLimit)
		if err != nil {
			err = fmt.Errorf("Invalid 'limit' value %s. Must be 'all', 'none', or a number.", lowerCaseLimit)
		} else if lf.maxRetrieval < 0 {
			lf.maxRetrieval = math.MaxInt
		}
	}
	return
}
