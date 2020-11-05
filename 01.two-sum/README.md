### 解题思路
通过哈希表记录以及查询过的索引和值，如果target - 新值，
在哈希表中可以查到则返回结果

时间复杂度 O(n) 数组长度
空间复杂度 O(n) map的开销，数组长度


### 代码
```golang
func twoSum(nums []int, target int) []int {
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
```