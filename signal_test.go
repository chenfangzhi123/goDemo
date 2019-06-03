package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

var w sync.WaitGroup

//处理进程信号
var signalChan = make(chan os.Signal)

var ctx, cancel = context.WithCancel(context.Background())

func processSignal() {
	//处理退出信号，优雅停机
	signal.Notify(signalChan, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChan:
		cancel()
		print("\ndone\n")
		time.Sleep(time.Second * 2)
		os.Exit(0)
	}
}

func TestSignal(t *testing.T) {
	w.Add(1)
	go processSignal()
	ticker := time.NewTicker(time.Second)
	go func() {
		defer w.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%s", ctx.Err())
				//这里不能使用break，break只是跳出select，还是在for循环，只能用return或者标签break
				return
			case <-ticker.C:
				fmt.Printf("p...\n")
			}
		}
	}()
	fmt.Scanf("%s")
	w.Wait()
}
