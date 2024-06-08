package selectionsort

import "fmt"

type SelectionSort struct {
	list []int
}

func New(list []int) *SelectionSort {
	return &SelectionSort{list: list}
}

func (ss *SelectionSort) Start() (string, error) {
	for i := range ss.list {
		if i == len(ss.list) {
			break
		}
		index_smallest := ss.findSmallest(i)
		ss.list[i], ss.list[index_smallest] = ss.list[index_smallest], ss.list[i]
	}
	return fmt.Sprint(ss.list), nil
}

func (ss *SelectionSort) findSmallest(offset int) int {
	smallest := ss.list[offset]
	index_smallest := offset
	for i, el := range ss.list[offset+1:] {
		if el < smallest {
			smallest = el
			index_smallest = offset + 1 + i
		}
	}
	return index_smallest
}
