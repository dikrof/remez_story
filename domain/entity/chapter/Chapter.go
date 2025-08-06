package chapter

// Chapter — сущность главы. Служит контейнером для узлов и метаинформации.
type Chapter struct {
	ID          ChapterID
	Title       string
	Description string
}
