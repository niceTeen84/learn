package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/panjf2000/ants/v2"
)

const (
	EMPTY         = ""
	KEY   ctxInfo = "a"
)

func DoReqest(ctx context.Context) (result string, e error) {
	result, e = "", nil
	// 声明一个通道，用于 goroutine 的数据交换
	ch := make(chan string)
	// 并行启动 goroutine
	go func() {
		// 启动新的 goroutine, 耗时 10 秒
		time.Sleep(10 * time.Second)
		select {
		case ch <- "abc":
		default:
			return
		}
	}()

	// 阻塞等待通道的信号或者等待到超时
	select {
	case <-ctx.Done():
		e = ctx.Err()
		return
	case result = <-ch:
		return
	}
}

type ctxFn func(context.Context) string

type ctxInfo string

func ContextDemo(ctx context.Context) (ret string) {
	ret = EMPTY
	// any 强制转换为 string
	val := ctx.Value(KEY).(string)
	fmt.Println(strings.ToUpper(val))
	return
}

func ContextDeadLine(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("err ", ctx.Err().Error())
	default:
		fmt.Println("you are so lucky")
		return
	}
}

func LoopUntilCanel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("canel")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("roaring...")
		}
	}
}

func UsePool() {
	pool, _ := ants.NewPool(1000, ants.WithExpiryDuration(10*time.Second), ants.WithMaxBlockingTasks(100))
	defer pool.Release()
}

func main() {
	// main goroutine root
	root := context.Background()
	// 声明一个新的 context 有超时时间

	ctx, cancel := context.WithTimeout(root, 1*time.Second)
	defer cancel()
	res, err := DoReqest(ctx)
	if err != nil {
		fmt.Println("do request failed ", err.Error())
	}
	fmt.Println("do request success ", res)

	// 声明一个带值 context
	// https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
	// 声明 key value 时会报
	// should not use basic type string as key in context.WithValue
	// 由于您定义了一个单独的类型，它永远不会发生冲突。即使你有两个包，pkg1.key(0) != pkg2.key(0).
	valCtx := context.WithValue(root, KEY, "abcd")
	ContextDemo(valCtx)

	deadLine := time.Date(2022, 9, 30, 20, 20, 20, 20, time.Local)

	ctx, cancel = context.WithDeadline(root, deadLine)
	defer cancel()
	ContextDeadLine(ctx)

	ctx, cancel = context.WithCancel(root)
	// 启动一个 goroutine 5 秒后发出取消命令
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	// 一直循环输出, 直到取消
	LoopUntilCanel(ctx)
}
