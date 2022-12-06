package main

import (
	"github.com/rsp9u/ar"
)

func main() {
	ar.ParseOptions()

	ar.DecideIgnores(".")
	rwg := ar.RunReplaceWorker(4)
	pwg := ar.RunPrintWorker()

	ar.ScanDirectory(".", ar.Ignore)
	ar.DoneScan()

	rwg.Wait()
	ar.DoneReplace()

	pwg.Wait()
}
