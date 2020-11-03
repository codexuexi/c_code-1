package kc

import (
	"github.com/chromedp/chromedp"
	"testing"
)

func TestDelWebdriver(t *testing.T) {
	cc := KcBody{}
	cc.Up()
	chromedp.Run(cc.Ctx, chromedp.Navigate("http://www.baidu.com"))
}
