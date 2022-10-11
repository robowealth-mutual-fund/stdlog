package errorCode

func Code() string {
	return ""
}

func Subject() string {
	return ""
}

type Interface interface {
	Code() string
	Subject() string
}
