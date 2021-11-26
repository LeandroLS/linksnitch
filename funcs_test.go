package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/LeandroLS/valoop"
	"golang.org/x/net/html"
)

func TestGetHtmlTags(t *testing.T) {
	sReader := strings.NewReader(`
	<a href="test1.golang">
	<a href="test2.golang">
`)
	doc, _ := html.Parse(sReader)
	tags := getHtmlTags(doc, "a", "href", nil)
	expected := 2
	if !valoop.IsSameValue(len(tags), expected) {
		t.Errorf("result '%d', expected '%d'", len(tags), expected)
	}
}

func TestLogBadLinksFound(t *testing.T) {
	links := []BadLink{
		{"https://badlink.com", 404},
	}
	buffer := bytes.Buffer{}
	logBadLinksFound(&buffer, links)
	expected := `Site                 Status Code
----                 -----------
https://badlink.com  404
`
	if !valoop.IsSameValue(buffer.String(), expected) {
		t.Errorf("result %s\n", buffer.String())
		t.Errorf("expected %s", expected)
	}
}

func TestGetAllowedStatusCodes(t *testing.T) {
	os.Setenv("INPUT_ALLOWEDSTATUSCODES", "[200]")
	statusCodes := getAllowedStatusCodes()
	if !valoop.IsSameValue(statusCodes[0], 200) {
		t.Errorf("result '%d', expected '%d'", statusCodes[0], 200)
	}
}
