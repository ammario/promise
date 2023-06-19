// Package promise implements a basic promise construct. Its goal is
// to increase concurrency by making it easy to model lazy computations.
package promise

import "fmt"

// Promise represents a value that will be computed in the future.
type Promise[T any] struct {
	doneCh chan struct{}
	v      T
	err    error
}

// Resolve blocks until the promise is resolved, then returns the value and
// error.
func (p *Promise[T]) Resolve() (T, error) {
	<-p.doneCh
	return p.v, p.err
}

// Instant returns a promise that is already resolved. It is useful in testing
// or when creating an optionally lazy API.
func Instant[T any](v T, err error) *Promise[T] {
	ch := make(chan struct{})
	close(ch)
	return &Promise[T]{
		doneCh: ch,
		v:      v,
		err:    err,
	}
}

// Go calls fn in a goroutine and promises that the value will be available
// in the future via Resolve.
func Go[T any](fn func() (T, error)) *Promise[T] {
	p := &Promise[T]{
		doneCh: make(chan struct{}),
	}
	go func() {
		defer close(p.doneCh)
		defer func() {
			if r := recover(); r != nil {
				p.err = fmt.Errorf("panic: %v", r)
			}
		}()
		p.v, p.err = fn()
	}()
	return p
}
