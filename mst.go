package mst

import "errors"

type Vertex struct {
	ID    int
	Name  string
	Data  any
	Edges []*Edge
}

func NewVertex(ID int, Name string, Data any, Edges []*Edge) (*Vertex, error) {
	if len(Edges) == 0 {
		return nil, errors.New("Vertex has no any edge")
	}
	return &Vertex{
		ID:    ID,
		Name:  Name,
		Data:  Data,
		Edges: Edges,
	}, nil
}

type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
	Data   any
}

func NewEdge(From *Vertex, To *Vertex, Weight int, Data any) (*Edge, error) {
	if From == nil && To == nil {
		return nil, errors.New("Edge has no any Vertex")
	}
	return &Edge{
		From:   From,
		To:     To,
		Weight: Weight,
		Data:   Data,
	}, nil
}
