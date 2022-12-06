package ar

import (
	"io/fs"
	"path/filepath"
	"time"
)

func ScanDirectory(root string, ignore func(string) bool) {
	startTime := time.Now()
	startTime = startTime.Add(time.Second * -1)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if ignore(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.IsDir() {
			// check whether the file has been modified or not after loop is started
			finfo, err := d.Info()
			if err != nil {
				return nil
			}
			mtime := finfo.ModTime()
			if mtime.After(startTime) {
				return nil
			}

			// throw to stream
			replaceTargets <- path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
