package crypto

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"hash"
	"io/ioutil"
)

//UtilCrypto encrypt and decrypt lib
type UtilCrypto struct {
	PrivKey string
	PubKey  string
	hash    hash.Hash
}

func (c *UtilCrypto) getHash() hash.Hash {
	if c.hash == nil {
		c.hash = sha256.New()
	}
	return c.hash
}

func (c *UtilCrypto) loadKey(privkey bool) ([]byte, error) {
	var keyFile string
	if privkey {
		keyFile = c.PrivKey
	} else {
		keyFile = c.PubKey
	}
	bytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getPrivKey(c *UtilCrypto) (*rsa.PrivateKey, error) {
	privKey, err := c.loadKey(true)
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

func getPubKey(c *UtilCrypto) (*rsa.PublicKey, error) {
	pubKey, err := c.loadKey(false)
	if err != nil {
		return nil, err
	}
	var pubRsa *rsa.PublicKey
	pubKeyBlock, _ := pem.Decode(pubKey)
	pubInterface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	pubRsa = pubInterface.(*rsa.PublicKey)

	return pubRsa, nil
}
