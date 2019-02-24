package feed

import "strings"

// ExtractUsnTitle is a function to get USN-XXXX-X string.
func ExtractUsnTitle(s string) string {
	tmp := strings.Replace(s, "<![CDATA[", "", -1)
	return strings.Split(tmp, ":")[0]
}
