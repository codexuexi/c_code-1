package lib

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ChangeIP_Info(api_url, kdname, kduser, kdpwd, parent_ip string) (now_ip string) {

	RasdialDis(kdname)
	RasdialConnect(kdname, kduser, kdpwd)
	//获取IP
	ic := 0
	for {
		ic++
		if ic > 10 {
			now_ip = ""
			return
		}
		_url := api_url
		req, _ := http.NewRequest("GET", _url, nil)
		client := http.Client{
			Timeout: time.Millisecond * 5000,
		}

		res, e := client.Do(req)
		if e != nil {
			continue
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		now_ip = gjson.Get(string(body), "ip").String()
		if now_ip == parent_ip || now_ip == "" {
			RasdialDis(kdname)
			RasdialConnect(kdname, kduser, kdpwd)
		} else {
			fmt.Println("ParentIP-->", parent_ip, "now_ip-->", string(body))
			break
		}
	}
	return
}
func ChangeIP(kdname, kduser, kdpwd, parent_ip string) (now_ip string) {
	KillChrome()
	RasdialDis(kdname)
	RasdialConnect(kdname, kduser, kdpwd)
	//获取IP
	ic := 0
	for {
		log.Println("切换IP第", ic)
		ic++
		if ic > 10 {
			now_ip = ""
			return
		}
		_url := "http://myexternalip.com/raw"
		req, _ := http.NewRequest("GET", _url, nil)
		client := http.Client{
			Timeout: time.Millisecond * 1000,
		}

		res, e := client.Do(req)
		if e != nil {
			continue
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		now_ip = string(body)
		if now_ip == parent_ip || now_ip == "" {
			RasdialDis(kdname)
			RasdialConnect(kdname, kduser, kdpwd)
		} else {
			fmt.Println("ParentIP-->", parent_ip, "now_ip-->", now_ip)
			break
		}
	}
	return
}
