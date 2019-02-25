package feed

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetCveURL(t *testing.T) {
	url := GetCveURL("CVE-2019-6465")
	if url != "https://people.canonical.com/~ubuntu-security/cve/2019/CVE-2019-6465.html" {
		t.Fatal("failed test")
	}
}

func TestGetPriorityFromDoc(t *testing.T) {
	file, err := os.Open("./test_cve.html")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}
	priority := GetPriorityFromDoc(doc)
	if priority != "Medium" {
		t.Fatal("failed test")
	}
}
