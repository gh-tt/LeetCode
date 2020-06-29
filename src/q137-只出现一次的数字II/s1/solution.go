package main

import "fmt"

func main()  {
	nums :=[]int{0,1,0,1,0,1,99}
	fmt.Println(singleNumber(nums))
}
func singleNumber(nums []int) int {
	res :=0
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	for k, v := range m {
		if v == 1 {
			res = k
		}
	}
	return res
}
