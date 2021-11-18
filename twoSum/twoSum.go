package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		fmt.Println(i)
		for j, next := range nums[i+1:] {
			fmt.Println(j)
			j += i + 1
			if num+next == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
