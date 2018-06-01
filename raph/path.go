package raph

import (
	"log"
)

// Path instance helps building a path from Dijkstra result.
type Path struct {
	From   string
	To     string
	PredsV map[string]string
	PredsE map[string]string
}

// NewPath returns a path instance.
func NewPath(from, to string, vertices, edges map[string]string) *Path {
	return &Path{from, to, vertices, edges}
}

// Get returns the path as a slice of strings.
func (p Path) Get() []string {
	path := []string{}

	// return empty path if to does not exist
	if _, ok := p.PredsV[p.To]; !ok {
		return path
	}

	// gather
	current := p.To
	ok := true
	for ok {
		path = append(path, current)
		path = append(path, p.PredsE[current])
		current, ok = p.PredsV[current]
	}

	// arrange
	path = path[:len(path)-1]
	Reverse(path)
	return path
}

// Append returns the path concatenated with the specified path.
func (p Path) Append(path2 []string) []string {
	path := p.Get()

	if len(path) == 0 {
		return path2
	}

	if len(path2) == 0 {
		return path
	}

	if path[len(path)-1] != path2[0] {
		log.Fatalln("Tried to compute", path, "+", path2, "but their ends differ")
	}

	path = append(path[:len(path)-1], path2...)
	return path
}