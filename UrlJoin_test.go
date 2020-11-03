package c_code

import (
	"testing"
)

func TestUrlJoin(t *testing.T) {
	tophost, e := UrlJoin2("https://www.mydeals.jp/category/books", "../../baidu.com")
	t.Error(tophost, e)
}
