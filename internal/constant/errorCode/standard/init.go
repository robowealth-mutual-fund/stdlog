package standard

type Value int

type ErrorCode struct {
	Value
}

const (
	INTERNAL Value = iota
	BAD_REQUEST
)
