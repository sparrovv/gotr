# gotr

Not officially supported way to access **google translate** from the command line. (Written in `go` for the speed and distribution ease)

**It can break at any point, so use it with caution.**

Said that, it has been working stable since the first release. (date of publication: 2014-01-06)


```
Usage:
  gotr <from> <to> <phrase>
  gotr [-s] <from> <to> <phrase>
  gotr (-h | --help)
  gotr (-l | --list)

Options:
  -h, --help     Show help
  -l, --list     List available languages
  -s, --speech   Enable speech synthesis
```

### Examples

```
$ gotr en pl equivocal
dwuznaczny
dwuznaczny, wymijający, podejrzany
```

```
$ gotr en zh "May the force be with you"
愿原力与你同在
```

#### Speech synthesis

**-s** flag enables speech synthesis

To check only the pronunciation, without translation, specify <to> to be same as <from>.
(Not all languages support speech though)

```
$ gotr -s en en obstreperous
obstreperous
```

#### TIP

Passing the `<from>` and `<to>` might be cumbersome, and usually we need translation for one language.
To ease up this dull task of passing <from> and <to>, you can consider creating aliases for your default languages:

`alias ep="gotr en pl"`
`alias pe="gotr pl en"`
`alias gees="gotr en en -s"`

You can download binaries from [releases](https://github.com/sparrovv/gotr/releases)

### Requirements

To use speech on OSX you need **afplay**.

### Other similar projects:

- https://github.com/pawurb/termit

### TODOS:

- [ ] enable speech on Linux.
- [ ] add timeout in case of the Network issues
- [ ] add brew formula
