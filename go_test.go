package main

import "testing"

func TestHappyMSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := mSort(input)

	for k, v := range s {
		if v != input[k] {
			t.Errorf("Found: %d, Expected: %d", v, input[k])
		}
	}
}
