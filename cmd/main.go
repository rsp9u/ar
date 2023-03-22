package main

import (
	"fmt"

	"github.com/rsp9u/ar"
)

func main() {
	err := ar.ParseOptions()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	ar.DecideIgnores(".")
	rwg := ar.RunReplaceWorker(4)
	pwg := ar.RunPrintWorker()

	ar.ScanDirectories(ar.Ignore)
	ar.DoneScan()

	rwg.Wait()
	ar.DoneReplace()

	pwg.Wait()
}
