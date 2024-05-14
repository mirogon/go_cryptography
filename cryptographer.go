package cryptography

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
)

type Cryptographer interface {
	GenerateSessionToken() string
	SaltAndHashPassword(pwRaw string) (string, string)
	SaltAndHashPasswordWithSalt(pwRaw string, salt string) string
	GenerateRandomConfirmationCode() string
	GenerateUniqueId() uint64
	GenerateUniqueIdStr() string
}

// Result is double the length because of the hex conversion
func GenerateRandomBytesAsHexString(bytes int) string {
	var token []byte
	for i := 0; i < bytes; i++ {
		r, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(256))
		if err != nil {
			i -= 1
			continue
		}
		token = append(token, byte(r.Int64()))
	}
	return fmt.Sprintf("%x", token)
}
