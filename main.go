package main

import (
	"errors"
	"fmt"
	"sort"
)

// Mean returns the average of the nums
func Mean(nums []float64) (float64, error) {

	if len(nums) == 0 {
		return 0, errors.New("no number(s) inputted")
	}

	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums)), nil

}

// Median returns the middle number
func Median(nums []float64) (float64, error) {

	if len(nums) == 0 {
		return 0, errors.New("no number(s) inputted")
	}
	//sort the numbers in an ascending order
	sort.Float64s(nums)

	if len(nums)%2 == 1 {
		//if its odd, return the middle number
		return float64(nums[len(nums)/2]), nil
	} else {
		//if its even return, add the two middle numbers and divide by two
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2, nil
	}
}

func main() {
	nums := []float64{2, 2, 3, 4, 5}

	fmt.Println(Mean(nums))

	fmt.Println(Median(nums))

}
