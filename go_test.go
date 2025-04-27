package main

import (
	"fmt"
	"testing"
)

func TestHappyMSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := mSort(input)

	for k, v := range s {
		if v != input[k] {
			t.Errorf("Found: %d, Expected: %d", v, input[k])
		}
	}
}

func TestIntMSort(t *testing.T) {
	input := []int{1, 3, 2, 6, 4, 5}
	expected := []int{1, 2, 3, 4, 5, 6} // Adjusted to match the sorted input
	s := mSort(input)                   // Assuming mSort is your sorting function
	fmt.Println(s)

	// Ensure the lengths are equal
	if len(s) != len(expected) {
		t.Fatalf("Expected length %d, but got %d", len(expected), len(s))
	}

	// Compare the sorted array with the expected values
	for k, v := range s {
		if v != expected[k] {
			t.Errorf("At index %d, Found: %d, Expected: %d", k, v, expected[k])
		}
	}
}

// func TestIntMSort(t *testing.T) {
// 	input := []int{1, 3, 2, 6, 4, 5}
// 	expected := []int{1, 2, 3, 4, 5}
// 	s := mSort(input)
// 	fmt.Println(s)
//
// 	for k, v := range s {
// 		if v != expected[k] {
// 			t.Errorf("Found: %d, Expected: %d", v, input[k-1])
// 		}
// 	}
// }

/*
func TestHappyQSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := qSort(input)

	for k, v := range s {
		if v != input[k] {
			t.Errorf("Found: %d, Expected: %d", v, input[k])
		}
	}
}
*/
