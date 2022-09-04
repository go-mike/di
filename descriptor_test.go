package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescriptorForType_OnInterface(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(typeOfTestServiceInterface, Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptorForType_OnStructPtr(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(
		typeOfTestServiceStructPtr,
		Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceStructPtr, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptorForType_OnStruct(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(
		typeOfTestServiceStruct,
		Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceStruct, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptor_OnInterface(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[testServiceInterface](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptor_OnStructPtr(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[*testServiceStruct](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceStructPtr, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptor_OnStruct(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[testServiceStruct](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceStruct, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewSingletonServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewSingletonServiceFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewSingletonServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewSingletonServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewSingletonFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewSingletonFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewSingletonFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForType(t *testing.T) {
	descriptor, err := NewSingletonStructForType(
		typeOfTestServiceInterface,
		typeOfTestServiceStruct)

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewSingletonStructForType(
		typeOfTestServiceInterface,
		typeOfInt)

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewSingletonStruct(t *testing.T) {
	descriptor, err := NewSingletonStruct[testServiceInterface, testServiceStruct]()

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewSingletonStruct[testServiceInterface, int64]()

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewScopedServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewScopedServiceFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewScopedServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewScopedServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewScopedFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewScopedFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewScopedFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForType(t *testing.T) {
	descriptor, err := NewScopedStructForType(
		typeOfTestServiceInterface,
		typeOfTestServiceStruct)

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewScopedStructForType(
		typeOfTestServiceInterface,
		typeOfInt)

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewScopedStruct(t *testing.T) {
	descriptor, err := NewScopedStruct[testServiceInterface, testServiceStruct]()

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewScopedStruct[testServiceInterface, int64]()

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewTransientServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewTransientServiceFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewTransientServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewTransientServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewTransientFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewTransientFactoryForType(
		typeOfTestServiceInterface,
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewTransientFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForType(t *testing.T) {
	descriptor, err := NewTransientStructForType(
		typeOfTestServiceInterface,
		typeOfTestServiceStruct)

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewTransientStructForType(
		typeOfTestServiceInterface,
		typeOfInt)

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewTransientStruct(t *testing.T) {
	descriptor, err := NewTransientStruct[testServiceInterface, testServiceStruct]()

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewTransientStruct[testServiceInterface, int64]()

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewInstanceForType(t *testing.T) {
	instance := &testServiceStruct{}
	descriptor, err := NewInstanceForType(
		typeOfTestServiceInterface,
		instance)

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewInstanceForType_OnNil(t *testing.T) {
	descriptor, err := NewInstanceForType(
		typeOfTestServiceInterface,
		nil)

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}

func TestNewInstance(t *testing.T) {
	instance := &testServiceStruct{}
	descriptor, err := NewInstance[testServiceInterface](instance)

	assert.NoError(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, typeOfTestServiceInterface, descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewInstance_OnNil(t *testing.T) {
	descriptor, err := NewInstance[testServiceInterface](nil)

	assert.Error(t, err)
	assert.Nil(t, descriptor)
}
