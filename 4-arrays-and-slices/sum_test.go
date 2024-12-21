package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := Sum(nums)

	want := 10

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, nums)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})

	want := []int{3, 9}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %d want %d", got, want)
	// }

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})

	want := []int{2, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
