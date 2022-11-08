package otp

func (e value) Code() string {
	switch e {
	case INCORRECT:
		return "ROA_422_004_XXX"
	default:
		return "UNKNOWN"
	}
}
