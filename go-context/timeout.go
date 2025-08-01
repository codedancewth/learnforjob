package learnforjob

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 设置超时时长context.WithTimeout，超时3s后自动结束goroutine

// BuildContextWithTimeOut 尝试构建context设置超时时间和value
func BuildContextWithTimeOut() context.Context {
	// 初始化ctx
	ctx := context.Background()

	// 设置ctx的value值
	ctx = context.WithValue(ctx, "test", "timeout")

	// 设置ctx的超时时间
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	return ctx
}

// Print1000WithTimeOut 设置ctx的超时时间，查看一下运用的方式
func Print1000WithTimeOut(ctx context.Context, synG *sync.WaitGroup) {
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

// RunTimeOut 执行
func RunTimeOut() {
	ctx := BuildContextWithTimeOut()
	group := &sync.WaitGroup{}
	group.Add(1)

	go Print1000WithTimeOut(ctx, group)
	group.Wait()

	fmt.Println("Done")
}
