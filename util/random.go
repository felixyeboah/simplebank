package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var stringBuilder strings.Builder
	k := len(alpabets)

	for i := 0; i < n; i++ {
		c := alpabets[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}

	return stringBuilder.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
