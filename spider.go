package main

import (
	"github.com/anaskhan96/soup"
	"os"
	"fmt"
)

func main() {

	resp, err := soup.Get("https://www.metal-archives.com/browse/country")

	if (err != nil) {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "countryCol").FindAll("a")

	for _, link := range links {

		fmt.Println(link)
	}
}
