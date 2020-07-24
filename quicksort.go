package main

import "fmt"

func quickSort(a []int, p, r int) []int {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q - 1)
		quickSort(a, q + 1, r)
	}
	return a
}

// loop invariant:
// at the beginning of each iteration of the loop of lines 28-31,
// for any array index k,
//
// 1. if p <= k <= i, then A[k] <= x. because i only increments when
// j <= the pivot element.
// 2. if i+1 <= k <= j-1, then A[k] > x. if k is within the partition
// then A[k] is greater than the pivot.
// 3. if k = r, then A[k] = x. if k is the last element i.e, the pivot.

func partition(a []int, p, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] <= x {
			i = i + 1 // tradional notation
			// exchange i with j
			a[i], a[j] = a[j], a[i]
			// swap done
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func main() {
	a := []int{9, 8, 7, 1, 3, 5, 6, 4}
	p := 0
	r := len(a) - 1
	sorted := quickSort(a, p, r)
	fmt.Println(sorted)
}
