package learnforjob

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 设置超时时间context.WithDeadline，超时3s后自动结束goroutine

// BuildContextWithDeadLine 尝试构建context设置过期时间
func BuildContextWithDeadLine() context.Context {
	// 初始化ctx
	ctx := context.Background()

	// 设置ctx的value值
	ctx = context.WithValue(ctx, "test", "deadline")

	// 设置ctx的超时时间
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	return ctx
}

// Print1000WithDeadline 设置ctx的过期时间
func Print1000WithDeadline(ctx context.Context, synG *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		fmt.Println(fmt.Sprintf("ctx:%s value%d", ctx.Value("test"), i))

		// 获取超时时间
		deadline, ok := ctx.Deadline()
		// 如果存在超时的话,返回done
		if ok && deadline.Before(time.Now()) {
			fmt.Println(fmt.Sprintf("deadline:%v", deadline))
			synG.Done()
			return
		}

		time.Sleep(time.Second)
	}

	synG.Done()
}

// RunDeadLine 执行
func RunDeadLine() {
	ctx := BuildContextWithDeadLine()
	group := &sync.WaitGroup{}
	group.Add(1)

	go Print1000WithDeadline(ctx, group)
	group.Wait()

	fmt.Println("Done")
}
