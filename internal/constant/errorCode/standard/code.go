package standard

func (e *ErrorCode) Code() string {
	switch Value(e.Value) {
	case BAD_REQUEST:
		return "ROA_500_001_XXX"
	case INTERNAL:
		return "ROA_500_001_XXX"
	default:
		return ""
	}
}
