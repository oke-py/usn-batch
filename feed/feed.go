package feed

import "strings"

func ExtractUsnTitle(s string) string {
	tmp := strings.Replace(s, "<![CDATA[", "", -1)
	return strings.Split(tmp, ":")[0]
}
