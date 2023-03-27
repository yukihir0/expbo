package expbo

import (
	"math/rand"
	"time"
)

type RandomNumberGenerator interface {
	Intn(n int) int
}

// check implement RandomNumberGenerater interface
var _ RandomNumberGenerator = (*RandomNumberGeneratorImpl)(nil)

type RandomNumberGeneratorImpl struct {
	rng *rand.Rand
}

func NewRandomNumberGeneratorImpl() RandomNumberGeneratorImpl {
	return RandomNumberGeneratorImpl{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r RandomNumberGeneratorImpl) Intn(n int) int {
	return r.rng.Intn(n)
}
