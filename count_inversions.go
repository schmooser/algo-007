package main

import (
  "fmt"
  "flag"
  "github.com/schmooser/algo-007/inversions"
  "github.com/schmooser/algo-007/util"
)

import (
  "os"
  "bufio"
  "io"
  "strconv"
)


func main() {
  filename := flag.String("filename", "input.txt", "File with integers")
  output := flag.Int("output", 0, "Turn on output")
  flag.Parse()

  a, _ := util.ReadInts(*filename)

  var s int64 = 0

  for i := 0; i < len(a); i++ {
    s += a[i]
  }

  if *output == 1 {
    fmt.Printf("Array of length %d.\n", len(a))
    fmt.Printf("Sum of values is %d.\n", s)
    fmt.Printf("2nd elem is %d.\n", a[1])
    fmt.Println(inversions.CountInv(a))
  }
  fmt.Println("Done!")
}


