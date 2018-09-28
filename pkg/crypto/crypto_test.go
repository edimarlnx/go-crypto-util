package crypto

import (
	"fmt"
	"os"
	"testing"
)

var libTest *UtilCrypto

const DecryptedDataTest = "go crypto data"
const EncryptedDataTest = "FHZRYMU9aYhEQ5zHFBuVYsXsIcmSDPj9i/kqMLU76yS5jNrNb0WG1/MV3cLM5x7sAAB2kw2S20xS3XTnpi1bnXy9Hp92TJx3zjKjfzKEZ0nxR5VYtwjD1EwAsC/HC6oNa+ZBU5ivUV3APyBm58wC22Vf8XOy15+jLGDL4FS09ISY63uYH3eOlmAINMsW0XcIaZUZ/hmAHnSCio25MqNccMTZ9ZXmAGP+cva18uFOLgpVrVKUey9t2lHehdccTGOZanX4XihqL9BKCoKEPPAx/RIcxOFSr9qSB48OwT1Wnx4G+CFlHHfwO/Bn1/J8XsbJM5+26qZ84MhJxtDW1IeU1w=="

func TestDecrypt(t *testing.T) {
	value, err := libTest.Decrypt(EncryptedDataTest)
	if err != nil {
		fmt.Println(value)
		t.Error(err)
	}
	if DecryptedDataTest != value {
		t.Errorf("Expected '%s', actual '%s'", DecryptedDataTest, value)
	}
}

func TestEncrypt(t *testing.T) {
	value, err := libTest.Encrypt(DecryptedDataTest)
	if err != nil {
		fmt.Println(value)
		t.Error(err)
	}
	actualDec, _ := libTest.Decrypt(value)
	if DecryptedDataTest != actualDec {
		t.Errorf("Expected '%s', actual '%s'", DecryptedDataTest, actualDec)
	}
}

func TestMain(m *testing.M) {
	libTest = &UtilCrypto{
		PrivKey: "../../tests/privkey.pem",
		PubKey:  "../../tests/key.pub",
	}
	code := m.Run()
	os.Exit(code)
}
