package test

import "testing"

//运行前初始化，只运行一次
func TestMain(m *testing.M) {
	println("start")
	m.Run()
	println("end")
}

//测试普通函数和方法
func Test1(t *testing.T) {
	println("test1")

}

func Test2(t *testing.T) {
	println("test2")
}

func Test3(t *testing.T) {
	print(testing.Short())

}