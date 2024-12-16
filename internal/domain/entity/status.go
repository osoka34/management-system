package entity

type Status int

const (
	StatusUndefined Status = iota
	StatusCreate
	StatusClosed
	// StatusInProgress
	// StatusDone
)
