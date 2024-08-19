package auth

import "math/rand"

func makeLoginKey() uint32 {
	return uint32(rand.Intn(int(^uint32(0)>>1)) + 1)
}

func Login(username, password string, res chan uint32) {
	success := true
	if success {
		res <- makeLoginKey()
	} else {
		res <- 0
	}
}

func LoginByKey(key uint32, res chan bool) {
	res <- true
}
