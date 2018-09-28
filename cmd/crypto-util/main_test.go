package main

import (
	"flag"
	"os"
	"testing"
)

const (
	stringTest  = "teste !@"
	testPrivKey = "../../tests/privkey.pem"
	testPubKey  = "../../tests/key.pub"
)

func TestEncrypt(t *testing.T) {
	flag.Set("encrypt", stringTest)
	encData, err := appMain("encrypt", stringTest, testPrivKey, testPubKey)
	if err != nil {
		t.Error(err)
	}
	if encData == "" {
		t.Errorf("No data for process.")
	}

	decData, err := appMain("decrypt", encData, testPrivKey, testPubKey)
	if err != nil {
		t.Error(err)
	}
	if decData == "" {
		t.Errorf("No data for process.")
	}
	if stringTest != decData {
		t.Errorf("Expected %s, actual %s", stringTest, decData)
	}
}
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
