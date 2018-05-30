package raph

type Dijkstra struct {
	G              Graph
	Q              []string
	Costs          map[string]int
	PredsV, PredsE map[string]string
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MaxCost = MaxInt

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

func (d *Dijkstra) Reset() {
	*d = *NewDijkstra(d.G)
}

func (d Dijkstra) GetCost(v string) int {
	if cost, ok := d.Costs[v]; ok {
		return cost
	} else {
		return MaxCost
	}
}

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

// because there can be multiple edges between s1 & s2, we pass edge parameter to store which edge we went though
func (d *Dijkstra) UpdateDistances(s1, s2, edge string, s1s2Weight int) {
	cost := d.GetCost(s2)
	potentialCost := d.GetCost(s1) + s1s2Weight
	if potentialCost >= 0 && potentialCost < cost {
		d.Costs[s2] = potentialCost
		d.PredsV[s2] = s1
		d.PredsE[s2] = edge
	}
}

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
