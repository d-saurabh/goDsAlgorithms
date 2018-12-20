package main

func main()  {
	a:= []int{1,2,4,6,7,2,6,9}
	println(linerSerarch(a,9))
}

func linerSerarch(a []int, i int) int  {

	for index:=range a  {
		if a[index] == i{
			return index
		}
	}
	return -1
}
