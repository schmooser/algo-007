package util

/*
  File operations
*/

import (
  "os"
  "bufio"
  "io"
  "strconv"
)


func ReadLines(path string) (ints []int, err error) {
  var file *os.File
  if file, err = os.Open(path); err != nil {
    return
  }
  defer file.Close()

  var line string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line = scanner.Text()
    i, _ := strconv.ParseInt(line, 0, 0)
    ints = append(ints, i)
  }

  if err == io.EOF {
    err = nil
  }

  return
}
