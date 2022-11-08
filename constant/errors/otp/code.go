package otp

func (e value) Code() string {
	switch e {
	case Incorrect:
		return "ROA_422_004_XXX"
	default:
		return "UNKNOWN"
	}
}
