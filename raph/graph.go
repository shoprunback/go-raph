package raph

// Graph represents a graph instance.
type Graph struct {
	Vertices    map[string]*Vertex
	Edges       map[string]*Edge
	Connections map[string][]string // indexes connection between vertices and edges
}

// NewGraph returns a new graph initilizated with empty fields.
func NewGraph() *Graph {
	return &Graph{map[string]*Vertex{}, map[string]*Edge{}, map[string][]string{}}
}

// hasVertex returns whether or not the graph contains the vertex specified by its id.
func (g *Graph) hasVertex(id string) bool {
	_, ok := g.Vertices[id]
	return ok
}

// AddVertex adds a vertex to the graph.
func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices[v.ID] = v
}

// AddEdge adds an edge to the graph and connects it the specified vertices. It also stores the inverse connections.
// Adding an edge connecting vertices that does not exist raises an error.
func (g *Graph) AddEdge(e *Edge) {
	g.Edges[e.ID] = e

	for from := range e.Froms {
		if g.hasVertex(from) {
			g.Connect(from, e.ID, e.Label)
			g.Connect(e.ID, from, "~"+e.Label) // store inverse relation
		}
	}

	for to := range e.Tos {
		if g.hasVertex(to) {
			g.Connect(e.ID, to, e.Label)
			g.Connect(to, e.ID, "~"+e.Label) // store inverse relation
		}
	}
}

// GetConnections returns reachable vertices or edges from specified edge or vertex, respectively.
func (g *Graph) GetConnections(id, label string) []string {
	return g.Connections[id+":"+label]
}

// Connect adds specified connection to the graph indexed on label.
func (g *Graph) Connect(from, to, label string) {
	key := from + ":" + label
	g.Connections[key] = append(g.Connections[key], to)
}

// GetNeighbors retrieves neighbors of vertex under specified constraints.
func (g Graph) GetNeighbors(vertex string, constraint Constraint) map[string]bool {
	neighbors := map[string]bool{}

	// retrieve outgoing edges with label
	edges := g.GetConnections(vertex, constraint.Label)

	for _, edge := range edges {
		// assert that edge satifies constraint
		if g.Edges[edge].Satisfies(*constraint.Edge) {
			// retrieve edge ends
			vertices := g.GetConnections(edge, constraint.Label)

			for _, neighbor := range vertices {
				// assert that vertex satifies constraint
				if g.Vertices[neighbor].Satisfies(*constraint.Vertex) {
					// add vertex to neighbors
					neighbors[neighbor] = true
				}
			}
		}
	}

	return neighbors
}

// GetNeighborsWithCostsAndEdges returns reachable vertices. For each neighbor, it returns with the minimal cost and the crossed edge (in the case of multiedges).
func (g Graph) GetNeighborsWithCostsAndEdges(vertex string, constraint Constraint, minimize ...string) (map[string]float64, map[string]string) {
	weights := map[string]float64{}
	crossedEdges := map[string]string{}

	// retrieve outgoing edges with label
	edges := g.GetConnections(vertex, constraint.Label)

	for _, e := range edges {
		edge := g.Edges[e]

		// assert that edge satifies constraint
		if edge.Satisfies(*constraint.Edge) {
			vertices := g.GetConnections(edge.ID, constraint.Label)

			for _, n := range vertices {
				neighbor := g.Vertices[n]

				// assert that vertex satifies constraint
				if neighbor.Satisfies(*constraint.Vertex) {
					// compute potential cost of crossing vertex+edge
					potentialCost := 0.0
					for _, cost := range minimize {
						potentialCost += edge.Costs[cost] + neighbor.Costs[cost]
					}

					// compare potential cost to actual cost
					cost, ok := weights[neighbor.ID]
					if !ok || potentialCost < cost {
						weights[neighbor.ID] = potentialCost
						crossedEdges[neighbor.ID] = edge.ID
					}
				}
			}
		}
	}
	return weights, crossedEdges
}

// GetAccessibleVertices returns accessible vertices from vertex using private method getAccessibleVerticesRecursive. selectionConstraint applies only on the Vertex.
func (g Graph) GetAccessibleVertices(vertex string, traversalConstraint, selectionConstraint Constraint) map[string]bool {
	// no vertex is accessible at first
	accessibleVertices := map[string]bool{}

	// update accessibleVertices with recursive calls
	if g.hasVertex(vertex) {
		g.getAccessibleVerticesRecursive(vertex, traversalConstraint, selectionConstraint, accessibleVertices)
	}

	return accessibleVertices
}

// getAccessibleVerticesRecursive adds accessible vertices from vertex to accessibleVertices.
func (g Graph) getAccessibleVerticesRecursive(vertex string, traversalConstraint, selectionConstraint Constraint, accessibleVertices map[string]bool) {
	// inform that the vertex is accessible
	if g.Vertices[vertex].Satisfies(*selectionConstraint.Vertex) {
		accessibleVertices[vertex] = true
	}

	// recursive call on all neighbors if not done yet
	neighbors := g.GetNeighbors(vertex, traversalConstraint)
	for neighbor := range neighbors {
		if _, ok := accessibleVertices[neighbor]; !ok {
			g.getAccessibleVerticesRecursive(neighbor, traversalConstraint, selectionConstraint, accessibleVertices)
		}
	}
}
