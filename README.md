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

### Create edges & vertices

Graphs can store vertices, edges and multiedges. A multiedge is an edge that can have multiple origins and destinations. It is useful when you have edges that share the same properties, and more efficient for traversal queries.

### Prototypes

```go
func NewVertex(id string, props map[string]string) *Vertex
func NewEdge(id, label, from, to string, props map[string]string) *Edge
func NewMultiEdge(id, label string, froms, tos map[string]bool, props map[string]string) *Edge
```

When creating an `Edge`, you must reference its connected nodes by their id (from / to).

### Populate

You can add vertices and edges to a `Graph` instance.

```go
// create vertices
noProps := map[string]string{}
A := raph.NewVertex("Paris", noProps)
B := raph.NewVertex("Amsterdam", noProps)
C := raph.NewVertex("Beijing", noProps)

// create edges
D := raph.NewEdge("P->B", "flight", "Paris", "Beijing", map[string]string{"price": "500", "time": "11", "maxLuggageSize": "S"})
E := raph.NewEdge("P->A", "flight", "Paris", "Amsterdam", map[string]string{"price": "100", "time": "5", "maxLuggageSize": "L"})
F := raph.NewEdge("A->B", "flight", "Amsterdam", "Beijing", map[string]string{"price": "300", "time": "10", "maxLuggageSize": "L"})

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

`Constraints` object lets you customize path traversal to your needs. Simply use *_AddNodeConstraint(prop, value string)_* & *_AddEdgeConstraint(prop, value string)_* to filter out nodes & edges.

```go
// init dijkstra
d := raph.NewDijkstra(*g)

var constraints *raph.Constraints
var path []string
var cost int

// find shortest path between Paris and Beijing, minimizing time
constraints = raph.NewConstraints("flight")
path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraints)
fmt.Println(path, cost)
// => [Paris P->B Beijing] 11

// find shortest path between Paris and Beijing, minimizing price
constraints = raph.NewConstraints("flight")
path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraints)
fmt.Println(path, cost)
// => [Paris P->A Amsterdam A->B Beijing] 400

// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
constraints = raph.NewConstraints("flight")
constraints.AddEdgeConstraint("maxLuggageSize", "M")
constraints.AddEdgeConstraint("maxLuggageSize", "L")
path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraints)
fmt.Println(path, cost)
// => [Paris P->A Amsterdam A->B Beijing] 15

// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
constraints = raph.NewConstraints("flight")
constraints.AddEdgeConstraint("time", "10")
path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraints)
fmt.Println(path, cost)
// => [Paris P->B Beijing] 500
```

### Custom shortest path

You can implement your own `ShortestPath` algorithm would you need further customization. To do so, you need to declare a new _struct_ overriding the **_ShortestPath()_** method. This [working example](example/mydijkstra/main.go) can help you.
