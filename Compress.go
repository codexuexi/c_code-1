package c_code

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
)

func CompressHtml(b string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	s, _ := m.String("text/html", b)
	return s
}

func CompressCss(b string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	s, _ := m.String("text/css", b)
	return s
}
