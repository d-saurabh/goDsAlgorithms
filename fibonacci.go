package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	rand.Seed(63)
	ch := make(chan int)
	done:= make(chan bool)
	go func(){
		for {
			select {
				case ch <- rand.Int():
			case <-done:
				return
			default:
			}
		}
	}()

	go func() {
		for i:=0;i<5 ;i++  {
			fmt.Println(<-ch)
		}
		done <- true
		//return
	}()

	<- done

	//c := fib_generator()
	//for n := 0; n < 12; n++ {
	//	fmt.Println(<-c) // Print channel value
	//}
	//for i := 0; i < 5; i++ {
	//	go getRand(ch)
	//	go printRand(ch)
	//}
	//time.Sleep(10 * time.Second)
	//m:= map[int]string{
	//	1:"saurabh",
	//	3:"deshpande",
	//	5:"tina",
	//	2:"khawase",
	//}
	////for i,j:=range m{
	////	println(i)
	////	println(j)
	////}
	//keys:= make([]int,0, len(m))
	//for key:= range m{
	//	keys = append(keys,key)
	//}
	//
	//sort.Ints(keys)
	//for _,key:=range keys{
	//	println(m[key])
	//}
	//fmt.Print(m)
	//fmt.Print(m)
	//fmt.Print(m)
	//println(fib(3))

}

func fib(n int) int  {
	println(n)
	if n ==0 || n ==1{
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func getRand(ch chan int)  {
	ch <- rand.Int()
}

func printRand(ch chan int){
	println("%d",<- ch)
}

func fib_generator() chan int {
	c := make(chan int) // Create return channel
	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i // Send generated value into the channel
		}
	}()
	return c // return generator channel
}