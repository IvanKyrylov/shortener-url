package generator

import (
	"testing"
)

func TestIsUrlValid(t *testing.T) {
	testcase := []struct {
		url  string
		want bool
	}{
		{"https://golang.org/pkg/encoding/base64/", true},
		{"https://github.com/IvanKyrylov/shortener-url", true},
		{"https://ru.wikipedia.org/wiki/%D0%A5%D0%B5%D1%88-%D1%84%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D1%8F", true},
		{"https://blog.golang.org/strings", true},
		{"www.blog.golang.org/strings", false},
		{"golang.org/pkg/encoding/base64/", false},
		{"github.com", false},
		{"github", false},
	}

	for _, v := range testcase {
		valid := IsUrlValid(v.url)
		if valid != v.want {
			t.Errorf("IsValidUrl(%v): got: %v want: %v", v.url, valid, v.want)
		}
	}
}

func TestGenerateKey(t *testing.T) {
	key := GenerateKey()
	if len(key) != 7 {
		t.Error("ID is empty")
	}
}
