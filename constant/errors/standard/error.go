package standard

func (e value) Error() string {
	switch e {
	case BadRequest:
		return "ROA_400_001_XXX"
	case Unauthorized:
		return "ROA_401_001_XXX"
	case Forbidden:
		return "ROA_403_001_XXX"
	case NotFound:
		return "ROA_404_001_XXX"
	case UnsupportedMediaType:
		return "ROA_415_001_XXX"
	case UnprocessableEntity:
		return "ROA_422_001_XXX"
	case InternalServerError:
		return "ROA_500_001_XXX"
	case GatewayTimeout:
		return "ROA_504_001_XXX"
	default:
		return "UNKNOWN"
	}
}
