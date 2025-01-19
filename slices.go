package main

import "fmt"

func main() {

	fmt.Println("Slices Practice")
	nums := make([]int, 1, 5)
	nums[0] = 1
	nums = append(nums, 367)
	nums = append(nums, 36)
	nums[2] = 556
	nums = append(nums, 66)
	nums = append(nums, 5)
	nums = append(nums, 15)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	fmt.Println(nums)

	nums2 := [][]int{{5, 6}, {7, 8}}
	fmt.Println(nums2)
	fmt.Println(nums2[0][0])
}
