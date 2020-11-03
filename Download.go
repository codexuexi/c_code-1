package c_code

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filename string) (err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()

	_, err = io.Copy(f, r.Body)
	return
}
