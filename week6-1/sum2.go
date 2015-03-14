/*
2-SUM implementation
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N int64 = 999763 // 9973 // number of buckets
const M int = 1000000  // number of numbers

type HashTable struct {
	Buckets []*Bucket
}

type Bucket struct {
	Values []int64
}

func abs(input int64) int64 {
	if input < 0 {
		return -1 * input
	} else {
		return input
	}
}

// hashFunction takes an input and returns bucket number
func hashFunction(input int64) int64 {
	return abs(input % N)
}

func (h *HashTable) Lookup(val int64) bool {
	bucket := hashFunction(val)
	for i := range h.Buckets[bucket].Values {
		if h.Buckets[bucket].Values[i] == val {
			return true
		}
	}
	return false
}

// HashTableFromFile creates HashTable and loads there the content of provided
// file.
func HashTableFromFile(path string) (h *HashTable, nums []int64) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	h = new(HashTable)
	h.Buckets = make([]*Bucket, N)

	nums = make([]int64, M)

	for i := range h.Buckets {
		b := new(Bucket)
		b.Values = []int64{}
		h.Buckets[i] = b
	}

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.ParseInt(line, 0, 0)
		idx := hashFunction(val)
		h.Buckets[idx].Values = append(h.Buckets[idx].Values, val)
		nums[i] = val
		i++
	}

	fmt.Println("Hash Table created")

	return
}

func main() {
	fmt.Println("Hello, 2-SUM")
	filename := "../data/algo1-programming_prob-2sum.txt" // "../data/2sum_test.txt"
	h, nums := HashTableFromFile(filename)
	//fmt.Println(h)
	res := 0
	for t := int64(-10000); t <= 10000; t++ {
		for _, x := range nums {
			if x != t-x && h.Lookup(t-x) {
				res++
				fmt.Printf("%d can be represented as %d+%d\n", t, x, t-x)
				break
			}
		}
	}
	fmt.Println(h.Lookup(82329471587))
	fmt.Println(h.Lookup(82329471588))
	fmt.Println(res)
}
