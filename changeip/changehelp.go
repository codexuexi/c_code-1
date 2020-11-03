package lib

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"os/exec"
)

func RasdialConnect(kdname, kduser, kdpwd string) (err error) {
	cmd := exec.Command("rasdial", kdname, kduser, kdpwd)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
		return
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return
	}
	cmd.Wait()
	toByte := ConvertToByte(string(opBytes), "gbk", "utf8")
	fmt.Println(string(toByte))
	return
}

func KillChrome() {
	cmd := exec.Command("taskkill", "/F", "/im", "chrome.exe")
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err = cmd.Start(); err != nil {
		return
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return
	}
	toByte := ConvertToByte(string(opBytes), "gbk", "utf8")
	fmt.Println(string(toByte))
	return
}

func RasdialDis(kdname string) (err error) {
	cmd := exec.Command("rasdial", kdname, "/DISCONNECT")
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err = cmd.Start(); err != nil {
		return
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return
	}
	toByte := ConvertToByte(string(opBytes), "gbk", "utf8")
	fmt.Println(string(toByte))
	cmd.Wait()
	return
}

func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
