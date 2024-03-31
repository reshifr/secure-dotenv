package crypto

type CSPRNGError int

const (
	ErrReadEntropyFailed CSPRNGError = iota + 1
)

func (err CSPRNGError) Error() string {
	switch err {
	case ErrReadEntropyFailed:
		return "ErrReadEntropyFailed: " +
			"Failed to read a random value from the entropy sources."
	default:
		return "Error: unknown."
	}
}

type FnCSPRNG struct {
	Read func(b []byte) (n int, err error)
}

type CSPRNG interface {
	Read(block []byte) (err error)
	Make(blockLen int) (block []byte, err error)
}