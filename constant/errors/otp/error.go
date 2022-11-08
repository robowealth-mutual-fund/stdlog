package otp

import "fmt"

func (e value) Error() string {
	switch e {
	case Incorrect:
		return fmt.Sprintf("ROA_422_004_XXX %s", e.Subject())
	default:
		return "UNKNOWN"
	}
}
