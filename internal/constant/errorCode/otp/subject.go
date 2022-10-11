package otp

func (e Error) Subject() string {
	switch e {
	case INCORRECT:
		return "OTP_INCORRECT"
	default:
		return ""
	}
}
