package distribution

import "math/rand"

type Beta struct {
	alpha float64
	beta  float64
	rand  *rand.Rand
}
