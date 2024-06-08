package hashtable

import (
	"strings"

	ll "abusafia.com/algorithms/internal/service/list"
)

type HashT struct {
	primeNums []int
	store     []ll.List
	count     int
}

func New() *HashT {
	store := make([]ll.List, 40)
	return &HashT{store: store}
}

func (h *HashT) hash(key string) int {
	filter := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	if len(h.primeNums) == 0 {
		h.primeNums = make([]int, len(filter))
		primeNumber(307).Tour(func(n *ll.Node, i int) bool {
			h.primeNums[i] = ll.ToInt(n.Get())
			return false
		})
	}

	sum := 0
	for _, k := range key {
		index := strings.IndexRune(filter, k)
		sum += h.primeNums[index]
	}
	return sum % len(h.store)
}

func (h *HashT) Set(key string, value interface{}) {
	list := &h.store[h.hash(key)]

	isNew := true
	list.Tour(func(n *ll.Node, i int) bool {
		if vint, ok := (n.Get()).([]interface{}); ok {
			if k, ok := vint[0].(string); ok {
				if k == key {
					n.Set([]interface{}{key, value})
					isNew = false
					return true
				}
			}
		}
		return false
	})
	if isNew {
		list.AddValue([]interface{}{key, value})
		h.count++
	}
}

// func (h *HashT) set(store []ll.List, key string, value interface{}) {
// 	list := &store[h.hash(key)]

// 	detected := false
// 	list.Tour(func(n *ll.Node, i int) bool {
// 		if vint, ok := (n.Get()).([]interface{}); ok {
// 			if k, ok := vint[0].(string); ok {
// 				if k == key {
// 					n.Set([]interface{}{key, value})
// 					detected = true
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	})
// 	if !detected {

// 		list.AddValue([]interface{}{key, value})
// 		h.count++
// 	}
// }

func (h *HashT) Get(key string) interface{} {
	value, _ := h.GetOk(key)
	return value
}

func (h *HashT) GetOk(key string) (interface{}, bool) {
	node := h.get(key)
	if node != nil {
		data := (node.Get()).([]interface{})
		value := data[1]
		return value, true
	} else {
		return nil, false
	}
}

func (h *HashT) Remove(key string) bool {
	node := h.get(key)
	if node != nil {
		list := &h.store[h.hash(key)]
		list.DeleteNode(node)
		h.count--
		return true
	} else {
		return false
	}
}

func (h *HashT) get(key string) *ll.Node {
	list := &h.store[h.hash(key)]
	var value *ll.Node
	list.Tour(func(n *ll.Node, i int) bool {
		if vint, ok := (n.Get()).([]interface{}); ok {
			if k, ok := vint[0].(string); ok {
				if k == key {
					value = n
					return true
				}
			}
		}
		return false
	})
	return value
}

// func (h *HashT) resize() {
// 	if h.count*100/len(h.store) > 70 {
// 		store = make([]ll.List, 2*len(h.store))
// 		for _, bucket := range h.store {
// 			bucket.Tour(func(n *ll.Node, i int) bool {
// 				if data, ok := (n.Get()).([]interface{}); ok {
// 					key := data[0].(string)
// 					value := data[1]
// 				}

// 			})
// 		}
// 	}
// }

func primeNumber(n int) *ll.List {
	list := ll.New()
	for i := n; i > 1; i-- {
		list.AddValue(i)
	}
	// fmt.Printf("Length:%d\n\n", list.Length())

	list.Tour(func(node1 *ll.Node, i int) bool {
		// fmt.Printf("\n#%d-----------{%v}----------- Len:%d\n", i, node1, list.Length())
		list.Tour(func(node2 *ll.Node, j int) bool {
			// fmt.Printf("node1: %v\t|%v:%v|\t node2: %v\n", node1, list.Head(), list.Tail(), node2)
			v1 := ll.ToInt(node1.Get())
			v2 := ll.ToInt(node2.Get())

			if v2%v1 == 0 && v2 != v1 {
				// fmt.Printf("%d\t--> Del:\t%v\t\tLen:%d\n", j, node2, list.Length())
				list.DeleteNode(node2)
			} else {
				// fmt.Printf("%d\t--> NotDel:\t%v\t\tLen:%d\n", j, node2, list.Length())
				_ = "skip"
			}
			return false
		})
		return false
	})
	return list
}
