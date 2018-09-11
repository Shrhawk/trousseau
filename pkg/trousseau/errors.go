package trousseau

// Dial errors.
const (
	ErrSectionNotFound      = Error("section not found")
	ErrSectionAlreadyExists = Error("section already exists")
)

const (
	ErrItemNotFound      = Error("item not found")
	ErrItemAlreadyExists = Error("item already exists")
)

// Error represents a WTF error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
