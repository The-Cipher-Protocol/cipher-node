package main

import (
	_ "embed"

	"github.com/nnlgsakib/neth/command/root"
	"github.com/nnlgsakib/neth/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
