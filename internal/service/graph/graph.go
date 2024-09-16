package graph

import "fmt"

type Graph struct {
	nodes map[any][]Neighbor
}

type Neighbor struct {
	to    any
	coast any
}

func New() *Graph {
	g := &Graph{
		nodes: make(map[any][]Neighbor),
	}

	return g
}

func (g *Graph) AddNode(node any, nei ...Neighbor) {
	g.nodes[node] = make([]Neighbor, len(nei))
	copy(g.nodes[node], nei)
}

func (g *Graph) Nodes() {
	i := 0
	for n := range g.nodes {
		i++
		fmt.Printf("%d-[%v]\n", i, n)
	}
}

func (g *Graph) Neighbors() {
	for nod, nei := range g.nodes {
		fmt.Printf("-- %v\n", nod)
		for _, n := range nei {

			fmt.Printf("-- -- %v: %v\n", n.to, n.coast)
		}
		fmt.Println("")
	}
}
