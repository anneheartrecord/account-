package biz_test

import (
	"account/biz"
	"fmt"
	"testing"
)

func TestGetMd5(t *testing.T) {
	s := "test"
	fmt.Println(biz.GetMd5(s))
	s1 := "test2"
	fmt.Println(biz.GetMd5(s1))
}
