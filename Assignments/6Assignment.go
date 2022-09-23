/* Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import (
	"fmt"
)

var nums = []int{0, 1, 2, 3, 4, 6, 8, 9}
var target int = 9

func twoSum(nums []int, target int) (int, int) {
	if len(nums) <= 1 {
		return 0, 0
	}
	hdict := make(map[int]int)
	for i := 1; i < len(nums); i++ {
		if val, ok := hdict[nums[i+1]]; ok {
			return val, i + 1
		} else {
			hdict[target-nums[i+1]] = i + 1
		}
	}
	return 0, 0
}

func main() {
	fmt.Println(twoSum(nums, target))
}
