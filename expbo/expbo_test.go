package expbo_test

import (
	"context"
	"testing"
	"time"

	"github.com/yukihir0/expbo/expbo"
)

// check implement RandomNumberGenerater interface
var _ expbo.RandomNumberGenerator = (*RandomNumberGeneratorMock)(nil)

type RandomNumberGeneratorMock struct{}

func (rng *RandomNumberGeneratorMock) Intn(n int) int {
	return n
}

func TestGenerateExponentialBackoffFullJitter(t *testing.T) {
	rng := &RandomNumberGeneratorMock{}
	ctx := context.Background()
	maxRetryCount := 5
	e := expbo.NewExponentialBackoffJitter(
		ctx,
		maxRetryCount,
		expbo.WithAlgorithm(expbo.ExponentialBackoffFullJitter),
		expbo.WithBase(1_000),
		expbo.WithCap(10_000),
		expbo.WithRandomNumberGenerator(rng),
	)

	expected := []time.Duration{
		time.Duration(1_000),
		time.Duration(2_000),
		time.Duration(4_000),
		time.Duration(8_000),
		time.Duration(10_000),
	}

	i := 0
	for actual := range e.Retry() {
		if actual != expected[i] {
			t.Errorf("got: %v\nwant: %v", actual, expected[i])
		}
		i = i + 1
	}
}

func TestGenerateExponentialBackoffEqualJitter(t *testing.T) {
	rng := &RandomNumberGeneratorMock{}
	ctx := context.Background()
	maxRetryCount := 5
	e := expbo.NewExponentialBackoffJitter(
		ctx,
		maxRetryCount,
		expbo.WithAlgorithm(expbo.ExponentialBackoffEqualJitter),
		expbo.WithBase(1_000),
		expbo.WithCap(10_000),
		expbo.WithRandomNumberGenerator(rng),
	)

	expected := []time.Duration{
		time.Duration(1_000),
		time.Duration(2_000),
		time.Duration(4_000),
		time.Duration(8_000),
		time.Duration(10_000),
	}

	i := 0
	for actual := range e.Retry() {
		if actual != expected[i] {
			t.Errorf("got: %v\nwant: %v", actual, expected[i])
		}
		i = i + 1
	}
}

func TestGenerateExpoentialBackoffDecorrelatedJitter(t *testing.T) {
	rng := &RandomNumberGeneratorMock{}
	ctx := context.Background()
	maxRetryCount := 5
	e := expbo.NewExponentialBackoffJitter(
		ctx,
		maxRetryCount,
		expbo.WithAlgorithm(expbo.ExpoentialBackoffDecorrelatedJitter),
		expbo.WithBase(1_000),
		expbo.WithCap(10_000),
		expbo.WithRandomNumberGenerator(rng),
	)

	expected := []time.Duration{
		time.Duration(3_000),
		time.Duration(9_000),
		time.Duration(10_000),
		time.Duration(10_000),
		time.Duration(10_000),
	}

	i := 0
	for actual := range e.Retry() {
		if actual != expected[i] {
			t.Errorf("got: %v\nwant: %v", actual, expected[i])
		}
		i = i + 1
	}
}
