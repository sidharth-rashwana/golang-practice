package main

import "fmt"

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsUtil(start, visited)
	fmt.Println()
}

func (g *Graph) dfsUtil(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Print(v, " ")

	for _, n := range g.adj[v] {
		if !visited[n] {
			g.dfsUtil(n, visited)
		}
	}
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	g.DFS(1)
}
