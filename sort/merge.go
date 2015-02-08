/*
Merge sort implementation.
Author: Pavel Popov <pavelpopov@outlook.com>
*/

package sort

//import "fmt"


func MergeSort(numbers []int64) []int64 {
  /*
    Sort given array using merge sort algorithm.
  */
  n := len(numbers)

  if n==1 {
    return numbers
  }

  if n==2 {
    var a, b = sortPair(numbers[0], numbers[1])
    return []int64{a, b}
  }

  foo := MergeSort(numbers[:n/2])
  bar := MergeSort(numbers[n/2:])

  return merge(foo, bar)
}


func sortPair(n int64, m int64) (int64, int64)  {
  if n <= m {
    return n, m
  } else {
    return m, n
  }
}


func merge(num1 []int64, num2 []int64) []int64 {

  //fmt.Println("Merge")
  //fmt.Println(num1)
  //fmt.Println(num2)

  var i, j, l1, l2 = 0, 0, len(num1), len(num2)
  numbers := make([]int64, l1+l2)

  for k := 0; k < l1+l2; k++ {

    //fmt.Printf("i=%d, j=%d\n", i, j)

    if i == l1 {
      numbers[k] = num2[j]
      j++
      continue
    }

    if j == l2 {
      numbers[k] = num1[i]
      i++
      continue
    }

    if num1[i] <= num2[j] {
      numbers[k] = num1[i]
      i++
    } else {
      numbers[k] = num2[j]
      j++
    }

  }

  return numbers
}
