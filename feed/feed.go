package feed

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Notice is a class of each USN
type Notice struct {
	ID          string `dynamo:"usn_id"`
	Pkg         string `dynamo:"name"`
	CVEs        []string
	Priority    string    `dynamo:"severity"`
	Affects1604 bool      `dynamo:"affects_1604"`
	Affects1804 bool      `dynamo:"affects_1804"`
	Published   time.Time `dynamo:"published"`
	Updated     time.Time `dynamo:"updated"`
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

// Affects1604 is a function to evaluate Ubuntu 16.04 LTS is affected.
func Affects1604(entry *goquery.Selection) bool {
	return strings.Contains(entry.Find("ul").Text(), "Ubuntu 16.04 LTS")
}

// Affects1804 is a function to evaluate Ubuntu 18.04 LTS is affected.
func Affects1804(entry *goquery.Selection) bool {
	return strings.Contains(entry.Find("ul").Text(), "Ubuntu 18.04 LTS")
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

// GetNotice returns an instance of Notice type.
func GetNotice(entry *goquery.Selection) Notice {
	notice := Notice{
		ID:          GetID(entry),
		Pkg:         GetPackageName(entry),
		CVEs:        GetCves(entry),
		Affects1604: Affects1604(entry),
		Affects1804: Affects1804(entry),
		Published:   GetPublished(entry),
		Updated:     GetUpdated(entry),
	}
	for _, cve := range notice.CVEs {
		notice.Priority = GetHigherPriority(notice.Priority, GetPriority(cve))
	}
	return notice
}
