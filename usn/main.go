package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/oke-py/usn/feed"
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

	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("usn")

	doc.Find("entry").Each(func(_ int, s *goquery.Selection) {
		notice := feed.GetNotice(s)
		fmt.Println(notice)

		if err := table.Put(notice).Run(); err != nil {
			fmt.Println("err")
			panic(err.Error())
		}

		// s.Find("content dt").Each(func(_ int, s2 *goquery.Selection) {
		// 	os := s2.Text()
		// 	pkg := s2.Next().Find("a").First().Text()
		// 	ver := s2.Next().Find("a").Last().Text()
		// 	fmt.Printf("%v : %v-%v\n", os, pkg, ver)
		// })
	})
}
