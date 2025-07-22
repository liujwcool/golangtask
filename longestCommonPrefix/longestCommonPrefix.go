package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
		// fmt.Println("prefix = ", prefix)
		// fmt.Println("strs[i] = ", strs[i])
		// fmt.Println("--------------------------------1")
        // 不断缩短 prefix，直到它是 strs[i] 的前缀
        // for len(prefix) > 0 && (len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix ){
		for len(prefix) > 0 && len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			// fmt.Println("prefix = ", prefix)
			// fmt.Println("strs[i] = ", strs[i])
			// fmt.Println("--------------------------------2")
            prefix = prefix[:len(prefix)-1]
        }
		// fmt.Println("prefix = ", prefix)
		// fmt.Println("--------------------------------3")
    }

	if prefix == "" {
		return ""
	}

    return prefix
}

func main() {
    fmt.Println(longestCommonPrefix([]string{"flowereu", "flow", "floght"}))
    fmt.Println(longestCommonPrefix([]string{"rog", "racecar", "rar"}))
	fmt.Println(longestCommonPrefix([]string{"f", "flow", "floght"}))
	fmt.Println(longestCommonPrefix([]string{"foe", "foow", "floght"}))
	fmt.Println(longestCommonPrefix([]string{"fooght","foe", "foow" }))
	fmt.Println(longestCommonPrefix([]string{"a","bc", "foow" }))
}