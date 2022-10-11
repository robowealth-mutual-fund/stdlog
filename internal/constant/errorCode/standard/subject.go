package standard

func (e Value) Subject() string {
	switch e {
	case BAD_REQUEST:
		return "BAD_REQUEST"
	case INTERNAL:
		return "INTERNAL"
	default:
		return ""
	}
}
