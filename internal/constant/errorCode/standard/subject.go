package standard

func (e *ErrorCode) Subject() string {
	switch Value(e.Value) {
	case BAD_REQUEST:
		return "BAD_REQUEST"
	case INTERNAL:
		return "INTERNAL"
	default:
		return ""
	}
}
