package raph

import (
	"log"
)

// GetPath returns the path as a slice of strings.
func GetPath(from, to string, predsV, predsE map[string]string) []string {
	path := []string{}

	// return empty path if to does not exist
	if _, ok := predsV[to]; !ok {
		return path
	}

	// gather
	current := to
	ok := true
	for ok {
		path = append(path, current)
		path = append(path, predsE[current])
		current, ok = predsV[current]
	}

	// arrange
	path = path[:len(path)-1]
	Reverse(path)
	return path
}

// Concat returns the path concatenated with the specified path.
func Concat(path, path2 []string) []string {
	pathCopy := make([]string, len(path))
	copy(pathCopy, path)

	if len(path) == 0 {
		path2Copy := make([]string, len(path2))
		copy(path2Copy, path)
		return path2Copy
	}

	if len(path2) == 0 {
		return pathCopy
	}

	if path[len(path)-1] != path2[0] {
		log.Fatalln("Tried to compute", path, "+", path2, "but their ends differ")
	}

	pathCopy = append(pathCopy[:len(pathCopy)-1], path2...)
	return pathCopy
}

// GetDetailedPath returns a slice of objects corresponding to specified slice of ids and graph
func GetDetailedPath(path []string, g Graph) []map[string]interface{} {
	detailedPath := []map[string]interface{}{}
	for i := 0; i < len(path)/2; i++ {
		placeID := path[2*i]
		routeID := path[2*i+1]
		place := g.Vertices[placeID].ToJSON()
		route := g.Edges[routeID].ToJSON()
		detailedPath = append(detailedPath, place, route)
	}
	if len(path) > 0 {
		lastPlaceID := path[len(path)-1]
		lastPlace := g.Vertices[lastPlaceID].ToJSON()
		detailedPath = append(detailedPath, lastPlace)
	}
	return detailedPath
}
