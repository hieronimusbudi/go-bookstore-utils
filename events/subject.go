package eventsutils

type Subject int

const (
	ItemCreated Subject = iota
	ItemUpdated
)
