package main

func main()  {
	a:= []int{1,2,3,4,5,6,7,8,9}
	println(binarySearch(a,1))
}

func binarySearch(a []int,n int) int  {
	beg:=0
	end:=len(a)
	mid:=(beg+end)/2

	for i:=0;i<len(a) ;i++  {
		if n < a[mid]{
			end = mid -1
		} else if n > a[mid]{
			beg = mid + 1
		}
		if a[mid] == n{
			println("found the %d at %d",n,mid)
			return mid
		}
		mid =(beg+end)/2
	}
	return -1
}
