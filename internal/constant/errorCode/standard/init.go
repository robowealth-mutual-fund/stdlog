package standard

type Error int

const (
	INTERNAL Error = iota
	BAD_REQUEST
)
