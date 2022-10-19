package error

type Interface interface {
	Code() string
	Subject() string
	Error() string
}
