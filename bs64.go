package c_code

import (
	"encoding/base64"
	"strings"
)

func Bs64Encode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

func Bs64Decode(str string) ([]byte, error) {
	//if strings.ContainsAny(str, "+/") {
	//	return nil, errors.New("invalid base64url encoding")
	//}
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
