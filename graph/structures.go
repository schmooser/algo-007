/*
Some structures suitable for graphs.
Author: Pavel Popov <pavelpopov@outlook.com>
*/
package graph

type Edge struct {
	From, To int
}

type Vortex struct {
	X int
}

type Vortices []Vortex

type Edges []Edge
