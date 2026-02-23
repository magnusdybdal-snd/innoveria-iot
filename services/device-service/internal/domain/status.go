package domain

type Status int

const (
	StatusOnline Status = iota
	StatusNeverSeen
	StatusOffline
)
