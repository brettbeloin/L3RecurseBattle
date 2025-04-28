package main

import (
	//"fmt"
	"golang.org/x/exp/constraints"
)

//type generic interface {
//	~int | ~string | ~float64
//}

func main() {
	//fmt.Println(mSort(input))
}

/*
chat model: chat gpt 3.5
why: Used to help fix a bug with the split func
chat log: https://chatgpt.com/share/6807d1fa-49ec-8006-acd4-19cd99f08f20
*/

func split[T constraints.Ordered](arr []T) int {
	mid := len(arr) / 2
	return mid
}

/*
chat model: chat gpt 3.5
why: needed help fixing the compairtion logic for a generic type
chat log: https://chatgpt.com/share/6807d668-4890-8006-ae55-75415959763f
*/

func mSort[T constraints.Ordered](arr []T) []T {
	cpArr := make([]T, len(arr))
	copy(cpArr, arr)

	length := split(arr)
	left := cpArr[:length]
	right := cpArr[length:]

	if len(left) != 1 {
		left = mSort(left)
	}

	if len(right) != 1 {
		right = mSort(right)
	}

	leftIdx := 0
	rightIdx := 0

	for idx, _ := range arr {

		if len(right) == rightIdx {
			arr[idx] = left[leftIdx]
			leftIdx++
			continue
		}
		if len(left) == leftIdx {
			arr[idx] = right[rightIdx]
			rightIdx++
			continue
		}

		if len(right) != rightIdx && left[leftIdx] < right[rightIdx] {
			arr[idx] = left[leftIdx]
			leftIdx++
		}
		if len(left) != leftIdx && right[rightIdx] < left[leftIdx] {
			arr[idx] = right[rightIdx]
			rightIdx++
		}
	}
	return arr
}

func qSort[T constraints.Ordered](arr []T, tail, head int) []T {
	for {
		pivit := arr[0]
		pivitIdx := 0
		compare := len(arr) - 1
		compareIdx := 0
		temp := 0

		if arr[compare] < pivit {
			temp = pivitIdx
			arr[pivitIdx] = arr[compareIdx]
			compareIdx = temp
			pivitIdx++
		}

		if arr[compare] > pivit {

		}
		return arr
	}
}

/*
* ai model: GTP-4.1 mini
* what was generated: mSort psudocode
* chat: look at in the md folder on github: https://github.com/brettbeloin/L3RecurseBattle
*
* * function mSort(arr):
    cpArr = copy of arr

    length = split(arr)  // split returns the midpoint index
    left = cpArr[0:length]
    right = cpArr[length:end]

    if length of left is not 1:
        left = mSort(left)

    if length of right is not 1:
        right = mSort(right)

    leftIdx = 0
    rightIdx = 0

    for idx in range length of arr:
        if rightIdx equals length of right:
            arr[idx] = left[leftIdx]
            leftIdx = leftIdx + 1
            continue loop

        if leftIdx equals length of left:
            arr[idx] = right[rightIdx]
            rightIdx = rightIdx + 1
            continue loop

        if left[leftIdx] < right[rightIdx]:
            arr[idx] = left[leftIdx]
            leftIdx = leftIdx + 1

        else if right[rightIdx] < left[leftIdx]:
            arr[idx] = right[rightIdx]
            rightIdx = rightIdx + 1

    return arr

* Quick Sort:
* func qSort(tail, head, arr):
*		mid = Split(arr)
*		left = arr[mid]
*		right = arr[mid]
*
*		if left or right dont't have one elm:
*			qSort(left)
*			qsort(right)

		for idx range arr:

*/
