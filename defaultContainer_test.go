package di

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestNewSingletonScope_Empty(t *testing.T) {
// 	descriptors := []ServiceDescriptor{}
// 	scope, err := newDefaultContainer(descriptors, nil)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, scope)
// 	assert.False(t, scope.IsDisposed())

// 	provider := scope.Provider()
// 	assert.NotNil(t, provider)

// 	serviceInfo := provider.GetServiceInfo(typeOfTestServiceInterface)
// 	assert.Equal(t, UnknownLifetime, serviceInfo.Lifetime)
// 	assert.False(t, serviceInfo.IsInstantiated)
// 	assert.Equal(t, typeOfTestServiceInterface, serviceInfo.ServiceType)

// 	scope.Dispose()
// 	assert.True(t, scope.IsDisposed())
// }

// func TestNewSingletonScope_WithStructToInterface(t *testing.T) {
// 	descriptor, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
// 	descriptors := []ServiceDescriptor{
// 		descriptor,
// 	}
// 	scope, err := newDefaultContainer(descriptors, nil)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, scope)
// 	assert.False(t, scope.IsDisposed())

// 	provider := scope.Provider()
// 	assert.NotNil(t, provider)

// 	serviceInfo := provider.GetServiceInfo(typeOfTestServiceInterface)
// 	assert.Equal(t, Singleton, serviceInfo.Lifetime)
// 	assert.False(t, serviceInfo.IsInstantiated)
// 	assert.Equal(t, typeOfTestServiceInterface, serviceInfo.ServiceType)

// 	service, err := provider.GetService(typeOfTestServiceInterface)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, service)
// 	assert.IsType(t, &testServiceStruct{}, service)

// 	scope.Dispose()
// 	assert.True(t, scope.IsDisposed())

// 	serviceInfo = provider.GetServiceInfo(typeOfTestServiceInterface)
// 	assert.Equal(t, UnknownLifetime, serviceInfo.Lifetime)
// 	assert.False(t, serviceInfo.IsInstantiated)
// 	assert.Equal(t, typeOfTestServiceInterface, serviceInfo.ServiceType)
// }
