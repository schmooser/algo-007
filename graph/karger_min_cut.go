/*
Implementation of Karger Min Cut algorithm for finding min cut in unordered
graph.
Author: Pavel Popov <pavelpopov@outlook.com>
*/

package graph

import "fmt"

type Edge struct {
  From, To int
}


func MakeGraph(filename string) (edges []Edge) {
    /*
    input file contains data for a graph:
    1 2 4 5
    2 1
    3 4 5
    4 1 3
    5 1 3
    First column is a Vertex number
    Other is connected Vertexes
    */

    fmt.Printf("Importing %s\n", filename)

    var rows [][]int;
    rows = append(rows, []int{1, 2, 4, 5})
    rows = append(rows, []int{2, 1})
    rows = append(rows, []int{3, 4, 5})
    rows = append(rows, []int{4, 1, 3})
    rows = append(rows, []int{5, 1, 3})

    for i := range rows {
        fmt.Println(rows[i])
        row := rows[i]
        edge := row[0]
        for j := range row {
            if row[j] > edge {
                edges = append(edges, Edge{edge, row[j]})
            }
        }
    }

    for i := range edges {
        fmt.Println(edges[i])
    }

    return nil
}
