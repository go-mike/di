package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testServiceInterface interface{}
type testServiceFactoryFunc struct{}
func (*testServiceFactoryFunc) Create(provider ServiceProvider) (ServiceInstance, error) {
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

func TestNewDescriptorForTypeExplicitFactory(t *testing.T) {
	// Given
	requirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf((*error)(nil)).Elem(),
	}
	// Given
	displayName := "some service name"
	// When
	descriptor := NewDescriptorForTypeExplicitFactory(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(), Scoped,
		testFactoryFunc, requirements, displayName)

	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}


func TestNewSingletonForType(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewSingletonForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		&factory)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
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

func TestNewSingletonFactoryForType(t *testing.T) {
	// When
	descriptor := NewSingletonFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonFactory(t *testing.T) {
	// When
	descriptor := NewSingletonFactory[testServiceInterface](testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForType(t *testing.T) {
	// When
	descriptor, err := NewSingletonStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testStructWithFields{}))
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForTypeOnNil(t *testing.T) {
	// When
	descriptor, err := NewSingletonStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		nil)
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
}

func TestNewSingletonStruct(t *testing.T) {
	// Wh	, en
	descriptor, err := NewSingletonStruct[testServiceInterface, testStructWithFields]()
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructOnOtherType(t *testing.T) {
	// Wh	, en
	descriptor, err := NewSingletonStruct[testServiceInterface, int]()
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
}


func TestNewScopedForType(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewScopedForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
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


func TestNewTransientForType(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewTransientForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		&factory)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
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
