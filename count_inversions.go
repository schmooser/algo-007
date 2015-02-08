package main

import (
  "fmt"
  "flag"
  "github.com/schmooser/algo-007/inversions"
)

import (
  "os"
  "bufio"
  "io"
  "strconv"
)

func readLines(path string) (ints []int64, err error) {
  var file *os.File
  if file, err = os.Open(path); err != nil {
    return
  }
  defer file.Close()

  var line string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line = scanner.Text()
    i, _ := strconv.ParseInt(line, 0, 64)
    ints = append(ints, i)
  }

  if err == io.EOF {
    err = nil
  }

  return
}

func main() {
  filename := flag.String("filename", "input.txt", "File with integers")
  output := flag.Int("output", 0, "Turn on output")
  flag.Parse()

  a, _ := readLines(*filename)

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


