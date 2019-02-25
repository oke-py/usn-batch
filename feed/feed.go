package feed

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Notice is a class of each USN
type Notice struct {
	ID        string
	Pkg       string
	CVEs      []string
	Priority  string
	Published time.Time
	Updated   time.Time
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

// GetCves is a function to get CVE list related to a specific USN.
func GetCves(entry *goquery.Selection) []string {
	cves := []string{}
	entry.Find("#references").Next().Find("li a").Each(func(_ int, s *goquery.Selection) {
		if strings.HasPrefix(s.Text(), "CVE-") {
			cves = append(cves, s.Text())
		}
	})
	return cves
}

// GetPublished is a function to get USN published time.
func GetPublished(entry *goquery.Selection) time.Time {
	text := entry.Find("published").Text()
	t, _ := time.Parse("2006-01-02T15:04:05-07:00", text)
	return t
}

// GetUpdated is a function to get USN updated time.
func GetUpdated(entry *goquery.Selection) time.Time {
	text := entry.Find("updated").Text()
	t, _ := time.Parse("2006-01-02T15:04:05-07:00", text)
	return t
}
