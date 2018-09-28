package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// Encrypt encrypt data with RSA public key
func (c *UtilCrypto) Encrypt(dataToEncrypt string) (string, error) {
	pubRsa, err := getPubKey(c)
	if err != nil {
		return "", err
	}
	decryptedData, decryptErr := rsa.EncryptOAEP(c.getHash(), rand.Reader, pubRsa, []byte(dataToEncrypt), nil)
	if decryptErr != nil {
		fmt.Println("Decrypt data error")
		panic(decryptErr)
	}
	decodedData := base64.StdEncoding.EncodeToString(decryptedData)
	return decodedData, nil
}
