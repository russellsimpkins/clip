package clip

import (
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"math/rand"
	"time"
	"hash/crc32"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func AddToken(token Token) (err error, token1 Token) {
	return
}

func DeleteToken(token Token) (err error) {
	return
}

func GenerateToken() (token Token) {

	rndString := make([]byte, 10)
    for i:= 0 ; i < 10 ; i++ {
        rndString[i] = byte(randInt(65,90))
    }

	hash := sha256.New()
	hash.Write(rndString)
	md := hash.Sum(nil)
	token.StringValue = hex.EncodeToString(md)
	token.IntValue = crc32.ChecksumIEEE(md)
    return
}

