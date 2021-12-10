package DoublyLinkedList

type DLLNode struct { // defines a DLLNode in a doubly linked list
	data interface{} // the datum
	prev *DLLNode    // pointer to the previous DLLNode
	next *DLLNode    // pointer to the next DLLNode
}

type DLL struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func (dll *DLL) CheckIfEmpty() bool {
	if dll.size == 0 {
		return true
	}
	return false
}

// CheckIfEmptyAndAdd ... check if doubly link list is empty
func (dll *DLL) CheckIfEmptyAndAdd(newNode *DLLNode) bool {
	// check if list is empty
	if dll.size == 0 {
		// insert first node in doubly linked list
		dll.head = newNode
		dll.tail = newNode
		dll.size++
		return true
	}
	return false
}

// InsertBeginning ... insert in the beginning of doubly linked list
func (dll *DLL) InsertBeginning(data int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.head
		// update newnode links - prev and next
		newNode.next = head
		newNode.prev = nil
		// update head node
		head.prev = newNode
		// update dll start and length
		dll.head = newNode
		dll.size++
		return
	}
	return
}

// InsertEnd ... inserts a node in the end of doubly linked list
func (dll *DLL) InsertEnd(data int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.head
		for i := 0; i < dll.size; i++ {
			if head.next == nil {
				// update newnode links - prev and next
				newNode.prev = head
				newNode.next = nil
				//update head node
				head.next = newNode
				// update dll end and size
				dll.tail = newNode
				dll.size++
				break
			}
			head = head.next
		}
	}
	return
}

// Insert ... insert between two nodes in doubly linkedlist
func (dll *DLL) Insert(data int, loc int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.head
		for i := 1; i < dll.size; i++ {
			if i == loc {
				// update newnode links - prev and next
				newNode.prev = head.prev
				newNode.next = head
				// update head node
				head.prev.next = newNode
				head.prev = newNode
				//keep traversing till we find the location
				dll.size++
				return
			}
			head = head.next
		}
	}
	return
}

func (dll *DLL) DeleteFirst() interface{} {
	// DeleteFirst ... Delete last element
	if !(dll.CheckIfEmpty()) {
		head := dll.head
		if head.prev == nil {
			deletedNode := head.data
			dll.head = head.next
			// update doubly linked list
			dll.head.prev = nil
			dll.size--
			return deletedNode
		}
	}
	return -1
}

// DeleteLast ... deletes last element from doubly linked list
func (dll *DLL) DeleteLast() interface{} {
	if !(dll.CheckIfEmpty()) {
		// delete from last
		head := dll.head
		for {
			if head.next == nil {
				break
			}
			head = head.next
		}
		// update doubly linked list
		dll.tail = head.prev
		dll.tail.next = nil
		dll.size--
		return head.data
	}
	return -1
}

// Delete .. delete from any location
func (dll *DLL) Delete(pos int) interface{} {
	if !(dll.CheckIfEmpty()) {
		//delete from any position
		head := dll.head
		for i := 1; i <= pos; i++ {
			if head.next == nil && pos > i {
				//list is lesser than given position
				return -1
			} else if i == pos {
				// delete from this location
				head.prev.next = head.next
				head.next.prev = head.prev
				dll.size--
				return head.data
			}
			head = head.next
		}
	}
	return -1
}
