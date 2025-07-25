package main
import "fmt"
// RemoveDuplicates 移除有序数组中的重复元素，返回新长度
func RemoveDuplicates(nums []int) int {
    fmt.Printf("[line 6] 输入 nums: %v\n", nums)
    if len(nums) == 0 {
        fmt.Printf("[line 8] len(nums) == 0, return 0\n")
        return 0
    }
    i := 0
    fmt.Printf("[line 11] 初始化 i: %d\n", i)
    for j := 1; j < len(nums); j++ {
        fmt.Printf("[line 13] j: %d, i: %d, nums: %v\n", j, i, nums)
		fmt.Printf("[line 14] nums[j] = %d, nums[i] = %d\n", nums[j], nums[i])
        if nums[j] != nums[i] {
            fmt.Printf("[line 15] nums[j] != nums[i], nums[j]: %d, nums[i]: %d\n", nums[j], nums[i])
            i++
            fmt.Printf("[line 17] i 自增后: %d\n", i)
            nums[i] = nums[j]
            fmt.Printf("[line 19] nums[%d] 赋值为 nums[%d]: %v\n", i, j, nums)
        } else {
            fmt.Printf("[line 21] nums[j] == nums[i], nums[j]: %d, nums[i]: %d\n", nums[j], nums[i])
        }
    }
    fmt.Printf("[line 23] 返回 i+1: %d\n", i+1)
    return i + 1
}

func main() {
    nums := []int{1, 1, 2}
    length := RemoveDuplicates(nums)
    fmt.Println(length)      // 输出: 2
    fmt.Println(nums[:length]) // 输出: [1 2]
	fmt.Println(nums) // 输出: [1 2 2]
}