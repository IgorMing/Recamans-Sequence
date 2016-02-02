package main

import "fmt"

func main() {
	fmt.Println("%v", recaman(10))
}

func recaman(size int) []int{
	vec := []int{}
	vec[0] = 0
	vec[1] = 1
	
	for i := 2; i < size; i++ {
		if !elementExists(vec, i) && vec[i-1]-i > 0 {
			vec[i] = vec[i-1]-i
		} else {
			vec[i] = vec[i-1]+i
		}	
	}
	return vec
}

func elementExists(vec []int, elem int) bool{
	for i := 0; i < len(vec); i++ {
		if vec[i] == elem {
			return true
		}
	}
	return false
}