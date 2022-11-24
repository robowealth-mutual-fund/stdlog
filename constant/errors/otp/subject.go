package otp

func (e value) Subject() string {
	switch e {
	case INCORRECT:
		return "OTP_INCORRECT"
	default:
		return "UNKNOWN"
	}
}
