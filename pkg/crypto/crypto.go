package crypto

import (
	"io/ioutil"
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
)

//UtilCrypto encrypt and decrypt lib
type UtilCrypto struct {
	PrivKey string
}

func (c *UtilCrypto) loadPrivKey() ([]byte, error) {
	bytes, err := ioutil.ReadFile(c.PrivKey)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getPrivKey(c *UtilCrypto) (*rsa.PrivateKey, error) {
	privKey, err := c.loadPrivKey()
	if err != nil {
		return nil, err
	}
	privateKeyBlock, _ := pem.Decode(privKey)
	privRsa, parseErr := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if parseErr != nil {
		return nil, err
	}
	return privRsa, nil
}