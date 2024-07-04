package primenumbers

import (
	ll "abusafia.com/algorithms/internal/service/list"
)

func primeNumber(n int) []int {
	list := ll.New()
	for i := n; i > 1; i-- {
		list.AddValue(i)
	}
	list.Tour(func(node1 *ll.Node, i int) bool {
		v1 := ll.ToInt(node1.Get())
		list.Tour(func(node2 *ll.Node, j int) bool {
			v2 := ll.ToInt(node2.Get())
			if v2%v1 == 0 && v2 != v1 {
				list.DeleteNode(node2)
			}
			return false
		})
		return false
	})
	primeNums := make([]int, list.Length())
	list.Tour(func(n *ll.Node, i int) bool {
		primeNums[i] = ll.ToInt(n.Get())
		return false
	})
	return primeNums
}
