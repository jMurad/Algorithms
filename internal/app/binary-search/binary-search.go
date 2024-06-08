package binarysearch

import (
	"fmt"
	"time"

	randomindex "abusafia.com/algorithms/internal/service/random-index"
)

type BinarySearch struct {
	list []int
}

func New(list []int) *BinarySearch {
	bs := &BinarySearch{}
	bs.list = list
	return bs
}

func (bs *BinarySearch) Start() (string, error) {
	rnd := randomindex.New(time.Now().Unix())
	index := rnd.Random.Intn(len(bs.list))

	begin := 0
	end := len(bs.list) - 1
	item := bs.list[index]

	var steps int
	for steps = 1; begin <= end; steps++ {
		mid := (begin + end) / 2
		guess := bs.list[mid]
		if guess == item {
			return fmt.Sprintf("found number: %d; steps: %d", guess, steps), nil
		} else if guess > item {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return fmt.Sprintf("found number: %d; steps: %d", bs.list[begin], steps), nil
}
