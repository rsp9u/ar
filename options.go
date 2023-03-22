package ar

import (
	"errors"

	"github.com/juju/gnuflag"
)

type Options struct {
	Src      string
	Dst      string
	Targets  []string
	IsDryRun bool
}

var opts = Options{}

func ParseOptions() error {
	gnuflag.BoolVar(&opts.IsDryRun, "dry", false, "")
	gnuflag.BoolVar(&opts.IsDryRun, "d", false, "not to overwrite files, only output diff")
	gnuflag.Parse(true)
	args := gnuflag.Args()
	if len(args) < 2 {
		return errors.New("too few arguments.")
	}
	opts.Src = args[0]
	opts.Dst = args[1]
	if len(args) < 3 {
		opts.Targets = []string{}
	} else {
		opts.Targets = args[2:]
	}
	return nil
}
