package otp

func (e Value) Subject() string {
	switch e {
	case INCORRECT:
		return "OTP_INCORRECT"
	default:
		return ""
	}
}
