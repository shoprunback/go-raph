package raph

import (
	"math"
)

// Dijkstra instance is used to compute Dijkstra algorithm.
type Dijkstra struct {
	G      Graph
	Q      []string
	Costs  map[string]float64
	PredsV map[string]string
	PredsE map[string]string
}

// NewDijkstra initializes and returns a Dijkstra instance with graph g.
func NewDijkstra(g Graph) *Dijkstra {
	vertices := g.Vertices

	// append all vertices to queue
	q := make([]string, 0, len(vertices))
	for _, vertex := range vertices {
		q = append(q, vertex.ID)
	}

	// initialize all costs with +infinity
	costs := map[string]float64{}
	for _, vertex := range vertices {
		costs[vertex.ID] = math.Inf(0)
	}

	PredsV := map[string]string{}
	PredsE := map[string]string{}

	return &Dijkstra{g, q, costs, PredsV, PredsE}
}

// Reset resets the Dijkstra instance for further use.
func (d *Dijkstra) Reset() {
	*d = *NewDijkstra(d.G)
}

// PickVertexFromQ returns the vertex with minimal cost and removes it from the queue.
func (d *Dijkstra) PickVertexFromQ() string {
	min := d.Costs[d.Q[0]]
	vertex := d.Q[0]
	index := 0
	for i, v := range d.Q {
		cost := d.Costs[v]
		if cost < min {
			min = cost
			vertex = v
			index = i
		}
	}
	d.Q = Remove(d.Q, index)
	return vertex
}

// UpdateDistances updates the costs of s2 if it is not minimal. It also stores the edge crossed to get that minimal cost.
func (d *Dijkstra) UpdateDistances(s1, s2, edge string, s1s2Weight float64) {
	cost := d.Costs[s2]
	potentialCost := d.Costs[s1] + s1s2Weight
	if potentialCost < cost {
		d.Costs[s2] = potentialCost
		d.PredsV[s2] = s1
		d.PredsE[s2] = edge
	}
}

// ShortestPath returns a slice of ids with its cost. The value minimized is the sum of specified costs (minimize slice).
func (d *Dijkstra) ShortestPath(query Query) ([]map[string]interface{}, float64) {
	// init dijkstra
	d.Reset()
	d.Costs[query.From] = 0

	// run dijkstra until queue is empty
	for len(d.Q) > 0 {
		s1 := d.PickVertexFromQ()
		neighbors, edges := d.G.GetNeighborsWithCostsAndEdges(s1, *query.Constraint, query.Minimize...)
		for s2, cost := range neighbors {
			edge := edges[s2]
			d.UpdateDistances(s1, s2, edge, cost)
		}
	}

	// arrange return variables
	path := GetPath(query.From, query.To, d.PredsV, d.PredsE)
	detailedPath := GetDetailedPath(path, d.G)
	cost := d.Costs[query.To]

	return detailedPath, cost
}

// ShortestPathInverse returns the inverted shortest path (to -> from) defined in the query. It is used to compute ShortestPathOption.
func (d *Dijkstra) ShortestPathInverse(query Query) ([]map[string]interface{}, float64) {
	tmp := query.From
	query.From = query.To
	query.To = tmp
	query.Constraint.Label = "~" + query.Constraint.Label
	return d.ShortestPath(query)
}

// ShortestPathOption returns a path (slice of nodes) with its cost. One of the vertices of the path includes the option specified in the query.
func (d *Dijkstra) ShortestPathOption(query Query) ([]map[string]interface{}, float64) {
	// compute bi-directional shortest path
	d.ShortestPath(query)
	fromCosts, fromPredsV, fromPredsE := d.Costs, d.PredsV, d.PredsE
	d.ShortestPathInverse(query)
	toCosts, toPredsV, toPredsE := d.Costs, d.PredsV, d.PredsE

	// select best vertex
	cost := math.Inf(0)
	minVertex := "none"
	for vertexID, vertex := range d.G.Vertices {
		if vertexCost, ok := vertex.Costs[query.Option]; ok {
			pathCost := fromCosts[vertexID] + toCosts[vertexID]
			if pathCost > 0 && pathCost < cost {
				cost = pathCost + vertexCost
				minVertex = vertexID
			}
		}
	}

	// gather paths from->vertex & vertex->to
	path1 := GetPath(query.From, minVertex, fromPredsV, fromPredsE)
	path2 := GetPath(query.To, minVertex, toPredsV, toPredsE)
	Reverse(path2)
	path := Concat(path1, path2)
	detailedPath := GetDetailedPath(path, d.G)

	// arrange return variables
	for _, vertex := range detailedPath {
		if vertex["id"] == minVertex {
			vertex["option"] = query.Option
		}
	}

	return detailedPath, cost
}
