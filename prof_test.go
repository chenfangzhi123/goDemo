package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"testing"
)

//打印go语言运行堆栈的方法
//https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func TestProf(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		//监控信息地址：http://localhost:6060/debug/pprof/
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	w.Wait()
}
