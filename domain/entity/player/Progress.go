package player

import (
	"remez_story/domain/entity/choice"
	"remez_story/domain/entity/node"
	"time"
)

// DecisionRecord — запись о сделанном выборе (для аудита/повторов/аналитики).
type DecisionRecord struct {
	NodeID   node.NodeID
	ChoiceID choice.ChoiceID
	At       time.Time
}

// Progress — корневой агрегат прохождения игрока.
type Progress struct {
	PlayerID      PlayerID
	CurrentNodeID *node.NodeID
	State         State
	Decisions     []DecisionRecord

	StartedAt time.Time
	UpdatedAt time.Time
}

// Reset сбрасывает прохождение: очищает состояние/журнал и ставит на стартовый узел.
func (p *Progress) Reset(to node.NodeID) {
	p.CurrentNodeID = &to
	p.State = NewState()
	p.Decisions = nil
	now := time.Now()
	p.StartedAt, p.UpdatedAt = now, now
}
