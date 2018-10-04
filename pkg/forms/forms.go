package forms

import (
	"strings"
	"unicode/utf8"
)

type NewSnippet struct {
	Title    string
	Content  string
	Expires  string
	Failures map[string]string
}

func (f *NewSnippet) Valid() bool {
	f.Failures = make(map[string]string)

	if strings.TrimSpace(f.Title) == "" {
		f.Failures["Title"] = "Title is required"
	} else if utf8.RuneCountInString(f.Title) > 100 {
		f.Failures["Title"] = "Title cannot be longer than 100 characters"
	}

	if strings.TrimSpace(f.Content) == "" {
		f.Failures["Content"] = "Content is required"
	}

	permitted := map[string]bool{
		"3600":     true,
		"86400":    true,
		"31536000": true,
	}

	if strings.TrimSpace(f.Expires) == "" {
		f.Failures["Expires"] = "Expires is required"
	} else if !permitted[f.Expires] {
		f.Failures["Expires"] = "Expiry time must be 3600, 86400 or 31536000 seconds"
	}

	return len(f.Failures) == 0
}
