package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个带有缓冲的整数通道，缓冲区大小为20
	ch := make(chan int, 20)

	// 使用WaitGroup来等待协程完成
	var wg sync.WaitGroup

	// 生产者协程：快速发送100个整数到缓冲通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("生产者协程开始工作...")

		for i := 1; i <= 100; i++ {
			ch <- i // 发送数据到缓冲通道
			if i%20 == 0 {
				fmt.Printf("生产者已发送 %d 个数据，当前通道长度: %d，通道容量: %d\n",
					i, len(ch), cap(ch))
			}
		}

		// 发送完毕后关闭通道
		close(ch)
		fmt.Println("生产者协程完成工作，通道已关闭")
	}()

	// 消费者协程：从缓冲通道接收整数并打印
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("消费者协程开始工作...")

		count := 0
		for {
			// 从通道接收数据
			value, ok := <-ch
			if !ok {
				// 通道已关闭，退出循环
				fmt.Printf("通道已关闭，消费者协程退出，共处理了 %d 个数据\n", count)
				return
			}

			count++
			// 打印接收到的数据
			fmt.Printf("消费者接收: %d (通道长度: %d)\n", value, len(ch))

			// 每处理20个数据显示一次统计信息
			if count%20 == 0 {
				fmt.Printf("消费者已处理 %d 个数据，当前通道长度: %d\n", count, len(ch))
			}
		}
	}()

	// 主协程等待所有协程完成
	fmt.Println("主协程等待其他协程完成...")
	fmt.Printf("初始通道状态 - 长度: %d, 容量: %d\n", len(ch), cap(ch))
	wg.Wait()
	fmt.Println("程序结束")
}
