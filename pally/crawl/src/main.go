package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
)

type pallytask struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Standard string `json:"standard"`
}

func main() {
	var website string
	var pally string
	var standard string

	flag.StringVar(&website, "site", "http://localhost/", "Website URL")
	flag.StringVar(&pally, "service", "http://localhost:3000/tasks", "Pa11y web service URL")
	flag.StringVar(&standard, "standard", "WCAG2AA", "WCAG2A, WCAG2AA or WCAG2AAA")
	flag.Parse()

	// parse url from command line
	url, _ := url.Parse(website)

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only single domain
		colly.AllowedDomains(url.Host),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// start with title tag on each page
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text

		// compile Pa11y task
		task := pallytask{}
		task.Name = title
		task.URL = e.Request.URL.String()
		task.Standard = standard
		text, _ := json.Marshal(task)

		// send request to Pa11y web service
		req, _ := http.NewRequest("POST", pally, bytes.NewBuffer(text))
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		fmt.Print(".")
	})

	// Start crawling
	c.Visit(url.String())
	fmt.Println()

}
