package main

import (
	"flag"
	"fmt"
	"strings"
	"thinknetica_go/GoSearch/pkg/crawler"
	"thinknetica_go/GoSearch/pkg/crawler/spider"
)

func main() {

	var word string
	flag.StringVar(&word, "s", "", "search links")
	flag.Parse()

	pages := [2]string{"https://go.dev/", "https://golang.org/"}
	fmt.Println(pages)

	var doc []crawler.Document

	for _, page := range pages {
		result := scanPage(page, 2)
		doc = append(doc, result...)
	}

	links := searchLinks(doc, word)
	fmt.Println(links)
}

func scanPage(s string, depth int) []crawler.Document {
	sp := spider.New()
	result, err := sp.Scan(s, depth)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}

func searchLinks(d []crawler.Document, s string) []string {
	var links []string
	for _, link := range d {
		if strings.Contains(link.Title, s) {
			links = append(links, link.URL)
		}
	}
	return links
}
