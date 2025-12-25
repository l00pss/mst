# Minimum Spanning Tree (MST) in Go

A Go implementation of graph data structures and Minimum Spanning Tree algorithms.

## Features

- **Graph Data Structures**: Vertex, Edge, and Graph implementations supporting both directed and undirected graphs
- **Kruskal's Algorithm**: MST using Union-Find with path compression and union by rank
- **Prim's Algorithm**: MST using priority queue (min-heap)
- **Graph Utilities**: Connectivity checking, graph printing, and MST weight calculation

## Installation

```bash
go get github.com/l00pss/mst
```

## Quick Start

```go
package main

import "github.com/yourusername/mst"

func main() {
    // Create an undirected graph
    g := mst.NewGraph(false)
    
    // Create vertices
    v0 := mst.Vertex{ID: 0, Name: "A", Edges: make([]*mst.Edge, 0)}
    v1 := mst.Vertex{ID: 1, Name: "B", Edges: make([]*mst.Edge, 0)}
    v2 := mst.Vertex{ID: 2, Name: "C", Edges: make([]*mst.Edge, 0)}
    
    // Add edges
    g.AddEdge(mst.Edge{From: &v0, To: &v1, Weight: 4})
    g.AddEdge(mst.Edge{From: &v1, To: &v2, Weight: 2})
    g.AddEdge(mst.Edge{From: &v0, To: &v2, Weight: 3})
    
    // Find MST using Kruskal
    mstEdges, totalWeight := g.Kruskal()
    mst.PrintMST(mstEdges, totalWeight, "KRUSKAL")
    
    // Or use Prim starting from vertex 0
    mstEdges, totalWeight = g.Prim(0)
    mst.PrintMST(mstEdges, totalWeight, "PRIM")
}
```

## Algorithms

### Kruskal's Algorithm
- Time Complexity: O(E log E)
- Uses Union-Find for cycle detection
- Sorts all edges and greedily adds minimum weight edges

### Prim's Algorithm
- Time Complexity: O(E log V)
- Uses priority queue (min-heap)
- Grows MST from a starting vertex

## Running Tests

```bash
go test -v
```

Run benchmarks:
```bash
go test -bench=.
```

## License

MIT
