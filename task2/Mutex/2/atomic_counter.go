package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// AtomicCounter 结构体，使用原子操作实现无锁计数器
type AtomicCounter struct {
	value int64 // 使用 int64 类型，因为 atomic.AddInt64 需要 int64
}

// Increment 方法：使用原子操作递增计数器
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1) // 原子地增加1
}

// GetValue 方法：使用原子操作获取计数器值
func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value) // 原子地读取值
}

// SetValue 方法：使用原子操作设置计数器值
func (c *AtomicCounter) SetValue(newValue int64) {
	atomic.StoreInt64(&c.value, newValue) // 原子地设置值
}

func main() {
	// 创建原子计数器
	counter := &AtomicCounter{}

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	fmt.Println("=== 原子操作计数器测试 ===")
	fmt.Printf("初始计数器值: %d\n", counter.GetValue())

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
	expectedValue := int64(10 * 1000) // 10个协程 × 1000次操作

	fmt.Printf("\n=== 最终结果 ===\n")
	fmt.Printf("计数器最终值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedValue)

	if finalValue == expectedValue {
		fmt.Println("✅ 测试通过：计数器值正确！")
	} else {
		fmt.Printf("❌ 测试失败：计数器值不正确\n")
		fmt.Printf("丢失了 %d 次递增操作\n", expectedValue-finalValue)
	}

	// 展示原子操作的其他功能
	fmt.Println("\n=== 原子操作演示 ===")

	// 演示 CompareAndSwap 操作
	oldValue := counter.GetValue()
	newValue := oldValue + 100
	swapped := atomic.CompareAndSwapInt64(&counter.value, oldValue, newValue)
	fmt.Printf("CompareAndSwap: 旧值=%d, 新值=%d, 交换成功=%t\n", oldValue, newValue, swapped)

	// 演示原子减法
	atomic.AddInt64(&counter.value, -50)
	fmt.Printf("原子减法后: %d\n", counter.GetValue())

	fmt.Println("程序结束")
}
