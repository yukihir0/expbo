package expbo

type Algorithm int

const (
	ExponentialBackoffFullJitter Algorithm = iota
	ExponentialBackoffEqualJitter
	ExpoentialBackoffDecorrelatedJitter
)
