package crypto

const (
	ErrEmptyPassphrase = Error("Empty passphrase is not allowed")
	ErrEmptyPayload    = Error("Empty payload is not allowed")
)

// Error is a helper type to wrap crate specific errors
type Error string

func (e Error) Error() string {
	return string(e)
}
