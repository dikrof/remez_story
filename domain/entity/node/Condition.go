package node

import "remez_story/domain/entity/event"

// Condition описывает требования к набору событий игрока.
type Condition struct {
	RequireAll  []event.EventID `json:"require_all,omitempty"`
	RequireNone []event.EventID `json:"require_none,omitempty"`
}

// Effect описывает изменение набора событий игрока (добавить/удалить).
type Effect struct {
	Add    []event.EventID `json:"add,omitempty"`
	Remove []event.EventID `json:"remove,omitempty"`
}

// ConditionalEdge — ребро графа без выбора пользователя, срабатывающее по условию.
type ConditionalEdge struct {
	When     Condition
	ToNodeID NodeID
}
