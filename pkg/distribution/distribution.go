package distribution

import (
	"math"
	"math/rand"
	"time"
)

func NewBeta(success, failed uint, seed ...int64) *Beta {
	var s int64
	if len(seed) == 0 {
		s = time.Now().UnixNano()
	} else {
		s = seed[0]
	}
	b := &Beta{
		alpha: float64(success + 1),
		beta:  float64(failed + 1),
		rand:  rand.New(rand.NewSource(s)), //nolint:gosec // криптографическая стойкость не требуется
	}

	return b
}

func (b *Beta) Update(clicked bool) {
	if clicked {
		b.alpha++
	} else {
		b.beta++
	}
}

func (b *Beta) Success() uint {
	return uint(b.alpha - 1)
}

func (b *Beta) Failed() uint {
	return uint(b.beta - 1)
}

func (b *Beta) SetSuccess(success uint) {
	b.alpha = float64(success + 1)
}

func (b *Beta) SetFailed(failed uint) {
	b.beta = float64(failed + 1)
}

func (b *Beta) Sample() float64 {
	x := gammaRand(b.alpha, b.rand)
	y := gammaRand(b.beta, b.rand)
	return x / (x + y)
}

// Генерация случайного число из распределения Γамма(alpha, 1).
func gammaRand(alpha float64, src *rand.Rand) float64 {
	if alpha <= 0 {
		return math.NaN()
	}

	// Marsaglia & Tsang method (2000)
	d := alpha - 1.0/3.0
	c := 1.0 / math.Sqrt(9*d)

	for {
		var x, v float64
		for {
			x = src.NormFloat64()
			v = 1 + c*x
			if v > 0 {
				break
			}
		}
		v = v * v * v
		u := src.Float64()
		if u < 1-0.0331*x*x*x*x {
			return d * v
		}
		if math.Log(u) < 0.5*x*x+d*(1-v+math.Log(v)) {
			return d * v
		}
	}
}
