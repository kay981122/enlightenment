package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func EncryptionByMD5(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	arr := h.Sum(nil)
	newArr := fmt.Sprintf("%x", arr)
	return strings.ToTitle(newArr)
}
