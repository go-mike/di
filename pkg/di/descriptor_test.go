package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testServiceInterface any
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

func TestNewScopedFactoryForType(t *testing.T) {
	// When
	descriptor := NewScopedFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedFactory(t *testing.T) {
	// When
	descriptor := NewScopedFactory[testServiceInterface](testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForType(t *testing.T) {
	// When
	descriptor, err := NewScopedStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testStructWithFields{}))
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForTypeOnNil(t *testing.T) {
	// When
	descriptor, err := NewScopedStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		nil)
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
}

func TestNewScopedStruct(t *testing.T) {
	// Wh	, en
	descriptor, err := NewScopedStruct[testServiceInterface, testStructWithFields]()
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Scoped, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructOnOtherType(t *testing.T) {
	// Wh	, en
	descriptor, err := NewScopedStruct[testServiceInterface, int]()
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
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

func TestNewTransientFactoryForType(t *testing.T) {
	// When
	descriptor := NewTransientFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientFactory(t *testing.T) {
	// When
	descriptor := NewTransientFactory[testServiceInterface](testFactoryFunc)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForType(t *testing.T) {
	// When
	descriptor, err := NewTransientStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testStructWithFields{}))
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForTypeOnNil(t *testing.T) {
	// When
	descriptor, err := NewTransientStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		nil)
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
}

func TestNewTransientStruct(t *testing.T) {
	// Wh	, en
	descriptor, err := NewTransientStruct[testServiceInterface, testStructWithFields]()
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Transient, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructOnOtherType(t *testing.T) {
	// Wh	, en
	descriptor, err := NewTransientStruct[testServiceInterface, int]()
	// Then
	assert.NotNil(t, err)
	// Then
	assert.Nil(t, descriptor)
}


type testInstanceInterface any
type testInstanceImpl struct{}

func TestNewInstanceForType(t *testing.T) {
	// Given
	instance := testInstanceImpl{}
	// When
	descriptor, err := NewInstanceForType(
		reflect.TypeOf((*testInstanceInterface)(nil)).Elem(),
		&instance)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testInstanceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
	// When
	serviceInstance, err := descriptor.Factory().Create(nil)
	// Then
	assert.Nil(t, err)
	// Then
	assert.Same(t, &instance, serviceInstance.Instance)
}

func TestNewInstanceForTypeOnNilInstance(t *testing.T) {
	// When
	descriptor, err := NewInstanceForType(
		reflect.TypeOf((*testInstanceInterface)(nil)).Elem(),
		nil)
	// Then
	assert.Same(t, ErrInvalidInstance, err)
	// Then
	assert.Nil(t, descriptor)
}

func TestNewInstance(t *testing.T) {
	// Given
	instance := testInstanceImpl{}
	// When
	descriptor, err := NewInstance[testInstanceInterface](&instance)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, descriptor)
	// Then
	assert.Equal(t, Singleton, descriptor.Lifetime())
	// Then
	assert.Equal(t, reflect.TypeOf((*testInstanceInterface)(nil)).Elem(), descriptor.ServiceType())
	// Then
	assert.NotNil(t, descriptor.Factory())
	// When
	serviceInstance, err := descriptor.Factory().Create(nil)
	// Then
	assert.Nil(t, err)
	// Then
	assert.Same(t, &instance, serviceInstance.Instance)
}

func TestNewInstanceOnNilInstance(t *testing.T) {
	// When
	descriptor, err := NewInstance[testInstanceInterface](nil)
	// Then
	assert.Same(t, ErrInvalidInstance, err)
	// Then
	assert.Nil(t, descriptor)
}
