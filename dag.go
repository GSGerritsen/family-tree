package main

import (
	"fmt"
)

// Adjacency list implementation

type DAG struct {
	Pairs []*Pair
}

// Node to represent a person in the tree
type Node struct {
	name string
	sex  string
	b    string
	d    string
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

// Given a name, return the number of immediate children that person has. In the future, maybe have this return a map because there may be name duplicates
func (DAG *DAG) CountChildren(name string) int {
	for _, pair := range DAG.Pairs {
		if pair.male.name == name || pair.female.name == name {
			return len(pair.children)
		}
	}
	return 0
}

func (DAG DAG) PrintAdjacencyList() {
	for _, pair := range DAG.Pairs {
		fmt.Printf("%s + %s:\n=>%s\n", pair.male.StringName(), pair.female.StringName(), PrintChildren(pair.children))
	}
}

// Maybe separate into first and last names? Also, this is counting people twice.
func (DAG DAG) FindMostCommonName() (string, int) {
	max := 0
	mostCommonName := ""
	nameMap := make(map[string]int)
	for _, pair := range DAG.Pairs {
		nameMap[pair.male.StringName()]++
		nameMap[pair.female.StringName()]++
		for _, child := range pair.children {
			nameMap[child.StringName()]++
		}
	}

	for k, v := range nameMap {
		if v > max {
			max = v
			mostCommonName = k
		}
	}
	fmt.Printf("%s: %d\n", mostCommonName, max)
	return mostCommonName, max
}

func NewNode(name, sex, b, d string) *Node {
	return &Node{name, sex, b, d}
}

func AddChildren(nodes []*Node, node ...*Node) []*Node {
	for _, n := range node {
		nodes = append(nodes, n)
	}
	return nodes
}

func (n *Node) StringName() string {
	if n == nil {
		return "NIL"
	} else {
		return n.name
	}
}

func PrintChildren(nodes []*Node) string {
	result := ""
	for _, n := range nodes {
		result += n.name + ", "
	}
	return result
}

func main() {
	dag := NewDAG()
	// Move all this into separate file that populates graph
	var children []*Node
	var children1 []*Node
	var children2 []*Node
	var children3 []*Node
	var children4 []*Node
	//var children5 []*Node

	wernerusBovens := NewNode("Wernerus Bovens", "m", "n/a", "1716")
	mariaBraecken := NewNode("Maria Braecken", "f", "n/a", "1679")

	gertrudisBovens := NewNode("Gertrudis Bovens", "f", "1672", "n/a")
	mariaBovens := NewNode("Maria Bovens", "f", "1673", "n/a")
	guilielmusBovens := NewNode("Guilielmus Bovens", "m", "1675", "n/a")
	wilhelmusBovens := NewNode("Wilhelmus Bovens", "m", "1676", "n/a")
	catherinaBovens := NewNode("Catherina Bovens", "f", "1679", "n/a")

	children = AddChildren(children, gertrudisBovens, mariaBovens, guilielmusBovens, wilhelmusBovens, catherinaBovens)
	dag.AddPair(wernerusBovens, mariaBraecken, "n/a", children)

	arnoldusMiermans := NewNode("Arnoldus Miermans", "m", "n/a", "1718")

	guilielmusMiermans := NewNode("Guilielmus Miermans", "m", "1704", "n/a")
	wernerusMiermans := NewNode("Wernerus Miermans", "m", "1707", "n/a")
	gertrudisMiermans := NewNode("Gertrudis Miermans", "f", "1709", "n/a")
	arnoldusMiermans1 := NewNode("Arnoldus Miermans", "m", "1713", "n/a")
	joannesMiermans := NewNode("Joannes Miermans", "f", "1714", "n/a")
	mariaGMiermans := NewNode("Maria Gertrudis Miermans", "f", "1715", "n/a")
	arnoldusMiermans2 := NewNode("Arnoldus Miermans", "m", "1718", "n/a")
	mariaMiermans := NewNode("Maria Miermans", "f", "1722", "n/a")

	children1 = AddChildren(children1, guilielmusMiermans, wernerusMiermans, gertrudisMiermans, arnoldusMiermans1, joannesMiermans, mariaGMiermans, arnoldusMiermans2, mariaMiermans)
	dag.AddPair(arnoldusMiermans, catherinaBovens, "September 16, 1700", children1)

	joannesJanssen := NewNode("Joannes Janssen", "m", "n/a", "n/a")
	gertrudisPhilippens := NewNode("Gertrudis Philippens", "f", "n/a", "n/a")

	leonardusJanssen := NewNode("Leonardus Janssen", "m", "1678", "1746")
	children2 = AddChildren(children2, leonardusJanssen)
	dag.AddPair(joannesJanssen, gertrudisPhilippens, "n/a", children2)

	johannesJanssen := NewNode("Johannes Janssen", "m", "n/a", "1771")
	annaPenne := NewNode("Anna Penné", "f", "n/a", "1776")

	children3 = AddChildren(children3, johannesJanssen)
	dag.AddPair(leonardusJanssen, nil, "n/a", children3)

	annaJanssen := NewNode("Anna Janssen", "f", "1738", "n/a")
	mariaCatharinaJanssen := NewNode("Maria Catharina Janssen", "f", "1742", "n/a")
	leonardusJannsen1 := NewNode("Leonardus Janssen", "m", "1745", "n/a")

	children4 = AddChildren(children4, annaJanssen, mariaCatharinaJanssen, leonardusJannsen1)
	dag.AddPair(johannesJanssen, annaPenne, "October 18, 1737", children4)

	//annaMariaHechermans := NewNode("Anna Maria Hechermans", "f", "n/a", "n/a") // Marries guiliulmus miermans
	// next row to add: arnoluds miermans, susanna miermans etc

	dag.FindMostCommonName()
	dag.PrintAdjacencyList()

}
