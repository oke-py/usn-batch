package feed

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetPriority is a function to get priority (severity) for specific CVE.
func GetPriority(cve string) string {
	res, err := http.Get(GetCveURL(cve))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return GetPriorityFromDoc(res.Body)
}

// GetCveURL is a function to convert CVE number to ubuntu-security URL.
func GetCveURL(cve string) string {
	year := strings.Split(cve, "-")[1]
	return "https://people.canonical.com/~ubuntu-security/cve/" + year + "/" + cve + ".html"
}

// GetPriorityFromDoc is a function to extract priority (severity) from document.
func GetPriorityFromDoc(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("#body-card .card-body .item .field").First().Next().Find("a").Text()
}