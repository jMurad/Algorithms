package primenumbers

import (
	"fmt"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	arr := primeNumber(307)
	fmt.Print("{")
	for i, v := range arr {
		if i != len(arr)-1 {
			fmt.Printf("%d, ", v)
		} else {
			fmt.Printf("%d}\n", v)
		}
	}
	t.Error()
}
