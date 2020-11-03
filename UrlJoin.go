package c_code

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func UrlJoin(loca_url string, inhtmlhref string) (_url string, e error) {
	if regexp.MustCompile(`^[http://|https://]`).MatchString(inhtmlhref) {
		_url = inhtmlhref
		return
	}
	parse, e := url.Parse(loca_url)
	if e != nil {
		return
	}
	Path := parse.Path
	join := ""
	if strings.HasPrefix(inhtmlhref, "/") {
		join = inhtmlhref
	} else {
		join = regexp.MustCompile(`[\\+|\\]`).ReplaceAllString(filepath.Join(Path, inhtmlhref), "/")
	}

	_url = parse.Scheme + "://" + parse.Host + join
	return
}

func UrlJoin2(loca_url string, inhtmlhref string) (_url string, e error) {
	u, e := url.Parse(inhtmlhref)
	if e != nil {
		return
	}
	base, e := url.Parse(loca_url)
	if e != nil {
		return
	}
	_url = base.ResolveReference(u).String()
	if strings.HasPrefix(_url, "//") {
		_url = "http:" + _url
	}
	return
}
