package raph

import (
	"fmt"
)

type Storage struct {
	Map   map[string]string
	Lists map[string][]string
}

func NewStorage() *Storage {
	return &Storage{map[string]string{}, map[string][]string{}}
}

// get the property of an object
func (s Storage) GetProp(id, prop string) (string, bool) {
	key := fmt.Sprintf("%s:%s", id, prop)
	value, ok := s.Map[key]
	return value, ok
}

// set the property of an object
func (s *Storage) SetProp(id, prop, value string) {
	key := fmt.Sprintf("%s:%s", id, prop)
	s.Map[key] = value
}

// set the properties of an object
func (s *Storage) SetProps(id string, props map[string]string) {
	for key, value := range props {
		s.SetProp(id, key, value)
	}
}

// get the list of an id
func (s Storage) Pull(id, list string) []string {
	key := id + ":" + list
	return s.Lists[key]
}

// push an element into the list of an id
func (s *Storage) Push(id, list, el string) {
	key := id + ":" + list
	s.Lists[key] = append(s.Lists[key], el)
}
