package any

import "math/rand"

// Int returns any one of the int's provided, or a random
// one is none are provided.
func Int(list ...int) int {
	if len(list) < 1 {
		return rand.Int()
	}
	return list[rand.Intn(len(list))]
}

// Between generates a random int between min and max.
func Between(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		i := rand.Intn(min-max) + max
		return i * -1
	}
	return rand.Intn(max-min) + min
}
