package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// Encrypt encrypt data with RSA key
func (c *UtilCrypto) Encrypt(dataToEncrypt string) (string, error) {
	privRsa, err := getPrivKey(c)
	if err != nil {
		return "", err
	}
	publicRsa := &privRsa.PublicKey
	decryptedData, decryptErr := rsa.EncryptPKCS1v15(rand.Reader, publicRsa, []byte(dataToEncrypt))
	if decryptErr != nil {
		fmt.Println("Decrypt data error")
		panic(decryptErr)
	}
	decodedData := base64.StdEncoding.EncodeToString(decryptedData)
	
	return decodedData, nil
}
