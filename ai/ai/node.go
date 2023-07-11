package ai

import (
	"fmt"
	"math/rand"
)

type Action int8

const (
	_ Action = iota // not defined
	Not
	And
	Or
	Xor
	// If ?
	// Add ?
	// Sub ?
)

type Nodes []*Node

type Node struct {
	action   Action
	constant uint8
}

func NewNode() *Node {
	return &Node{
		action: Action(rand.Int31n(4) + 1),
		// constant: uint8(rand.Int31n(math.MaxInt8 + 1)),
		constant: uint8(rand.Int31n(2)), // 0 or 1
	}
}

func (n *Node) Act(v uint8) (uint8, error) {
	switch n.action {
	case Not:
		return ^v, nil
	case And:
		return v & n.constant, nil
	case Or:
		return v | n.constant, nil
	case Xor:
		return v ^ n.constant, nil
	}

	return 0, fmt.Errorf("unknown action: %d", n.action)
}
