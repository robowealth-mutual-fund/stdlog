package otp

func (e *ErrorCode) Code() string {
	switch Value(e.Value) {
	case INCORRECT:
		return "ROA_422_004_XXX"
	default:
		return ""
	}
}
