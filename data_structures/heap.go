package heap

type Heap struct {
	vals []int //values stored in the heap
}

/*
// isEven returns True if given int is even.
func isEven(i int) bool {
	return i%2 == 0
}

// isOdd returns True if given int is odd.
func isOdd(i int) bool {
	return !isEven(i)
}
*/

// parentPos returns position of the parent node.
func (h *Heap) parentPos(pos int) int {
	return pos / 2
}

// parentVal returns value of the parent node for given.
func (h *Heap) parentVal(pos int) int {
	return h.vals[h.parentPos(pos)]
}

// clildrenPos returns children positions
func (h *Heap) childrenPos(pos int) (l, r int) {
	l, r = -1, -1
	if len(h.vals) > 2*pos+1 {
		r = 2*pos + 1
	}
	if len(h.vals) > 2*pos {
		l = 2 * pos
	}
	return
}

// clildrenVal returns children values.
func (h *Heap) childrenVal(pos int) (l, r int) {
	lpos, rpos := h.childrenPos(pos)
	if lpos != -1 {
		l = h.vals[lpos]
	}
	if rpos != -1 {
		r = h.vals[rpos]
	}
	return
}

// Insert inserts value into the heap.
func (h *Heap) Insert(val int) {
	h.vals = append(h.vals, val)
	pos := len(h.vals) - 1 // position of inserted item
	for h.parentVal(pos) > h.vals[pos] {
		h.vals[pos], h.vals[h.parentPos(pos)] = h.vals[h.parentPos(pos)], h.vals[pos]
		pos = h.parentPos(pos)
	}
}
