package c_code

import (
	readability "github.com/go-shiori/go-readability"
	"strings"
)

func ExtContentReadability(content string) (readability.Article, error) {
	//if !regexp.MustCompile(`<html.*?>`).MatchString(content){
	//	return  readability.Article{},errors.New("Html格式错误")
	//}
	reader, err := readability.FromReader(strings.NewReader(content), "https://www.baidu.com/")
	return reader, err
}
