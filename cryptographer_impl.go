package cryptography

import (
	"fmt"
	"math/rand"
	"time"

	util "github.com/mirogon/go_util"
)

var CRYPTO Cryptographer = CryptographerImpl{}

type CryptographerImpl struct{}

func (tokenGenerator CryptographerImpl) GenerateSessionToken() string {
	return GenerateRandomBytesAsHexString(16)
}

func (tokenGenerator CryptographerImpl) SaltAndHashPassword(pwRaw string) (string, string) {
	salt := GenerateRandomBytesAsHexString(16)
	return CRYPTO.SaltAndHashPasswordWithSalt(pwRaw, salt), salt
}
func (crypto CryptographerImpl) SaltAndHashPasswordWithSalt(pwRaw string, salt string) string {
	pwSalted := pwRaw + salt
	hashedAndSaltedPw := util.HashString(pwSalted)
	return hashedAndSaltedPw
}

func (crypto CryptographerImpl) GenerateRandomConfirmationCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		randNum := rand.Intn(10)
		code += fmt.Sprintf("%d", randNum)
	}
	return code
}

func (crypto CryptographerImpl) GenerateUniqueId() uint64 {
	now := time.Now()
	var timePart uint32 = uint32(now.Month()<<24) + uint32(now.Day())<<16 + uint32(now.Hour())
	timePart = timePart & 0x000fffff
	lastPart := rand.Uint32()
	return uint64(timePart)<<32 | uint64(lastPart)
}
