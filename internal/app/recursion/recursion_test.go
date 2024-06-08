package recursion

import (
	"fmt"
	"testing"

	ll "abusafia.com/algorithms/internal/service/list"
)

var list []int = []int{14, 18, 15, 13, 111, 16, 14, 15}

func TestSmallest(t *testing.T) {
	ans := smallest(list)
	res := 3
	if ans != res {
		t.Errorf("smallest([%v]) = %d; want %d", list, ans, res)
	}

}

func TestSelsort(t *testing.T) {
	ans := selsort(list, 0)
	res := []int{13, 14, 14, 15, 15, 16, 18, 111}
	if len(res) == len(ans) {
		for i, v := range res {
			if v != ans[i] {
				t.Errorf("selsort([%v]) = %d; want %d", list, ans, res)
			}
		}
	} else {
		t.Errorf("selsort([%v]) = %d; want %d", list, ans, res)
	}
}

func TestQsort(t *testing.T) {
	ans := qsort(list)
	res := []int{13, 14, 14, 15, 15, 16, 18, 111}
	if len(res) == len(ans) {
		for i, v := range res {
			if v != ans[i] {
				t.Errorf("qsort([%v]) = %d; want %d", list, ans, res)
			}
		}
	} else {
		t.Errorf("qsort([%v]) = %d; want %d", list, ans, res)
	}

}

func TestQsortList(t *testing.T) {
	vals := []ll.NodeValue{5, 33, 4, 10, 6}
	list := ll.New(vals...)
	ans := qsortlist(list)
	str := ans.PrintAll()
	fmt.Println(str)

	t.Error()
}
