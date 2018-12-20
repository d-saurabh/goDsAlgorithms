package main

func main()  {
	s:= stack{}
	s.init(5)
	s.push(2)
	s.push(5)
	s.push(4)
	s.push(6)
	s.push(5)
	s.push(8)
	println(s.pop())
	println(s.pop())
	println(s.pop())
	println(s.pop())
	println(s.pop())
	println(s.pop())
}

type stack struct{
	arr []int
	top int
	item int
	length int
}

func (s *stack) init(l int)  {
	s.length = l
	s.arr = make([]int,0,l)
}
func (s *stack) push(item int)()  {
	if s.isFull(){
		println("stack full")
		return
	}
	s.arr = append(s.arr,item)
	s.top++
}

func (s *stack) pop() int  {
	if s.isEmpty(){
		println("stack is empty")
		return -1
	}
	i:= s.top -1
	s.top--
	item:= s.arr[i]
	s.arr = append(s.arr[:i])
	return item
}

func (s *stack) peek() int  {
	return s.arr[s.top]
}

func (s *stack) isFull() bool {
	return s.top == s.length
}

func (s *stack) isEmpty() bool {
	return s.top == 0
}