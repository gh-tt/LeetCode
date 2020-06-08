package main

import "fmt"

func main() {
	s := []string{"a==b", "c==d", "c!=e", "a==c"}
	fmt.Println(equationsPossible(s))
}

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	for _, v := range equations {
		if v[1] == '=' {
			index1 := int(v[0]) - int('a')
			index2 := int(v[3]) - int('a')
			union(parent, index1, index2)
		}
	}

	for _, v := range equations {
		if v[1] == '!' {
			index1 := int(v[0]) - int('a')
			index2 := int(v[3]) - int('a')

			if find(parent, index1) == find(parent, index2) {
				return false
			}
		}
	}
	return true
}

func union(parent []int, index1, index2 int) {
	index1 = find(parent, index1)
	index2 = find(parent, index2)
	if index1 != index2 {
		parent[index1] = index2
	}
}

func find(parent []int, index int) int {
	for parent[index] != index {
		index = parent[index]
	}
	return index
}
