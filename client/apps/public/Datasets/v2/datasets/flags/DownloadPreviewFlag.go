package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type DownloadPreviewFlag struct {
	FlagInterface
	downloadPreview bool
}

func NewDownloadPreviewFlag() *DownloadPreviewFlag {
	return &DownloadPreviewFlag{downloadPreview: false}
}

func (dp *DownloadPreviewFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&dp.downloadPreview, "preview", false, "Show information about the requested data package")
}

func (dp *DownloadPreviewFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (dp *DownloadPreviewFlag) IsPreview() bool {
	return dp.downloadPreview
}
