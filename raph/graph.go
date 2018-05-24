package raph

import (
	"strconv"
)

type Vertex struct {
	ID    string
	Props map[string]string
}

func NewVertex(id string, props map[string]string) *Vertex {
	return &Vertex{id, props}
}

type Edge struct {
	Vertex
	Label string
	Froms map[string]bool
	Tos   map[string]bool
}

func NewMultiEdge(id, label string, froms, tos map[string]bool, props map[string]string) *Edge {
	return &Edge{Vertex{id, props}, label, froms, tos}
}

func NewEdge(id, label, from, to string, props map[string]string) *Edge {
	froms := map[string]bool{from: true}
	tos := map[string]bool{to: true}
	return &Edge{Vertex{id, props}, label, froms, tos}
}

type Graph struct {
	Vertices []string
	Edges    []string
	DB       Storage
}

func NewGraph() *Graph {
	DB := NewStorage()
	return &Graph{[]string{}, []string{}, *DB}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices = append(g.Vertices, v.ID)
	g.DB.SetProps(v.ID, v.Props)
}

func (g *Graph) AddEdge(e *Edge) {
	g.Edges = append(g.Edges, e.ID)
	g.DB.SetProps(e.ID, e.Props)

	// connect all "from" vertices to edge
	for from := range e.Froms {
		g.Connect(from, e.ID, e.Label)
	}

	// connect edge to all "to" vertices
	for to := range e.Tos {
		g.Connect(e.ID, to, e.Label)
	}
}

func (g *Graph) Connect(from, to, label string) {
	g.DB.Push(from, label, to)
}

// retrieve neighbors of vertex
func (g Graph) GetNeighbors(vertex string, constraints Constraints) map[string]bool {
	neighbors := map[string]bool{}

	// retrieve outgoing edges with label
	edges := g.DB.Pull(vertex, constraints.EdgeLabel)

	for _, edge := range edges {
		// assert that edge satifies constraints
		if g.Satisfies(edge, constraints.EdgeProps) {
			// retrieve edge ends
			vertices := g.DB.Pull(edge, constraints.EdgeLabel)

			for _, neighbor := range vertices {
				// assert that vertex satifies constraints
				if g.Satisfies(neighbor, constraints.VertexProps) {
					// add vertex to neighbors
					neighbors[neighbor] = true
				}
			}
		}
	}

	return neighbors
}

// retrieve neighbors of vertexID through edgeLabel with edgeConstraints, minimizing weightProp
func (g Graph) GetNeighborsWithEdgesAndWeights(vertex, weightProp string, constraints Constraints) (map[string]string, map[string]int) {
	crossedEdges := map[string]string{}
	weights := map[string]int{}

	// retrieve outgoing edges with label
	edges := g.DB.Pull(vertex, constraints.EdgeLabel)

	for _, edge := range edges {
		// assert that edge satifies constraints
		if g.Satisfies(edge, constraints.EdgeProps) {
			// retrieve weight of edge
			if weightString, ok := g.DB.GetProp(edge, weightProp); ok {
				// the weight exists
				weight, _ := strconv.Atoi(weightString)
				vertices := g.DB.Pull(edge, constraints.EdgeLabel)

				for _, neighbor := range vertices {
					// assert that vertex satifies constraints
					if g.Satisfies(neighbor, constraints.VertexProps) {
						if currentWeight, exist := weights[neighbor]; !exist || weight < currentWeight {
							// update the path to the vertex
							crossedEdges[neighbor] = edge
							weights[neighbor] = weight
						}
					}
				}
			}
		}
	}
	return crossedEdges, weights
}

// retrieve accessible vertices from vertex
func (g Graph) GetAccessibleVertices(vertex string, constraints Constraints) map[string]bool {
	// retrieve neighbors of vertex
	neighbors := g.GetNeighbors(vertex, constraints)

	// only vertex is accessible at the beginning
	accessibleVertices := map[string]bool{vertex: true}
	for neighbor := range neighbors {
		// pass accessible vertices map to be updated
		g.getAccessibleVerticesRecursive(neighbor, constraints, accessibleVertices)
	}

	return accessibleVertices
}

// retrieve accessible vertices from vertexID through edgeLabel with edgeConstraints
func (g Graph) getAccessibleVerticesRecursive(vertex string, constraints Constraints, accessibleVertices map[string]bool) {
	// inform that the vertex is accessible
	accessibleVertices[vertex] = true

	// retrieve neighbors of vertex
	neighbors := g.GetNeighbors(vertex, constraints)

	// recursive call on all neighbors if not done yet
	for neighbor := range neighbors {
		if _, ok := accessibleVertices[neighbor]; !ok {
			g.getAccessibleVerticesRecursive(neighbor, constraints, accessibleVertices)
		}
	}
}

// check if id element statisfies props
func (g Graph) Satisfies(id string, props map[string][]string) bool {
	for prop, satisfyingValues := range props {
		// retrieve property value of id element
		propString, ok := g.DB.GetProp(id, prop)

		// if property does not exist
		if !ok {
			return false
		}

		// check if property can be converted to int
		propInt, notInt := strconv.Atoi(propString)

		if notInt != nil {
			// if propString is a string, we check if it is in acceptedValues
			if !Contains(satisfyingValues, propString) {
				return false
			}
		} else {
			// if prop is an integer, propInt should be grater or equal than constraint
			min, _ := strconv.Atoi(satisfyingValues[0])
			if !(propInt >= min) {
				return false
			}
		}
	}

	// all constraints are satisfied
	return true
}
