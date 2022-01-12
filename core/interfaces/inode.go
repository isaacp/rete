package interfaces

type INode interface {
	Id() string
	Name() string
	Neighbors() []INode
	AddNeighbor(INode)
}
