package sort

//import "fmt"

func MergeSort(numbers []uint64) []uint64 {
  n := len(numbers)

  if n==1 {
    return []uint64{numbers[0]}
  }

  if n==2 {
    var a, b = sortPair(numbers[0], numbers[1])
    return []uint64{a, b}
  }

  foo := MergeSort(numbers[0:n/2])
  bar := MergeSort(numbers[n/2:])

  return merge(foo, bar)
}


func sortPair(n uint64, m uint64) (uint64, uint64)  {
  if n <= m {
    return n, m
  } else {
    return m, n
  }
}


func merge(num1 []uint64, num2 []uint64) []uint64 {

  //fmt.Println("Merge")
  //fmt.Println(num1)
  //fmt.Println(num2)

  var numbers []uint64

  var i, j = 0, 0
  var l1, l2 = len(num1), len(num2)

  for !(i >= l1 && j >= l2) {

    //fmt.Printf("i=%d, j=%d\n", i, j)

    if i >= l1 {
      numbers = append(numbers, num2[j])
      j++
      continue
    }

    if j >= l2 {
      numbers = append(numbers, num1[i])
      i++
      continue
    }

    if num1[i] <= num2[j] {
      numbers = append(numbers, num1[i])
      i++
    } else {
      numbers = append(numbers, num2[j])
      j++
    }

  }

  return numbers
}
