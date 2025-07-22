package main

import "sort"
import "fmt"

func Merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    // 按区间起点排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
	fmt.Println(intervals)

    res := [][]int{}
    res = append(res, intervals[0])

    for i := 1; i < len(intervals); i++ {
        last := res[len(res)-1]
        curr := intervals[i]
		fmt.Println("curr[0] = ", curr[0], "last[1] = ", last[1])
        if curr[0] <= last[1] {
            // 有重叠，合并
            res[len(res)-1][1] = max(last[1], curr[1])
        } else {
            // 无重叠，直接加入
            res = append(res, curr)
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	intervals := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	fmt.Println(Merge(intervals))
}