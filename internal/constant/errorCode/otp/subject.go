package otp

func (e *ErrorCode) Subject() string {
	switch Value(e.Value) {
	case INCORRECT:
		return "OTP_INCORRECT"
	default:
		return ""
	}
}
