package feed

import (
	"os"
	"reflect"
	"testing"
	"time"

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

func TestGetCves(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}
	cves := GetCves(doc.Find("entry"))
	if !reflect.DeepEqual(cves, []string{"CVE-2018-5744", "CVE-2018-5745", "CVE-2019-6465"}) {
		t.Fatal("failed test")
	}
}

func TestAffects1604(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}

	if !Affects1604(doc.Find("entry")) {
		t.Fatal("failed test")
	}
}

func TestAffects1804(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}

	if !Affects1804(doc.Find("entry")) {
		t.Fatal("failed test")
	}
}

func TestGetPublished(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}

	published := GetPublished(doc.Find("entry"))
	expectd := time.Date(2019, time.February, 22, 8, 8, 13, 0, time.UTC)

	if !published.Equal(expectd) {
		t.Fatal("failed test")
	}
}

func TestGetUpdated(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}

	updated := GetUpdated(doc.Find("entry"))
	expectd := time.Date(2019, time.February, 22, 8, 8, 13, 0, time.UTC)

	if !updated.Equal(expectd) {
		t.Fatal("failed test")
	}
}

func TestGetNotice(t *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		t.Fatal("cannot open file")
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal("cannot read file")
	}

	expectd := Notice{
		ID:          "USN-3893-1",
		Pkg:         "bind9",
		CVEs:        []string{"CVE-2018-5744", "CVE-2018-5745", "CVE-2019-6465"},
		Priority:    "Medium",
		Affects1604: true,
		Affects1804: true,
		Published:   time.Date(2019, time.February, 22, 8, 8, 13, 0, time.UTC),
		Updated:     time.Date(2019, time.February, 22, 8, 8, 13, 0, time.UTC),
	}
	actual := GetNotice(doc.Find("entry"))

	if actual.ID != expectd.ID {
		t.Fatal("failed test")
	}
	if actual.Pkg != expectd.Pkg {
		t.Fatal("failed test")
	}
	if !reflect.DeepEqual(actual.CVEs, expectd.CVEs) {
		t.Fatal("failed test")
	}
	if actual.Priority != expectd.Priority {
		t.Fatal("failed test")
	}
	if actual.Affects1604 != expectd.Affects1604 {
		t.Fatal("failed test")
	}
	if actual.Affects1804 != expectd.Affects1804 {
		t.Fatal("failed test")
	}
	if !actual.Published.Equal(expectd.Published) {
		t.Fatal("failed test")
	}
	if !actual.Updated.Equal(expectd.Updated) {
		t.Fatal("failed test")
	}
}
