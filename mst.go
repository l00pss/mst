package mst

import (
	"container/heap"
	"errors"
	"fmt"
	"sort"
)

type Vertex struct {
	ID    int
	Name  string
	Data  any
	Edges []*Edge
}

func (v *Vertex) String() string {
	return fmt.Sprintf("ID : %d", v.ID)
}

func NewVertex(id int, name string, data any, edges []*Edge) (*Vertex, error) {
	if len(edges) == 0 {
		return nil, errors.New("Vertex has no any edge")
	}
	return &Vertex{
		ID:    id,
		Name:  name,
		Data:  data,
		Edges: edges,
	}, nil
}

type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
	Data   any
}

func NewEdge(From *Vertex, To *Vertex, weight int, data any) (*Edge, error) {
	if From == nil && To == nil {
		return nil, errors.New("Edge has no any Vertex")
	}
	return &Edge{
		From:   From,
		To:     To,
		Weight: weight,
		Data:   data,
	}, nil
}

func (e *Edge) Reverse() *Edge {
	return &Edge{
		From:   e.To,
		To:     e.From,
		Weight: e.Weight,
		Data:   e.Data,
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s ---> %d ---> %s", e.From.String(), e.Weight, e.To.String())
}

func (e *Edge) Compare(other *Edge) int {
	if e.Weight < other.Weight {
		return -1
	} else if e.Weight > other.Weight {
		return 1
	} else {
		return 0
	}
}

type Graph struct {
	Vertices map[int]Vertex
	Edges    []*Edge
	Directed bool
}

func NewGraph(directed bool) Graph {
	return Graph{
		Vertices: make(map[int]Vertex),
		Edges:    make([]*Edge, 0),
		Directed: directed,
	}
}

func (g *Graph) GetVertex(id int) (*Vertex, bool) {
	v, exists := g.Vertices[id]
	return &v, exists
}

func (g *Graph) AddVertex(vertex Vertex) *Vertex {
	if v, exists := g.GetVertex(vertex.ID); exists {
		return v
	} else {
		g.Vertices[vertex.ID] = vertex
		return &vertex
	}
}

func (g *Graph) AddEdge(edge Edge) *Edge {
	from, fromExists := g.GetVertex(edge.From.ID)
	to, toExists := g.GetVertex(edge.To.ID)

	if !fromExists {
		from = g.AddVertex(*edge.From)
	}
	if !toExists {
		to = g.AddVertex(*edge.To)
	}

	// Add edge to graph
	newEdge := &Edge{
		From:   from,
		To:     to,
		Weight: edge.Weight,
		Data:   edge.Data,
	}
	g.Edges = append(g.Edges, newEdge)

	// Add edge to From vertex
	fromVertex := g.Vertices[from.ID]
	fromVertex.Edges = append(fromVertex.Edges, newEdge)
	g.Vertices[from.ID] = fromVertex

	// If undirected graph, add reverse edge as well
	if !g.Directed {
		reverseEdge := newEdge.Reverse()
		toVertex := g.Vertices[to.ID]
		toVertex.Edges = append(toVertex.Edges, reverseEdge)
		g.Vertices[to.ID] = toVertex
	}

	return newEdge
}

// VertexCount returns the total number of vertices
func (g *Graph) VertexCount() int {
	return len(g.Vertices)
}

// EdgeCount returns the total number of edges
func (g *Graph) EdgeCount() int {
	return len(g.Edges)
}

// Print displays the graph to the console
func (g *Graph) Print() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║            GRAPH INFORMATION           ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("Vertex Count: %d\n", g.VertexCount())
	fmt.Printf("Edge Count: %d\n", g.EdgeCount())
	if g.Directed {
		fmt.Println("Type: Directed Graph")
	} else {
		fmt.Println("Type: Undirected Graph")
	}
	fmt.Println("\nVertices and Edges:")
	for id, vertex := range g.Vertices {
		fmt.Printf("  [%d] %s -> ", id, vertex.Name)
		if len(vertex.Edges) == 0 {
			fmt.Println("(no edges)")
		} else {
			for i, edge := range vertex.Edges {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%s(w:%d)", edge.To.String(), edge.Weight)
			}
			fmt.Println()
		}
	}
}

// ==================== UNION-FIND DATA STRUCTURE ====================

// UnionFind is a data structure used for cycle detection
type UnionFind struct {
	parent map[int]int
	rank   map[int]int
}

// NewUnionFind creates a new UnionFind structure
func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

// MakeSet creates a new set for a vertex
func (uf *UnionFind) MakeSet(x int) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}

