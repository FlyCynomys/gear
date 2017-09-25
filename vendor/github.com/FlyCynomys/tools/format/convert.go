package format

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func TransToPinyin(params string) (string, bool) {
	if len(params) <= 0 {
		return "", false
	}
	trans := pinyin.NewArgs()
	data := pinyin.LazyPinyin(params, trans)
	if len(data) <= 0 {
		return "", false
	}
	return strings.Join(data, "-"), true
}

func LittleEndianOrBigEndian() bool {
	var value uint32
	value = 0x12345678
	bytevalue := int8(value)
	if bytevalue == 0x12 {
		return false //big
	}
	return true //little
}

func Reverse(src string) string {
	var temp = make([]byte, len(src), len(src))
	length := len(src) - 1
	for index, value := range src {
		temp[length-index] = byte(value)
	}
	return string(temp)
}
