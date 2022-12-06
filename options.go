package ar

import (
	"github.com/juju/gnuflag"
)

type Options struct {
	Src      string
	Dst      string
	IsDryRun bool
}

var opts = Options{}

func ParseOptions() {
	gnuflag.BoolVar(&opts.IsDryRun, "dry", false, "")
	gnuflag.BoolVar(&opts.IsDryRun, "d", false, "not to overwrite files, only output diff")
	gnuflag.Parse(true)
	args := gnuflag.Args()
	if len(args) < 2 {
		panic("too few arguments.")
	}
	opts.Src = args[0]
	opts.Dst = args[1]
}
