package main

import "testing"

func TestMedian1(t *testing.T) {
	if Median([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) != 30 {
		t.Fail()
	}
}

func TestMedian2(t *testing.T) {
	if Median([]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}) != 39 {
		t.Fail()
	}
}

func TestMedian3(t *testing.T) {
	if Median([]int{1, 2, 3}) != 4 {
		t.Fail()
	}
}

func TestMedian4(t *testing.T) {
	if Median([]int{3, 2, 1}) != 7 {
		t.Fail()
	}
}

func TestMedian5(t *testing.T) {
	if Median([]int{10, 9, 1, 2, 3, 4, 5, 6, 7, 8}) != 54 {
		t.Fail()
	}
}
