package main

import (
	"github.com/docopt/docopt-go"
	"github.com/sparrovv/gotr/cli"
)

var usage string = `google translate in terminal

Usage:
  gotr <from> <to> <phrase>
  gotr [-sdi] <from> <to> <phrase>
  gotr -h | --help
  gotr --history
  gotr --list

Options:
  -h, --help        Show help
  -s, --speech      Enable speech synthesis
  -d, --disablelog  Disable query loging into ~/.gotr_history
  -i, --ignore      Ignore extra meanings
  --list            Print list of available languages
  --history         Print log history in JSON
`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "Gotr", false)
	cli.Run(args, usage)
}
