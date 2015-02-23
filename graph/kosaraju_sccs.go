/*
Kosaraju algorithm for computing Strongly Connected Components in directed graph.
Author: Pavel Popov <pavelpopov@outlook.com>
*/
package graph

import (
	//"errors"
	"fmt"
)

type Direction int

const (
	Forward Direction = iota
	Backward
)

// Graph is the collection of Nodes. Only directed graph is implemented.
type Graph struct {
	Nodes []*Node
	fv    []int // f(v) for Kosaraju
	t     int   // counter
	s     int   // current leader
}

// Node is the vortex in the Graph
type Node struct {
	arcsOut  []arc // output arcs
	arcsIn   []arc // input arcs
	index    int
	leader   int  // leader index value
	explored bool // flag if node is explored
}

// Pointer to specified Node
type arc struct {
	weight int
	end    *Node
}

// MakeNode makes the new node, adds it to the graph. Returns pointer to created
// node.
func (g *Graph) MakeNode() *Node {
	newNode := &Node{index: len(g.Nodes)}
	g.Nodes = append(g.Nodes, newNode)
	return newNode
}

// MakeArc calls MakeArcWeight with a weight of 0 and returns an error if either
// of the nodes do not belong in the graph. Calling MakeArc multiple times on
// the same nodes will not create multiple edges.
func (g *Graph) MakeArc(from, to *Node) error {
	return g.MakeArcWeight(from, to, 0)
}

// MakeArcWeight add arc between nodes. If such arc already exists it will
// update weight of the arc.
func (g *Graph) MakeArcWeight(from, to *Node, weight int) error {
	/*
		if from.index >= len(g.Nodes) || g.Nodes[from.index] != from {
			return errors.New("Node From does not belong to this graph")
		}
		if to.index >= len(g.Nodes) || g.Nodes[to.index] != to {
			return errors.New("Node To does not belong to this graph")
		}
	*/

	// check if such arc already exists
	for i, arc := range from.arcsOut {
		if arc.end == to {
			from.arcsOut[i].weight = weight
			return nil
		}
	}

	// doesn't exists, create new arc
	from.arcsOut = append(from.arcsOut, arc{weight, to})
	to.arcsIn = append(to.arcsIn, arc{0, from})

	return nil
}

// DFS_Loop
func (g *Graph) DFS_Loop(direction Direction) {
	fmt.Printf("DFS_Loop with direction %d\n", direction)
	g.t = 0                    // number of nodes processed so far
	g.s = -1                   // current source node
	if direction == Backward { // first call
		for i := len(g.Nodes) - 1; i >= 0; i-- {
			//fmt.Printf("Processing %dth node\n", i)
			if !g.Nodes[i].explored {
				g.DFS(g.Nodes[i], direction)
			}
		}
	} else { // second call
		//fmt.Println("Processing times:")
		//fmt.Println(g.fv)
		for i := len(g.Nodes) - 1; i >= 0; i-- {
			//fmt.Printf("Processing %dth node\n", g.fv[i])
			if !g.Nodes[g.fv[i]].explored {
				g.s = g.fv[i]
				g.DFS(g.Nodes[g.fv[i]], direction)
			}
		}
	}
}

// DFS implements Depth-First search of the graph starting from given Node
func (g *Graph) DFS(node *Node, direction Direction) {
	fmt.Printf("DFS on node %d with direction %d\n", node.index, direction)
	node.explored = true

	if direction == Forward {
		node.leader = g.s
	}

	var arcs []arc
	if direction == Forward {
		arcs = node.arcsOut
	} else {
		arcs = node.arcsIn
	}

	for _, arc := range arcs {
		if !arc.end.explored {
			g.DFS(arc.end, direction)
		}
	}

	g.t++
	if direction == Backward {
		//fmt.Printf("Set f(%d) to %d\n", node.index, g.t-1)
		g.fv[g.t-1] = node.index
	}
}

// ComputeSCCS computes Strongly Connected Components in directed graph.
// It takes directed graph as input and returns slice with directed graphs
// which are strongly connected components from original graph.
// From computing SCCS Kosaraju algorithm is used.
func (g *Graph) StronglyConnectedComponents() {
	g.fv = make([]int, len(g.Nodes))
	g.DFS_Loop(Backward)

	//fmt.Println("fv after run on the Reversed graph")
	//fmt.Println(g.fv)

	//mark all nodes unexplored
	for _, node := range g.Nodes {
		node.explored = false
	}

	g.DFS_Loop(Forward)
}

// SizeSCCS computes sizes of Strongly connected components of the graph.
func (g *Graph) SizeSCCS() []int {
	sizes := make(map[int]int)
	for _, node := range g.Nodes {
		if _, ok := sizes[node.leader]; !ok {
			sizes[node.leader] = 1
		} else {
			sizes[node.leader]++
		}
	}
	v := make([]int, 0, len(sizes))
	for _, value := range sizes {
		v = append(v, value)
	}
	return v
}
