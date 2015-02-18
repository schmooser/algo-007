/*
Implementation of Karger Min Cut algorithm for finding min cut in unordered
graph.
Author: Pavel Popov <pavelpopov@outlook.com>
*/

package graph

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	//rand.Seed(13)
}

func MakeGraph(rows [][]int) (vortices Vortices, edges Edges) {
	/*
	   input rows contain data for a graph:
	       1 2 4 5
	       2 1
	       3 4 5
	       4 1 3
	       5 1 3
	   First column is a Vertex number.
	   Other is connected Vertexes.
	*/

	/*
	   var rows [][]int;
	   rows = append(rows, []int{1, 2, 4, 5})
	   rows = append(rows, []int{2, 1})
	   rows = append(rows, []int{3, 4, 5})
	   rows = append(rows, []int{4, 1, 3})
	   rows = append(rows, []int{5, 1, 3})
	*/

	//rows, _ := util.ReadGraph(filename)

	for _, row := range rows {
		vortex := Vortex{row[0]}
		vortices = append(vortices, vortex)
		for _, x := range row {
			if x > vortex.X {
				edges = append(edges, Edge{vortex.X, x})
			}
		}
	}

	return
}

func (vortices Vortices) RemoveVortex(k int) Vortices {
	l := len(vortices) - 1
	_, vortices[k] = vortices[k], vortices[l]
	return vortices[:l]
}

func KargerContraction(vortices Vortices, edges []Edge) (Vortices, []Edge) {

	if len(vortices) == 2 {
		return vortices, edges
	}

	var print bool = false

	if len(vortices) < 0 {
		print = true
	}

	if print {
		fmt.Println("KargerContraction invoked")
		fmt.Printf("Len(vortices) = %d\n", len(vortices))

		fmt.Println("Vortices")
		fmt.Println(vortices)

		fmt.Println("Edges")
		fmt.Println(edges)
	}

	//fmt.Println(edges)
	k := rand.Intn(len(edges))

	// swap chosen edge with the last in slice, x - contracted edge
	var x Edge

	edges[k], x = edges[len(edges)-1], edges[k]
	edges = edges[:len(edges)-1]

	if print {
		fmt.Println("Selected edge")
		fmt.Println(x)
	}

	for i, v := range vortices {
		if v.X == x.To { // remove vortex
			vortices = vortices.RemoveVortex(i)
			break
		}
	}

	//fmt.Println("Before merge")

	// merge into super vertex
	for i := range edges {
		if edges[i].From == x.To {
			edges[i].From = x.From
		}

		if edges[i].To == x.To {
			edges[i].To = x.From
		}
	}

	if print {
		fmt.Println("Edges after contraction")
		fmt.Println(edges)
	}

	//fmt.Println("After merge")

	// Remove self-loops
	cross_edges := make(Edges, 0)

	for _, edge := range edges {
		if edge.From != edge.To {
			cross_edges = append(cross_edges, edge)
		}
	}

	if print {
		fmt.Println("Crossing edges")
		fmt.Println(cross_edges)
	}

	return KargerContraction(vortices, cross_edges)
}
