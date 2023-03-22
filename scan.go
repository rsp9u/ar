package ar

import (
	"io/fs"
	"path/filepath"
	"time"
)

func ScanDirectories(ignore func(string) bool) {
	startTime := time.Now()
	startTime = startTime.Add(time.Second * -1)
	scannedAll := []string{}

	for _, dir := range(opts.Targets) {
		scanned := scanDirectory(dir, ignore, startTime, scannedAll)
		scannedAll = append(scannedAll, scanned...)
	}
}

func scanDirectory(root string, ignore func(string) bool, startTime time.Time, scannedAll []string) []string {
	scanned := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if ignore(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.IsDir() {
			if contains(scannedAll, path) {
				return nil
			}
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
			scanned = append(scanned, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return scanned
}
