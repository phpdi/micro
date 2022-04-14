package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumk1j9HD2B8sVy = []byte("k1j9HD2B8sVy")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesk1j9HD2B8sVy(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumk1j9HD2B8sVy
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
