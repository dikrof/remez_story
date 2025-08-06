package node

import (
	"encoding/json"
	"errors"
	"strings"
)

type NodeKind string

const (
	NodeNarration NodeKind = "NARRATION"
	NodeDialogue  NodeKind = "DIALOGUE"
	NodeChoice    NodeKind = "CHOICE"
)

func ParseNodeKind(s string) (NodeKind, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	switch NodeKind(s) {
	case NodeNarration, NodeDialogue, NodeChoice:
		return NodeKind(s), nil
	default:
		return "", errors.New("invalid NodeKind")
	}
}
func (k NodeKind) IsValid() bool {
	_, err := ParseNodeKind(string(k))
	return err == nil
}

func (k NodeKind) MarshalJSON() ([]byte, error) { return json.Marshal(string(k)) }
func (k *NodeKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	val, err := ParseNodeKind(s)
	if err != nil {
		return err
	}
	*k = val
	return nil
}
func (k NodeKind) String() string { return string(k) }
