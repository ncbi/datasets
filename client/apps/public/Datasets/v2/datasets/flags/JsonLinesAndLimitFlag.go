package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type JsonLinesAndLimitFlag struct {
	FlagInterface
	limitFlag *LimitFlag
	jsonLines bool
}

func NewJsonLineAndLimitFlag(limitType string) *JsonLinesAndLimitFlag {
	return &JsonLinesAndLimitFlag{
		limitFlag: NewLimitFlag(limitType),
		jsonLines: false,
	}
}

func (jll *JsonLinesAndLimitFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&jll.jsonLines, "as-json-lines", false, "Output results in JSON Lines format")
	jll.limitFlag.RegisterFlags(flags)
}

func (jll *JsonLinesAndLimitFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	// Add error checking here.
	err = jll.limitFlag.PreRunE(cmd, args)
	if err != nil {
		return err
	}
	if strings.ToLower(jll.limitFlag.LimitRaw()) == "none" && jll.jsonLines {
		err = fmt.Errorf("Setting 'limit' to 0 is not compatible with 'as-json-lines'")
		return
	}
	return nil
}

func (jll *JsonLinesAndLimitFlag) JsonLines() bool {
	return jll.jsonLines
}

func (jll *JsonLinesAndLimitFlag) RetrievalCount() int {
	return jll.limitFlag.maxRetrieval
}

func (jll *JsonLinesAndLimitFlag) CountOnly() bool {
	return jll.limitFlag.countOnly
}

func (jll *JsonLinesAndLimitFlag) LimitRaw() string {
	return jll.limitFlag.limitRaw
}
