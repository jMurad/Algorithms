package hashtable

import (
	"fmt"
	"testing"

	ll "abusafia.com/algorithms/internal/service/list"
)

func TestPrimeNumber(t *testing.T) {
	list := primeNumber(307)
	t.Errorf("\n%v\n", list.PrintAll())
}

func TestNew(t *testing.T) {
	mapa := New()

	fmt.Println(mapa)
	t.Error()
}

func TestSetGet(t *testing.T) {
	mapa := New()
	keys := []string{"EH", "FG", "DI", "HE", "GF", "ID", "Dagestan"}
	for i, key := range keys {
		mapa.Set(key, 100*(i+1))
	}

	for _, key := range keys {
		// if val, ok := mapa.GetOk(key); ok {
		// 	fmt.Printf("%s - #%d, value: %v\n", key, mapa.hash(key), val)
		// }
		fmt.Printf("%s - #%d, value: %v\n", key, mapa.hash(key), mapa.Get(key))
	}
	fmt.Println()

	mapa.Set("Dagestan", []int{12, 22, 100})

	for _, key := range keys {
		// if val, ok := mapa.GetOk(key); ok {
		// 	fmt.Printf("%s - #%d, value: %v\n", key, mapa.hash(key), val)
		// }
		fmt.Printf("%s - #%d, value: %v\n", key, mapa.hash(key), mapa.Get(key))
	}

	fmt.Println()
	fmt.Println(mapa.store)

	t.Error()
}

func TestN(t *testing.T) {
	l := ll.New(1, 2, 3, 4, 5)
	val := ll.NodeValue(32)
	l.Tour(func(n *ll.Node, i int) bool {
		fmt.Println(n.Get(), i, l.Length())
		return n.Equal(val)

	})
	t.Error()
}

func TestGett(t *testing.T) {
	mapa := New()
	keys := []string{"EH", "FG", "DI", "HE", "GF", "ID", "Dagestan"}
	for i, key := range keys {
		mapa.Set(key, 100*(i+1))
	}
	fmt.Println(mapa.get("EH"), mapa.count)
	fmt.Println(mapa.get("FG"), mapa.count)
	fmt.Println(mapa.get("DI"), mapa.count)
	fmt.Println(mapa.get("HE"), mapa.count)
	fmt.Println(mapa.get("GF"), mapa.count)
	fmt.Println(mapa.get("ID"), mapa.count)
	fmt.Println(mapa.get("Dagestan"), mapa.count)

	fmt.Println(mapa.store)

	t.Error()
}
