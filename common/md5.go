package common

import (
	"crypto/md5"
	"io"
	"strconv"
)

func MD5(s string)string {
	h := md5.New()
	io.WriteString(h, s)
	var x string
	for _, v := range h.Sum(nil) {
		x += strconv.FormatInt(int64(v), 16)
	}
	return x
}
