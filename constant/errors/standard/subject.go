package standard

func (e value) Subject() string {
	switch e {
	case BadRequest:
		return "BAD_REQUEST"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "NOT_FOUND"
	case UnsupportedMediaType:
		return "UNSUPPORTED_MEDIA_TYPE"
	case UnprocessableEntity:
		return "UNPROCESSABLE_ENTITY"
	case InternalServerError:
		return "INTERNAL_SERVER_ERROR"
	case GatewayTimeout:
		return "GATEWAY_TIMEOUT"
	default:
		return "UNKNOWN"
	}
}
