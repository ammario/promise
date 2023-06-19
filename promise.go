// Package promise implements a basic promise construct. Its goal is
// to increase concurrency by making it easy to model lazy computations.
package promise

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

// Go calls fn in a goroutine and promises that the value will be available
// in the future via Resolve.
func Go[T any](fn func() (T, error)) *Promise[T] {
	p := &Promise[T]{
		doneCh: make(chan struct{}),
	}
	go func() {
		p.v, p.err = fn()
		close(p.doneCh)
	}()
	return p
}
