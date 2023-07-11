package ai

import (
	"fmt"
)

type Net struct {
	nodes Nodes // TODO: [][]Node
}

func NewNet(size int) Net {
	nodes := make(Nodes, size)

	for i := 0; i < size; i++ {
		nodes[i] = NewNode()
	}

	return Net{nodes: nodes}
}

func (n Net) Learn(data Data) {
	for i, node := range n.nodes {
		var try int

	label: // TODO: func
		try++

		fmt.Printf("learning node #%d(%+v), try #%d\n", i, node, try)

		for _, d := range data {
			got, err := node.Act(d.Input)
			if err != nil {
				fmt.Printf("learn: got error: %s\n", err)
				node = NewNode()
				n.nodes[i] = node

				goto label
			}

			if got != d.Want {
				node = NewNode()
				n.nodes[i] = node

				goto label
			}
		}
	}

}

func (n Net) Test(data Data) {
	passed := true
	for i, node := range n.nodes {
		for _, d := range data {
			got, err := node.Act(d.Input)
			if err != nil {
				fmt.Printf("test node #%d: got error: %s\n", i, err)
				passed = false
			}

			if got != d.Want {
				fmt.Printf("test node #%d(%+v): expected '%d', but got '%d'\n", i, node, d.Want, got)
				passed = false
			}
		}
	}

	fmt.Println("test passed:", passed)
}
