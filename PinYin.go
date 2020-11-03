package c_code

import (
	"github.com/axgle/pinyin"
	"strings"
)

func Pinyin(to_pin_yin string) string {
	return strings.ToLower(pinyin.Convert(to_pin_yin))
}
