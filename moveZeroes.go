package main

import "fmt"

func main() {
	array := []int{1, 2, 4, 0, 2}
	fmt.Println(moveZeroes(array))
}

func moveZeroes(nums []int) []int {
	zero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
	}
	//fmt.Printf("co %d so 0 \n", zero)
	res := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res = append(res, nums[i])
		}
	}
	fmt.Println(len(res))
	for i := zero; i != 0; i-- {
		res = append(res, 0)
	}
	fmt.Println(len(res))
	/*
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	*/
	fmt.Println(len(nums), len(res))
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
	fmt.Println(len(nums))
	//
	return nums
}
