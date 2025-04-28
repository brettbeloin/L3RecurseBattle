# Pseudocode for mSort function
Created: 4/27/2025, 10:32:51 PM
Last Updated: 4/27/2025, 10:32:55 PM
---

### User

create psuedocode for this:
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


---

### Assistant (gpt-4.1-mini)

Here's the pseudocode for your `mSort` function, which implements a merge sort algorithm:

```
function mSort(arr):
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
```

Let me know if you'd like me to explain any part of it!


---
