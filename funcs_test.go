package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetHtmlTags(t *testing.T) {
	sReader := strings.NewReader(`
	<a href="test1.golang">
	<a href="test2.golang">
`)
	doc, _ := html.Parse(sReader)
	tags := GetHtmlTags(doc, "a", "href", nil)
	expected := 2
	if len(tags) < expected {
		t.Errorf("result '%d', expected '%d'", len(tags), expected)
	}
}
