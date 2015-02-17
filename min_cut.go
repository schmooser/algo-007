package main

import (
    "fmt"
    "github.com/schmooser/algo-007/graph"
    "github.com/schmooser/algo-007/util"
    //"math/rand"
)


func main() {
    filename := "data/kargerMinCut.txt"

    fmt.Printf("Importing %s\n", filename)
    rows, _ := util.ReadRows(filename, " ")

    /*
    var rows [][]int;
    rows = append(rows, []int{1, 2, 3})
    rows = append(rows, []int{2, 1, 4, 5})
    rows = append(rows, []int{3, 1, 4})
    rows = append(rows, []int{4, 3, 2, 6})
    rows = append(rows, []int{5, 2, 7, 6})
    rows = append(rows, []int{6, 5, 4, 8})
    rows = append(rows, []int{7, 5, 8})
    rows = append(rows, []int{8, 6, 5})
    */

    fmt.Println("MakeGraph")
    vortices, edges := graph.MakeGraph(rows)

    //fmt.Println(vortices)
    //fmt.Println(edges)

    fmt.Println("KargerContraction")

    //rand.Seed(int64(164))
    //mc_vortices, mc_edges := graph.KargerContraction(vortices, edges)

    m := len(edges)
    var vv graph.Vortices
    var ee graph.Edges

    for i := 1; i <= len(vortices)*2; i++ {
        //fmt.Printf("Iteration %d\n", i)

        //rand.Seed(int64(i))

        vv = make([]graph.Vortex, len(vortices))
        copy(vv, vortices)

        ee = make([]graph.Edge, len(edges))
        copy(ee, edges)

        //mc_vortices, mc_edges := graph.KargerContraction(vv, ee)
        _, mc_edges := graph.KargerContraction(vv, ee)

        //fmt.Printf("MinCut on iteration %d = %d\n", i, len(mc_edges))

        if len(mc_edges) < m {
            m = len(mc_edges)
        }

        //fmt.Println(mc_vortices)
        //fmt.Println(mc_edges)

    }

    //fmt.Println(mc_vortices)
    //fmt.Println(mc_edges)

    fmt.Printf("MinCut: %d\n", m)
    fmt.Println("Done!")
}

