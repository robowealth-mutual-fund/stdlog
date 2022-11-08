package standard

func (e value) Subject() string {
	switch e {
	case BAD_REQUEST:
		return "BAD_REQUEST"
	case UNAUTHORIZED:
		return "UNAUTHORIZED"
	case FORBIDDEN:
		return "FORBIDDEN"
	case NOT_FOUND:
		return "NOT_FOUND"
	case UNSUPPORTED_MEDIA_TYPE:
		return "UNSUPPORTED_MEDIA_TYPE"
	case UNPROCESSABLE_ENTITY:
		return "UNPROCESSABLE_ENTITY"
	case INTERNAL_SERVER_ERROR:
		return "INTERNAL_SERVER_ERROR"
	case GATEWAY_TIMEOUT:
		return "GATEWAY_TIMEOUT"
	default:
		return "UNKNOWN"
	}
}
