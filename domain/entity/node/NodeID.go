package node

import "remez_story/common/domainPrimitive/primitive/id"

// NodeID — доменный примитив идентификатора узла.
type NodeID struct{ id.EntityID }

func NewNodeID(v int64) (NodeID, error) {
	x, err := id.NewEntityID(v)
	return NodeID{EntityID: x}, err
}

func NodeIDFromString(s string) (NodeID, error) {
	x, err := id.EntityIDFrom(s)
	return NodeID{EntityID: x}, err
}
