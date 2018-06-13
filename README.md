# Go-raph

go-raph is a highly customizable shortest path finder written in Go. It implements the Dijkstra algorithm with very detailed filters over vertices and edges properties.

You can find a detailed documentation [here](https://godoc.org/github.com/shoprunback/go-raph/raph).

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

#### Properties

Properties are sets of strings. They can be added to vertices and edges.

```go
A := raph.NewVertex("Quentin", "people")
A.AddProp("drinks", "water")
A.AddProp("drinks", "coffee", "juice")
A.SetCost("brothers", 2)
fmt.Println(A.Props["drinks"])
fmt.Println(A.Costs["brothers"])
// => [water coffee juice]
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
A := raph.NewVertex("Paris")
B := raph.NewVertex("Amsterdam")
C := raph.NewVertex("Beijing")

// create edges
D := raph.NewEdge("P->B", "flight", "Paris", "Beijing")
D.AddProp("luggageSize", "S")
D.SetCost("price", 500)
D.SetCost("time", 11)
E := raph.NewEdge("P->A", "flight", "Paris", "Amsterdam")
E.AddProp("luggageSize", "L")
E.SetCost("price", 100)
E.SetCost("time", 5)
F := raph.NewEdge("A->B", "flight", "Amsterdam", "Beijing")
F.AddProp("luggageSize", "L")
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

You can compute shortest paths with a `Query` instance. If no path exists, `cost` will be `-1`.

Queries are expressed in JSON format:
- `from` origin vertex ID
- `to` destination vertex ID
- `constraint`
    - `vertex` constraint over edge props/costs
    - `edge` constraint over vertex props/costs
    - `label` edge label to go through
- `minimize` array of costs to minimize (vertices & edges)

```go
query = raph.NewQuery(`
    {
        "from": "Paris",
        "to": "Beijing",
        "constraint": {
            "edge": {
                "props": {
                    "luggageSize": ["M", "L"] // M or L
                },
                "costs": {
                    "time": 10 // at least 10
                }
            },
            "label": "flight"
        },
        "minimize": ["price"]
    }
`)
res = query.Run(*g)
```

If a vertex/edge does not contain the property specified by the constraint, it will not be filtered out.

If a vertex/edge does not contain the cost to minimize, passing though it will cost `0`.

Cost acts like a threshold.

If `minimize` is set to `["price", "price", "time"]` the equation `2 * price + time` will be minimized.

Find more examples [here](example/flight/main.go).

### Custom shortest path

You can implement your own `ShortestPath` algorithm would you need further customization. To do so, you need to declare a new _struct_ overriding the original **_ShortestPath(q Query)_** method. This [working example](example/mydijkstra/main.go) can help you.
