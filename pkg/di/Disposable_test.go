package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDisposable(t *testing.T) {
	// Given
	disposedCount := 0
	// Given
	dispose := func() {
		disposedCount++
	}
	// Given
	disposable := NewDisposable(dispose)
	// Then
	assert.Equal(t, 0, disposedCount)
	// When
	disposable.Dispose()
	// Then
	assert.Equal(t, 1, disposedCount)
	// When
	disposable.Dispose()
	// Then
	assert.Equal(t, 1, disposedCount)
}

func TestNewNoopDisposable(t *testing.T) {
	// Given
	disposable := NewNoopDisposable()
	// When
	disposable.Dispose()
	// Then
	// Nothing happens
}
