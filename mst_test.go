package mst

import (
	"fmt"
	"testing"
)

// TestBasicGraph tests basic graph operations
func TestBasicGraph(t *testing.T) {
	fmt.Println("\n=== BASIC GRAPH TEST ===")

	// Create undirected graph
	g := NewGraph(false)

	// Create vertices
	v0 := Vertex{ID: 0, Name: "A", Data: nil, Edges: make([]*Edge, 0)}
	v1 := Vertex{ID: 1, Name: "B", Data: nil, Edges: make([]*Edge, 0)}
	v2 := Vertex{ID: 2, Name: "C", Data: nil, Edges: make([]*Edge, 0)}

	// Add edges
	g.AddEdge(Edge{From: &v0, To: &v1, Weight: 4, Data: nil})
	g.AddEdge(Edge{From: &v1, To: &v2, Weight: 2, Data: nil})
	g.AddEdge(Edge{From: &v0, To: &v2, Weight: 3, Data: nil})

	g.Print()

	if g.VertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", g.VertexCount())
	}

	if g.EdgeCount() != 3 {
		t.Errorf("Expected 3 edges, got %d", g.EdgeCount())
	}
}

// TestKruskal tests Kruskal's algorithm
func TestKruskal(t *testing.T) {
	fmt.Println("\n=== KRUSKAL ALGORITHM TEST ===")

	g := NewGraph(false)

	// Create sample graph: 6 vertices, 9 edges
	vertices := make([]*Vertex, 6)
	for i := 0; i < 6; i++ {
		vertices[i] = &Vertex{
			ID:    i,
			Name:  fmt.Sprintf("V%d", i),
			Data:  nil,
			Edges: make([]*Edge, 0),
		}
	}

	// Add edges
	edges := []struct{ from, to, weight int }{
		{0, 1, 4},
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 5},
		{2, 3, 8},
		{2, 4, 10},
		{3, 4, 2},
		{3, 5, 6},
		{4, 5, 3},
	}

	for _, e := range edges {
		g.AddEdge(Edge{
			From:   vertices[e.from],
			To:     vertices[e.to],
			Weight: e.weight,
			Data:   nil,
		})
	}

	g.Print()

	// Calculate MST
	mst, totalWeight := g.Kruskal()

	PrintMST(mst, totalWeight, "KRUSKAL")

	// Verify results
	expectedEdges := 5 // 6 vertices need 5 edges
	if len(mst) != expectedEdges {
		t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(mst))
	}

	// Expected total weight: 1+2+2+3+5 = 13
	expectedWeight := 13
	if totalWeight != expectedWeight {
		t.Errorf("Expected MST weight %d, got %d", expectedWeight, totalWeight)
	}
}

// TestPrim tests Prim's algorithm
func TestPrim(t *testing.T) {
	fmt.Println("\n=== PRIM ALGORITHM TEST ===")

	g := NewGraph(false)

	vertices := make([]*Vertex, 6)
	for i := 0; i < 6; i++ {
		vertices[i] = &Vertex{
			ID:    i,
			Name:  fmt.Sprintf("V%d", i),
			Data:  nil,
			Edges: make([]*Edge, 0),
		}
	}

	edges := []struct{ from, to, weight int }{
		{0, 1, 4},
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 5},
		{2, 3, 8},
		{2, 4, 10},
		{3, 4, 2},
		{3, 5, 6},
		{4, 5, 3},
	}

	for _, e := range edges {
		g.AddEdge(Edge{
			From:   vertices[e.from],
			To:     vertices[e.to],
			Weight: e.weight,
			Data:   nil,
		})
	}

	g.Print()

	// Calculate MST starting from vertex 0
	mst, totalWeight := g.Prim(0)

	PrintMST(mst, totalWeight, "PRIM")

	expectedEdges := 5
	if len(mst) != expectedEdges {
		t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(mst))
	}

	expectedWeight := 13
	if totalWeight != expectedWeight {
		t.Errorf("Expected MST weight %d, got %d", expectedWeight, totalWeight)
	}
}

// TestKruskalVsPrim tests that both algorithms produce the same result
func TestKruskalVsPrim(t *testing.T) {
	fmt.Println("\n=== KRUSKAL vs PRIM COMPARISON TEST ===")

	g := NewGraph(false)

	// Create larger graph
	vertices := make([]*Vertex, 10)
	for i := 0; i < 10; i++ {
		vertices[i] = &Vertex{
			ID:    i,
			Name:  fmt.Sprintf("Node%d", i),
			Data:  nil,
			Edges: make([]*Edge, 0),
		}
	}

	// Chain connections + some cross connections
	edges := []struct{ from, to, weight int }{
		{0, 1, 5}, {1, 2, 3}, {2, 3, 7}, {3, 4, 2},
		{4, 5, 8}, {5, 6, 4}, {6, 7, 6}, {7, 8, 1},
		{8, 9, 9}, {0, 5, 10}, {2, 7, 12}, {4, 9, 15},
	}

	for _, e := range edges {
		g.AddEdge(Edge{
			From:   vertices[e.from],
			To:     vertices[e.to],
			Weight: e.weight,
			Data:   nil,
		})
	}

	// Kruskal
	mstKruskal, weightKruskal := g.Kruskal()
	fmt.Printf("Kruskal - Edges: %d, Weight: %d\n", len(mstKruskal), weightKruskal)

	// Prim
	mstPrim, weightPrim := g.Prim(0)
	fmt.Printf("Prim    - Edges: %d, Weight: %d\n", len(mstPrim), weightPrim)

	// Both algorithms must give the same total weight
	if weightKruskal != weightPrim {
		t.Errorf("Algorithms gave different results. Kruskal: %d, Prim: %d", weightKruskal, weightPrim)
	}

	fmt.Println("✓ Both algorithms found the same total weight!")
}

