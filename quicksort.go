package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	//a:= []int{54,26,93, 17,77,31,44,55,20}
	a:= []int{54,26,93, 17}
	fmt.Print(quicksort(a))
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left := 0
	right := len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
