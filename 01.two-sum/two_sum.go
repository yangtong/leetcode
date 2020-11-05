package leetcode

func TwoSum(nums []int, target int) []int {
	tempMap := map[int]int{}
	for n, i := range nums {
		m, ok := tempMap[target-i]
		if ok {
			return []int{m, n}
		}
		tempMap[i] = n
	}
	return []int{}
}
