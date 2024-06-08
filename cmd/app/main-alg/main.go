package main

import (
	"log"

	recursion "abusafia.com/algorithms/internal/app/recursion"
	"abusafia.com/algorithms/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Printf("%v\n", err)
	}

	data := []int{4, 6, 2, 9, 4, 5, 6, 2, 49, 5}

	// bs := binarysearch.New(data)
	// ss := selectionsort.New(data)
	// rc := recursion.New(data, "binsearch")
	// rc := recursion.New(data, "selsort")
	// rc := recursion.New(data, "qsort")
	rc := recursion.New(data, "qsortlist")

	err = a.Run(rc)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
