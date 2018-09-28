package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/edimarlnx/go-crypto-util/pkg/crypto"
)

var (
	flgMode        string
	flgData        string
	flgPrivKeyFile string
	flgPubKeyFile  string
)

func parseFlags() {
	flag.StringVar(&flgMode, "mode", "", "Use -mode=decrypt to decrypt data or -mode=encrypt")
	flag.StringVar(&flgData, "data", "", "Use -data=DATA_TO_PROCESS to pass a data")
	flag.StringVar(&flgPrivKeyFile, "priv-key-file", "", "Use -priv-key-file=/file/to/private/rsa/key ")
	flag.StringVar(&flgPubKeyFile, "pub-key-file", "", "Use -priv-key-file=/file/to/private/rsa/key ")
	flag.Parse()
}

func appMain(mode, data, privKey, pubKey string) (string, error) {

	cryptoApp := &crypto.UtilCrypto{
		PrivKey: privKey,
		PubKey:  pubKey,
	}
	var dataOut string
	var err error
	switch mode {
	case "decrypt":
		dataOut, err = cryptoApp.Decrypt(data)
	case "encrypt":
		dataOut, err = cryptoApp.Encrypt(data)
	}

	if err != nil {
		return "", err
	}
	return dataOut, nil
}

func main() {
	parseFlags()
	if flgMode == "" || flgData == "" {
		panic("No -mode=encrypt|decrypt and -data=DATA passed")
	}

	privKeyFile := flgPrivKeyFile
	if flgPrivKeyFile == "" {
		privKeyFile = "/privkey.pem"
	}

	pubKeyFile := flgPubKeyFile
	if flgPrivKeyFile == "" {
		pubKeyFile = "/key.pub"
	}
	data, err := appMain(flgMode, flgData, privKeyFile, pubKeyFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "%s\n", data)
}
