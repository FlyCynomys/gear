package randomstring

import (
	"math/rand"
	"time"
)

const charlist = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() []byte {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	length := len(charlist)
	var data []byte
	for index := 0; index < 16; index++ {
		num := r.Intn(length)
		data = append(data, charlist[num])
	}
	return data
}
