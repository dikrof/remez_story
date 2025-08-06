package node

import (
	"remez_story/domain/entity/chapter"
	"remez_story/domain/entity/choice"
)

// Node — узел графа сюжета (реплика/повествование/выбор).
// Линейное продолжение задаётся через NextID. Ветвления — через Choices
// или Conditional (отложенные последствия).
type Node struct {
	ID         NodeID
	ChapterID  chapter.ChapterID
	SceneLabel SceneLabel // «мягкая сцена» для UI/аналитики

	Kind    NodeKind
	Speaker string // имя персонажа; для повествования может быть пустым
	Text    string

	NextID      *NodeID
	Choices     []choice.Choice
	Conditional []ConditionalEdge

	// Version — для оптимистичной блокировки.
	Version int
}
