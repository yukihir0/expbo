package expbo

// Functional Option Pattern
type Option func(e *ExponentialBackoffJitter)

func WithAlgorithm(a Algorithm) Option {
	return func(e *ExponentialBackoffJitter) {
		e.algorithm = a
	}
}

func WithBase(base int) Option {
	return func(e *ExponentialBackoffJitter) {
		e.base = base
	}
}

func WithCap(cap int) Option {
	return func(e *ExponentialBackoffJitter) {
		e.cap = cap
	}
}

func WithRandomNumberGenerator(rng RandomNumberGenerator) Option {
	return func(e *ExponentialBackoffJitter) {
		e.rng = rng
	}
}
