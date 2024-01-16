package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type List struct {
	head *Node
}

func main() {
	var l List
	l.addToLinkedList(3)
	l.addToLinkedList(10)
	l.addToLinkedList(1)
	l.addToLinkedList(4)
	l.addToLinkedList(6)
	l.addToLinkedList(7)
	l.addToLinkedList(2)
	l.display()
	l.findNode(2)
	l.findMiddleElement()

	//Add an element
	//Show all elements
	//Find length of the LL
	//Check if a particular value is present or not
	//Find kth index value from last
	//Find the middle element
	//Reverse the LL
	//Given two LL, sort it
	
}
func (list *List) addToLinkedList(newNode int) {
	n := &Node{data: newNode, next: nil}
	if list.head == nil {
		list.head = n
	} else {
		p := list.head
		for p.next != nil {
			p = p.next
		}
		p.next = n
	}
}
func (list *List) display() {
	p := list.head
	for p.next != nil {
		fmt.Println("Data :", p.data)
		p = p.next
	}
	fmt.Println("Data :", p.data)
}
func (node *List) findNode(k int) {
	c := 0
	p := node.head
	for p.next != nil {
		c++
		p = p.next
	}
	c++
	fmt.Println("Total no.of values:", c)

	temp := node.head
	newC := 0
	for temp.next != nil {
		if c-k+1 == newC {
			fmt.Println("Middle element:", temp.data)
			break
		} else {
			temp = temp.next
			newC++
		}
	}

	// iPtr, jPtr := node.head, node.head

	// //with two-pointer
	// for jPtr.next != nil {
	// 	if jPtr.next.next == nil {
	// 		iPtr = iPtr.next
	// 		break
	// 	}else{
	// 		iPtr = iPtr.next
	// 		jPtr = jPtr.next.next
	// 	}
	// }
	// fmt.Println("Got Kth element:", iPtr.data)
}
func (l *List) findMiddleElement() {
	// p := l.head
	// c := 0
	// for p.next != nil {
	// 	c++
	// 	p = p.next
	// }
	// c++
	// fmt.Println("Count:", c)

	// temp := l.head
	// mid, t := c/2, 0
	// for temp.next != nil {
	// 	t++
	// 	temp = temp.next
	// 	if t == mid {
	// 		fmt.Println("Mid element:", temp.data)
	// 		break
	// 	}
	// }

	fPtr, sPtr := l.head, l.head
	speed := 0
	for sPtr.next != nil {
		fmt.Println(fPtr.data,sPtr.data)
		speed=2*speed
		fPtr = fPtr.next
		fast:=speed
		for fast > 0 {
			sPtr = sPtr.next
			fast--
		}
	}
	fmt.Println(fPtr.data)
}
