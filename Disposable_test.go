package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDisposable(t *testing.T) {
	disposedCount := 0
	dispose := func() {
		disposedCount++
	}
	disposable := NewDisposable(dispose)
	assert.Equal(t, 0, disposedCount)
	disposable.Dispose()
	assert.Equal(t, 1, disposedCount)
	disposable.Dispose()
	assert.Equal(t, 1, disposedCount)
}

func TestNewNoopDisposable(t *testing.T) {
	disposable := NewNoopDisposable()
	disposable.Dispose()
	// Nothing happens
}
