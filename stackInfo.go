package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)
//打印go语言运行堆栈的方法
//https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Scanf("%s")
}
