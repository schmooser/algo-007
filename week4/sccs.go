/*
Compute Stronly Connected Components in the Directed Graph
*/
package main

import (
	"fmt"
	"github.com/schmooser/algo-007/graph"
	"github.com/schmooser/algo-007/sort"
)

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var err error
var errs []error

func printGraph(g graph.Graph) {
	fmt.Println("Graph")
	fmt.Println(g)
	fmt.Println("Nodes")
	for _, node := range g.Nodes {
		fmt.Println(node)
	}
}

func greatest(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Read file and create graph from it
func main() {
	path := "../data/SCC.txt"
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)
	g := graph.Graph{}

	for scanner.Scan() {
		line = scanner.Text()
		s := strings.Split(line, " ")
		from64, _ := strconv.ParseInt(s[0], 0, 0)
		to64, _ := strconv.ParseInt(s[1], 0, 0)
		from := int(from64)
		to := int(to64)
		idx := greatest(from, to)
		l := len(g.Nodes)
		//fmt.Printf("idx=%d len=%d\n", idx, l)
		if idx > l {
			//fmt.Printf("Adding %d nodes\n", idx-l)
			for i := 0; i < idx-l; i++ {
				//fmt.Printf("Loop Number of Nodes = %d\n", len(g.Nodes))
				g.MakeNode()
			}
		}
		//fmt.Printf("Number of Nodes = %d\n", len(g.Nodes))
		//fmt.Printf("Making Arc between %d and %d nodes\n", from-1, to-1)
		g.MakeArc(g.Nodes[from-1], g.Nodes[to-1])
	}

	if err == io.EOF {
		err = nil
	}

	g.StronglyConnectedComponents()

	//printGraph(g)

	sccs := sort.QuickSort(g.SizeSCCS())
	//fmt.Println(sccs)
	fmt.Println(sccs[len(sccs)-6:])

}

func main2() {

	errs = make([]error, 0)

	fmt.Println("Hello!")
	G := graph.Graph{}

	a := G.MakeNode()
	b := G.MakeNode()
	c := G.MakeNode()

	d := G.MakeNode()
	e := G.MakeNode()
	f := G.MakeNode()
	g := G.MakeNode()

	h := G.MakeNode()

	x := G.MakeNode()
	y := G.MakeNode()
	z := G.MakeNode()

	//printGraph(G)

	errs = append(errs, G.MakeArc(a, b))
	errs = append(errs, G.MakeArc(b, c))
	errs = append(errs, G.MakeArc(c, a))

	errs = append(errs, G.MakeArc(c, d))
	errs = append(errs, G.MakeArc(c, g))

	errs = append(errs, G.MakeArc(b, h))
	errs = append(errs, G.MakeArc(h, x))
	errs = append(errs, G.MakeArc(h, y))

	errs = append(errs, G.MakeArc(d, x))
	errs = append(errs, G.MakeArc(e, z))

	errs = append(errs, G.MakeArc(g, d))
	errs = append(errs, G.MakeArc(d, e))
	errs = append(errs, G.MakeArc(e, f))
	errs = append(errs, G.MakeArc(f, g))
	errs = append(errs, G.MakeArc(d, f))

	errs = append(errs, G.MakeArc(x, y))
	errs = append(errs, G.MakeArc(y, z))
	errs = append(errs, G.MakeArc(z, x))

	exit := false
	for _, e := range errs {
		if e != nil {
			fmt.Println(e)
			exit = true
		}
	}
	if exit {
		return
	}

	G.StronglyConnectedComponents()

	//printGraph(G)

	sccs := sort.QuickSort(G.SizeSCCS())
	fmt.Println(sccs)
	fmt.Println(sccs[len(sccs)-2:])
}
