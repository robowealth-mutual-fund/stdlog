package errorCode

type Value int

func Code() string {
	return ""
}

func Subject() string {
	return ""
}

type ErrorCode interface {
	Code() string
	Subject() string
}
