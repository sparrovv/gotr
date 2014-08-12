# gotr

Not officially supported way to access **google translate** from the command line. (Written in `go` for the speed and distribution ease)

**It can break at any point, so use it with caution.**

Said that, it has been working stable since the first release. (date of publication: 2014-01-06)


```
Usage:
  gotr <from> <to> <phrase>
  gotr [-l] [-s] <from> <to> <phrase>
  gotr -h | --help
  gotr --history
  gotr --list

Options:
  -h, --help    Show help
  -s, --speech  Enable speech synthesis
  -l, --log     Log query into ~/.gotr_history
  --list        List available languages
  --history     Print log history in JSON to the SDTOUT
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

To check only the pronunciation, without translation, specify `<to>` to be same as `<from>`.
(Not all languages support speech though)

```
$ gotr -s en en obstreperous
obstreperous
```

#### TIP

Passing the `<from>` and `<to>` might be cumbersome.
To ease up this dull task of passing <from> and <to>, you can consider creating aliases for your default languages:

`alias ep="gotr -l en pl"`
`alias pe="gotr -l pl en"`
`alias gees="gotr -l en en -s"`

You can download binaries from [releases](https://github.com/sparrovv/gotr/releases)

### Requirements

To play audio from command line, install **afplay**(OSX) or **mpg123**.

### Other similar projects:

- https://github.com/pawurb/termit

### TODOS:

- [ ] add timeout in case of the Network issues
