package main

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/cli/browser"
	"github.com/lukasmwerner/pick-link/links"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err.Error())
		return
	}
	l := links.ExtractLinks(string(b))

	if len(os.Args) == 2 {
		index := os.Args[1]
		i, err := strconv.Atoi(index)
		if err != nil {
			log.Println(err.Error())
			return
		}
		link := l[i+1]
		err = browser.OpenURL(link)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	opts := []huh.Option[string]{}
	for _, link := range l {
		opts = append(opts, huh.NewOption(link, link))
	}

	var link string

	err = huh.NewSelect[string]().Title("Pick a Link").Options(opts...).Value(&link).Run()
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = browser.OpenURL(link)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
