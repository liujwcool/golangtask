package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, ch := range s {
        switch ch {
        case '(', '{', '[':
            stack = append(stack, ch)
        case ')', '}', ']':
			// fmt.Printf("len(stack)-1 = %d ch = %c\n", len(stack)-1, ch)
			// fmt.Printf("stack[len(stack)-1] = %c pairs[ch] = %c\n", stack[len(stack)-1], pairs[ch])
			// fmt.Println("--------------------------------")
            if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

// 查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        // 不断缩短 prefix，直到它是 strs[i] 的前缀
        for len(prefix) > 0 && (len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix) {
            prefix = prefix[:len(prefix)-1]
        }
        if prefix == "" {
            return ""
        }
    }
    return prefix
}

func main() {
    fmt.Println(isValid("()"))      // true
    fmt.Println(isValid("()[]{}"))  // true
    fmt.Println(isValid("(]"))      // false
    fmt.Println(isValid("([)]"))    // false
    fmt.Println(isValid("{[]}"))    // true
    fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
    fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}