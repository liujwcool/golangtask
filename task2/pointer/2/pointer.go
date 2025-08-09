// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

//     考察点 ：指针运算、切片操作。

package main

import "fmt"

// doubleSliceElements 接收一个整数切片的指针，将切片中的每个元素乘以2
func doubleSliceElements(slicePtr *[]int) {
	// 通过指针访问切片
	slice := *slicePtr

	// 遍历切片中的每个元素，将其乘以2
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	// 测试用例1：普通切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", numbers)

	// 传递切片的指针给函数
	doubleSliceElements(&numbers)
	fmt.Println("乘以2后:", numbers)

	// 测试用例2：空切片
	emptySlice := []int{}
	fmt.Println("\n空切片:", emptySlice)
	doubleSliceElements(&emptySlice)
	fmt.Println("乘以2后:", emptySlice)

	// 测试用例3：包含0和负数的切片
	mixedNumbers := []int{0, -1, 10, -5}
	fmt.Println("\n混合数字切片:", mixedNumbers)
	doubleSliceElements(&mixedNumbers)
	fmt.Println("乘以2后:", mixedNumbers)

	// 测试用例4：单元素切片
	singleElement := []int{7}
	fmt.Println("\n单元素切片:", singleElement)
	doubleSliceElements(&singleElement)
	fmt.Println("乘以2后:", singleElement)
}
