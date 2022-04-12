package sliceUtils

import (
	"math/rand"
)

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a
}

func ReverseArray(a []int) []int {
	if len(a) == 0 {
		panic(any("Empty Slices Not Allowed"))
		return nil
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
