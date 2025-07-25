package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
		fmt.Println("digits[i] = ", digits[i], "i = ", i)
		fmt.Println("--------------------------------")
        if digits[i] < 9 {
			//假如小于9就不用进位了，所以只需要执行一次，然后返回
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    // 如果所有位都是9，则需要进位，所以需要返回一个1，然后digits前面加一个0
    return append([]int{1}, digits...)
}

func main() {
	// fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	// fmt.Println(plusOne([]int{4, 3, 2, 1}))
	// fmt.Println(plusOne([]int{9}))
	// fmt.Println(plusOne([]int{9, 9, 9}))
}