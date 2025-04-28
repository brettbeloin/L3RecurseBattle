package main

import (
	"testing"
)

func TestSplitM(t *testing.T) {
	input := []int{1, 4, 5, 6, 7}
	expect := 2
	s := split(input)

	if s != expect {
		t.Errorf("Expected: %v, found: %v", expect, s)
	}
}

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
	expected := []int{1, 2, 3, 4, 5, 6}
	s := mSort(input)

	for k, v := range s {
		if v != expected[k] {
			t.Errorf("Found: %d, Expected: %d", v, input[k-1])
		}
	}
}

func TestStrMSort(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "orange"}
	expect := []string{"apple", "banana", "cherry", "orange"}
	s := qSort(input)

	for k, v := range s {
		if v != expect[k] {
			t.Errorf("Expected: %v, found: %v", expect[k], v)
		}
	}
}

func TestDupsQSort(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "orange", "banana"}
	expect := []string{"apple", "banana", "cherry", "orange", "banana"}
	s := qSort(input)

	for k, v := range s {
		if v != expect[k] {
			t.Errorf("Expected: %v, found: %v", expect[k], v)
		}
	}
}

func TestHappyQSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := qSort(input)

	for k, v := range s {
		if v != input[k] {
			t.Errorf("Found: %d, Expected: %d", v, input[k])
		}
	}
}

func TestIntQSort(t *testing.T) {
	input := []int{1, 3, 2, 6, 4, 5}
	expected := []int{1, 2, 3, 4, 5, 6}
	s := qSort(input)

	for k, v := range s {
		if v != expected[k] {
			t.Errorf("At index %d, Found: %d, Expected: %d", k, v, expected[k])
		}
	}
}

func TestStrQSort(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "orange"}
	expect := []string{"apple", "banana", "cherry", "orange"}
	s := mSort(input)

	for k, v := range s {
		if v != expect[k] {
			t.Errorf("Expected: %v, found: %v", expect[k], v)
		}
	}
}

func TestDupsMSort(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "orange", "banana"}
	expect := []string{"apple", "banana", "cherry", "orange", "banana"}
	s := mSort(input)

	for k, v := range s {
		if v != expect[k] {
			t.Errorf("Expected: %v, found: %v", expect[k], v)
		}
	}
}

func TestSingleElemQSort(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	s := qSort(input)

	for k, v := range s {
		if v != expected[k] {
			t.Errorf("At index %d, Found: %d, Expected: %d", k, v, expected[k])
		}
	}
}
