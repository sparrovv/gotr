# gotr

Hacky way to access **google translate** from the command line.  (Written in `go` for the speed)

**It can break at any point, so use it with caution.**

Said that, it has been working stable for couple months now. (date of publication: 06/01/2014)

### How to use it:

```
$ gotr --from=en --to=pl equivocal
dwuznaczny
dwuznaczny, wymijający, podejrzany
```

Passing the `--from` and `--to` is a bit cumbersome so I recommend creating aliases for your default languages:

`alias ep="gotr --from=en --to=pl"`

`alias pe="gotr --from=pl --to=en"`

You can download binaries from here: <link>

