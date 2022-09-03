package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewDescriptorForType_OnInterface(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptorForType_OnStructPtr(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(
		reflect.TypeOf((*testServiceStruct)(nil)),
		Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceStruct)(nil)), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptorForType_OnStruct(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptorForType(
		reflect.TypeOf(testServiceStruct{}),
		Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf(testServiceStruct{}), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}


func TestNewDescriptor_OnInterface(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[testServiceInterface](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptor_OnStructPtr(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[*testServiceStruct](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceStruct)(nil)), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewDescriptor_OnStruct(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewDescriptor[testServiceStruct](Scoped, factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf(testServiceStruct{}), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}


func TestNewSingletonServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewSingletonServiceFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewSingletonServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewSingletonServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewSingletonFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewSingletonFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewSingletonFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForType(t *testing.T) {
	descriptor, err := NewSingletonStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testServiceStruct{}))

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewSingletonStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(0))

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}

func TestNewSingletonStruct(t *testing.T) {
	descriptor, err := NewSingletonStruct[testServiceInterface, testServiceStruct]()

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewSingletonStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewSingletonStruct[testServiceInterface, int64]()

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}


func TestNewScopedServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewScopedServiceFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewScopedServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewScopedServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewScopedFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewScopedFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewScopedFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForType(t *testing.T) {
	descriptor, err := NewScopedStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testServiceStruct{}))

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewScopedStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(0))

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}

func TestNewScopedStruct(t *testing.T) {
	descriptor, err := NewScopedStruct[testServiceInterface, testServiceStruct]()

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Scoped, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewScopedStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewScopedStruct[testServiceInterface, int64]()

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}


func TestNewTransientServiceFactoryForType(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewTransientServiceFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewTransientServiceFactory(t *testing.T) {
	factory := &testServiceFactoryFunc{}
	descriptor := NewTransientServiceFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.Equal(t, factory, descriptor.Factory())
}

func TestNewTransientFactoryForType(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewTransientFactoryForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientFactory(t *testing.T) {
	factory := func(provider ServiceProvider) (any, error) {
		panic("not needed for test")
	}
	descriptor := NewTransientFactory[testServiceInterface](factory)

	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForType(t *testing.T) {
	descriptor, err := NewTransientStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(testServiceStruct{}))

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStructForType_OnNonStruct(t *testing.T) {
	descriptor, err := NewTransientStructForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		reflect.TypeOf(0))

		assert.NotNil(t, err)
		assert.Nil(t, descriptor)
	}

func TestNewTransientStruct(t *testing.T) {
	descriptor, err := NewTransientStruct[testServiceInterface, testServiceStruct]()

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Transient, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewTransientStruct_OnNonStruct(t *testing.T) {
	descriptor, err := NewTransientStruct[testServiceInterface, int64]()

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}


func TestNewInstanceForType(t *testing.T) {
	instance := &testServiceStruct{}
	descriptor, err := NewInstanceForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		instance)

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewInstanceForType_OnNil(t *testing.T) {
	descriptor, err := NewInstanceForType(
		reflect.TypeOf((*testServiceInterface)(nil)).Elem(),
		nil)

		assert.NotNil(t, err)
		assert.Nil(t, descriptor)
	}

func TestNewInstance(t *testing.T) {
	instance := &testServiceStruct{}
	descriptor, err := NewInstance[testServiceInterface](instance)

	assert.Nil(t, err)
	assert.NotNil(t, descriptor)
	assert.Equal(t, Singleton, descriptor.Lifetime())
	assert.Equal(t, reflect.TypeOf((*testServiceInterface)(nil)).Elem(), descriptor.ServiceType())
	assert.NotNil(t, descriptor.Factory())
}

func TestNewInstance_OnNil(t *testing.T) {
	descriptor, err := NewInstance[testServiceInterface](nil)

	assert.NotNil(t, err)
	assert.Nil(t, descriptor)
}
