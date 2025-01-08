//go:generate ${GO_PATH}/bin/goversioninfo -manifest=../../resource/goversioninfo.exe.manifest -64 -product-version=${VER} -file-version=${VER}

package main

import (
	"datasets_cli/v2/datasets"
)

func main() {
	datasets.Execute()
}
