package ar

import (
	"flag"
)

type Options struct {
	Src      string
	Dst      string
	IsDryRun bool
}

var opts = Options{}

func ParseOptions() {
	flag.BoolVar(&opts.IsDryRun, "dry", false, "not to rewrite files, only output diff (long)")
	flag.BoolVar(&opts.IsDryRun, "d", false, "not to rewrite files, only output diff (short)")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		panic("too few arguments.")
	}
	opts.Src = args[0]
	opts.Dst = args[1]
}
