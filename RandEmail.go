package c_code

import "math/rand"
import "strings"
import "time"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandEmail(n int) string {
	zz := letterBytes
	b := make([]byte, n)
	emial_last := []string{"@163.com", "@qq.com", "@sina.com", "@126.com"}
	e2 := emial_last[Rand(len(emial_last))]
	if e2 == "@qq.com" {
		zz = "0123456789"
	}
	for i := range b {
		b[i] = zz[rand.Intn(len(zz))]
	}

	return strings.ToLower(string(b) + e2)
}
func Random(n int) string {
	rand.Seed(time.Now().Unix())
	zz := letterBytes
	b := make([]byte, n)

	for i := range b {
		b[i] = zz[rand.Intn(len(zz))]
	}

	return strings.ToLower(string(b))
}
