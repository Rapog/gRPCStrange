package main

import "fmt"

func main() {
	target := 21
	var nums = []int{0, 1, 2, 3, 9, 11, 15, 18, 20, 200}

	var solve []int = ind(nums, target)

	fmt.Println(solve)
}

// дан целочисленный массив и целое число (target). вернуть индексы двух чисел, сумма которых равна target
//Input: nums = [0,1,2,3,9,11,15,18,20, 200], target = 21

func ind(nums []int, target int) []int {
	var solve []int
	for ind, val := range nums {
		for ind1, val1 := range nums {
			if val+val1 == target {
				solve = append(solve, ind)
				solve = append(solve, ind1)
				return solve
			}
		}
	}
	return solve
}
