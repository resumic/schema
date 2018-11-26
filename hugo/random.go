package hugo

import "math/rand"

const randomLetters = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	b := make([]byte, n)
	l := len(randomLetters)
	for i := range b {
		b[i] = randomLetters[rand.Intn(l)]
	}
	return string(b)
}
