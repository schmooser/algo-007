package main

import (
  "fmt"
  "github.com/schmooser/algo-007/sort"
)

func main() {
  fmt.Printf("Hello, world.\n")
  a := []uint64{1, 3, 5, 2351235123, 12341234, 12341, 2, 4}
  fmt.Println(a)
  fmt.Println(sort.MergeSort(a))
}
