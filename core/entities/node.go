package entities

import (
	"github.com/google/uuid"
	"github.com/isaacp/rete/core/interfaces"
)

type Node struct {
	id        string
	name      string
	neighbors []interfaces.INode
}

func NewNode(name string) interfaces.INode {
	identifier := uuid.New()
	return &Node{
		id:   identifier.String(),
		name: name,
	}
}

func (node *Node) AddNeighbor(neighbor interfaces.INode) {
	node.neighbors = append(node.neighbors, neighbor)
}

func (node *Node) Id() string {
	return node.id
}
func (node *Node) Name() string {
	return node.name
}
func (node *Node) Neighbors() []interfaces.INode {
	return node.neighbors
}
