package main

import (
	"fmt"

	"github.com/rsp9u/ar"
)

func main() {
	ar.ParseOptions()

	ar.DecideIgnores(".")
	entries := ar.ScanDirectory(".", ar.Ignore)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
