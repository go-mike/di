package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testServiceInterface interface{}
type testServiceFactoryFunc struct{}
func (*testServiceFactoryFunc) Create(provider ServiceProvider) (interface{}, error) {
	panic("unimplemented")
}
func (*testServiceFactoryFunc) DisplayName() string {
	panic("unimplemented")
}
func (*testServiceFactoryFunc) Requirements() []reflect.Type {
	panic("unimplemented")
}

func TestNewDescriptorForType(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewDescriptorForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		Scoped,
		&factory)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.Equal(t, &factory, descriptor.Factory())
}

func TestNewDescriptor(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewDescriptor[testServiceInterface](Scoped, &factory)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.Equal(t, &factory, descriptor.Factory())
}

func TestNewSingleton(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewSingleton[testServiceInterface](&factory)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.Equal(t, &factory, descriptor.Factory())
}

func TestNewScoped(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewScoped[testServiceInterface](&factory)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.Equal(t, &factory, descriptor.Factory())
}

func TestNewTransient(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewTransient[testServiceInterface](&factory)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.Equal(t, &factory, descriptor.Factory())
}
