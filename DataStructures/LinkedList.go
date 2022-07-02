package main

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}



/*The length() function takes a linked list as input and counts the number of nodes in the list. To find the length of
a linked list, one would normally have to traverse all nodes, and count them. This operation takes O( ) time. In
order to have this operation done in constant time, we can add the size field in linked list structure definition.
Hence, finding the list size is trivial:
*/
// with extra size field
func (ll *LinkedList) Length() int {
	return ll.size
}

// length returns the linked list size
func (ll *LinkedList) Length2() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

//insert at the beginning

func (ll *LinkedList) InsertBeginning(data interface{}) {
	node := &ListNode{data: data}
	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
	return
}

// insert at the end

func (ll *LinkedList) InsertEnd(data interface{}) {
	node := &ListNode{data: data}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
	return
}

// Insert adds an item at position i
func (ll *LinkedList) Insert(position int, data interface{}) error {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	newNode := &ListNode{data, nil}
	var prev, current *ListNode
	prev = nil
	current = ll.head
	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}
	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}
	ll.size++
	return nil
}

func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteLast: List is empty")
	}
	var prev *ListNode
	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.data, nil
}

// delete removes an element at position i
func (ll *LinkedList) Delete(position int) (interface{}, error) {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	var prev, current *ListNode
	prev = nil
	current = ll.head
	pos := 0
	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}
		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}

func main() {
	ln := ListNode{12, nil}
	ln2 := ListNode{12, nil}
	ln.next = &ln2
	ll := LinkedList{&ln, 5}
	ll.Display()
}
