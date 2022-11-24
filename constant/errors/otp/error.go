package otp

import "fmt"

func (e value) Error() string {
	switch e {
	case INCORRECT:
		return fmt.Sprintf("ROA_422_004_XXX %s", e.Subject())
	default:
		return "UNKNOWN"
	}
}
