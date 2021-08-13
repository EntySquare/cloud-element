package spec

import (
	"math/rand"
	"strconv"
	"time"
)

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func StrToUInt64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(i)
}