// TestCityNetwork tests a practical city network example
func TestCityNetwork(t *testing.T) {
	fmt.Println("\n=== CITY NETWORK TEST ===")

	g := NewGraph(false)

	// Cities
	cities := []struct {
		id   int
		name string
	}{
		{0, "Istanbul"},
		{1, "Ankara"},
		{2, "Izmir"},
		{3, "Bursa"},
		{4, "Antalya"},
	}

	vertices := make([]*Vertex, len(cities))
	for i, city := range cities {
		vertices[i] = &Vertex{
			ID:    city.id,
			Name:  city.name,
			Data:  map[string]any{"population": 1000000 * (i + 1)},
			Edges: make([]*Edge, 0),
		}
	}

	// Distances (km)
	distances := []struct{ from, to, km int }{
		{0, 1, 450}, // Istanbul-Ankara
		{0, 2, 330}, // Istanbul-Izmir
		{0, 3, 150}, // Istanbul-Bursa
		{1, 2, 550}, // Ankara-Izmir
		{2, 3, 380}, // Izmir-Bursa
		{2, 4, 500}, // Izmir-Antalya
		{3, 4, 450}, // Bursa-Antalya
	}

	for _, d := range distances {
		g.AddEdge(Edge{
			From:   vertices[d.from],
			To:     vertices[d.to],
			Weight: d.km,
			Data:   map[string]any{"distance_km": d.km, "type": "highway"},
		})
	}

	g.Print()

	// Calculate MST using Kruskal
	mst, totalDistance := g.Kruskal()

	PrintMST(mst, totalDistance, "KRUSKAL - CITY NETWORK")

	fmt.Printf("\n📊 Analysis:\n")
	fmt.Printf("   • City Count: %d\n", g.VertexCount())
	fmt.Printf("   • Minimum Road Network: %d km\n", totalDistance)
	fmt.Printf("   • Roads Used: %d\n", len(mst))

	if !g.IsConnected() {
		t.Error("Graph should be connected")
	}
}

// TestIsConnected tests graph connectivity
func TestIsConnected(t *testing.T) {
	fmt.Println("\n=== CONNECTIVITY TEST ===")

	// Connected graph
	g1 := NewGraph(false)
	v0 := &Vertex{ID: 0, Name: "A", Edges: make([]*Edge, 0)}
	v1 := &Vertex{ID: 1, Name: "B", Edges: make([]*Edge, 0)}
	v2 := &Vertex{ID: 2, Name: "C", Edges: make([]*Edge, 0)}

	g1.AddEdge(Edge{From: v0, To: v1, Weight: 1})
	g1.AddEdge(Edge{From: v1, To: v2, Weight: 1})

	if !g1.IsConnected() {
		t.Error("Graph should be connected")
	}
	fmt.Println("✓ Graph 1 is connected")

	// Disconnected graph
	g2 := NewGraph(false)
	v3 := &Vertex{ID: 0, Name: "A", Edges: make([]*Edge, 0)}
	v4 := &Vertex{ID: 1, Name: "B", Edges: make([]*Edge, 0)}
	v5 := &Vertex{ID: 2, Name: "C", Edges: make([]*Edge, 0)}
	v6 := &Vertex{ID: 3, Name: "D", Edges: make([]*Edge, 0)}

	g2.AddEdge(Edge{From: v3, To: v4, Weight: 1})
	g2.AddEdge(Edge{From: v5, To: v6, Weight: 1})

	if g2.IsConnected() {
		t.Error("Graph should be disconnected")
	}
	fmt.Println("✓ Graph 2 is disconnected (2 components)")
}

// BenchmarkKruskal benchmarks Kruskal's algorithm
func BenchmarkKruskal(b *testing.B) {
	g := NewGraph(false)

	vertices := make([]*Vertex, 100)
	for i := 0; i < 100; i++ {
		vertices[i] = &Vertex{
			ID:    i,
			Name:  fmt.Sprintf("V%d", i),
			Edges: make([]*Edge, 0),
		}
	}

	for i := 0; i < 99; i++ {
		g.AddEdge(Edge{
			From:   vertices[i],
			To:     vertices[i+1],
			Weight: i + 1,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Kruskal()
	}
}

// BenchmarkPrim benchmarks Prim's algorithm
func BenchmarkPrim(b *testing.B) {
	g := NewGraph(false)

	vertices := make([]*Vertex, 100)
	for i := 0; i < 100; i++ {
		vertices[i] = &Vertex{
			ID:    i,
			Name:  fmt.Sprintf("V%d", i),
			Edges: make([]*Edge, 0),
		}
	}

	for i := 0; i < 99; i++ {
		g.AddEdge(Edge{
			From:   vertices[i],
			To:     vertices[i+1],
			Weight: i + 1,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Prim(0)
	}
}
