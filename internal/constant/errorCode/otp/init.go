package otp

type Value int
type ErrorCode struct {
	Value
}

const (
	INCORRECT Value = iota
)
