## 两数之和解法

+ 暴力循环法，时间复杂度O(n<sup>2</sup>),空间复杂度O(1)
```
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```
+ map法,时间复杂度O(n),空间复杂度O(n)
```
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for k, v := range nums {
		if _, ok := m[v]; ok {
			return []int{m[v], k}
		}
		m[target-v] = k
	}

	return nil
}
```