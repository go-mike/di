package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSingletonScope_Empty(t *testing.T) {
	descriptors := []ServiceDescriptor{}
	scope, err := newSingletonScope(descriptors)

	assert.NoError(t, err)
	assert.NotNil(t, scope)
	assert.Equal(t, Singleton, scope.Lifetime())
	assert.False(t, scope.IsDisposed())

	provider := scope.Provider()
	assert.NotNil(t, provider)
	serviceInfo := provider.GetServiceInfo(typeOfTestStructWithFields)
	assert.Equal(t, Transient, serviceInfo.Lifetime)

	scope.Dispose()
	assert.True(t, scope.IsDisposed())
}
