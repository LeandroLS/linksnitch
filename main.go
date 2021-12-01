package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/LeandroLS/valoop"
	"github.com/gomarkdown/markdown"
	"golang.org/x/net/html"
)

type BadLink struct {
	url        string
	statusCode int
}

func main() {
	markDown, err := os.ReadFile("README.md")
	handleErr(err)
	htmlBytes := markdown.ToHTML(markDown, nil, nil)
	bReader := bytes.NewReader(htmlBytes)
	htmlParsed, err := html.Parse(bReader)
	handleErr(err)
	links := getHtmlTags(htmlParsed, "a", "href", nil)
	if len(links) < 1 {
		fmt.Println("No links found.")
		return
	}
	badLinks := getBadLinks(links)
	if len(badLinks) < 1 {
		fmt.Println("All links works.")
		return
	}
	logBadLinksFound(os.Stdout, badLinks)
	os.Exit(1)
}

func getAllowedStatusCodes() []int {
	statusCodesJson := os.Getenv("INPUT_ALLOWEDSTATUSCODES")
	var statusCodeArr []int
	err := json.Unmarshal([]byte(statusCodesJson), &statusCodeArr)
	handleErr(err)
	statusCodeArr = append(statusCodeArr, 200)
	return statusCodeArr
}

func getBadLinks(links []string) []BadLink {
	statusCodesArr := getAllowedStatusCodes()
	var badLinks []BadLink
	ms, _ := time.ParseDuration("0.35s")
	for i := 0; i < len(links); i++ {
		resp, err := http.Get(links[i])
		if err != nil {
			badLinks = append(badLinks, BadLink{links[i], 0})
			continue
		}
		defer resp.Body.Close()

		if !valoop.IntSliceContains(statusCodesArr, resp.StatusCode) {
			badLinks = append(badLinks, BadLink{links[i], resp.StatusCode})
		}
		time.Sleep(ms)
	}
	return badLinks
}

func getHtmlTags(n *html.Node, rawHtmlTag string, htmlTagKey string, tags []string) []string {
	if n.Type == html.ElementNode && n.Data == rawHtmlTag {
		for _, a := range n.Attr {
			if a.Key == htmlTagKey {
				tags = append(tags, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tags = getHtmlTags(c, rawHtmlTag, htmlTagKey, tags)
	}
	return tags
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func logBadLinksFound(writer io.Writer, links []BadLink) {
	w := tabwriter.NewWriter(writer, 0, 8, 2, ' ', 0)
	format := "%v\t%v\n"
	fmt.Fprintf(w, format, "Site", "Status Code")
	fmt.Fprintf(w, format, "----", "-----------")

	for _, link := range links {
		fmt.Fprintf(w, format, link.url, link.statusCode)
	}
	w.Flush()
}
