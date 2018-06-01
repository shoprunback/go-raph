# raph
--
    import "github.com/shoprunback/go-raph/raph"


## Usage

```go
const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MaxCost = MaxInt
)
```
Sets MaxCost to MaxInt.

#### func  Contains

```go
func Contains(s []string, e string) bool
```

#### func  ContainsAll

```go
func ContainsAll(s []string, e []string) bool
```

#### func  ContainsOne

```go
func ContainsOne(s []string, e []string) bool
```

#### func  RandSeq

```go
func RandSeq(n int) string
```

#### func  Remove

```go
func Remove(s []string, i int) []string
```

#### func  Reverse

```go
func Reverse(s []string)
```

#### type Constraint

```go
type Constraint struct {
	Vertex
}
```

Constraint is an instance used to filter out nodes. It inheritates from Vertex
structure because it behaves more or less like a Vertex against which we will
compare vertices and edges. Its ID is not important.

#### func  NewConstraint

```go
func NewConstraint(label string) *Constraint
```
NewConstraint returns a constraint with specified label.

#### func (Constraint) Copy

```go
func (c Constraint) Copy() *Constraint
```
Copy returns a copy of the constraint.

#### type Dijkstra

```go
type Dijkstra struct {
	G      Graph
	Q      []string
	Costs  map[string]int
	PredsV map[string]string
	PredsE map[string]string
}
```

Dijkstra instance is used to compute Dijkstra algorithm.

#### func  NewDijkstra

```go
func NewDijkstra(g Graph) *Dijkstra
```

#### func (Dijkstra) GetCost

```go
func (d Dijkstra) GetCost(id string) int
```
GetCost returns the cost of speficied vertex. If it has not been set by
UpdateDistances yet, it returns MaxCost.

#### func (*Dijkstra) PickVertexFromQ

```go
func (d *Dijkstra) PickVertexFromQ() string
```
PickVertexFromQ returns the vertex with minimal cost and removes it from the
queue.

#### func (*Dijkstra) Reset

```go
func (d *Dijkstra) Reset()
```
Reset resets the Dijkstra instance for further use.

#### func (*Dijkstra) ShortestPath

```go
func (d *Dijkstra) ShortestPath(from, to string, constraint Constraint, minimize ...string) ([]string, int)
```
ShortestPath returns a slice of ids with its cost.

#### func (*Dijkstra) ShortestPathDetailed

```go
func (d *Dijkstra) ShortestPathDetailed(from, to string, constraint Constraint, minimize ...string) ([]map[string]interface{}, int)
```
ShortestPathDetailed returns detailed shortest path with its cost.

#### func (*Dijkstra) UpdateDistances

```go
func (d *Dijkstra) UpdateDistances(s1, s2, edge string, s1s2Weight int)
```
UpdateDistances updates the costs of s2 if it is not minimal. It also stores the
edge crossed to get that minimal cost.

#### type Edge

```go
type Edge struct {
	Vertex
	Froms map[string]bool `json:"-"` // list of vertices from which the edge is reachable
	Tos   map[string]bool `json:"-"` // list of vertices that the edge can reach
}
```

Edge represents an edge instance. It inheritates from Vertex structure. Froms &
Tos fields are removed from JSON Marshaling to be added manually as slices (cf
ToJSON).

#### func  NewEdge

```go
func NewEdge(id, label, from, to string) *Edge
```
NewEdge returns a new edge.

#### func  NewMultiEdge

```go
func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge
```
NewMultiEdge returns a new multiedge.

#### func (Edge) ToJSON

```go
func (e Edge) ToJSON() map[string]interface{}
```
ToJSON formats the vertex to JSON. Froms and Tos fields are converted to slice.

#### type Graph

```go
type Graph struct {
	Vertices    map[string]*Vertex
	Edges       map[string]*Edge
	Connections map[string][]string // indexes connection between vertices and edges
}
```

Graph represents a graph instance.

#### func  NewGraph

```go
func NewGraph() *Graph
```
NewGraph returns a new graph initilizated with empty fields.

#### func (*Graph) AddEdge

```go
func (g *Graph) AddEdge(e *Edge)
```
AddEdge adds an edge to the graph and connects it the specified vertices. It
also stores the inverse connections. Adding an edge connecting vertices that
does not exist raises an error.

#### func (*Graph) AddVertex

```go
func (g *Graph) AddVertex(v *Vertex)
```
AddVertex adds a vertex to the graph.

#### func (*Graph) Connect

```go
func (g *Graph) Connect(from, to, label string)
```
Connect adds specified connection to the graph indexed on label.

#### func (Graph) GetAccessibleVertices

```go
func (g Graph) GetAccessibleVertices(vertex string, constraint Constraint) map[string]bool
```
GetAccessibleVertices returns accessible vertices from vertex using private
method getAccessibleVerticesRecursive.

#### func (*Graph) GetConnections

```go
func (g *Graph) GetConnections(id, label string) []string
```
GetConnections returns reachable vertices or edges from specified edge or
vertex, respectively.

#### func (Graph) GetNeighbors

```go
func (g Graph) GetNeighbors(vertex string, constraint Constraint) map[string]bool
```
GetNeighbors retrieves neighbors of vertex under specified constraints.

#### func (Graph) GetNeighborsWithCostsAndEdges

```go
func (g Graph) GetNeighborsWithCostsAndEdges(vertex string, constraint Constraint, costs ...string) (map[string]int, map[string]string)
```
GetNeighborsWithCostsAndEdges returns reachable vertices. For each neighbor, it
returns with the minimal cost and the crossed edge (in the case of multiedges).

#### type Path

```go
type Path struct {
	From   string
	To     string
	PredsV map[string]string
	PredsE map[string]string
}
```

Path instance helps building a path from Dijkstra result.

#### func  NewPath

```go
func NewPath(from, to string, vertices, edges map[string]string) *Path
```
NewPath returns a path instance.

#### func (Path) Append

```go
func (p Path) Append(path2 []string) []string
```
Append returns the path concatenated with the specified path.

#### func (Path) Get

```go
func (p Path) Get() []string
```
Get returns the path as a slice of strings.

#### type Vertex

```go
type Vertex struct {
	ID    string              `json:"id"`
	Label string              `json:"label"`
	Props map[string][]string `json:"props"`
	Costs map[string]int      `json:"costs"`
}
```

Vertex represents a vertex instance.

#### func  NewVertex

```go
func NewVertex(id, label string) *Vertex
```
NewVertex returns a new vertex.

#### func (*Vertex) AddProp

```go
func (v *Vertex) AddProp(prop string, values ...string)
```
AddProp adds a list of property values to the vertex.

#### func (Vertex) Satisfies

```go
func (v Vertex) Satisfies(constraint Constraint) bool
```
Satisfies returns whether or not the vertex satisfies the props and costs
(threshold) of the constraint.

#### func (*Vertex) SetCost

```go
func (v *Vertex) SetCost(cost string, value int)
```
SetCost sets a cost value for the vertex.

#### func (Vertex) ToJSON

```go
func (v Vertex) ToJSON() map[string]interface{}
```
ToJSON formats the vertex to JSON.
