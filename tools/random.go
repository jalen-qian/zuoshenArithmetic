package tools

import (
	"crypto/rand"
	"math/big"
)

func GetRandom(max int) int {
	random, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	return int(random.Int64())
}
