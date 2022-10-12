package standard

func (e value) Subject() string {
	switch e {
	case BAD_REQUEST:
		return "BAD_REQUEST"
	case INTERNAL:
		return "INTERNAL"
	default:
		return ""
	}
}
