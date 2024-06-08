package list

import (
	"fmt"
	"reflect"
)

type NodeValue interface{}

type List struct {
	length int
	head   *Node
	tail   *Node
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

// New returns a pointer to a List initialized to the concrete value stored in the interface NodeValue. New() returns a pointer to a zero List.
func New(val ...NodeValue) (list *List) {
	list = &List{}
	list.AddValue(val...)
	return
}

// AddValue appends elements to the head of a List.
//
//	list := &List{}
//	list.AddValue(elem1, elem2, elem3)
//	vals := []NodeValue{elem1, elem2, elem3}
//	list.AddValue(vals...)
func (list *List) AddValue(val ...NodeValue) {
	last := len(val) - 1
	for i := range val {
		v := val[last-i]

		if list.length == 0 {
			list.head = &Node{value: v}
			list.tail = list.head
			list.length++
		} else if list.length == 1 {
			temp := list.head
			list.head = &Node{value: v}
			list.tail = temp
			list.head.next = list.tail
			list.tail.prev = list.head
			list.length++
		} else {
			temp := list.head
			list.head = &Node{value: v}
			temp.prev = list.head
			list.head.next = temp
			list.length++
		}
	}
}

// AddList appends elements of Lists to the end of a current List
//
//	list.AddList(list1, list2, list3)
func (list *List) AddList(lists ...*List) {
	for _, l := range lists {
		list.addList(l)
	}
}

// addList appends elements of a List to the end of a current List
//
//	list.addList(newlist)
func (list *List) addList(addedList *List) {
	if list.length == 0 && addedList.length != 0 {
		list.head = addedList.head
		list.tail = addedList.tail
		list.length += addedList.length
	} else if list.length != 0 && addedList.length != 0 {
		list.tail.next = addedList.head
		addedList.head.prev = list.tail
		list.length += addedList.length
		list.tail = addedList.tail
	}
}

// Contains reports whether val is present in a current List.
func (list *List) Contains(val NodeValue) (bool, *Node) {
	current := list.head
	for i := 0; i < list.length; i++ {
		if current.Equal(val) {
			return true, current
		} else {
			current = current.next
		}
	}
	return false, nil
}

// Remove deletes the element with the val from the List. If List is empty or there is no such element, delete is a no-op.
func (list *List) Remove(val NodeValue) bool {
	for i := 0; i < list.length; i++ {
		if ok, deleted := list.Contains(val); ok {
			prev := deleted.prev
			next := deleted.next

			switch deleted {
			case list.head:
				list.head = next
			case list.tail:
				list.tail = prev
			}

			if prev != nil {
				prev.next = next
			}

			if next != nil {
				next.prev = prev
			}

			list.length--

			return true
		}
	}
	return false
}

func (list *List) DeleteNode(node *Node) bool {
	if node == nil || list.length == 0 {
		return false
	} else if node == list.head && node == list.tail && list.length == 1 {
		list.head, list.tail = nil, nil
		list.length--
		return true
	} else if node == list.head {
		list.head = node.next
		node.next.prev = node.prev
		list.length--
		return true
	} else if node == list.tail {
		list.tail = node.prev
		node.prev.next = node.next
		list.length--
		return true
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		list.length--
		return true
	}
}

func (list *List) PrintAll() string {
	current := list.head
	if list.length == 0 {
		return fmt.Sprintf("List: %v\n", list)
	}

	str := fmt.Sprintf("List: %v\n", list)

	for i := 0; i < list.length; i++ {
		if current == nil {
			break
		}
		str += fmt.Sprintf("\tNode #%d: [%p] %v\n", i, current, current)
		current = current.next
	}
	return str
}

// Head returns head node in the List
func (list *List) Head() *Node {
	return list.head
}

// Tail returns tail node in the List
func (list *List) Tail() *Node {
	return list.tail
}

// Next returns the next node relative to the node specified in the argument
func (list *List) Next(node *Node) *Node {
	if node != list.tail {
		return node.next
	} else {
		return list.head
	}
}

// Prev returns the previous node relative to the node specified in the argument
func (list *List) Prev(node *Node) *Node {
	if node != list.head {
		return node.prev
	} else {
		return list.tail
	}
}

// The Tour goes through the entire list from the head node to the tail node and
// for each node calls a function to which it passes the node value and the iteration number.
// If the function returns true, the loop is interrupted.
//
//	list.Tour(func(val NodeValue, i int) bool {
//		fmt.Println(val, i)
//		if i == 2 {
//			return true
//		}
//	})
func (list *List) Tour(f func(*Node, int) bool) {
	if list.length < 1 {
		return
	}
	current := list.head
	end := false
	for i := 0; !end; i++ {
		if current == list.tail {
			end = true
		}
		if f(current, i) {
			return
		}
		current = current.next
	}
}

// Length returns the length of the List
func (list *List) Length() int {
	return list.length
}

// Get returns the value of the node
func (n *Node) Get() NodeValue {
	return n.value
}

// Set sets the value for the node
func (n *Node) Set(val NodeValue) {
	n.value = val
}

// Equal compares the value of nodes, if the values are equal returns true otherwise false
func (n *Node) Equal(val NodeValue) bool {
	return reflect.ValueOf(n.value).Equal(reflect.ValueOf(val))
}

// ToInt converts node values to int type
func ToInt(val NodeValue) int {
	if dec, ok := val.(int); ok {
		return dec
	} else {
		panic("cannot convert to int")
	}
}
