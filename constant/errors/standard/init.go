package standard

type value int

const (
	BadRequest value = iota
	Unauthorized
	Forbidden
	NotFound
	UnsupportedMediaType
	UnprocessableEntity
	InternalServerError
	GatewayTimeout
)
