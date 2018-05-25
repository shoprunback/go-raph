# Go-raph

go-raph is a highly customizable shortest path finder written in Go. It implements the Dijkstra algorithm with very detailed filters over vertices and edges properties.

## Usage

```go
import "github.com/shoprunback/go-raph/raph"
```

### Initialize

To initialize a graph, you can use the built-in **_NewGraph()_** function.

```go
g := raph.NewGraph()
```

### Vertices and Edges

#### Prototypes

```go
func NewVertex(id string) *Vertex {}
func NewEdge(id, label, from, to string) *Edge {}
```

When creating an `Edge`, you must reference its connected nodes (from / to) by their id.

#### Properties

Properties are sets of strings. They can be added to vertices and edges.

```go
A := raph.NewVertex("Quentin")
A.AddProp("job", "student")
A.AddProp("job", "intern")
fmt.Println(A.Props["job"])
// => [student intern]
```

#### Costs

Costs are integer values. They can be added to vertices and edges.

They can be used during shortest path computation:
- as the value to minimize
- to filter vertices and edges

```go
A := raph.NewEdge("US Route 66", "route", "Chicago", "Santa Monica")
A.SetCost("length", 3940)
fmt.Println(A.Costs["length"])
// => 3940
```

### Populate

You can add vertices and edges to a `Graph` instance.

```go
// create graph
g := raph.NewGraph()

// create vertices
A := raph.NewVertex("Paris")
B := raph.NewVertex("Amsterdam")
C := raph.NewVertex("Beijing")

// create edges
D := raph.NewEdge("P->B", "flight", "Paris", "Beijing")
D.AddProp("maxLuggageSize", "S")
D.SetCost("price", 500)
D.SetCost("time", 11)
E := raph.NewEdge("P->A", "flight", "Paris", "Amsterdam")
E.AddProp("maxLuggageSize", "L")
E.SetCost("price", 100)
E.SetCost("time", 5)
F := raph.NewEdge("A->B", "flight", "Amsterdam", "Beijing")
F.AddProp("maxLuggageSize", "L")
F.SetCost("price", 300)
F.SetCost("time", 10)

// populate graph
g.AddVertex(A)
g.AddVertex(B)
g.AddVertex(C)
g.AddEdge(D)
g.AddEdge(E)
g.AddEdge(F)
```

## Shortest path

You can compute shortest paths with a `Dijkstra` instance. If no path exists, `cost` will be `-1`.

`Constraint` objects let you customize path traversals to your needs.

To filter out nodes and edges, simply use:
- `AddVertexConstraint(prop, value string)`
- `AddEdgeConstraint(prop, value string)`
- `SetMinCostConstraint(prop string, value int)` (will filter edges & vertices)

```go
// init dijkstra
d := raph.NewDijkstra(*g)

var constraint *raph.Constraint
var path []string
var cost int

// find shortest path between Paris and Beijing, minimizing time
constraint = raph.NewConstraint("flight")
path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraint)
fmt.Println(path, cost)
// => [Paris Beijing] 11

// find shortest path between Paris and Beijing, minimizing price
constraint = raph.NewConstraint("flight")
path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraint)
fmt.Println(path, cost)
// => [Paris Amsterdam Beijing] 400

// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
constraint = raph.NewConstraint("flight")
constraint.AddEdgeConstraint("maxLuggageSize", "M")
constraint.AddEdgeConstraint("maxLuggageSize", "L")
path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraint)
fmt.Println(path, cost)
// => [Paris Amsterdam Beijing] 15

// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
constraint = raph.NewConstraint("flight")
constraint.SetMinCostConstraint("time", 10)
path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraint)
fmt.Println(path, cost)
// => [Paris Beijing] 500
```

### Custom shortest path

You can implement your own `ShortestPath` algorithm would you need further customization. To do so, you need to declare a new _struct_ overriding the original **_ShortestPath()_** method. This [working example](example/mydijkstra/main.go) can help you.
