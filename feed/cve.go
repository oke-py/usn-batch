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
	return GetPriorityFromReader(res.Body)
}

// GetCveURL is a function to convert CVE number to ubuntu-security URL.
func GetCveURL(cve string) string {
	year := strings.Split(cve, "-")[1]
	return "https://people.canonical.com/~ubuntu-security/cve/" + year + "/" + cve + ".html"
}

// GetPriorityFromReader is a function to extract priority (severity) from document.
func GetPriorityFromReader(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("#body-card .card-body .item .field").First().Next().Find("a").Text()
}

// GetHigherPriority returns a higher priority.
// https://people.canonical.com/~ubuntu-security/cve/priority.html
func GetHigherPriority(p1 string, p2 string) string {
	if p1 == "Critical" || p2 == "Critical" {
		return "Critical"
	}
	if p1 == "High" || p2 == "High" {
		return "High"
	}
	if p1 == "Medium" || p2 == "Medium" {
		return "Medium"
	}
	if p1 == "Low" || p2 == "Low" {
		return "Low"
	}

	return "Unknown"
}
