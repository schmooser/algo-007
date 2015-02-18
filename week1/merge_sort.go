package main

import (
	"flag"
	"fmt"
	"github.com/schmooser/algo-007/sort"
	"math/rand"
	"time"
)

func main() {
	n := flag.Int("n", 42, "Number of entries in array")
	output := flag.Int("output", 1, "Turn on output")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	var a []int

	for i := int(0); i < *n; i++ {
		a = append(a, rand.Intn(*n))
	}

	if *output == 1 {
		fmt.Printf("Array of length %d.\n", *n)
		fmt.Println(a)
		fmt.Println(sort.MergeSort(a))
	}
	fmt.Println("Done!")
}
