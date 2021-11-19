package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func checkIfIsExpected(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func TestGetHtmlTags(t *testing.T) {
	sReader := strings.NewReader(`
	<a href="test1.golang">
	<a href="test2.golang">
`)
	doc, _ := html.Parse(sReader)
	tags := getHtmlTags(doc, "a", "href", nil)
	expected := 2
	if len(tags) < expected {
		t.Errorf("result '%d', expected '%d'", len(tags), expected)
	}
}

func TestLogBadLinksFound(t *testing.T) {
	links := []string{"https://badlink.com1"}
	buffer := bytes.Buffer{}
	logBadLinksFound(&buffer, links)
	expected := `-------------------
Bad links found

https://badlink.com1

-------------------
`
	checkIfIsExpected(t, buffer.String(), expected)
}
