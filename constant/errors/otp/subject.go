package otp

func (e value) Subject() string {
	switch e {
	case Incorrect:
		return "OTP_INCORRECT"
	default:
		return "UNKNOWN"
	}
}
