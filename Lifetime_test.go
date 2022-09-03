package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifetimeString(t *testing.T) {
	// Then
	assert.Equal(t, "Singleton", Singleton.String())
	assert.Equal(t, "Scoped", Scoped.String())
	assert.Equal(t, "Transient", Transient.String())
	assert.Equal(t, "Unknown", (Lifetime(-1)).String())
}