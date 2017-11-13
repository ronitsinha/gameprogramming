package main 

import "fmt"

func max(nums ...int) int {
	var maxValue int
	for i := 0; i < len (nums); i++ {
		if (i != 0) {
			pastValue := nums [i-1]
			if (nums[i] >= pastValue) {
				maxValue = nums[i]
			}
		} else {
			maxValue = nums [i]
		}
	}

	return maxValue
}

func min(nums ...int) int {
	var minValue int
	for i := 0; i < len (nums); i++ {
		if (i != 0) {
			pastValue := nums [i-1]
			if (nums[i] <= pastValue && nums[i] <= nums[0]) {
				minValue = nums[i]
			}
		} else {
			minValue = nums [i]
		}
	}

	return minValue
}

func main() {
	fmt.Println ("Max of 2, 10, 100, 40 is", max(2, 10, 100, 40))
	fmt.Println ("Min of 2, 10, 100, 40 is", min(2, 10, 100, 40))

}