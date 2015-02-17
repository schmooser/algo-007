package util

/*
  File operations
*/

import (
    //"fmt"
    "os"
    "bufio"
    "io"
    "strings"
    "strconv"
)


func ReadInts(path string) (ints []int, err error) {
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
        ints = append(ints, int(i))
    }

    if err == io.EOF {
        err = nil
    }

    return
}

func ReadRows(path string, dlm string) (rows [][]int, err error) {

    var file *os.File
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    var line string
    var row []int

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line = scanner.Text()
        row = make([]int, 0)
        for _, elem := range strings.Split(line, dlm) {
            num, err1 := strconv.ParseInt(elem, 0, 0)
            if err1 == nil {
                row = append(row, int(num))
            }
        }
        rows = append(rows, row)
    }

    if err == io.EOF {
        err = nil
    }

    return
}
