package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type testpair struct {
		vals []NodeValue
	}

	var tests []testpair = []testpair{
		{[]NodeValue{}},
		{[]NodeValue{1, "2", 3.0}},
		{[]NodeValue{1, "2", 3.0, 1, 1}},
		{[]NodeValue{1, "2", 3.0, "2", "2"}},
		{[]NodeValue{1, "2", 3.0, 3.3, 3.0}},
	}

	for _, test := range tests {
		list := New(test.vals...)
		current := list.head
		for _, v := range test.vals {
			if !current.Equal(v) {
				t.Errorf("%v is not added to list(%v)", v, test.vals)
			}
			current = current.next
		}
	}
}

func TestContains(t *testing.T) {
	type testpair struct {
		vals []NodeValue
		ans  NodeValue
		exp  bool
	}

	var tests = []testpair{
		{[]NodeValue{2, 4, 6, 8, 10}, 1, false},
		{[]NodeValue{2, 4, 6, 8, 10}, 2, true},
		{[]NodeValue{2.0, 4.1, 0.26, 8.0001, 1.100}, 0.26, true},
		{[]NodeValue{"2", 4, "6", 8, 10}, "2", true},
		{[]NodeValue{"2", 4.01, "0.6", 8, 0.104}, 4.01, true},
		{[]NodeValue{2, 4, 6, 8, 10}, 10, true},
	}
	for _, v := range tests {
		if ok, _ := New(v.vals...).Contains(v.ans); ok != v.exp {
			t.Error("Error")
		}
	}
}

func TestRemove(t *testing.T) {
	type testpair struct {
		vals      []NodeValue
		deletedes []NodeValue
		isdeleted []bool
	}

	var tests = []testpair{
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{1}, []bool{true}},
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{1, 44}, []bool{true, true}},
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{"45", 45, 20}, []bool{true, false, true}},
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{1, 23}, []bool{true, true}},
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{23, 4.4}, []bool{true, true}},
		{[]NodeValue{1, 44, "45", 4.4, "4.4", 10.1, 20, 23}, []NodeValue{1, 4.4, 44, "4.4", 10.1, 4.4, 23}, []bool{true, true, true, true, true, false, true}},
		{[]NodeValue{1, 44, "45", "45", "4.4", 10.1, 20, 23}, []NodeValue{1, 4.4, 44, "45", 10.1, 4.4, 23}, []bool{true, false, true, true, true, false, true}},
	}

	for _, test := range tests {
		list := New(test.vals...)
		length := list.length
		for i, del := range test.deletedes {
			if list.Remove(del) != test.isdeleted[i] {
				t.Errorf("{%v} not deleted from %v", del, test.vals)
			}
			if test.isdeleted[i] {
				length--
			}
			if list.length != length {
				t.Errorf("{%v} deleted from %v, but LENGTH is not decremented %d != %d", del, test.vals, list.length, length)
			}
		}
		fmt.Println(list.PrintAll())
	}
}

func TestNext(t *testing.T) {
	type testpair []NodeValue

	var tests []testpair = []testpair{
		[]NodeValue{1, 2, 3, 4, 5},
		[]NodeValue{"1", "2", "3", "4", "5"},
		[]NodeValue{1.0, 2.1, 3.2, 4.3, 5.4},
		[]NodeValue{1, "2", 3.2, "4", 5.9},
	}

	for i, test := range tests {
		list := New(test...)
		current := list.head
		for j := range test {
			if !current.Equal(test[j]) {
				t.Errorf("Test #%d, node #%d: Next value = %v", i, j, current.value)
			}
			current = list.Next(current)
		}
	}
}

func TestPrev(t *testing.T) {
	type testpair []NodeValue

	var tests []testpair = []testpair{
		[]NodeValue{1, 2, 3, 4, 5},
		[]NodeValue{"1", "2", "3", "4", "5"},
		[]NodeValue{1.0, 2.1, 3.2, 4.3, 5.4},
		[]NodeValue{1, "2", 3.2, "4", 5.9},
	}

	for i, test := range tests {
		list := New(test...)
		current := list.tail
		for j := range test {
			if !current.Equal(test[list.length-1-j]) {
				t.Errorf("Test #%d, node #%d: Next value = %v", i, j, current.value)
			}
			current = list.Prev(current)
		}
	}
}

func TestTour(t *testing.T) {
	type testpair []NodeValue

	var tests []testpair = []testpair{
		[]NodeValue{},
		[]NodeValue{222},
		[]NodeValue{10, 20, 30, 40, 50},
		[]NodeValue{10, "20"},
		[]NodeValue{"100", 200},
	}

	for _, test := range tests {
		list := New(test...)
		list.Tour(func(node *Node, i int) bool {
			if !reflect.ValueOf(node.Get()).Equal(reflect.ValueOf(test[i])) {
				t.Errorf("Test[%d]: %v != %v", i, node.Get(), test[i])
			}
			return false
		})
	}
}

func TestDeleteNode(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	head := list.Head()
	middle := list.Next(list.Next(head))
	tail := list.Tail()

	fmt.Println(1, list.PrintAll())

	list.DeleteNode(head)
	fmt.Println(2, list.PrintAll())

	list.DeleteNode(middle)
	fmt.Println(3, list.PrintAll())

	list.DeleteNode(tail)
	fmt.Println(4, list.PrintAll())

	list.DeleteNode(list.Head())
	fmt.Println(5, list.PrintAll())

	list.DeleteNode(list.Tail())
	fmt.Println(6, list.PrintAll())

	t.Error()
}
