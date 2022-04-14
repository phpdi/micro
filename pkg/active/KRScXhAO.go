package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumKRScXhAO = []byte("KRScXhAO")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesKRScXhAO(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumKRScXhAO
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}
