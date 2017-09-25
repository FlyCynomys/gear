package format

import (
	"encoding/binary"
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	num, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		return 0
	}
	return int(num)
}

func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	num, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func FormatTimeFromTi(template string, timepoint time.Time) string {
	if template == "" {
		return time.Now().Format(timeFormat)
	}
	return time.Now().Format(template)
}

func FormatTimeFromUnix(template string, timepoint int64) string {
	ti := time.Unix(timepoint, 0)
	if template == "" {
		return ti.Format(timeFormat)
	}
	return ti.Format(template)
}

func FormatTimeToUnix(timepoint string) int64 {
	ti, err := time.Parse(timeFormat, timepoint)
	if err != nil {
		return -1
	}
	return ti.Unix()
}

func BytesToUint64(src []byte) uint64 {
	return binary.BigEndian.Uint64(src)
}

func Uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}
