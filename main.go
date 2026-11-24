package main

import "fmt"

// FindRedundantConnection returns the edge whose removal turns the graph
// back into a tree — the LAST edge in input order that closes a cycle.
// Union-find processes edges in order; the first union that fails is it.
func FindRedundantConnection(edges [][2]int) [2]int {
	n := len(edges) // nodes are 1..n and there are exactly n edges
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		ra, rb := find(e[0]), find(e[1])
		if ra == rb {
			return e // endpoints already connected: this edge closes the cycle
		}
		if rank[ra] < rank[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		if rank[ra] == rank[rb] {
			rank[ra]++
		}
	}
	return [2]int{} // unreachable for valid input (tree + one extra edge)
}

func main() {
	fmt.Println(FindRedundantConnection([][2]int{{1, 2}, {1, 3}, {2, 3}}))
}
