# Go-raph

go-raph is a highly customizable shortest path finder written in Go. It implements the Dijkstra algorithm with very detailed filters over vertices and edges properties.

You can find a detailed documentation [here](raph)

## Usage

```go
import "github.com/shoprunback/go-raph/raph"
```

### Initialize

To initialize a graph, you can use the built-in **_NewGraph()_** function.

```go
g := raph.NewGraph()
```

### Vertices & Edges

Graphs can store vertices, edges and multiedges.

A multiedge is an edge that can have multiple origins and destinations. It is useful when you have edges that share the same properties, and more efficient for traversal queries. When creating an `Edge`, you must reference its connected nodes (from / to) by their id.

#### Prototypes

```go
func NewVertex(id, label string) *Vertex {}
func NewEdge(id, label, from, to string) *Edge {}
func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge {}
```

#### Properties & Costs

Properties are sets of strings. They can be added to vertices and edges.

```go
A := raph.NewVertex("Quentin")
A.AddProp("job", "student")
A.AddProp("job", "intern")
A.SetCost("coffee", 2)
fmt.Println(A.Props["job"])
fmt.Println(A.Costs["coffee"])
// => [student intern]
// => 2
```

#### Costs

Costs are integer values. They can be set on vertices and edges.

They are used during shortest path computation as:
- the value to minimize
- a threshold to filter out vertices and edges

```go
A := raph.NewEdge("US Route 66", "route", "Chicago", "Santa Monica")
A.AddProp("states", "Illinois")
A.AddProp("states", "Missouri", "Kansas", "Oklahoma")
A.AddProp("states", "Texas", "New Mexico", "Arizona", "California")
A.SetCost("length", 3940)
fmt.Println(A.Props["states"])
fmt.Println(A.Costs["length"])
// => [Illinois Missouri Kansas Oklahoma Texas New Mexico Arizona California]
// => 3940
```

### Populate

You can add vertices and edges to a `Graph` instance.

```go
// create graph
g := raph.NewGraph()

// create vertices
A := raph.NewVertex("Paris", "city")
B := raph.NewVertex("Amsterdam", "city")
C := raph.NewVertex("Beijing", "city")

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

`Constraint` objects let you customize path traversals to fit your needs.

To filter out nodes and edges, simply use:
- `Constraint.AddProp(prop, value string)` The node/edge will be filtered out if there is no intersection between the props of the vertex/edge and the props of the constraint. If a vertex/edge does not have the property specified by the constraint, it will not be filtered out.

- `Constraint.SetCost(prop string, value int)` acts like a threshold: will dodge the vertex/edge if the property value is less than specified value. If a vertex/edge does not contain the cost to minimize, passing though it will cost `0`.

#### Prototypes

```go
func (d *Dijkstra) ShortestPath(from, to string, constraint Constraint, minimize ...string) ([]string, int) // slice of ids
func (d *Dijkstra) ShortestPathDetailed(from, to string, constraint Constraint, minimize ...string) ([]map[string]interface{}, int) // slice of detailed objects
```

```go
// init dijkstra
d := raph.NewDijkstra(*g)

var constraint *raph.Constraint
var path []string
var cost int

// find shortest path between Paris and Beijing, minimizing time
constraint = raph.NewConstraint("flight")
path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "time")
fmt.Println(path, cost)
// => [Paris Beijing] 11

// find shortest path between Paris and Beijing, minimizing price
constraint = raph.NewConstraint("flight")
path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "price")
fmt.Println(path, cost)
// => [Paris Amsterdam Beijing] 400

// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
constraint = raph.NewConstraint("flight")
constraint.AddProp("maxLuggageSize", "M")
constraint.AddProp("maxLuggageSize", "L")
path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "time")
fmt.Println(path, cost)
// => [Paris Amsterdam Beijing] 15

// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
constraint = raph.NewConstraint("flight")
constraint.SetCost("time", 10)
path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "price")
fmt.Println(path, cost)
// => [Paris Beijing] 500
```

You can minimize more than a single cost by giving a list of costs at the end of **_ShortestPath()_** method. For instance, if you want to minimize **_2 x time + price_**, call `d.ShortestPath("Paris", "Beijing", *constraint, "time", "time", "price")`.

### Custom shortest path

You can implement your own `ShortestPath` algorithm would you need further customization. To do so, you need to declare a new _struct_ overriding the original **_ShortestPath()_** method. This [working example](example/mydijkstra/main.go) can help you.
