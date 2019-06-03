package test

import "testing"

//运行前初始化，只运行一次
func TestMain(m *testing.M) {
	println("start")
	m.Run()
	println("end")
}

func Test1(t *testing.T) {
	println("test1")

}

func Test2(t *testing.T) {
	println("test2")
}
