package main

import "golang.org/x/exp/constraints"

//type generic interface {
//	~int | ~string | ~float64
//}

func main() {

}

/*
chat model: chat gpt 3.5
why: Used to help fix a bug with the split func
chat log: https://chatgpt.com/share/6807d1fa-49ec-8006-acd4-19cd99f08f20
*/

func split[T any](arr []T) int {
	mid := len(arr) / 2
	return mid
}

/*
chat model: chat gpt 3.5
why: needed help fixing the compairtion logic for a generic type
chat log: https://chatgpt.com/share/6807d668-4890-8006-ae55-75415959763f
*/

func mSort[T constraints.Ordered](arr []T) []T {
	length := split(arr)
	left := arr[:length]
	right := arr[length:]
	if len(left) != 1 || len(right) != 1 {
		mSort(left)
		mSort(right)
	}

	leftIdx := left[0]
	rightIdx := right[0]

	for idx, _ := range arr {
		if leftIdx < rightIdx {
			arr[idx] = leftIdx
			leftIdx = left[idx+1]
		}

		if rightIdx < leftIdx {
			arr[idx] = rightIdx
			rightIdx = right[idx+1]
		}
	}
	return arr
}

/*
* Merge sort:
* func mSort(tail, head, arr):
* 	splitHalf = Split(arr)
* 	anotherSplit = Split(splitHalf)
*
 */
