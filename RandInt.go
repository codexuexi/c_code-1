package c_code

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

//生成真随机数
func Rand(size int) int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
	ints := fmt.Sprintf("%d", result)
	i, _ := strconv.Atoi(ints)
	return i
}

//获取2个值之间的随机数
func RandDom(min, max int) int {
	randNum := Rand(max-min) + min
	return randNum
}