// Find finds the root vertex of a vertex (with path compression)
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges two sets (with union by rank)
// If two vertices are in different sets, it merges them and returns true
// If they are in the same set (cycle would be formed), it returns false
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in the same set, cycle would be formed
	}

	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

// ==================== KRUSKAL ALGORITHM ====================

// Kruskal finds MST using Kruskal's algorithm
// Sorts edges by weight and adds them without forming cycles
func (g *Graph) Kruskal() ([]*Edge, int) {
	if g.Directed {
		panic("Kruskal algorithm only works for undirected graphs")
	}

	mst := make([]*Edge, 0)
	totalWeight := 0

	// Sort edges by weight
	edges := make([]*Edge, len(g.Edges))
	copy(edges, g.Edges)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Create Union-Find structure
	uf := NewUnionFind()
	for id := range g.Vertices {
		uf.MakeSet(id)
	}

	// Check each edge
	for _, edge := range edges {
		// If edge doesn't form a cycle, add it
		if uf.Union(edge.From.ID, edge.To.ID) {
			mst = append(mst, edge)
			totalWeight += edge.Weight

			// MST should have V-1 edges
			if len(mst) == g.VertexCount()-1 {
				break
			}
		}
	}

	return mst, totalWeight
}

// ==================== PRIORITY QUEUE (FOR PRIM) ====================

// PriorityQueue is a min-heap priority queue for edges
type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// ==================== PRIM ALGORITHM ====================

// Prim finds MST using Prim's algorithm
// Starting from a vertex, at each step it adds the nearest vertex to the current tree
func (g *Graph) Prim(startID int) ([]*Edge, int) {
	if g.Directed {
		panic("Prim algorithm only works for undirected graphs")
	}

	start, exists := g.Vertices[startID]
	if !exists {
		return nil, 0
	}

	mst := make([]*Edge, 0)
	totalWeight := 0
	visited := make(map[int]bool)

	// Create priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Mark starting vertex
	visited[start.ID] = true

	// Add edges from starting vertex
	for _, edge := range start.Edges {
		heap.Push(pq, edge)
	}

	// Build MST
	for pq.Len() > 0 && len(mst) < g.VertexCount()-1 {
		edge := heap.Pop(pq).(*Edge)

		// Skip if target vertex is already visited
		if visited[edge.To.ID] {
			continue
		}

		// Add edge to MST
		mst = append(mst, edge)
		totalWeight += edge.Weight
		visited[edge.To.ID] = true

		// Add edges from the new vertex
		toVertex := g.Vertices[edge.To.ID]
		for _, nextEdge := range toVertex.Edges {
			if !visited[nextEdge.To.ID] {
				heap.Push(pq, nextEdge)
			}
		}
	}

	return mst, totalWeight
}

// ==================== HELPER FUNCTIONS ====================

// IsConnected checks if the graph is connected (using DFS)
func (g *Graph) IsConnected() bool {
	if g.VertexCount() == 0 {
		return true
	}

	// Start from the first vertex
	var startID int
	for id := range g.Vertices {
		startID = id
		break
	}

	visited := make(map[int]bool)
	g.dfs(startID, visited)

	return len(visited) == g.VertexCount()
}

// dfs Depth-First Search
func (g *Graph) dfs(nodeID int, visited map[int]bool) {
	visited[nodeID] = true
	vertex := g.Vertices[nodeID]

	for _, edge := range vertex.Edges {
		if !visited[edge.To.ID] {
			g.dfs(edge.To.ID, visited)
		}
	}
}

// GetMSTWeight returns the total weight of the MST
func GetMSTWeight(mst []*Edge) int {
	weight := 0
	for _, edge := range mst {
		weight += edge.Weight
	}
	return weight
}

// PrintMST prints the MST in a formatted way
func PrintMST(mst []*Edge, totalWeight int, algorithmName string) {
	fmt.Println("\n╔════════════════════════════════════════════════╗")
	fmt.Printf("║    MINIMUM SPANNING TREE - %-19s ║\n", algorithmName)
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Printf("\nEdge Count: %d\n", len(mst))
	fmt.Println("\nMST Edges:")
	for i, edge := range mst {
		fmt.Printf("  %2d. [%d:%s] --%d--> [%d:%s]\n",
			i+1,
			edge.From.ID, edge.From.Name,
			edge.Weight,
			edge.To.ID, edge.To.Name)
	}
	fmt.Printf("\n✓ Total Weight: %d\n", totalWeight)
	fmt.Println("════════════════════════════════════════════════")
}
