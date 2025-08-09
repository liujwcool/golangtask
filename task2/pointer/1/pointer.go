// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

//     考察点 ：指针的使用、值传递与引用传递的区别。

package main

import "fmt"

// increaseByTen 接收一个整数指针作为参数，将该指针指向的值增加10
func increaseByTen(ptr *int) {
	// 通过指针修改原值
	*ptr += 10
}

func main() {
	// 定义一个整数变量
	value := 25
	fmt.Printf("修改前的值: %d\n", value)

	// 调用函数，传递变量的地址（指针）
	increaseByTen(&value)

	// 输出修改后的值
	fmt.Printf("修改后的值: %d\n", value)

	// 验证指针确实修改了原值
	fmt.Printf("主函数中的值: %d\n", value)

	// 额外演示：展示指针的地址
	fmt.Printf("变量value的地址: %p\n", &value)
}
