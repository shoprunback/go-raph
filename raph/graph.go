package raph

type Graph struct {
	Vertices map[string]*Vertex
	Edges    map[string]*Edge
	Connections map[string][]string
}

func NewGraph() *Graph {
	return &Graph{map[string]*Vertex{}, map[string]*Edge{}, map[string][]string{}}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices[v.ID] = v
}

func (g *Graph) AddEdge(e *Edge) {
	g.Edges[e.ID] = e
	for from := range e.Froms {
		g.Connect(from, e.ID, e.Label)
	}
	for to := range e.Tos {
		g.Connect(e.ID, to, e.Label)
	}
}

func (g *Graph) Connect(from, to, label string) {
	key := from+":"+label
	g.Connections[key] = append(g.Connections[key], to)
}

func (g *Graph) GetConnections(id, label string) ([]string) {
	return g.Connections[id+":"+label]
}

// retrieve neighbors of vertex
func (g Graph) GetNeighbors(vertex string, constraint Constraint) map[string]bool {
	neighbors := map[string]bool{}

	// retrieve outgoing edges with label
	edges := g.GetConnections(vertex, constraint.Label)

	for _, e := range edges {
		edge := g.Edges[e]
		// assert that edge satifies constraint
		if edge.Satisfies(constraint) {
			// retrieve edge ends
			vertices := g.GetConnections(edge.ID, constraint.Label)

			for _, neighbor := range vertices {
				// assert that vertex satifies constraint
				if g.Vertices[neighbor].Satisfies(constraint) {
					// add vertex to neighbors
					neighbors[neighbor] = true
				}
			}
		}
	}

	return neighbors
}

// retrieve neighbors & costs of vertex under constraint
func (g Graph) GetNeighborsWithCosts(vertex, cost string, constraint Constraint) map[string]int {
	weights := map[string]int{}

	// retrieve outgoing edges with label
	edges := g.GetConnections(vertex, constraint.Label)

	for _, e := range edges {
		edge := g.Edges[e]
		// assert that edge satifies constraint
		if edge.Satisfies(constraint) {
			vertices := g.GetConnections(edge.ID, constraint.Label)
			for _, n := range vertices {
				neighbor := g.Vertices[n]
				// assert that vertex satifies constraint
				if neighbor.Satisfies(constraint) {
					// append cost of vertex
					potentialCost := edge.Costs[cost] + neighbor.Costs[cost]
					actualCost, ok := weights[neighbor.ID]
					if !ok || potentialCost < actualCost {
						weights[neighbor.ID] = potentialCost
					}
				}
			}
		}
	}
	return weights
}

// retrieve accessible vertices from vertex
func (g Graph) GetAccessibleVertices(vertex string, constraint Constraint) map[string]bool {
	// retrieve neighbors of vertex
	neighbors := g.GetNeighbors(vertex, constraint)

	// only vertex is accessible at the beginning
	accessibleVertices := map[string]bool{vertex: true}
	for neighbor := range neighbors {
		// pass accessible vertices map to be updated
		g.getAccessibleVerticesRecursive(neighbor, constraint, accessibleVertices)
	}

	return accessibleVertices
}

// retrieve accessible vertices from vertexID through edgeLabel with edgeConstraint
func (g Graph) getAccessibleVerticesRecursive(vertex string, constraint Constraint, accessibleVertices map[string]bool) {
	// inform that the vertex is accessible
	accessibleVertices[vertex] = true

	// retrieve neighbors of vertex
	neighbors := g.GetNeighbors(vertex, constraint)

	// recursive call on all neighbors if not done yet
	for neighbor := range neighbors {
		if _, ok := accessibleVertices[neighbor]; !ok {
			g.getAccessibleVerticesRecursive(neighbor, constraint, accessibleVertices)
		}
	}
}
