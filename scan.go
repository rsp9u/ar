package ar

import (
	"io/fs"
	"log"
	"path/filepath"
)

func ScanDirectory(root string, ignore func(string) bool) []string {
	ret := []string{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if ignore(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.IsDir() {
			// ret = append(ret, path)
			ret, err := Replace(path, opts.Src, opts.Dst)
			if err == nil {
				log.Println(ret[0])
				log.Println(ret[1])
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return ret
}
