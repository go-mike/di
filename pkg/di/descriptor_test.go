package di

import (
	"reflect"
	"testing"
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
	if descriptor == nil {
		t.Error("descriptor is nil")
	}
	// Then
	if descriptor.Lifetime() != Scoped {
		t.Error("descriptor lifetime is not Scoped")
	}
	// Then
	if descriptor.ServiceType().Name() != "testServiceInterface" {
		t.Error("descriptor service type is not testServiceInterface but", descriptor.ServiceType().Name())
	}
	// Then
	if descriptor.Factory() != &factory {
		t.Error("descriptor factory is not &factory but", descriptor.Factory())
	}
}

func TestNewDescriptor(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewDescriptor[testServiceInterface](Scoped, &factory)

	// Then
	if descriptor == nil {
		t.Error("descriptor is nil")
	}
	// Then
	if descriptor.Lifetime() != Scoped {
		t.Error("descriptor lifetime is not Scoped")
	}
	// Then
	if descriptor.ServiceType().Name() != "testServiceInterface" {
		t.Error("descriptor service type is not testServiceInterface but", descriptor.ServiceType().Name())
	}
	// Then
	if descriptor.Factory() != &factory {
		t.Error("descriptor factory is not &factory but", descriptor.Factory())
	}
}

func TestNewSingleton(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewSingleton[testServiceInterface](&factory)

	// Then
	if descriptor == nil {
		t.Error("descriptor is nil")
	}
	// Then
	if descriptor.Lifetime() != Singleton {
		t.Error("descriptor lifetime is not Singleton")
	}
	// Then
	if descriptor.ServiceType().Name() != "testServiceInterface" {
		t.Error("descriptor service type is not testServiceInterface but", descriptor.ServiceType().Name())
	}
	// Then
	if descriptor.Factory() != &factory {
		t.Error("descriptor factory is not &factory but", descriptor.Factory())
	}
}

func TestNewScoped(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewScoped[testServiceInterface](&factory)

	// Then
	if descriptor == nil {
		t.Error("descriptor is nil")
	}
	// Then
	if descriptor.Lifetime() != Scoped {
		t.Error("descriptor lifetime is not Scoped")
	}
	// Then
	if descriptor.ServiceType().Name() != "testServiceInterface" {
		t.Error("descriptor service type is not testServiceInterface but", descriptor.ServiceType().Name())
	}
	// Then
	if descriptor.Factory() != &factory {
		t.Error("descriptor factory is not &factory but", descriptor.Factory())
	}
}

func TestNewTransient(t *testing.T) {
	// Given
	factory := testServiceFactoryFunc{}
	// When
	descriptor := NewTransient[testServiceInterface](&factory)

	// Then
	if descriptor == nil {
		t.Error("descriptor is nil")
	}
	// Then
	if descriptor.Lifetime() != Transient {
		t.Error("descriptor lifetime is not Transient")
	}
	// Then
	if descriptor.ServiceType().Name() != "testServiceInterface" {
		t.Error("descriptor service type is not testServiceInterface but", descriptor.ServiceType().Name())
	}
	// Then
	if descriptor.Factory() != &factory {
		t.Error("descriptor factory is not &factory but", descriptor.Factory())
	}
}
