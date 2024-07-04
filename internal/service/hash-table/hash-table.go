package hashtable

import (
	"crypto/sha256"
	"fmt"

	ll "abusafia.com/algorithms/internal/service/list"
)

type HashT struct {
	store []ll.List
	count int
}

func New() *HashT {
	store := make([]ll.List, 100)
	return &HashT{store: store}
}

func (h *HashT) Set(key string, value interface{}) {
	if set(h.store, key, value) {
		h.count++
		h.resize()
	}
}

func (h *HashT) Get(key string) interface{} {
	value, _ := h.GetOk(key)
	return value
}

func (h *HashT) GetOk(key string) (interface{}, bool) {
	node := getNode(h.store, key)
	if node != nil {
		data := (node.Get()).([]interface{})
		value := data[1]
		return value, true
	} else {
		return nil, false
	}
}

func (h *HashT) Remove(key string) bool {
	node := getNode(h.store, key)
	if node != nil {
		list := &h.store[h.hash(key)]
		if list.DeleteNode(node) {
			h.count--
			return true
		}
		return false
	} else {
		return false
	}
}

func (h *HashT) hash(key string) int {
	return hash256(key, len(h.store))
}

func (h *HashT) resize() {
	fmt.Println()
	if h.count*100/len(h.store) > 70 {
		store := make([]ll.List, 2*len(h.store))

		for _, bucket := range h.store {
			bucket.Tour(func(n *ll.Node, i int) bool {
				if data, ok := (n.Get()).([]interface{}); ok {
					key := data[0].(string)
					value := data[1]
					fmt.Printf("\t\tHash[\"%s\"]:%d\n", key, hash256(key, len(store)))
					set(store, key, value)
				}
				return false
			})
		}
		h.store = store
	}
}

func hash256(key string, l int) int {
	h := sha256.Sum256([]byte(key))
	sum := 0
	for _, v := range h {
		sum += int(v)
	}
	return sum % l
}

func set(store []ll.List, key string, value interface{}) bool {
	list := &store[hash256(key, len(store))]

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
		return true
	} else {
		return false
	}
}

func getNode(store []ll.List, key string) *ll.Node {
	list := &store[hash256(key, len(store))]
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
