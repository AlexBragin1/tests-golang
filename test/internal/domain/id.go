package domain

import "math/rand"

const alphabet = "0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}
func NewID() string {
	return randString(4)
}
