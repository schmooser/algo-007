/*
Inversions calculation implementation.
Author: Pavel Popov <pavelpopov@outlook.com>
*/

package inversions

import (
	//"fmt"
	"github.com/schmooser/algo-007/sort"
)

func CountInv(numbers []int) int {
	/*
	   Calculate number of inversions (n[i] > n[j] for i < j) for given array.
	*/
	n := len(numbers)
	if n == 1 {
		return 0
	}

	if n == 2 {
		a, b := numbers[0], numbers[1]
		if a < b {
			return 0
		} else {
			return 1
		}
	}

	left := numbers[:n/2]
	right := numbers[n/2:]

	left_inv := CountInv(left)
	right_inv := CountInv(right)

	split_inv := countSplitInv(left, right)

	return left_inv + right_inv + split_inv
}

func countSplitInv(left []int, right []int) int {
	l := sort.MergeSort(left)
	r := sort.MergeSort(right)

	len_l, len_r := len(l), len(r)

	var i, j, n int = 0, 0, 0

	//fmt.Println("countSplitInv")
	//fmt.Println(left)
	//fmt.Println(right)

	for i+j < len_l+len_r {

		//fmt.Printf("i=%d, j=%d\n", i, j)

		if l[i] < r[j] {
			i++
		} else {
			n += len_l - i
			//fmt.Printf("Found inversion. Total inversions: %d\n", n)
			j++
		}

		if i >= len_l || j >= len_r {
			return n
		}
	}

	return n
}
