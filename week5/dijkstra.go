/*
Dijkstra short path searh algorithm implementation
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	V []*Vertex
	A []int     //computed shortest path distances
	X []*Vertex // vertices processed so far
}

type Vertex struct {
	E     []*Edge
	index int  // position in array of vertices
	inX   bool // true if Vertex is processed and placed into X
}

type Edge struct {
	l   int     //length
	end *Vertex // end point
}

// graphFromFile reads file into a graph structure. It takes filename to read
// and returns pointer to a graph.
func graphFromFile(path string) *Graph {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	g := new(Graph)
	g.V = make([]*Vertex, 200) // todo: remove hardcode
	//g.V = make([]*Vertex, 4) // todo: remove hardcode

	for i := 0; i < len(g.V); i++ {
		g.V[i] = new(Vertex)
		g.V[i].index = i
	}

	fmt.Println("Vertices were added to the graph.")
	//fmt.Println(g.V)

	scanner := bufio.NewScanner(file)

	var k int = 0

	for scanner.Scan() {
		line := scanner.Text()
		for i, edge := range strings.Split(line, "\t") {
			if i > 0 && len(edge) > 0 { //only edges
				s := strings.Split(edge, ",")
				//fmt.Println(s)
				v, _ := strconv.ParseInt(s[0], 0, 0) // vortex index
				l, _ := strconv.ParseInt(s[1], 0, 0) // length
				e := new(Edge)
				e.l, e.end = int(l), g.V[int(v)-1] // 1-base indexed array in the data
				g.V[k].E = append(g.V[k].E, e)
			}
		}
		k++
	}

	return g
}

// vertexByIndex returns Vertex from the graph by corresponding index.
func (g *Graph) vertexByIndex(index int) *Vertex {
	return g.V[index]
}

// DijkstraShortestPaths returns shortest pathes from given source vortex to
// other vertices. If there is no path then it return 1000000 at the
// corresponding index.
func (g *Graph) DijkstraShortestPaths(s *Vertex) []int {
	var gc int    // Dijkstra greedy criteria
	var w *Vertex // Picked vertex that minimizes gc

	// initialization
	g.X = append(g.X, s)
	s.inX = true
	g.A = make([]int, len(g.V))
	for i, _ := range g.A {
		if g.V[i] == s {
			g.A[i] = 0
		} else {
			g.A[i] = 1000000
		}
	}

	fmt.Println("Dijkstra Search is initialized!")
	fmt.Println(g)

	for len(g.X) < len(g.V) {

		gc = 1000000
		for _, v := range g.X {
			//fmt.Println(v)
			//fmt.Println(g.A)
			for _, e := range v.E {
				if !e.end.inX && g.A[v.index]+e.l < gc {
					gc = g.A[v.index] + e.l
					w = e.end
				}
			}
		}

		if w != nil { // path is found
			g.X = append(g.X, w)
			w.inX = true
			for i, _ := range g.A {
				if g.V[i] == w {
					g.A[i] = gc
				}
			}
		}

		fmt.Printf("Processed %d vertices so far\n", len(g.X))
		//fmt.Println(g.X)
	}
	return g.A
}

// printArr prints value from array at given indices. It expects that given
// indices are 1-based, not 0-based.
func printArr(paths []int, indices ...int) {
	for _, i := range indices {
		fmt.Printf("%d,", paths[i-1]) // 1-based index
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("Hello, Dijkstra!")
	G := graphFromFile("../data/dijkstraData.txt")
	//G := graphFromFile("../data/dijkstraData-test.txt")
	fmt.Println("Graph is built!")
	fmt.Println(G)
	s := G.vertexByIndex(0)
	paths := G.DijkstraShortestPaths(s)
	fmt.Println(G)
	fmt.Println(paths)
	fmt.Println("Answer is:")
	printArr(paths, 7, 37, 59, 82, 99, 115, 133, 165, 188, 197)
}
