package main

import "fmt"

func main()  {
	a:= []int{8,5,6,2,1}
	fmt.Print(bubbleSort(a))
}

func bubbleSort(A []int) []int  {
	if len(A) == 1{
		return A
	}

	for i:=0;i<len(A) ;i++  {
		for j:=i+1;j<len(A) ;j++  {
			if A[i] >A[j]{
				temp:=A[i]
				A[i] = A[j]
				A[j] = temp
			}
		}
	}
	return A
}