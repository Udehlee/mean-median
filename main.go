package main

import (
	"errors"
	"fmt"
	"sort"
)

func Mean(nums []float64) (float64, error) {

	lengthOfNums := len(nums)

	if lengthOfNums == 0 {
		return 0, errors.New("no number(s) inputted")
	}
	var sum float64
	// sum = 0

	for _, v := range nums {
		sum += v
	}

	return sum / float64(lengthOfNums), nil

}

func Median(nums []float64) (float64, error) {

	lengthOfNums := len(nums)

	if lengthOfNums == 0 {
		return 0, errors.New("no number(s) inputted")
	}
	//sort the numbers in an ascending order
	sort.Float64s(nums)

	if lengthOfNums%2 == 1 {
		//if its odd, return the middle number
		return float64(nums[lengthOfNums/2]), nil
	} else {
		//if its even return, add the two middle numbers and divide by two
		return float64(nums[lengthOfNums/2-1]+nums[lengthOfNums/2]) / 2, nil
	}
}
func main() {
	xi := []float64{1, 2, 3, 4, 5, 1}

	//mean
	fmt.Println(Mean(xi))

	//median
	fmt.Println(Median(xi))

}
