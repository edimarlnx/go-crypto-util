package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// Decrypt decrypt data with RSA key
func (c *UtilCrypto) Decrypt(encryptedData string) (string, error) {
	privRsa, err := getPrivKey(c)
	if err != nil {
		return "", err
	}
	decodedData, _ := base64.StdEncoding.DecodeString(encryptedData)
	decryptedData, decryptErr := rsa.DecryptPKCS1v15(rand.Reader, privRsa, decodedData)
	if decryptErr != nil {
		fmt.Println("Decrypt data error")
		panic(decryptErr)
	}
	return string(decryptedData), nil
}
