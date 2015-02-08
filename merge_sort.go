package main

import (
  "fmt"
  "flag"
  "math/rand"
  "time"
  "github.com/schmooser/algo-007/sort"
)


func main() {
  n := flag.Int64("n", 42, "Number of entries in array")
  output := flag.Int("output", 1, "Turn on output")
  flag.Parse()

  rand.Seed(time.Now().UTC().UnixNano())

  var a []int64

  for i := int64(0); i < *n; i++ {
    a = append(a, rand.Int63n(*n))
  }

  if *output == 1 {
    fmt.Printf("Array of length %d.\n", *n)
    fmt.Println(a)
    fmt.Println(sort.MergeSort(a))
  }
  fmt.Println("Done!")
}
