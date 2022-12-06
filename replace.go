package ar

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"golang.org/x/tools/godoc/util"
)

const (
	ERR_BINARY_FILE = "it's a binary file"
	ERR_NO_REPLACE  = "no replacement"
)

func Replace(filepath string, src, dst string) ([2]string, error) {
	empty_return := [2]string{"", ""}

	f, err := os.Open(filepath)
	if err != nil {
		return empty_return, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return empty_return, err
	}

	if !util.IsText(bytes) {
		return empty_return, errors.New(ERR_BINARY_FILE)
	}
	content := string(bytes)

	match := strings.Index(content, src)
	if match == -1 {
		return empty_return, errors.New(ERR_NO_REPLACE)
	}

	replaced := strings.ReplaceAll(content, src, dst)
	ret := [2]string{content, replaced}

	if !opts.IsDryRun {
		wf, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return ret, err
		}
		defer wf.Close()
		_, err = wf.WriteString(replaced)
		if err != nil {
			return ret, err
		}
	}

	return ret, err
}

func RunReplaceWorker(workerNum int) *sync.WaitGroup {
	var wg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case path := <-replaceTargets:
					ret, err := Replace(path, opts.Src, opts.Dst)
					if err == nil {
						printableDiffs <- [3]string{path, ret[0], ret[1]}
					}
				default:
					if isDoneScanning {
						return
					}
				}
			}
		}()
	}
	return &wg
}
