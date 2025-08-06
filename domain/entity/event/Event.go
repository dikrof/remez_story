package event

// Event — элемент справочника событий, влияющих на развилки.
type Event struct {
	ID          EventID
	Code        EventCode
	Title       string
	Description string
	Deprecated  bool
}
