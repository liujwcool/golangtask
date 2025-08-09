package main

import (
	"fmt"
	"sync"
)

func main() {
	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup

	// 启动第一个协程：打印奇数 1,3,5,7,9
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("奇数协程开始执行...")
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
		}
		fmt.Println("奇数协程执行完成")
	}()

	// 启动第二个协程：打印偶数 2,4,6,8,10
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("偶数协程开始执行...")
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
		}
		fmt.Println("偶数协程执行完成")
	}()

	// 等待所有协程完成
	fmt.Println("主协程等待子协程完成...")
	wg.Wait()
	fmt.Println("所有协程执行完成，程序结束")
}
