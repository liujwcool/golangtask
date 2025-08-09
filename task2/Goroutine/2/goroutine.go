package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task 表示一个任务，包含任务ID、名称和执行函数
type Task struct {
	ID       int
	Name     string
	Function func() error
}

// TaskResult 表示任务执行结果
type TaskResult struct {
	TaskID    int
	TaskName  string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
	Success   bool
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	mutex   sync.RWMutex
	wg      sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(task Task) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	ts.tasks = append(ts.tasks, task)
}

// ExecuteTasks 并发执行所有任务
func (ts *TaskScheduler) ExecuteTasks() {
	ts.mutex.RLock()
	taskCount := len(ts.tasks)
	ts.mutex.RUnlock()

	if taskCount == 0 {
		fmt.Println("没有任务需要执行")
		return
	}

	fmt.Printf("开始执行 %d 个任务...\n", taskCount)

	// 为每个任务启动一个协程
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(task)
	}

	// 等待所有任务完成
	ts.wg.Wait()
	fmt.Println("所有任务执行完成！")
}

// executeTask 执行单个任务
func (ts *TaskScheduler) executeTask(task Task) {
	defer ts.wg.Done()

	startTime := time.Now()

	// 执行任务
	err := task.Function()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// 创建结果
	result := TaskResult{
		TaskID:    task.ID,
		TaskName:  task.Name,
		StartTime: startTime,
		EndTime:   endTime,
		Duration:  duration,
		Error:     err,
		Success:   err == nil,
	}

	// 保存结果
	ts.mutex.Lock()
	ts.results = append(ts.results, result)
	ts.mutex.Unlock()

	// 打印任务执行状态
	if result.Success {
		fmt.Printf("✅ 任务 %s (ID: %d) 执行成功，耗时: %v\n",
			task.Name, task.ID, duration)
	} else {
		fmt.Printf("❌ 任务 %s (ID: %d) 执行失败，耗时: %v，错误: %v\n",
			task.Name, task.ID, duration, err)
	}
}

// GetResults 获取所有任务执行结果
func (ts *TaskScheduler) GetResults() []TaskResult {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	// 复制结果切片避免竞态条件
	results := make([]TaskResult, len(ts.results))
	copy(results, ts.results)
	return results
}

// PrintSummary 打印任务执行摘要
func (ts *TaskScheduler) PrintSummary() {
	results := ts.GetResults()
	if len(results) == 0 {
		fmt.Println("没有执行结果")
		return
	}

	fmt.Println("\n=== 任务执行摘要 ===")

	var totalDuration time.Duration
	var successCount, failureCount int

	for _, result := range results {
		totalDuration += result.Duration
		if result.Success {
			successCount++
		} else {
			failureCount++
		}
	}

	fmt.Printf("总任务数: %d\n", len(results))
	fmt.Printf("成功任务数: %d\n", successCount)
	fmt.Printf("失败任务数: %d\n", failureCount)
	fmt.Printf("总执行时间: %v\n", totalDuration)
	fmt.Printf("平均执行时间: %v\n", totalDuration/time.Duration(len(results)))

	// 按执行时间排序显示
	fmt.Println("\n任务执行时间排序:")
	for i, result := range results {
		fmt.Printf("%d. %s (ID: %d): %v",
			i+1, result.TaskName, result.TaskID, result.Duration)
		if !result.Success {
			fmt.Printf(" [失败: %v]", result.Error)
		}
		fmt.Println()
	}
}

// 示例任务函数
func simulateTask1() error {
	// 模拟任务1：随机延迟
	delay := time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(delay)

	// 模拟偶尔失败
	if rand.Float32() < 0.1 {
		return fmt.Errorf("任务1随机失败")
	}
	return nil
}

func simulateTask2() error {
	// 模拟任务2：固定延迟
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func simulateTask3() error {
	// 模拟任务3：快速完成
	time.Sleep(500 * time.Millisecond)
	return nil
}

func simulateTask4() error {
	// 模拟任务4：长时间运行
	time.Sleep(4000 * time.Millisecond)
	return nil
}

func main() {
	// Go 1.20+ automatically seeds the global random source

	// 创建任务调度器
	scheduler := NewTaskScheduler()

	// 添加任务
	scheduler.AddTask(Task{ID: 1, Name: "数据处理任务", Function: simulateTask1})
	scheduler.AddTask(Task{ID: 2, Name: "文件下载任务", Function: simulateTask2})
	scheduler.AddTask(Task{ID: 3, Name: "缓存更新任务", Function: simulateTask3})
	scheduler.AddTask(Task{ID: 4, Name: "数据库备份任务", Function: simulateTask4})

	// 执行所有任务
	scheduler.ExecuteTasks()

	// 打印执行摘要
	scheduler.PrintSummary()
}
