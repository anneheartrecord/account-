package biz

import (
	"account/errorCode"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
)

func GetMd5(s string) (string, error) {
	hash := md5.New()
	writeString, err := io.WriteString(hash, s)
	if writeString <= 0 {
		return "", errors.New(errorCode.HashFailed)
	}
	return hex.EncodeToString(hash.Sum(nil)), err
}
