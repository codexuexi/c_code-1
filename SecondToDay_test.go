package c_code_test

import (
	"fmt"
	"github.com/123456/c_code"
	"testing"
)

func TestSecondsToDay(t *testing.T) {
	s1 := "1"
	for _, v := range s1 {
		println(v, string(v))
	}
	s := "【１Зち-电-б４９９薇８１９б】"
	for _, v := range s {
		println(v, string(v))
	}
	return
	fmt.Println(c_code.SecondsToDay(60*60*24*180 + 464))
}

func TestSecondsToDay2222(t *testing.T) {

	acctoken := "34_Z1X0ybx8T37_qrgHwnB-PHSwmTWQ0-27vjRhYXZn2ju3sM2h8IwHXa_AQr-yrfYH_eGKwrpqZ8TMO3Csd_9QVDdOq7rOqUkGGvmSSGQ2kz5xw6xyNOC_RgsR6Gq91dDlNui-OajIQ8LX5PQwIBUbAEAXQJ"
	values := map[string]string{
		"content": `亚洲-无码-专区`,
	}
	err, data := c_code.CPostJson("https://api.weixin.qq.com/wxa/msg_sec_check?access_token="+acctoken, values)
	fmt.Println(err, data)
}
