package clip

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
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
	Token.StringValue = hex.EncodeToString(md)
	Token.IntValue = crc32.ChecksumIEEE(md)
    fmt.Printf("This is my value %d\n", Token.IntValue)
    return
}

