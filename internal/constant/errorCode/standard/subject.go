package standard

func (e Error) Subject() string {
	switch e {
	case BAD_REQUEST:
		return "BAD_REQUEST"
	case INTERNAL:
		return "INTERNAL"
	default:
		return ""
	}
}
