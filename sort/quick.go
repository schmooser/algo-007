/*
Merge sort implementation.
Author: Pavel Popov <pavelpopov@outlook.com>
*/

package sort


import (
  //"fmt"
  "math/rand"
  "time"
)

var nums []int = []int{4, 8, 6, 1, 2, 5, 9, 7, 3}

func init() {
  rand.Seed(time.Now().UTC().UnixNano())
}

func QuickSort(numbers []int) (n int) {
    /* sort array, return number of comparisons */

  //fmt.Println("Input:")
  //fmt.Println(numbers)
  //fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

  if len(numbers) <= 1 {
    //fmt.Println("It's time to return")
    return 0
  }

  n = len(numbers)-1

  pos := choosePivot(numbers)
  //fmt.Printf("Chosen %d (%d) as a pivot\n", pos, numbers[pos])

  swap(numbers, 0, pos)
  pos = 0

  var i, j int = 0, 0

  for j < len(numbers) {

    if j == pos {
      i++
      j++
      continue
    }

    if j != pos && numbers[j] < numbers[pos] {
      swap(numbers, i, j)
      i++
    }

    j++
  }
  swap(numbers, pos, i-1)

  //fmt.Println("After partitioning:")
  //fmt.Println(numbers)

  n += QuickSort(numbers[:i-1])

  //fmt.Println("Intermediate result:")
  //fmt.Println(numbers)

  if len(numbers) > i {
    n += QuickSort(numbers[i:])
  }

  //fmt.Println("Result:")
  //fmt.Println(numbers)

  return n
}

func choosePivot(numbers []int) int {
  /* Choose pivot from given array */
  //return rand.Intn(len(numbers))
  //return 0 // Question 1 - 162085
  //return len(numbers)-1 // Question 2 - 164123

  // median // Question 3 - 138382
  l := len(numbers)

  var m int

  if l % 2 == 0 {
    m = (l-1)/2
  } else {
    m = l/2
  }

  x := numbers[0]
  y := numbers[m]
  z := numbers[l-1]

  //fmt.Println(numbers)
  //fmt.Println(x, y, z)

  if x >= y && y >= z || x <= y && y <= z {
    return m
  }

  if y >= x && x >= z || y <= x && x <= z {
    return 0
  }

  if x >= z && z >= y || x <= z && z <= y {
    return l-1
  }

  return -1

}

func swap(numbers []int, i int, j int) {
  //fmt.Printf("Swapping %d's and %d's\n", i, j)
  numbers[i], numbers[j] = numbers[j], numbers[i]
}

