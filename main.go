package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"golang.org/x/net/html"
)

func main() {
	markDown, err := os.ReadFile("README.md")
	handleErr(err)
	htmlBytes := markdown.ToHTML(markDown, nil, nil)
	bReader := bytes.NewReader(htmlBytes)
	htmlParsed, err := html.Parse(bReader)
	handleErr(err)
	links := GetHtmlTags(htmlParsed, "a", "href", nil)
	ms, _ := time.ParseDuration("0.35s")
	var badLinks []string
	statusCodesJson := os.Getenv("INPUT_ALLOWEDSTATUSCODES")
	var statusCodeArr []int
	err = json.Unmarshal([]byte(statusCodesJson), &statusCodeArr)
	handleErr(err)
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

			if !contains(statusCodeArr, resp.StatusCode) {
				badLinks = append(badLinks, links[i])
			}
			time.Sleep(ms)
		}
	}
	if len(badLinks) >= 1 {
		logBadLinksFound(os.Stdout, badLinks)
		os.Exit(1)
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

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func logBadLinksFound(writer io.Writer, links []string) {
	templateStr := `-------------------
Bad links found
{{range $val := .}}
{{$val}}
{{end}}
-------------------
`
	tmpl, err := template.New("LogMessage").Parse(templateStr)
	handleErr(err)
	err = tmpl.Execute(writer, links)
	handleErr(err)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
