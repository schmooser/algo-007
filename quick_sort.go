package main

import (
  "fmt"
  "flag"
  //"math/rand"
  //"time"
  "github.com/schmooser/algo-007/sort"
  "github.com/schmooser/algo-007/util"
)


func main() {
  filename := flag.String("filename", "input.txt", "File with integers")
  flag.Parse()

  a, _ := util.ReadLines(*filename)
  comparisons := sort.QuickSort(a)
  fmt.Printf("%d comparisons\n", comparisons)
}
