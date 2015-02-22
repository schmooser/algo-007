/*
Compute Stronly Connected Components in the Directed Graph
*/
package main

import (
	"fmt"
	"github.com/schmooser/algo-007/graph"
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

func main() {

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

	printGraph(G)

}
