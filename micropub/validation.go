package micropub

// ContentType the various content types available
type ContentType int

const (
	// WwwForm form content type
	WwwForm ContentType = iota

	// JSON json content type
	JSON

	// MultiPart multi-part form type
	MultiPart

	// UnsupportedType content type not supported
	UnsupportedType
)
