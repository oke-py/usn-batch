package feed

import "testing"

func TestExtractUsnTitle(t *testing.T) {
	s := ExtractUsnTitle("<![CDATA[USN-3891-1: systemd vulnerability ]]>")
	if s != "USN-3891-1" {
		t.Fatal("failed test")
	}
}
