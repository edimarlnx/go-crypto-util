package crypto
import (
	"fmt"
	"testing"
	"os"
)

var libTest *UtilCrypto
const DecryptedDataTest = "go crypto data"
const EncryptedDataTest = "s/R0h1ptoPuN+RW4aMXg5VP/QIWSUDYLc+p9CeTU4eFHuyDi84xJFWGwWGVu6z26CpMiTUtDfOHbC9Bfvl9YHW+/YDi1+o1hpYriKxQT9jyGkHsoI0t9KiCdL9dtnI4AhyIncuhvJUuT+d3EkZkijkFNIhF/VEZfVfPT9O6glKPRxhWOCNeCG+5XxsvFACQWCrw2VzUV5vmMXDIYDEPvrTh/oPPeU/rciyQu6VF6jOKKljDNIhVhkmk7U1+5SN2MfzUjWMt3DCZq6EjCfovweBAv/1NHFT7Et7iL93szlwOYxzbSo34map3vTCtqPE1Rw5gnnj4/DU3dqJ1BipR4yQ=="

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
	if DecryptedDataTest !=  actualDec {
		t.Errorf("Expected '%s', actual '%s'", DecryptedDataTest, actualDec)
	}
}


func TestMain(m *testing.M) {
	libTest = &UtilCrypto{
		PrivKey: "../../tests/private_nopass_test.pem",
	}
	code := m.Run()
	os.Exit(code)
}