package any

import "math/rand"

func One(list ...interface{}) interface{} {
	if len(list) < 1 {
		return nil
	}
	return list[rand.Intn(len(list))]
}
