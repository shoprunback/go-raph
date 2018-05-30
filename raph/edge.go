package raph

type Edge struct {
	Vertex
	Label      string
	Froms, Tos map[string]bool
}

func NewEdge(id, label, from, to string) *Edge {
	froms := map[string]bool{from: true}
	tos := map[string]bool{to: true}
	return NewMultiEdge(id, label, froms, tos)
}

func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge {
	return &Edge{*NewVertex(id), label, froms, tos}
}
