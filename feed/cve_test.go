package feed

import (
	"os"
	"testing"
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
	priority := GetPriorityFromDoc(file)
	if priority != "Medium" {
		t.Fatal("failed test")
	}
}
