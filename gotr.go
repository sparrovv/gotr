package main

import (
	"github.com/docopt/docopt-go"
	"github.com/sparrovv/gotr/cli"
)

var usage string = `google translate in terminal

Usage:
  gotr <from> <to> <phrase>
  gotr [-sp] <from> <to> <phrase>
  gotr (-h | --help)
  gotr (-l | --list)

Options:
  -h, --help     Show help
  -l, --list     List available languages
  -s, --speech   Enable speech synthesis
  -p, --persist  Persist query into ~/.gotr_history
`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "Gotr", false)
	cli.Run(args, usage)
}
