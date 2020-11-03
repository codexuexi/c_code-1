package c_code

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func CGet(_url string) (b string, err error) {

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	//生成要访问的url

	//提交请求
	reqest, err := http.NewRequest("GET", _url, nil)

	//增加header选项
	//reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	//reqest.Header.Add("X-Requested-With", "xxxx")
	parse, err := url.Parse(_url)
	if err != nil {
		return
	}
	reqest.Header.Add("Host", parse.Host)
	//if err != nil {
	//	panic(err)
	//}
	//处理返回结果

	resp, err := client.Do(reqest)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
		return
	}
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	//处理编码
	//判断头部是否已给出编码
	content_type := resp.Header.Get("content-type")
	lower := strings.ToLower(content_type)
	bck := []byte{}
	bk := false
	switch {
	case strings.Count(lower, "big5") > 0:
		r := transform.NewReader(bytes.NewReader(by), traditionalchinese.Big5.NewDecoder())
		bck, _ = ioutil.ReadAll(r)
		bk = true
		break
	case strings.Count(lower, "gb") > 0:
		_decode := simplifiedchinese.GBK.NewDecoder()
		r := transform.NewReader(bytes.NewReader(by), _decode)
		bck, _ = ioutil.ReadAll(r)
		bk = true
		break
		//utf编码不处理
	case strings.Contains(lower, "utf"):
		bk = true
		bck = by
		break
	default:
		bck = by
		break
	}

	if bk == false {
		encoding, name, _, e := DetermineEncodingFromReader(bytes.NewReader(by), len(by))
		if e != nil {
			return
		}
		_decode := encoding.NewDecoder()
		switch name {
		case "windows-1252":
			_decode = simplifiedchinese.GBK.NewDecoder()
		}
		r := transform.NewReader(bytes.NewReader(by), _decode)
		bck, _ = ioutil.ReadAll(r)
	}

	b = string(bck)
	return
}

func CGetRef(_url string) (ref string, b string, err error) {

	client := &http.Client{
		Timeout: time.Second * 60,
	}
	//生成要访问的url

	//提交请求
	reqest, err := http.NewRequest("GET", _url, nil)

	//增加header选项
	//reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	//reqest.Header.Add("X-Requested-With", "xxxx")
	parse, err := url.Parse(_url)
	if err != nil {
		return
	}
	reqest.Header.Add("Host", parse.Host)
	if err != nil {
		panic(err)
	}
	//处理返回结果

	resp, err := client.Do(reqest)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New("not 200")
		return
	}
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	locaUrl := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host + resp.Request.URL.Path
	if resp.Request.URL.RawQuery != "" {
		locaUrl += "?" + resp.Request.URL.RawQuery
	}
	ref = locaUrl

	//处理编码
	//判断头部是否已给出编码
	content_type := resp.Header.Get("content-type")
	lower := strings.ToLower(content_type)
	bck := []byte{}
	bk := false
	switch {
	case strings.Count(lower, "big5") > 0:
		r := transform.NewReader(bytes.NewReader(by), traditionalchinese.Big5.NewDecoder())
		bck, _ = ioutil.ReadAll(r)
		bk = true
		break
	case strings.Count(lower, "gb") > 0:
		_decode := simplifiedchinese.GBK.NewDecoder()
		r := transform.NewReader(bytes.NewReader(by), _decode)
		bck, _ = ioutil.ReadAll(r)
		bk = true
		break
		//utf编码不处理
	case strings.Contains(lower, "utf"):
		bk = true
		bck = by
		break
	default:
		bck = by
		break
	}

	if bk == false {
		encoding, name, _, e := DetermineEncodingFromReader(bytes.NewReader(by), len(by))
		if e != nil {
			return
		}
		_decode := encoding.NewDecoder()
		switch name {
		case "windows-1252":
			_decode = simplifiedchinese.GBK.NewDecoder()
		}
		r := transform.NewReader(bytes.NewReader(by), _decode)
		bck, _ = ioutil.ReadAll(r)
	}

	b = string(bck)
	return
}

func DetermineEncodingFromReader(r io.Reader, peek int) (e encoding.Encoding, name string, certain bool, err error) {
	if peek >= 1024 {
		peek = 1024
	}
	b, err := bufio.NewReader(r).Peek(peek)
	if err != nil {
		return
	}
	e, name, certain = charset.DetermineEncoding(b, "")
	return
}

func CPostJson(_url string, post_data interface{}) (e2 error, data string) {
	var e error
	client := http.Client{
		Timeout: time.Second * 15,
	}

	otype := reflect.TypeOf(post_data)
	var marshal []byte
	if otype.Name() != "string" {
		ma, e5 := json.Marshal(post_data)
		marshal = ma
		if e5 != nil {
			return
		}
	} else {
		s := post_data.(string)
		marshal = []byte(s)
	}

	//fmt.Printf("%s\n",marshal)
	request, e := http.NewRequest("POST", _url, bytes.NewReader(marshal))
	if e != nil {
		e2 = e
		return
	}
	//发送请求
	response, e := client.Do(request)
	if e != nil {
		e2 = e
		return
	}
	defer response.Body.Close()
	b, e := ioutil.ReadAll(response.Body)
	if e != nil {
		e2 = e
		return
	}
	data = string(b)
	return
}
func CPost(_url string, d2 url.Values) (e2 error, data string) {
	var e error
	client := http.Client{
		Timeout: time.Second * 15,
	}
	post_data := d2.Encode()
	if e2 != nil {
		return
	}
	//fmt.Printf("%s \n %s \n", _url, post_data)
	request, e := http.NewRequest("POST", _url, strings.NewReader(post_data))
	if e != nil {
		e2 = e
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")
	//发送请求
	response, e := client.Do(request)
	if e != nil {
		e2 = e
		return
	}
	defer response.Body.Close()
	b, e := ioutil.ReadAll(response.Body)
	if e != nil {
		e2 = e
		return
	}
	data = string(b)
	return
}
