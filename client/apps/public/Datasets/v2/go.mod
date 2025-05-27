module datasets_cli/v2

go 1.22.0

toolchain go1.23.4

require (
	bou.ke/monkey v1.0.2
	github.com/antihax/optional v1.0.0
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/docker/go-units v0.5.0
	github.com/gosuri/uiprogress v0.0.1
	// gosuri/uilive v0.0.4 causes hang on Apple silicon macs with Go >=1.21, see
	// - https://github.com/ncbi/datasets/issues/490
	// - https://github.com/gosuri/uiprogress/issues/53
	// It is actually only an indirect dependency, but we need to pin it due to bug
	github.com/gosuri/uilive v0.0.3
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.9.0
	github.com/thediveo/enumflag v0.10.1
	gitlab.com/metakeule/fmtdate v1.2.2
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c
	golang.org/x/text v0.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
