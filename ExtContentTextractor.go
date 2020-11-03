package c_code

import (
	"github.com/gloomyzerg/textractor"
)

func ExtContentTextractor(content string) (*textractor.Text, error) {
	return textractor.Extract(content)
}
