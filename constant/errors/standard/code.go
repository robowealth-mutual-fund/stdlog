package standard

func (e value) Code() string {
	switch e {
	case BAD_REQUEST:
		return "ROA_400_001_XXX"
	case UNAUTHORIZED:
		return "ROA_401_001_XXX"
	case FORBIDDEN:
		return "ROA_403_001_XXX"
	case NOT_FOUND:
		return "ROA_404_001_XXX"
	case UNSUPPORTED_MEDIA_TYPE:
		return "ROA_415_001_XXX"
	case UNPROCESSABLE_ENTITY:
		return "ROA_422_001_XXX"
	case INTERNAL_SERVER_ERROR:
		return "ROA_500_001_XXX"
	case GATEWAY_TIMEOUT:
		return "ROA_504_001_XXX"
	default:
		return "UNKNOWN"
	}
}
