package feed

import (
	"os"
	"testing"
)

func TestGetPriority(t *testing.T) {
	if GetPriority("CVE-2019-6465") != "Medium" {
		t.Fatal("failed test")
	}
}

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
	priority := GetPriorityFromReader(file)
	if priority != "Medium" {
		t.Fatal("failed test")
	}
}

func TestGetHigherPriority(t *testing.T) {
	if GetHigherPriority("", "") != "Unknown" {
		t.Fatal("failed test")
	}
	if GetHigherPriority("", "Low") != "Low" {
		t.Fatal("failed test")
	}
	if GetHigherPriority("Low", "Low") != "Low" {
		t.Fatal("failed test")
	}
	if GetHigherPriority("Medium", "Low") != "Medium" {
		t.Fatal("failed test")
	}
	if GetHigherPriority("High", "Low") != "High" {
		t.Fatal("failed test")
	}
	if GetHigherPriority("High", "Critical") != "Critical" {
		t.Fatal("failed test")
	}
}
