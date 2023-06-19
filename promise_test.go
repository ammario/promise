package promise

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	t.Parallel()

	beforeGo := time.Now()
	p := Go(func() (int, error) {
		time.Sleep(time.Second)
		return 1000, nil
	})
	// Should launch instantaneously
	require.WithinDuration(t, beforeGo, time.Now(), time.Millisecond)

	v, err := p.Resolve()
	require.Equal(t, 1000, v)
	require.NoError(t, err)
	afterResolve := time.Now()

	// Should launch instantaneously
	require.WithinDuration(t, afterResolve, beforeGo.Add(time.Second), time.Millisecond*10)
}

func TestPromise_CatchPanic(t *testing.T) {
	t.Parallel()

	p := Go(func() (int, error) {
		panic("oops")
	})

	v, err := p.Resolve()
	require.Equal(t, 0, v)
	require.Error(t, err)
	require.Contains(t, err.Error(), "panic: oops")
}

func TestInstant(t *testing.T) {
	t.Parallel()

	p := Instant(1000, nil)
	v, err := p.Resolve()
	require.Equal(t, 1000, v)
	require.NoError(t, err)
}
