package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oke-py/usn/feed"

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
	doc.Find("entry").Each(func(_ int, s *goquery.Selection) {
		usn := feed.ExtractUsnTitle(s.Find("title").Text())
		fmt.Println(usn)

		s.Find("content dt").Each(func(_ int, s2 *goquery.Selection) {
			os := s2.Text()
			pkg := s2.Next().Find("a").First().Text()
			ver := s2.Next().Find("a").Last().Text()
			fmt.Printf("%v : %v-%v\n", os, pkg, ver)
		})
	})
}
