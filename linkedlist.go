package main

import "fmt"

func main()  {
	l:= LinkedList{}
	l.append(2)
	l.append(3)
	l.append(4)

	l.insertAt(0,1)
	l.insertAt(2,100)

	l.removeAt(1)
}

type Node struct {
	next *Node
	item int
}

type LinkedList struct {
	head *Node
	size int
}

//Append an item to the linked list
func (l *LinkedList) append(item int)  {
	node := Node{item:item}

	if l.head==nil{
		l.head = &node
		l.size++
		return
	}
	last:=l.head
	for{
		if last.next == nil{
			break
		}
		last = last.next
	}
	last.next = &node
	l.size++
}

//insert an item at specified position
func (l *LinkedList) insertAt(i int, item int)  error {
	if i < 0 || i> l.size{
		return fmt.Errorf("index out of bound")
	}
	newNode:= Node{item:item,next:nil}
	if i == 0{
		newNode.next = l.head
		l.head = &newNode
		l.size++
		return nil
	}

	node:=l.head
	j:=0
	for j<i-1{
		node = node.next
		j++
	}
	newNode.next=node.next
	node.next =&newNode
	l.size++
	return nil
}

//remove an item  from a position
func (l *LinkedList) removeAt(i int) (int,error)  {
	if i< 0 || i > l.size{
		return 0,fmt.Errorf("index out of bound")
	}

	if i ==0{
		node:= l.head
		l.head =node.next
		l.size--
		return node.item,nil
	}

	node:= l.head
	j:=0
	for j< i-1{
		node = node.next
		j++
	}
	remove:= node.next
	node.next = remove.next
	l.size--
	return remove.item,nil
}

//index of item
func (l *LinkedList) indexOf(item int) int  {
	if l.head == nil{
		return -1
	}
	node:= l.head
	for j:=0;j<l.size ;j++  {
		if node.item == item{
			return j
		} else if node.next == nil{
			return -1
		} else{
			node = node.next
		}
	}
	return -1
}

func (l *LinkedList) isEmpty() bool  {
	return l.head==nil
}



