package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	node *Node
}

type LL interface {
	addDataToLL(n int)
	findLength() int
	display()
	delete(n int)
	checkValExists(n int)
	findKthIndex(n int)
	findMiddleElement()
	reverseLL()
	sortTwoLL(ll1 LinkedList, ll2 LinkedList)
}

func main() {
	//Reverse the LL
	//Given two LL, sort it
	var l = &LinkedList{}
	//Add an element
	l.addDataToLL(2)
	l.addDataToLL(3)
	l.addDataToLL(1)
	l.addDataToLL(9)
	l.addDataToLL(5)
	l.addDataToLL(7)
	//Show all elements

	l.display()

	//Find length of the LL
	fmt.Println("Total no.of values:", l.findLength())

	//Delete an element from LL
	//l.delete(9)

	//Show after delete elements
	//l.display()

	//Check if a particular value is present or not
	l.checkValExists(20)

	//Find kth index value from last
	//1st approach - find length - O(N) + find N-k-1 index value
	//2nd approach - two pointer, both starts at 1st index and 2nd jumps to k-1 index in every iteration.
	   //so moment 2nd ptr reaches end, return 1st index vale
	l.findKthIndex(4)

	//Find kth index value from last
	//1st approach - find length - O(N) + find N/2 index value
	//2nd approach - slow fast pointer, both starts at 1st index and 2nd jumps twice to k-1 index in every iteration.
	   //so moment 2nd ptr reaches end, return 1st index vale
	l.findMiddleElement()

	//Reverse the LL
	l.reverseLL()
	l.display()

	//

	//Clone the LL

}
func (l *LinkedList) addDataToLL(inp int) {
	n := &Node{data: inp, next: nil}
	if l.node == nil {
		l.node = n
	} else {
		p := l.node
		for p.next != nil {
			p = p.next
		}
		p.next = n
	}
}
func (l *LinkedList) display() {
	p := l.node
	for p.next != nil {
		fmt.Println("Data:", p.data)
		p = p.next
	}
	fmt.Println("Data:", p.data)
}
func (l *LinkedList) findLength() int {
	c := 0
	p := l.node
	for p.next != nil {
		c++
		p = p.next
	}
	c++
	return c
}
func (l *LinkedList) delete(n int) {
	p := l.node
	var prev *Node
	for p.next != nil {
		if p.data == n && prev != nil {
			prev.next = p.next
			return
		} else if p.data == n && prev == nil {
			l.node = p.next
			return
		}
		prev = p
		p = p.next
	}
	if p.data == n {
		prev.next = nil
	}
}
func (l *LinkedList) checkValExists(n int) {
	p := l.node
	for p.next != nil {
		if p.data == n {
			fmt.Println(n, ": value is present")
			return
		}
		p = p.next
	}
	if p.data == n {
		fmt.Println(n, ": value is present")
		return
	}
	fmt.Println(n, ": value is not present")
}
func (l *LinkedList) findKthIndex(n int) {
	count := l.findLength()
	p := l.node
	ind := 0
	for p.next != nil {
		if count-n+1 == ind {
			fmt.Println("Kth index value:", p.data)
			return
		}
		p = p.next
		ind++
	}
	if count-n+1 == ind {
		fmt.Println("Kth index value:", p.data)
		return
	}
}
func (l *LinkedList) findMiddleElement() {

	//slow,fast:=l.node,l.node
	mid := l.findLength() / 2
	p := l.node
	c := 0
	fmt.Println("Mid index:", mid)
	for p.next != nil {
		if mid == c {
			fmt.Println("Mid element:", p.data)
			return
		}
		p = p.next
		c++
	}
}

// Put all data into stack and frame LL freshly
func (l *LinkedList) reverseLL() {
	p := l.node
	stack := make([]int, 0)
	for p.next != nil {
		stack = append(stack, p.data)
		p = p.next
	}
	stack = append(stack, p.data)
	fmt.Println(stack)
	node := &Node{data: stack[len(stack) - 1], next: nil}
	l.node = node
	p = l.node
	for i := len(stack) - 2; i >= 0; i-- {
		node = &Node{data: stack[i], next: nil}
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
	p.next = node
}
func (l *LinkedList) sortTwoLL(ll1 LinkedList, ll2 LinkedList) {

}