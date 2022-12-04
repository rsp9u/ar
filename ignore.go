package ar

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ignore_files = []string{
	".ignore",
	".gitignore",
	".git/info/exclude",
}

var ignore_targets_default = []string{
	"node_modules",
}

var ignore_targets = []string{}

// https://git-scm.com/docs/gitignore
func ParseIgnorePattern(line string) (string, error) {
	if len(line) == 0 {
		return "", errors.New("skip")
	}
	if strings.HasPrefix(line, "#") {
		return "", errors.New("skip")
	}
	if strings.HasPrefix(line, "/") || strings.HasPrefix(line, "\\") {
		line = line[1:]
	}
	return line, nil
}

func DecideIgnores(root string) {
	for _, t := range ignore_targets_default {
		ignore_targets = append(ignore_targets, t)
	}
	for _, ignore_file := range ignore_files {
		path := filepath.Join(root, ignore_file)
		f, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		sc := bufio.NewScanner(f)
		sc.Split(bufio.ScanLines)
		for sc.Scan() {
			pattern, err := ParseIgnorePattern(sc.Text())
			if err != nil {
				continue
			}
			ignore_targets = append(ignore_targets, pattern)
		}
	}
}

func Ignore(name string) bool {
	// ignore hidden files
	if len(name) > 1 && strings.HasPrefix(name, ".") {
		return true
	}
	// ignore listed files
	for _, ignore_target := range ignore_targets {
		if name == ignore_target {
			return true
		}
	}
	// don't ignore
	return false
}
