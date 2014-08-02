# gotr

Not officially supported way to access **google translate** from the command line. (Written in `go` for the speed and distribution ease)

**It can break at any point, so use it with caution.**

Said that, it has been working stable since the first release. (date of publication: 2014-01-06)

### How to use it:

```
$ gotr --from=en --to=pl equivocal
dwuznaczny
dwuznaczny, wymijający, podejrzany
```

#### Speech synthesis

**-s** flag enables speech synthesis

```
$ gotr -s --from=en --to=pl equivocal
dwuznaczny
dwuznaczny, wymijający, podejrzany
```

If you need only speech, specify --to flag to be same as --from.

```
$ gotr -s --from=en --to=en equivocal
equivocal
```

#### Available languages

To get a list of available languages' codes run:

```
$ gotr --list
```

#### TIP

Passing the `--from` and `--to` is cumbersome, and usually we only need translation for one language.
To ease up this dull task of specifying options, you can create aliases for your default languages:

`alias ep="gotr --from=en --to=pl"`

`alias pe="gotr --from=pl --to=en"`

`alias gees="gotr --from=en --to=en -s"`

You can download binaries from [releases](https://github.com/sparrovv/gotr/releases)

### Requirements

To use speech on OSX you need **afplay**.

### Other similar projects:

- https://github.com/pawurb/termit

### TODOS:

- [ ] enable speech on Linux.
- [ ] add timeout in case of the Network issues
- [ ] add brew formula
