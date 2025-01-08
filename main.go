package main

import (
	_ "embed"

	"github.com/The-Cipher-Protocol/cipher-node/cmd/root"
)

func main() {
	//licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
