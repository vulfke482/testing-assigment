package rand

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"

	"github.com/valyala/fastrand"
)

// FastRand return definite rand in range
func FastRand(min, max uint32) uint32 {
	if min > max {
		return 0
	}
	return fastrand.Uint32n(max-min) + min
}

// FastRandFloat32 return float32 rand in range
func FastRandFloat32(precision, min, max float32) float32 {
	if precision <= 0 {
		return 0
	}
	return float32(FastRand(uint32(min*precision), uint32(max*precision))) / precision
}

// FastRandFloat64 return float64 rand in range
func FastRandFloat64(precision, min, max float64) float64 {
	if precision <= 0 {
		return 0
	}
	return float64(FastRand(uint32(min*precision), uint32(max*precision))) / precision
}

// CryptoRandomNumber return crypto random int64
func CryptoRandomNumber(min, max int64) int64 {
	if min > max {
		return 0
	}
	nBig, err := crand.Int(crand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return n + min
}

// RandomNumber return random int64
func RandomNumber(min, max int64) int64 {
	if min > max {
		return 0
	}
	mrand.Seed(time.Now().UnixNano())
	return mrand.Int63n(max-min) + min
}

// RandomDigitString return random digit string
func RandomDigitString(length int) (result string) {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = uint8(FastRand(48, 57))
	}
	return string(b)
}
