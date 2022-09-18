package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceCollection(t *testing.T) {
	services := NewServiceCollection()
	descriptors := services.ListDescriptors()
	assert.Equal(t, []ServiceDescriptor{}, descriptors)
}

func TestServiceCollection_Add(t *testing.T) {
	services := NewServiceCollection()
	descriptor, _ := NewInstance(&testStructWithFields{})
	newServices := services.Add(descriptor)
	descriptors := services.ListDescriptors()
	assert.Equal(t, []ServiceDescriptor{descriptor}, descriptors)
	assert.Same(t, services, newServices)
}

func TestServiceCollection_AddMultiple(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	newServices := services.Add(descriptor1)
	newServices = newServices.Add(descriptor2)
	newServices = newServices.Add(descriptor3)
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2, descriptor3}
	descriptors := services.ListDescriptors()
	assert.Equal(t, expectedDescriptors, descriptors)
	assert.Same(t, services, newServices)
}

func TestServiceCollection_AddRange(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	newServices := services.AddRange(descriptor1, descriptor2, descriptor3)
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2, descriptor3}
	descriptors := services.ListDescriptors()
	assert.Equal(t, expectedDescriptors, descriptors)
	assert.Same(t, services, newServices)
}

func TestServiceCollection_FindDescriptors(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	descriptors := services.FindDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestStructWithFieldsPtr
		})
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyDisposablePtr
		})
	expectedDescriptors = []ServiceDescriptor{descriptor3}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyNonDisposablePtr
		})
	expectedDescriptors = []ServiceDescriptor{}
	assert.Equal(t, expectedDescriptors, descriptors)
}

func TestServiceCollection_FindFirstDescriptor(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	actualDescriptor := services.FindFirstDescriptor(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestStructWithFieldsPtr
		})
	expectedDescriptor := descriptor1
	assert.Equal(t, expectedDescriptor, actualDescriptor)

	actualDescriptor = services.FindFirstDescriptor(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyDisposablePtr
		})
	expectedDescriptor = descriptor3
	assert.Equal(t, expectedDescriptor, actualDescriptor)

	actualDescriptor = services.FindFirstDescriptor(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyNonDisposablePtr
		})
	assert.Nil(t, actualDescriptor)
}

func TestServiceCollection_FindDescriptorsForType(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	descriptors := services.FindDescriptorsForType(typeOfTestStructWithFieldsPtr)
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptorsForType(typeOfTestDummyDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor3}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptorsForType(typeOfTestDummyNonDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{}
	assert.Equal(t, expectedDescriptors, descriptors)
}

func TestServiceCollection_FindFirstDescriptorForType(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	actualDescriptor := services.FindFirstDescriptorForType(typeOfTestStructWithFieldsPtr)
	expectedDescriptor := descriptor1
	assert.Equal(t, expectedDescriptor, actualDescriptor)

	actualDescriptor = services.FindFirstDescriptorForType(typeOfTestDummyDisposablePtr)
	expectedDescriptor = descriptor3
	assert.Equal(t, expectedDescriptor, actualDescriptor)

	actualDescriptor = services.FindFirstDescriptorForType(typeOfTestDummyNonDisposablePtr)
	assert.Nil(t, actualDescriptor)
}

func TestServiceCollection_UpdateDescriptors(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	descriptor4, _ := NewInstance(&testStructWithFields{Field1: 4})
	descriptor5, _ := NewInstance(&testDummyDisposable{})
	descriptor6, _ := NewInstance(&testDummyDisposable{})
	descriptor7, _ := NewInstance(&testDummyNonDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	newServices := services.UpdateDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestStructWithFieldsPtr
		},
		func(found []ServiceDescriptor) []ServiceDescriptor {
			return []ServiceDescriptor{descriptor4}
		})
	descriptors := newServices.FindDescriptorsForType(typeOfTestStructWithFieldsPtr)
	expectedDescriptors := []ServiceDescriptor{descriptor4}
	assert.Equal(t, expectedDescriptors, descriptors)
	assert.Same(t, services, newServices)

	services.UpdateDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyDisposablePtr
		},
		func(found []ServiceDescriptor) []ServiceDescriptor {
			return []ServiceDescriptor{descriptor5, descriptor6}
		})
	descriptors = services.FindDescriptorsForType(typeOfTestDummyDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor5, descriptor6}
	assert.Equal(t, expectedDescriptors, descriptors)

	services.UpdateDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == typeOfTestDummyNonDisposablePtr
		},
		func(found []ServiceDescriptor) []ServiceDescriptor {
			return []ServiceDescriptor{descriptor7}
		})
	descriptors = services.FindDescriptorsForType(typeOfTestDummyNonDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor7}
	assert.Equal(t, expectedDescriptors, descriptors)
}

func TestServiceCollection_TryAdd(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	descriptor4, _ := NewInstance(&testStructWithFields{Field1: 4})
	descriptor5, _ := NewInstance(&testDummyNonDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	newServices := services.TryAdd(descriptor4)
	assert.Same(t, services, newServices)
	descriptors := services.FindDescriptorsForType(typeOfTestStructWithFieldsPtr)
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2}
	assert.Equal(t, expectedDescriptors, descriptors)

	services.TryAdd(descriptor5)
	descriptors = services.FindDescriptorsForType(typeOfTestDummyNonDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor5}
	assert.Equal(t, expectedDescriptors, descriptors)
	descriptors = services.FindDescriptorsForType(typeOfTestDummyDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor3}
	assert.Equal(t, expectedDescriptors, descriptors)
}

func TestServiceCollection_TryAddRange(t *testing.T) {
	services := NewServiceCollection()
	descriptor1, _ := NewInstance(&testStructWithFields{Field1: 1})
	descriptor2, _ := NewInstance(&testStructWithFields{Field1: 2})
	descriptor3, _ := NewInstance(&testDummyDisposable{})
	descriptor4, _ := NewInstance(&testStructWithFields{Field1: 4})
	descriptor5, _ := NewInstance(&testDummyNonDisposable{})
	services.AddRange(descriptor1, descriptor2, descriptor3)

	newServices := services.TryAddRange(descriptor4, descriptor5)
	assert.Same(t, services, newServices)

	descriptors := services.FindDescriptorsForType(typeOfTestStructWithFieldsPtr)
	expectedDescriptors := []ServiceDescriptor{descriptor1, descriptor2}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptorsForType(typeOfTestDummyNonDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor5}
	assert.Equal(t, expectedDescriptors, descriptors)

	descriptors = services.FindDescriptorsForType(typeOfTestDummyDisposablePtr)
	expectedDescriptors = []ServiceDescriptor{descriptor3}
	assert.Equal(t, expectedDescriptors, descriptors)
}

func TestServiceCollection_Build(t *testing.T) {
	services := NewServiceCollection()

	scope, err := services.Build()
	assert.NoError(t, err)
	assert.NotNil(t, scope)
}
