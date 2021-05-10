package any

import (
	"math/rand"
	"time"
)

// set the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}
