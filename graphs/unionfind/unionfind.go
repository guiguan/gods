package unionfind

import (
	"fmt"
	"strings"
)

type Graph struct {
	count   int   // Number of connected components
	parents []int // Stores parent of each tree node at i. This stores the forest
	sizes   []int // Stores size/weight of each tree rooted at i
}

// New instantiates a new graph with n nodes for union-find algorithm
func New(n int) *Graph {
	g := &Graph{
		count:   n,
		parents: make([]int, n),
		sizes:   make([]int, n),
	}

	g.Clear()

	return g
}

func (g *Graph) Union(p, q int) {
	rootP, rootQ := g.findRoot(p), g.findRoot(q)

	if rootP == rootQ {
		return
	}

	// adding smaller tree as a child of the larger one to get a more balanced tree
	if g.sizes[rootP] < g.sizes[rootQ] {
		g.parents[rootP] = rootQ
		g.sizes[rootQ] += g.sizes[rootP]
	} else {
		g.parents[rootQ] = rootP
		g.sizes[rootP] += g.sizes[rootQ]
	}

	g.count--
}

func (g *Graph) Connected(p, q int) bool {
	return g.findRoot(p) == g.findRoot(q)
}

func (g *Graph) Count() int {
	return g.count
}

func (g *Graph) findRoot(x int) int {
	for g.parents[x] != x {
		// compress path
		g.parents[x] = g.parents[g.parents[x]]
		x = g.parents[x]
	}

	return x
}

// Empty returns true if graph does not contain any connections.
func (g *Graph) Empty() bool {
	return g.Count() == len(g.parents)
}

// Size returns number of connections of the graph.
func (g *Graph) Size() int {
	return len(g.parents) - g.Count()
}

// Clear removes all connections from the graph.
func (g *Graph) Clear() {
	g.count = len(g.parents)

	for i := range g.parents {
		g.parents[i] = i
		g.sizes[i] = 1
	}
}

// Values returns all graph nodes' parents.
func (g *Graph) Values() []interface{} {
	parents := make([]interface{}, len(g.parents))
	for i, v := range g.parents {
		parents[i] = v
	}
	return parents
}

// String returns a string representation of container
func (g *Graph) String() string {
	str := "UnionFind\n"
	values := []string{}
	for _, value := range g.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
