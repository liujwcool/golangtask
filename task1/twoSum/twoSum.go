package main

import (
    "fmt"
)

// TwoSum 在 nums 中找到和为 target 的两个数的下标。
func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil // 按题意不会到这里
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := TwoSum(nums, target)
    fmt.Println("结果下标:", result)
}
