package feed

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Notice is a class of each USN
type Notice struct {
	ID   string
	Pkg  string
	CVEs []string
}

// ExtractUsnTitle is a function to get USN-XXXX-X string.
func ExtractUsnTitle(s string) string {
	tmp := strings.Replace(s, "<![CDATA[", "", -1)
	return strings.Split(tmp, ":")[0]
}

// GetID is a function to get USN-XXXX-X string.
func GetID(entry *goquery.Selection) string {
	return ExtractUsnTitle(entry.Find("title").Text())
}

// GetPackageName is a function to get vulnerable package name.
func GetPackageName(entry *goquery.Selection) string {
	name := entry.Find("#software-description").Next().Find("li").Text()
	return strings.Split(name, " ")[0]
}

// GetCVEs is a function to get CVE list related to a specific USN.
func GetCVEs(entry *goquery.Selection) []string {
	cves := []string{}
	entry.Find("#references").Next().Find("li a").Each(func(_ int, s *goquery.Selection) {
		if strings.HasPrefix(s.Text(), "CVE-") {
			cves = append(cves, s.Text())
		}
	})
	return cves
}
