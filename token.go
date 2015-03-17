package clip

import (
	"crypto/sha256"
	"encoding/hex"
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


/*
 * Taken right from the docs. This is how you populate a map of maps
 */
func add(m map[string]map[string]int, env, feature string, val int) {
    mm, ok := m[env]
    if !ok {
        mm = make(map[string]int)
        m[env] = mm
    }
    mm[feature] = val
}
