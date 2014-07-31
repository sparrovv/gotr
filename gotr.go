package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/sparrovv/gotr/googletranslate"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotr"
	app.Usage = "google translate in the terminal"
	app.Flags = []cli.Flag{
		cli.StringFlag{"from, f", "", "translate from"},
		cli.StringFlag{"to, t", "", "translate to"},
		cli.BoolFlag{"list, l", "list of languages"},
		cli.BoolFlag{"speech, s", "download and execute run audio file"},
	}
	app.Action = func(c *cli.Context) {
		from := strings.TrimSpace(c.String("from"))
		to := strings.TrimSpace(c.String("to"))
		term := strings.TrimSpace(strings.Join(c.Args(), " "))

		if c.Bool("list") == true {
			fmt.Println(`Supported languages:`)
			fmt.Println(googletranslate.ListLanguages())
			os.Exit(1)
		}

		if len(from) == 0 || len(to) == 0 || len(term) == 0 {
			fmt.Println("  Usage: " + usage())
			os.Exit(1)
		}

		phrase, err := googletranslate.Translate(from, to, term)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println(phrase.Translation)
		fmt.Println(strings.Join(phrase.ExtraMeanings, ", "))

		if c.Bool("speech") == true {
			path := "/tmp/gotr.speech.file.mpg"
			err := googletranslate.FetchSoundFile(to, phrase.Translation, path)
			if err == nil {
				fmt.Println("\nSpeech support on OSX:")
				fmt.Println("afplay " + path)
			}
		}

	}
	app.Run(os.Args)
}

func usage() string {
	return `
    gotr --from=en to=pl phrase
    gotr --list - returns list of available languages
  `
}
