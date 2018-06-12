package raph

import (
	"encoding/json"
	"log"
)

type Query struct {
	From       string      `json:"from"`
	To         string      `json:"to"`
	Constraint *Constraint `json:"constraint"`
	Minimize   []string    `json:"minimize"`
	Option     string      `json:"option"`
}

func NewQuery(b []byte) *Query {
	query := Query{Constraint: &Constraint{Vertex: &Component{}, Edge: &Component{}}}
	err := json.Unmarshal(b, &query)
	if err != nil {
		log.Fatalln(err)
	}
	return &query
}

func (q *Query) AddMinimize(costs ...string) {
	q.Minimize = append(q.Minimize, costs...)
}
