# gotr

**Google translate** in the command line.

Disclaimer:

This is not an officially supported way of accessing Google Translate. **It can break at any time, so don't use it on production systems**. (That said, it has been stable since the first release: 2014-01-06.)


```
Usage:
  gotr <from> <to> <phrase>
  gotr [-l] [-s] [-i] <from> <to> <phrase>
  gotr -h | --help
  gotr --history
  gotr --list

Options:
  -h, --help    Show help
  -s, --speech  Enable speech synthesis
  -l, --log     Log query into ~/.gotr_history
  -i, --ignore  Ignore extra meanings
  --list        List available languages
  --history     Print log history in JSON to the SDTOUT
```

### How to install

You can download binaries from [releases](https://github.com/sparrovv/gotr/releases).

#### Homebrew

```
brew tap sparrovv/tap
brew install gotr
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

Unfortunately it's not available to all languages and some less popular ones don't have it yet.

#### Tips

- To only check the pronunciation, specify `<to>` to be same as `<from>`.

```
$ gotr -s en en obstreperous
obstreperous
```

- Passing `<from>` and `<to>` might be cumbersome.
You might consider creating aliases for your default languages:

`alias ep="gotr -l en pl"`

`alias pe="gotr -l pl en"`

`alias gees="gotr -l en en -s"`

### Requirements

To play audio, install **afplay**(OSX) or **mpg123**.

### Other similar projects:

- https://github.com/pawurb/termit (written in ruby)
