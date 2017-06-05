package random

import (
	"math/rand"
	"sync"
	"testing"
	"github.com/stretchr/testify/require"
	"io"
)

// for go test -v -race
func TestConcurrency(t *testing.T) {
	rnd := rand.New(NewSource())
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rnd.Int63()
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestReader_Read(t *testing.T) {
	var Reader io.Reader = &reader{rnd: Rand}
	var s []byte = make([]byte, 5, 5)
	n, err := Reader.Read(s)

	require.Equal(t, n, 5)
	require.Nil(t, err)
}