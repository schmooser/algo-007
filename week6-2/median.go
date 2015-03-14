/*
Median implementation
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

type Direction int

const (
	Max Direction = iota
	Min
)

type Heap struct {
	Type Direction
	Vals []int
}

// Cmp defines order of in Heap.
func (h *Heap) Cmp(child, parent int) bool {
	if h.Type == Max {
		return parent >= child
	} else {
		return parent <= child
	}
}

// Insert inserts value in Heap.
func (h *Heap) Insert(val int) {
	h.Vals = append(h.Vals, val)

	i_idx := len(h.Vals) - 1 // item index and value

	p_idx, p_val := h.Parent(i_idx)

	for i_idx != 0 && !h.Cmp(val, p_val) { // Heap condition is violated - bubbling it up
		h.Vals[p_idx], h.Vals[i_idx] = h.Vals[i_idx], h.Vals[p_idx]
		i_idx = p_idx // item index and value
		p_idx, p_val = h.Parent(i_idx)
	}
}

// Parent returns parent's position and value for given element (by index).
func (h *Heap) Parent(i int) (pos, val int) {
	pos = i / 2
	val = h.Vals[pos]
	return
}

// Get gets item from the Heap. If there is no such element then ok = False and
// val = 0.
func (h *Heap) Get(pos int) (val int, ok bool) {
	val = 0
	if pos > len(h.Vals)-1 {
		ok = false
		return
	}
	val = h.Vals[pos]
	ok = true
	return
}

// Next returns next value from the Heap and reorganises Heap in order to it's
// satisfying Heap's property.
func (h *Heap) Next() int {
	x := h.Vals[0]
	h.Vals[0], h.Vals[len(h.Vals)-1] = h.Vals[len(h.Vals)-1], h.Vals[0] // swap root with the last element
	h.Vals = h.Vals[:len(h.Vals)-1]                                     // delete last element (former root)
	var (
		i     int // current item index
		c_idx int // comparing item index
		c_val int // comparing item value
		l_idx int // left-child item index
		l_val int // left-child item value
		r_idx int // right-child item index
		r_val int // right-child item value
		ok    bool
	)
	for {
		l_idx, r_idx = 2*(i+1)-1, 2*(i+1)
		//log.Debug("i=%d, l_idx=%d, r_idx=%d", i, l_idx, r_idx)

		l_val, ok = h.Get(l_idx)
		if !ok { // no left child - Heap property is satisfied - exiting
			//log.Debug("No left child - done!")
			break
		}

		r_val, ok = h.Get(r_idx)
		if !ok || !h.Cmp(l_val, r_val) { // no right child - checking only with left child
			c_idx, c_val = l_idx, l_val
		} else {
			c_idx, c_val = r_idx, l_val
		}

		//log.Debug("i=%d, val=%d", i, h.Vals[i])
		//log.Debug("c_idx=%d, c_val=%d", c_idx, c_val)

		if h.Cmp(c_val, h.Vals[i]) { // if Heap property is satisfied
			break
		} else { // bubbling down
			h.Vals[i], h.Vals[c_idx] = h.Vals[c_idx], h.Vals[i]
			i = c_idx
		}
	}
	return x
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -1 * a
	}
}

func Median(input []int) int64 {

	h_low := new(Heap) // elements lesser than current
	h_low.Type = Max

	h_high := new(Heap) // elements greater than current
	h_high.Type = Min

	var i int
	var median int64

	for _, n := range input {
		i++

		//fmt.Printf("\nStep %d, inserted %d\n", i, n)

		if i == 1 { // first step
			h_low.Insert(n)
		} else {
			if i == 2 { // second step
				h_high.Insert(n)
				if h_high.Vals[0] < h_low.Vals[0] {
					h_high.Vals[0], h_low.Vals[0] = h_low.Vals[0], h_high.Vals[0]
				}
			} else {
				if n <= h_low.Vals[0] {
					h_low.Insert(n)
				} else {
					h_high.Insert(n)
				}
			}
		}

		if i > 2 && len(h_low.Vals) < len(h_high.Vals) {
			h_low.Insert(h_high.Next())
		}

		if i > 2 && len(h_low.Vals) > len(h_high.Vals)+1 {
			h_high.Insert(h_low.Next())
		}

		//fmt.Printf("Median=%d\n", h_low.Vals[0])
		median += int64(h_low.Vals[0])

		//fmt.Println(h_low)
		//fmt.Println(h_high)

	}

	return median
}

func main3() {
	m := Median([]int{3, 2, 1})
	fmt.Printf("Running median = %d\n", m)
}

func main() {
	fmt.Println("Running median")
	//path := "../data/Median_test.txt"
	path := "../data/Median.txt"
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	data := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		n, _ := strconv.ParseInt(line, 0, 0)
		data = append(data, int(n))

	}

	fmt.Printf("Running median = %d\n", Median(data))

	//fmt.Println(h_low)
	//fmt.Println(h_high)
}
