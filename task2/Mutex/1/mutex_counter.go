package main

import (
	"fmt"
	"sync"
)

// Counter 结构体，包含一个共享的计数器和互斥锁
type Counter struct {
	value int
	mutex sync.Mutex
}

// Increment 方法：安全地递增计数器
func (c *Counter) Increment() {
	c.mutex.Lock()         // 获取锁
	defer c.mutex.Unlock() // 确保在函数结束时释放锁

	c.value++ // 递增计数器
}

// GetValue 方法：安全地获取计数器值
func (c *Counter) GetValue() int {
	c.mutex.Lock()         // 获取锁
	defer c.mutex.Unlock() // 确保在函数结束时释放锁

	return c.value // 返回计数器值
}

func main() {
	// 创建共享计数器
	counter := &Counter{}

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			fmt.Printf("协程 %d 开始工作...\n", workerID)

			// 每个协程对计数器进行1000次递增操作
			for j := 1; j <= 1000; j++ {
				counter.Increment()

				// 每100次操作显示一次进度
				if j%100 == 0 {
					fmt.Printf("协程 %d 已完成 %d 次递增操作\n", workerID, j)
				}
			}

			fmt.Printf("协程 %d 完成工作\n", workerID)
		}(i)
	}

	// 等待所有协程完成
	fmt.Println("主协程等待所有工作协程完成...")
	wg.Wait()

	// 输出最终结果
	finalValue := counter.GetValue()
	expectedValue := 10 * 1000 // 10个协程 × 1000次操作

	fmt.Printf("\n=== 最终结果 ===\n")
	fmt.Printf("计数器最终值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedValue)

	if finalValue == expectedValue {
		fmt.Println("✅ 测试通过：计数器值正确！")
	} else {
		fmt.Printf("❌ 测试失败：计数器值不正确，可能存在竞态条件\n")
	}

	fmt.Println("程序结束")
}
