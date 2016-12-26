package main

import (
	"fmt"
)

// Adjacency list implementation
// https://en.wikipedia.org/wiki/Adjacency_list

type DAG struct {
	Pairs []*Pair
}

// Node to represent a person in the tree
type Node struct {
	name   string
	gender string
	b      string
	d      string
}

// subject to change: the male/female pair that produced a set of children, increasing graph depth. Could only have one person.
// Ex: female is unknown but some male contributed children that were relevant to rest of graph.
type Pair struct {
	male        Node
	female      Node
	dateMarried string
	children    []*Node
}

func NewDAG() *DAG {
	return &DAG{Pairs: nil}
}

func (DAG *DAG) AddPair(male, female Node, dateMarried string, children []*Node) {
	newPair := &Pair{male, female, dateMarried, children}
	DAG.Pairs = append(DAG.Pairs, newPair)
}

func main() {
	fmt.Println("Working")
}
