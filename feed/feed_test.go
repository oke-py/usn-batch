package feed

import (
	"os"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractUsnTitle(t *testing.T) {
	s := ExtractUsnTitle("<![CDATA[USN-3891-1: systemd vulnerability ]]>")
	if s != "USN-3891-1" {
		t.Fatal("failed test")
	}
}

func TestGetID(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}
	ID := GetID(doc.Find("entry"))
	if ID != "USN-3893-1" {
		t.Fatal("failed test")
	}
}

func TestGetPackageName(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}
	name := GetPackageName(doc.Find("entry"))
	if name != "bind9" {
		t.Fatal("failed test")
	}
}

func TestGetCVEs(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}
	cves := GetCVEs(doc.Find("entry"))
	if !reflect.DeepEqual(cves, []string{"CVE-2018-5744", "CVE-2018-5745", "CVE-2019-6465"}) {
		t.Fatal("failed test")
	}
}
