package hashtable

import (
	"fmt"
	"reflect"
	"testing"

	ll "abusafia.com/algorithms/internal/service/list"
)

func TestNew(t *testing.T) {
	ht := New()
	if len(ht.store) != 100 {
		t.Errorf("The \"store\" length is not equal to 100")
	}
	if ht.count != 0 {
		t.Errorf("The \"count\" is not equal to 0")
	}
}

func TestSet(t *testing.T) {
	mapa := New()
	key := "A"
	value := 10
	mapa.Set(key, value)
	index := mapa.hash(key)
	bucket := mapa.store[index]
	found := false
	keyval := []interface{}{key, value}
	count := 0
	bucket.Tour(func(n *ll.Node, i int) bool {
		val := (n.Get()).([]interface{})
		if val[0].(string) == keyval[0].(string) {
			if reflect.ValueOf(val[1]).Equal(reflect.ValueOf(keyval[0])) {
				found = true
			}
			count++
		}
		return false
	})
	if !found {
		t.Errorf("The Set don't added value: %v", value)
	}

	if count > 1 {
		t.Errorf("Keys (%v) is too many", key)
	}
}

func TestSetGet(t *testing.T) {
	mapa := New()
	keys := []string{"EH", "FG", "DI", "HE", "GF", "ID", "Da", "ge", "st",
		"an", "sw", "er", "w", "ds", "fkb", "wkn", "fwn", "erj", "fnv", "wen",
		"fln", "we", "rov", "oer", "ivn"}
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
	for _, v := range mapa.store {
		if v.Length() > 1 {
			fmt.Printf("%v \n\n", v.PrintAll())
		}
	}

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
	keys := []string{"EH", "FG", "DI", "HE", "GF", "ID", "D", "a", "g", "e", "s", "t", "a", "n", "s", "dd", "df", "ssd", "sddds", "saq", "qwe"}
	for _, key := range keys {
		fmt.Printf("HASH[\"%s\"]=%d, len store: %d  ===\n", key, mapa.hash(key), len(mapa.store))
		mapa.Set(key, 0)

		for i, v := range mapa.store {
			if v.Length() != 0 {
				fmt.Printf("\tstore[%d]: %v\n\n", i, v.PrintAll())
			}
		}
		// fmt.Println()
		// fmt.Printf("store: %v\n\n", mapa.store)
	}
	fmt.Printf("length mapa: %v\nlength store: %v\n", mapa.count, len(mapa.store))

	t.Error()
}
