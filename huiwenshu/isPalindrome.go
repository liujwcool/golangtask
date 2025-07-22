package main

import "fmt"

// isPalindrome 判断一个整数是否为回文数。
// 如果整数x是回文数，返回true，否则返回false。
// 负数不是回文数。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed := 0
	original := x

	for x > 0 {
		fmt.Println("x = ", x, "reversed = ", reversed)
		reversed = reversed * 10 + x % 10
		fmt.Println("reversed = ", reversed)
		x /= 10
		fmt.Println("x = ", x, "reversed = ", reversed)
		fmt.Println("--------------------------------")
	}

	return reversed == original
}

func main() {
	fmt.Println(isPalindrome(123))
}