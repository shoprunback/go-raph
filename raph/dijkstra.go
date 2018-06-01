package raph

// Dijkstra instance is used to compute Dijkstra algorithm.
type Dijkstra struct {
	G      Graph
	Q      []string
	Costs  map[string]int
	PredsV map[string]string
	PredsE map[string]string
}

// Sets MaxCost to MaxInt.
const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MaxCost = MaxInt
)

func NewDijkstra(g Graph) *Dijkstra {
	vertices := g.Vertices

	// append all vertices to queue
	q := make([]string, 0, len(vertices))
	for _, vertex := range vertices {
		q = append(q, vertex.ID)
	}

	// initialize all costs at MaxCost
	costs := map[string]int{}
	for _, vertex := range vertices {
		costs[vertex.ID] = MaxCost
	}

	PredsV := map[string]string{}
	PredsE := map[string]string{}

	return &Dijkstra{g, q, costs, PredsV, PredsE}
}

// Reset resets the Dijkstra instance for further use.
func (d *Dijkstra) Reset() {
	*d = *NewDijkstra(d.G)
}

// GetCost returns the cost of speficied vertex. If it has not been set by UpdateDistances yet, it returns MaxCost.
func (d Dijkstra) GetCost(id string) int {
	if cost, ok := d.Costs[id]; ok {
		return cost
	} else {
		return MaxCost
	}
}

// PickVertexFromQ returns the vertex with minimal cost and removes it from the queue.
func (d *Dijkstra) PickVertexFromQ() string {
	min := d.GetCost(d.Q[0])
	vertex := d.Q[0]
	index := 0
	for i, v := range d.Q {
		cost := d.GetCost(v)
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
func (d *Dijkstra) UpdateDistances(s1, s2, edge string, s1s2Weight int) {
	cost := d.GetCost(s2)
	potentialCost := d.GetCost(s1) + s1s2Weight
	if potentialCost >= 0 && potentialCost < cost {
		d.Costs[s2] = potentialCost
		d.PredsV[s2] = s1
		d.PredsE[s2] = edge
	}
}

// ShortestPathDetailed returns detailed shortest path with its cost.
func (d *Dijkstra) ShortestPathDetailed(from, to string, constraint Constraint, minimize ...string) ([]map[string]interface{}, int) {
	path, cost := d.ShortestPath(from, to, constraint, minimize...)

	// gather path properties
	detailedPath := []map[string]interface{}{}
	for i := 0; i < len(path)/2; i++ {
		placeID := path[2*i]
		routeID := path[2*i+1]
		place := d.G.Vertices[placeID].ToJSON()
		route := d.G.Edges[routeID].ToJSON()
		detailedPath = append(detailedPath, place, route)
	}
	if len(path) > 0 {
		lastPlaceID := path[len(path)-1]
		lastPlace := d.G.Vertices[lastPlaceID].ToJSON()
		detailedPath = append(detailedPath, lastPlace)
	}

	return detailedPath, cost
}

// ShortestPath returns a slice of ids with its cost.
func (d *Dijkstra) ShortestPath(from, to string, constraint Constraint, minimize ...string) ([]string, int) {
	d.Reset()

	// init dijkstra with distance 0 for first vertex
	d.Costs[from] = 0

	// run dijkstra until queue is empty
	for len(d.Q) > 0 {
		s1 := d.PickVertexFromQ()
		neighbors, edges := d.G.GetNeighborsWithCostsAndEdges(s1, constraint, minimize...)
		for s2, cost := range neighbors {
			edge := edges[s2]
			d.UpdateDistances(s1, s2, edge, cost)
		}
	}

	// arrange return variables
	path := NewPath(from, to, d.PredsV, d.PredsE).Get()
	cost := d.GetCost(to)
	if from == to {
		cost = 0
	} else if len(path) == 0 || cost == MaxCost {
		cost = -1
	}

	return path, cost
}
