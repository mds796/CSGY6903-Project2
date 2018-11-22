package crypto

type FingerPrint interface {
	Sign(data []byte) ([]byte, error)
}
