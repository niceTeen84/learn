package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed hello.txt
var s string

func main() {
	fmt.Println("test embed")
	sql := fmt.Sprintf(s, 11)
	fmt.Println(sql)

	fmt.Println("test select multi channel")

	ch := make(chan int, 1024)
	defer close(ch)
	// 启动一个协程后台读取通道内的值
	go func(ch chan int) {
		for {
			if val, ok := <-ch; ok {
				fmt.Printf("read %d from channel\n", val)
			}
		}
	}(ch)
	// 新建一个 timer
	timer := time.NewTicker(1 * time.Second)
	defer timer.Stop()

	for i := 0; i < 30; i++ {
		select {
		case ch <- i:
		}

		select {
		// case ch <- i:
		// timer 触发器结构体成员
		case <-timer.C:
			fmt.Printf("%d <- timer.C\n", i)
		default:
		}

		time.Sleep(200 * time.Millisecond)
	}

}
