package expbo

import (
	"context"
	"math"
	"time"
)

const (
	BASE = 1_000
	CAP  = 10_000
)

type ExponentialBackoffJitter struct {
	algorithm     Algorithm
	base          int
	cap           int
	ctx           context.Context
	maxRetryCount int
	prev          int
	retryCount    int
	rng           RandomNumberGenerator
}

func NewExponentialBackoffJitter(
	ctx context.Context,
	maxRetryCount int,
	options ...Option,
) ExponentialBackoffJitter {
	e := ExponentialBackoffJitter{
		algorithm:     ExponentialBackoffFullJitter,
		base:          BASE,
		cap:           CAP,
		ctx:           ctx,
		maxRetryCount: maxRetryCount,
		prev:          BASE,
		retryCount:    0,
		rng:           NewRandomNumberGeneratorImpl(),
	}

	// Functional Option Pattern
	for _, opt := range options {
		opt(&e)
	}

	return e
}

func (e *ExponentialBackoffJitter) Retry() <-chan time.Duration {
	ch := make(chan time.Duration)

	go func() {
		defer close(ch)

		for i := 0; i < e.maxRetryCount; i++ {
			select {
			case <-e.ctx.Done():
				return
			default:
				t := e.generate()
				e.prev = t
				e.retryCount = e.retryCount + 1
				ch <- time.Duration(t)
			}
		}
	}()

	return ch
}

func (e ExponentialBackoffJitter) generate() int {
	var t int

	switch e.algorithm {
	case ExponentialBackoffFullJitter:
		t = e.generateExponentialBackoffFullJitter()
	case ExponentialBackoffEqualJitter:
		t = e.generateExponentialBackoffEqualJitter()
	case ExpoentialBackoffDecorrelatedJitter:
		t = e.generateExponentialBackoffDecorrelatedJitter()
	default:
		t = e.base
	}

	return t
}

// Exponential Backoff And Full Jitter
// https://zenn.dev/sinozu/articles/5c0457876be42e#exponential-backoff-and-full-jitter
//
// sleep = random_between(0 min(cap, base * 2 ** attempt))
func (e ExponentialBackoffJitter) generateExponentialBackoffFullJitter() int {
	t := e.base * int(math.Pow(2, float64(e.retryCount)))
	if t > e.cap {
		t = e.cap
	}

	return e.rng.Intn(t)
}

// Exponential Backoff And Equal Jitter
// https://zenn.dev/sinozu/articles/5c0457876be42e#exponential-backoff-and-equal-jitter
//
// temp = min(cap, base * 2 ** attempt)
// sleep = temp / 2 + random_between(0, temp / 2)
func (e ExponentialBackoffJitter) generateExponentialBackoffEqualJitter() int {
	t := e.base * int(math.Pow(2, float64(e.retryCount)))
	if t > e.cap {
		t = e.cap
	}

	return t/2 + e.rng.Intn(t/2)
}

// Exponential Backoff And Decorrelated Jitter
// https://zenn.dev/sinozu/articles/5c0457876be42e#exponential-backoff-and-decorrlated-jitter
//
// sleep = min(cap, random_between(base, sleep * 3))
func (e ExponentialBackoffJitter) generateExponentialBackoffDecorrelatedJitter() int {
	t := e.rng.Intn((e.prev*3)-e.base) + e.base
	if t > e.cap {
		t = e.cap
	}

	return t
}
