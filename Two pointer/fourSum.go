package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println("Quadruplets summing to", target, ":", result)
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)
	
	for i:=0; i<n-3; i++ {
		if i>0 && nums[i] == nums[i-1]{
			continue
		}

		for j:=i+1; j<n-2; j++ {
			if j>i+1 && nums[j]==nums[j-1]{
				continue
			}

			left,right := j+1,n-1

			for left < right {
				sum := nums[i]+nums[j]+nums[left]+nums[right]
				switch {
				case sum < target:
					left++
				case sum > target:
					right--
				default:
					result = append(result, []int{nums[i],nums[j],nums[left],nums[right]})
					for left<right && nums[left]==nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1]{
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return result
}