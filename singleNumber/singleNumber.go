package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // 异或运算，相同为0，不同为1
	}
	return result
}

func main() {
	// 测试用例
	testCases := [][]int{
		{2, 2, 1},           // 应该返回 1
		{4, 1, 2, 1, 2},     // 应该返回 4
		{1},                 // 应该返回 1
		{1, 1, 2, 2, 3},     // 应该返回 3
	}

	for i, nums := range testCases {
		result := singleNumber(nums)
		fmt.Printf("测试用例 %d: nums = %v, 结果 = %d\n", i+1, nums, result)
	}
}