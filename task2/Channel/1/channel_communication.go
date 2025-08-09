package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲的整数通道
	ch := make(chan int)

	// 启动生产者协程：生成1到10的整数并发送到通道
	go func() {
		fmt.Println("生产者协程开始工作...")
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者发送: %d\n", i)
			ch <- i                            // 发送数据到通道
			time.Sleep(100 * time.Millisecond) // 稍微延迟以便观察
		}
		close(ch) // 发送完毕后关闭通道
		fmt.Println("生产者协程完成工作，通道已关闭")
	}()

	// 启动消费者协程：从通道接收整数并打印
	go func() {
		fmt.Println("消费者协程开始工作...")
		for {
			// 从通道接收数据
			value, ok := <-ch
			if !ok {
				// 通道已关闭，退出循环
				fmt.Println("通道已关闭，消费者协程退出")
				return
			}
			fmt.Printf("消费者接收: %d\n", value)
		}
	}()

	// 主协程等待一段时间，让其他协程完成工作
	fmt.Println("主协程等待其他协程完成...")
	time.Sleep(2 * time.Second)
	fmt.Println("程序结束")
}
