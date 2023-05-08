package pkcs_7

import "bytes"

const (
	blockSize = 32
)

type PKCS7Encoder struct {
}

func (p *PKCS7Encoder) Encode(count int) []byte {
	amountToPad := blockSize - (count % blockSize)
	if amountToPad == 0 {
		amountToPad = blockSize
	}

	padChr := byte(amountToPad)
	padBytes := bytes.Repeat([]byte{padChr}, amountToPad)
	return padBytes
}

func (p *PKCS7Encoder) Decode(decrypted []byte) []byte {
	pad := int(decrypted[len(decrypted)-1])
	if pad < 1 || pad > blockSize {
		pad = 0
	}
	return decrypted[:len(decrypted)-pad]
}
