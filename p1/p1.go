package p1

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func p1() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
