package ar

import (
	"strings"
	"sync"

	"github.com/fatih/color"
	godiffpatch "github.com/sourcegraph/go-diff-patch"
)

func colorPrint(diff string) {
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "@@") {
			color.Cyan(line)
		} else if strings.HasPrefix(line, "---") || strings.HasPrefix(line, "+++") {
			color.Yellow(line)
		} else if strings.HasPrefix(line, "-") {
			color.Red(line)
		} else if strings.HasPrefix(line, "+") {
			color.Green(line)
		} else {
			color.White(line)
		}
	}
}

func PrintDiff(filepath, before, after string) {
	diff := godiffpatch.GeneratePatch(filepath, before, after)
	colorPrint(diff)
}

func RunPrintWorker() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case diff := <-printableDiffs:
				PrintDiff(diff[0], diff[1], diff[2])
			default:
				if isDoneReplacing {
					return
				}
			}
		}
	}()
	return &wg
}
