package player

import (
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
)

// State хранит множество ID событий игрока.
type State struct {
	Events map[event.EventID]struct{}
}

// NewState создаёт пустое состояние игрока.
func NewState() State {
	return State{Events: map[event.EventID]struct{}{}}
}

// Has возвращает true, если событие присутствует в состоянии.
func (s *State) Has(id event.EventID) bool { _, ok := s.Events[id]; return ok }

// Add добавляет событие в состояние.
func (s *State) Add(id event.EventID) { s.Events[id] = struct{}{} }

// Remove удаляет событие из состояния.
func (s *State) Remove(id event.EventID) { delete(s.Events, id) }

// ApplyEffect применяет события к состоянию (добавляет/удаляет события).
func (s *State) ApplyEffect(e node.Effect) {
	for _, id := range e.Add {
		s.Add(id)
	}
	for _, id := range e.Remove {
		s.Remove(id)
	}
}
