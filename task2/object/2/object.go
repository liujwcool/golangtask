package main

import "fmt"

// Person 结构体，包含 Name 和 Age 字段
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段
type Employee struct {
	Person     // 组合 Person 结构体
	EmployeeID string
}

// PrintInfo 方法，输出员工的信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
}

func main() {
	// 创建一个 Person 实例
	person := Person{
		Name: "张三",
		Age:  30,
	}

	// 创建一个 Employee 实例，组合 Person
	employee := Employee{
		Person:     person,
		EmployeeID: "EMP001",
	}

	// 调用 PrintInfo 方法输出员工信息
	employee.PrintInfo()

	// 也可以直接访问组合的字段
	fmt.Printf("\n直接访问字段:\n")
	fmt.Printf("姓名: %s\n", employee.Name)
	fmt.Printf("年龄: %d\n", employee.Age)
	fmt.Printf("员工ID: %s\n", employee.EmployeeID)
}
