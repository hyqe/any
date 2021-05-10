package any

import (
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
)

// String returns any string
func String(list ...string) string {
	if len(list) < 1 {
		return ""
	}
	return list[rand.Intn(len(list))]
}

// Bytes generates a random byte array with a random length.
// if no length is specified, a random length will be generated.
func Bytes(lengths ...int) []byte {
	var x []byte
	if len(lengths) > 0 {
		x = make([]byte, len(lengths))
	} else {
		x = make([]byte, Between(20, 100))
	}

	crand.Read(x)
	return x
}

func Hex(lengths ...int) string {
	return hex.EncodeToString(Bytes(lengths...))
}

func Base64(lengths ...int) string {
	return base64.StdEncoding.EncodeToString(Bytes(lengths...))
}
