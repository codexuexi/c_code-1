package c_code

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if first {
				first = false
				return unicode.ToUpper(r)
			}
			return r
		},
		s)
}
