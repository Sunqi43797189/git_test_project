package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func funcRes() {
	for _, url := range os.Args[1:] {
		fmt.Println(os.Args[1:])
		links, err := findLinks(url)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(links)
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a"{
		for _, a := range n.Attr {
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}
	return visit(visit(links, n.FirstChild), n.NextSibling)
}