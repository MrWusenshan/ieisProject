package util

import (
	"math/rand"
	"time"
)

//生成随机字符串作为姓名
func RandomString(n int) string {
	var latter = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = latter[rand.Intn(len(latter))]
	}

	return string(result)
}
