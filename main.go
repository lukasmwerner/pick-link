package main

import (
	"io"
	"log"
	"os"

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
