package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://usn.ubuntu.com/atom.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("entry content dt").Each(func(_ int, s *goquery.Selection) {
		os := s.Text()
		pkg := s.Next().Find("a").First().Text()
		ver := s.Next().Find("a").Last().Text()
		fmt.Printf("%v : %v-%v\n", os, pkg, ver)
	})
}
