package recursion

import (
	"fmt"
	"time"

	ll "abusafia.com/algorithms/internal/service/list"
	randomindex "abusafia.com/algorithms/internal/service/random-index"
)

type Recursion struct {
	list    []int
	typerec string
}

func New(list []int, typerec string) *Recursion {
	return &Recursion{list: list, typerec: typerec}
}

func (r *Recursion) Start() (string, error) {
	switch r.typerec {
	case "sum":
		return fmt.Sprintf("sum:\n%d", sum(r.list)), nil
	case "count":
		return fmt.Sprintf("count:\n%d", count(r.list)), nil
	case "max":
		return fmt.Sprintf("max:\n%d", max(r.list)), nil
	case "binsearch":
		return fmt.Sprintf("binsearch:\nsteps: %d", binsearchdecor(r.list)), nil
	case "selsort":
		return fmt.Sprintf("selsort:\nlist: %d", selsort(r.list, 0)), nil
	case "qsort":
		return fmt.Sprintf("qsort:\nlist: %d", qsort(r.list)), nil
	case "qsortlist":
		return fmt.Sprintf("qsortlist:\nlist: %v", qsortlistdecor(r.list)), nil
	default:
		return "", nil
	}
}

func sum(list []int) int {
	if len(list) == 0 {
		return 0
	} else if len(list) == 1 {
		return list[0]
	}
	return list[0] + sum(list[1:])
}

func count(list []int) int {
	if len(list) == 0 {
		return 0
	} else if len(list) == 1 {
		return 1
	}
	return 1 + count(list[1:])
}

func max(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	fmax := max(list[1:])
	if list[0] > fmax {
		return list[0]
	} else {
		return fmax
	}
}

func binsearchdecor(list []int) int {
	rnd := randomindex.New(time.Now().Unix())
	index := rnd.Random.Intn(len(list))
	item := list[index]

	res := binsearch(list, item)
	return res
}

func binsearch(list []int, item int) int {
	if len(list) == 0 {
		return 0
	}
	mid := (len(list) - 1) / 2
	guess := list[mid]
	if guess == item {
		return 1
	}
	if guess > item {
		return 1 + binsearch(list[:mid], item)
	} else {
		return 1 + binsearch(list[mid+1:], item)
	}
}

func selsort(list []int, offset int) []int {
	if len(list)-1 == offset {
		return list
	}

	index_smallest := offset + smallest(list[offset:])
	list[offset], list[index_smallest] = list[index_smallest], list[offset]
	offset++
	return selsort(list, offset)
}

func smallest(list []int) int {
	l := len(list) - 1

	if l == 0 {
		return 0
	}

	small := smallest(list[:l])
	if list[l] < list[small] {
		return l
	} else {
		return small
	}
}

func qsort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	var left, right []int
	rnd_ind := randomindex.New(time.Now().UnixNano()).Random.Intn(len(list))
	pivot := list[rnd_ind]

	for i, v := range list {
		if i == rnd_ind {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(qsort(left), append([]int{pivot}, qsort(right)...)...)
}

func qsortlistdecor(list []int) []int {
	nv := []ll.NodeValue{}

	for _, v := range list {
		nv = append(nv, v)
	}
	l := ll.New(nv...)

	res := qsortlist(l)

	arr := make([]int, res.Length())
	res.Tour(func(node *ll.Node, i int) bool {
		arr[i] = ll.ToInt(node.Get())
		return false
	})
	return arr
}

func qsortlist(list *ll.List) *ll.List {
	if list.Length() <= 1 {
		return list
	}

	left, right := ll.New(), ll.New()
	pivot := list.Head().Get()

	list.Tour(func(node *ll.Node, i int) bool {
		if i != 0 {
			if ll.ToInt(node.Get()) < ll.ToInt(pivot) {
				left.AddValue(node.Get())
			} else {
				right.AddValue(node.Get())
			}
		}
		return false
	})

	out := ll.New()
	out.AddList(qsortlist(left), ll.New(pivot), qsortlist(right))
	return out
}
