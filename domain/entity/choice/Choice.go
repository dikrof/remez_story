package choice

import "remez_story/domain/entity/node"

// Choice — вариант выбора игрока: отображаемый текст, эффекты и целевой узел.
type Choice struct {
	ID       ChoiceID
	Text     string
	Effects  []node.Effect
	ToNodeID node.NodeID
}
