package flags

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type DebugFlag struct {
	FlagInterface
	RunDebug bool
}

func (df *DebugFlag) useEnv(envVarName, argName string) (val string) {
	val = os.Getenv(envVarName)
	return
}

func (df *DebugFlag) useEnvBool(envVarName, argName string, defaultVal bool) bool {
	if val, err := strconv.ParseBool(df.useEnv(envVarName, argName)); err == nil {
		return val
	}
	return defaultVal
}

func NewDebugFlag() *DebugFlag {
	return &DebugFlag{}
}

func (df *DebugFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&df.RunDebug, "debug", df.useEnvBool("DATASETS_DEBUG", "debug", false), "Emit debugging info")
}

func (df *DebugFlag) Debug() bool {
	return df.RunDebug
}

func (df *DebugFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
