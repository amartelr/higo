package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    colly domain.org ...\n")
		os.Exit(1)
	}

	argsWithoutProg := os.Args[1]

	c := colly.NewCollector(
		//colly.AllowedDomains(argsWithoutProg),
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://" + argsWithoutProg)
}
