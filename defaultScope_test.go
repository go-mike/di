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
	assert.Equal(t, UnknownLifetime, serviceInfo.Lifetime)
	assert.False(t, serviceInfo.IsInstantiated)
	assert.Equal(t, typeOfTestStructWithFields, serviceInfo.ServiceType)

	scope.Dispose()
	assert.True(t, scope.IsDisposed())
}

func TestNewSingletonScope_WithStruct(t *testing.T) {
	descriptor, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor {
		descriptor,
	}
	scope, err := newSingletonScope(descriptors)

	assert.NoError(t, err)
	assert.NotNil(t, scope)
	assert.Equal(t, Singleton, scope.Lifetime())
	assert.False(t, scope.IsDisposed())

	provider := scope.Provider()
	assert.NotNil(t, provider)

	serviceInfo := provider.GetServiceInfo(typeOfTestServiceInterface)
	assert.Equal(t, Singleton, serviceInfo.Lifetime)
	assert.False(t, serviceInfo.IsInstantiated)
	assert.Equal(t, typeOfTestServiceInterface, serviceInfo.ServiceType)

	scope.Dispose()
	assert.True(t, scope.IsDisposed())
}
