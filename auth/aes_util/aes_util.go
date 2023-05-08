package aes_util

import (
	"bytedance_hertz_mod/auth/pkcs_7"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesUtil struct{}

func (p *AesUtil) Decrypt(text, aesKey string) (string, error) {
	aesKeyByte, err := base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(aesKeyByte)
	if err != nil {
		return "", err
	}

	iv := aesKeyByte[:16]
	cipherText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	pkcs7 := pkcs_7.PKCS7Encoder{}
	decrypted := pkcs7.Decode(cipherText)

	networkOrder := decrypted[16:20]
	xmlLength := recoverNetworkBytesOrder(networkOrder)
	xmlContent := string(decrypted[20 : 20+xmlLength])

	return xmlContent, nil
}

func recoverNetworkBytesOrder(orderBytes []byte) int {
	sourceNumber := 0
	for i := 0; i < 4; i++ {
		sourceNumber <<= 8
		sourceNumber |= int(orderBytes[i]) & 0xff
	}
	return sourceNumber
}
