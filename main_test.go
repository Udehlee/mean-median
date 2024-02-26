package main

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	nums := []float64{2, 5, 4, 3, 2}

	want := float64(3.2)
	got, err := Mean(nums)

	if err != nil {
		t.Error("error getting the mean", err)
	}

	if got != want {

		t.Errorf("got %v,  want %v", got, want)
	}

}

func TestMedian(t *testing.T) {
	nums := []float64{2, 5, 4, 3, 2}

	want := float64(3)

	got, err := Median(nums)

	if err != nil {
		t.Error("Error getting the median", err)
	}

	if got != want {

		t.Errorf("got %v,  want %v", got, want)
	}

}

func ExampleMean() {
	nums := []float64{2, 5, 4, 3, 2}
	fmt.Println(Mean(nums))
	//output
	//3.2

}

func ExampleMedian() {
	nums := []float64{2, 5, 4, 3, 2}
	fmt.Println(Median(nums))
	//output
	//3

}

func BenchmarkMean(b *testing.B) {
	nums := []float64{2, 5, 4, 3, 2}

	for i := 0; i < b.N; i++ {
		Mean(nums)
	}
}

func BenchmarkMedian(b *testing.B) {
	nums := []float64{2, 5, 4, 3, 2}

	for i := 0; i < b.N; i++ {
		Median(nums)
	}
}
