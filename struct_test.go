package main

import (
	"fmt"
	"testing"
)

type s struct {
	a int
	b string
}

/**
测试接口赋值是引用复制还是拷贝，结果为拷贝
*/
func TestStruct(t *testing.T) {
	type i interface{}
	v := s{a: 2}
	var j i = &v
	v.b = "str"
	fmt.Printf("%p\n", &v)
	fmt.Printf("%p\n", j)
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", j)
}
