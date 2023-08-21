package cryptography_test

import (
	"testing"

	cryptography "github.com/mirogon/go_cryptography"
	util "github.com/mirogon/go_util"
)

func TestGenerateSessionToken(t *testing.T) {
	crypt := cryptography.CryptographerImpl{}
	token := crypt.GenerateSessionToken()
	if len(token) != 16*2 {
		t.Error()
	}
}

func TestHashSessionToken(t *testing.T) {
	crypt := cryptography.CryptographerImpl{}
	sessionToken := crypt.GenerateSessionToken()
	result := util.HashString(sessionToken)
	if len(result) != 32*2 {
		t.Error()
	}
}
func TestSaltAndHashPw(t *testing.T) {
	pwSaltedHashed, salt := cryptography.CRYPTO.SaltAndHashPassword("Hello")
	manual := util.HashString("Hello" + salt)
	if manual != pwSaltedHashed {
		t.Error()
	}
}

func TestGenerateRandomString(t *testing.T) {
	r := cryptography.GenerateRandomBytesAsHexString(16)
	if len(r) != 32 {
		t.Error()
	}
}

func TestGenerateUniqueId(t *testing.T) {
	id := cryptography.CRYPTO.GenerateUniqueId()
	if id > 9007199254740991 {
		t.Error(id)
	}
	//t.Error(id)
}

func TestWhatev(t *testing.T) {
	var lastPart uint32 = 4111111111
	lastPart = lastPart & 0xffffff00
	//t.Error(lastPart)
}
