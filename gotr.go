package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/sparrovv/gotr/googletranslate"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotr"
	app.Usage = "gotr --from=en to=pl phrase"
	app.Flags = []cli.Flag{
		cli.StringFlag{"from, f", "", "translate from"},
		cli.StringFlag{"to, t", "", "translate to"},
	}
	app.Action = func(c *cli.Context) {
		from := strings.TrimSpace(c.String("from"))
		to := strings.TrimSpace(c.String("to"))
		term := strings.TrimSpace(strings.Join(c.Args(), " "))

		if len(from) == 0 || len(to) == 0 || len(term) == 0 {
			fmt.Println("Usage:\n  gotr --from=en to=pl phrase")
			os.Exit(1)
		}

		phrase, err := googletranslate.Translate("https://translate.google.com/translate_a/t", from, to, term)
		if err != nil {
			panic(err)
		}
		fmt.Println(phrase.Translation)
		fmt.Println(strings.Join(phrase.ExtraMeanings, ", "))
	}
	app.Run(os.Args)
}
