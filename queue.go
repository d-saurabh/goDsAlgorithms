package main

func main()  {
	q:= queue{}
	q.init(5)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)

	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
}

type queue struct{
	arr []int
	length int
}

func (q * queue) enqueue(item int)  {
	if q.isFull(){
		println("queue is full")
		return
	}
	q.arr = append(q.arr,item)
}


func (q * queue) dequeue() int {
	if q.isEmpty(){
		println("queue is empty")
		return -1
	}
	i:= q.arr[0]
	q.arr = append(q.arr[1:])
	return i
}


func (q * queue) isFull() bool {
   return len(q.arr) == q.length
}


func (q * queue) isEmpty() bool {
	return len(q.arr) == 0
}


func (q * queue) init(l int)  {
	q.length= l
	q.arr = make([]int,0,l)
}