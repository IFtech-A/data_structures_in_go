package graph

import "fmt"

type Vertex struct {
	Value    interface{}
	Visited  bool
	Neigbors []*Vertex
}

type Graph struct {
	Vertices    []*Vertex
	VerticesMap map[interface{}]*Vertex
}

func New() *Graph {
	return &Graph{
		VerticesMap: make(map[interface{}]*Vertex),
	}
}

func (g *Graph) AddVertex(v interface{}) {

	newVertex := &Vertex{Value: v}
	g.Vertices = append(g.Vertices, newVertex)
	g.VerticesMap[v] = newVertex
}

func (g *Graph) AddEdge(a, b interface{}) {
	aVertex, ok := g.VerticesMap[a]
	if !ok {
		return
	}
	bVertex, ok := g.VerticesMap[b]
	if !ok {
		return
	}

	aVertex.Neigbors = append(aVertex.Neigbors, bVertex)
}

func (g *Graph) BFS(startVertex interface{}) {
	startV, ok := g.VerticesMap[startVertex]
	if !ok {
		return
	}

	frontier := make([]*Vertex, 0)
	frontier = append(frontier, startV)
	startV.Visited = true
	level := 0
	for i := 0; i < len(frontier); i++ {
		fmt.Println("Level ", level)
		fmt.Println("BFS ", frontier[i].Value)
		next := make([]*Vertex, 0, len(frontier[i].Neigbors))
		next = append(next, frontier[i].Neigbors...)
		for _, v := range next {
			if !v.Visited {
				v.Visited = true
				fmt.Println(v.Value)
				frontier = append(frontier, v)
			}
		}
		level++
	}
}

func (g *Graph) dfs(node interface{}) {
	nodeVertex, ok := g.VerticesMap[node]
	if !ok {
		return
	}
	nodeVertex.Visited = true

	fmt.Println("DFS ", nodeVertex.Value)
	for _, neighbor := range nodeVertex.Neigbors {
		if !neighbor.Visited {
			g.dfs(neighbor.Value)
		}
	}
}
func (g *Graph) DFS(startNode interface{}) {
	g.dfs(startNode)
}

func (g *Graph) DFSAll() {
	for _, v := range g.Vertices {
		g.dfs(v.Value)
	}
}
