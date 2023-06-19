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
