package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"golang.org/x/net/html"
)

func main() {
	markDown, err := os.ReadFile("README.md")
	check(err)
	htmlBytes := markdown.ToHTML(markDown, nil, nil)
	bReader := bytes.NewReader(htmlBytes)
	htmlParsed, err := html.Parse(bReader)
	check(err)
	links := GetHtmlTags(htmlParsed, "a", "href", nil)
	ms, _ := time.ParseDuration("0.35s")
	var badLinks []string
	if len(links) < 1 {
		fmt.Println("No links found.")
	} else {
		for i := 0; i < len(links); i++ {
			resp, err := http.Get(links[i])
			if err != nil {
				badLinks = append(badLinks, links[i])
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				badLinks = append(badLinks, links[i])
			}
			time.Sleep(ms)
		}
	}
	if len(badLinks) >= 1 {
		log.Fatalf("Bad Links %s", badLinks)
	} else {
		fmt.Println("All links works.")
	}
}

func GetHtmlTags(n *html.Node, rawHtmlTag string, htmlTagKey string, tags []string) []string {
	if n.Type == html.ElementNode && n.Data == rawHtmlTag {
		for _, a := range n.Attr {
			if a.Key == htmlTagKey {
				tags = append(tags, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tags = GetHtmlTags(c, rawHtmlTag, htmlTagKey, tags)
	}
	return tags
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
