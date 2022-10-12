package standard

func (e Value) Code() string {
	switch e {
	case BAD_REQUEST:
		return "ROA_500_001_XXX"
	case INTERNAL:
		return "ROA_500_001_XXX"
	default:
		return ""
	}
}
