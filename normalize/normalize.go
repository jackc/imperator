package normalize

import (
	"strings"
	"unicode"
)

// TextField performed common normalization for a single line text field. It performs the following operations:
//
// Remove any invalid UTF-8
// Replace non-printable characters with standard space
// Remove spaces from left and right
func TextField(s string) string {
	s = strings.ToValidUTF8(s, "")
	s = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		} else {
			return ' '
		}
	}, s)
	s = strings.TrimSpace(s)

	return s
}
