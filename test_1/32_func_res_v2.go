package main

import (
	"os"
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func funcResV2() {
	words, images, _ := CountWordsAndImages(os.Args[1])
	fmt.Println(words, images)
}

func CountWordsAndImages(url string) (words, images int, err error){
	resp, err := http.Get(url)
	if err != nil {
		return 
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordAndImages(doc)
	return

}

func countWordAndImages(n *html.Node)(words, images int){
    tests, images := visit3(nil, 0, n)
    for _, v := range tests {
    	v = strings.Trim(strings.TrimSpace(v), "\r\n")
    	if v == "" {
    		continue
    	}
    	words += strings.Count(v, "")
    }
    return
}

func visit3(tests []string, imgs int, n *html.Node)([]string, int){
    if n.Type == html.TextNode {
    	tests = append(tests, n.Data)
    }
    if n.Type == html.ElementNode && (n.Data == "img") {
    	imgs ++
    }
    for c := n.FirstChild; c != nil; c= c.NextSibling {
    	if c.Data == "script" || c.Data == "style" {
    		continue
    	}
    	tests, imgs = visit3(tests, imgs, c)
    }
    return tests, imgs
}