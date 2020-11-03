package c_code

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"strings"
)

func WindowsCmd(cmd_line string) (cmd_result_str string, err error) {
	cmdArgs := strings.Fields(cmd_line)
	cmd, err := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...).Output()
	if err != nil {
		return
	}
	//fmt.Println(cmd, err)
	cmd_result_str = string(cmd)
	encoding, name, _, err := DetermineEncodingFromReader(strings.NewReader(cmd_result_str), len(cmd_result_str))
	if err != nil {
		return
	}
	_decode := encoding.NewDecoder()
	switch name {
	case "windows-1252":
		_decode = simplifiedchinese.GBK.NewDecoder()
	}
	r := transform.NewReader(strings.NewReader(cmd_result_str), _decode)
	bck, _ := ioutil.ReadAll(r)
	cmd_result_str = string(bck)
	return
}
