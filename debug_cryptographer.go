package cryptography

import (
	util "github.com/mirogon/go_util"
)

type DebugCryptographer struct {
}

func (tokenGenerator DebugCryptographer) GenerateSessionToken() string {
	return "123"
}
func (tokenGenerator DebugCryptographer) SaltAndHashPassword(pwRaw string) (string, string) {
	return "password", "salt"
}
func (crypto DebugCryptographer) SaltAndHashPasswordWithSalt(pwRaw string, salt string) string {
	pwSalted := pwRaw + salt
	hashedAndSaltedPw := util.HashString(pwSalted)
	return hashedAndSaltedPw
}

func (crypto DebugCryptographer) GenerateRandomConfirmationCode() string {
	return "123456"
}

func (crypto DebugCryptographer) GenerateUniqueId() uint64 {
	return uint64(10)
}
