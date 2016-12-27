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
	male        *Node
	female      *Node
	dateMarried string
	children    []*Node
}

func NewDAG() *DAG {
	return &DAG{Pairs: nil}
}

func (DAG *DAG) AddPair(male, female *Node, dateMarried string, children []*Node) {
	newPair := &Pair{male, female, dateMarried, children}
	DAG.Pairs = append(DAG.Pairs, newPair)
}

// Given a name, return the number of immediate children that person has
func (DAG *DAG) CountChildren(name string) int {
	for _, pair := range DAG.Pairs {
		if pair.male.name == name || pair.female.name == name {
			return len(pair.children)
		}
	}
	return 0
}

func NewNode(name, gender, b, d string) *Node {
	return &Node{name, gender, b, d}
}

func addChildren(nodes []*Node, node ...*Node) []*Node {
	for _, n := range node {
		nodes = append(nodes, n)
	}
	return nodes
}

func main() {
	dag := NewDAG()
	var children []*Node

	wernerusBovens := NewNode("Wernerus Bovens", "m", "n/a", "1716")
	mariaBraecken := NewNode("Maria Braecken", "f", "n/a", "1679")

	gertrudisBovens := NewNode("Gertrudis Bovens", "f", "1672", "n/a")
	mariaBovens := NewNode("Maria Bovens", "f", "1673", "n/a")
	guilielmusBovens := NewNode("Guilielmus Bovens", "m", "1675", "n/a")
	wilhelmusBovens := NewNode("Wilhelmus Bovens", "m", "1676", "n/a")
	catherinaBovens := NewNode("Catherina Bovens", "f", "1679", "n/a")

	children = addChildren(children, gertrudisBovens, mariaBovens, guilielmusBovens, wilhelmusBovens, catherinaBovens)

	dag.AddPair(wernerusBovens, mariaBraecken, "n/a", children)
	fmt.Printf("%d", dag.CountChildren("Wernerus Bovens"))

}
