package main

func sortColors(nums []int) {
	if len(nums) < 1 {
		return
	}
	tmp := make([]int, 3)

	for _, v := range nums {
		if v == 0 {
			tmp[0]++
		} else if v == 1 {
			tmp[1]++
		} else {
			tmp[2]++
		}
	}

	i := 0
	for k, v := range tmp {
		if v > 0 {
			for v > 0 {
				nums[i] = k
				v--
				i++
			}
		}
	}
}
