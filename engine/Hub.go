package engine

type Status int

const (
	SENDMESSAGE Status = iota
	SETSTATUSSENDER
)

type Hub struct {
	Data   string
	Status Status
}
