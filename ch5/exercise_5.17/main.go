// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	for _, node := range ElementsByTagName(doc, "h1", "img") {
		fmt.Printf("%s\n", node.Data)
	}

	return nil
}

func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {
	queriedTagNode := func(n *html.Node) bool {
		if n == nil {
			return false
		}
		nodeTag := n.Data
		for _, data := range tags {
			if nodeTag == data {
				return true
			}
		}
		return false
	}
	return forEachNode(doc, queriedTagNode)
}

var count int

func forEachNode(n *html.Node, nodeTagChecker func(*html.Node) bool) []*html.Node {
	nodes := []*html.Node{}
	if nodeTagChecker(n) {
		nodes = append(nodes, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, forEachNode(c, nodeTagChecker)...)
	}
	return nodes
}
