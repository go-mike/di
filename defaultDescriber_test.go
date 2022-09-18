package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultDescriber_Empty(t *testing.T) {
	descriptors := []ServiceDescriptor{}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_WithStructAsInterface(t *testing.T) {
	descriptor, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)

	actualDescriptor := describer.GetServiceDescriptor(typeOfTestServiceInterface)
	assert.Same(t, descriptor, actualDescriptor)

	actualDescriptor = describer.GetServiceDescriptor(typeOfTestStructWithFieldsPtr)
	assert.Nil(t, actualDescriptor)
}

type testStructWithDependency struct {
	Dependency testServiceInterface
}

type testStructWithOtherDependency struct {
	Dependency *testStructWithDependency
}

func TestNewDefaultDescriber_WithMissingDependency(t *testing.T) {
	descriptor, _ := NewSingletonStructPtr[testStructWithDependency]()
	descriptors := []ServiceDescriptor{descriptor}

	describer, err := newDefaultDescriber(descriptors)
	assert.Nil(t, describer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Singleton] *di.testStructWithDependency =(not found)=> di.testServiceInterface")
}

type testStructWithDependencySlice struct {
	Dependency []testServiceInterface
}

type testStructWithOtherDependencySlice struct {
	Dependency []*testStructWithDependencySlice
}

func TestNewDefaultDescriber_WithMissingDependencies(t *testing.T) {
	descriptor, _ := NewSingletonStructPtr[testStructWithDependencySlice]()
	descriptors := []ServiceDescriptor{descriptor}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_SingletonToSingleton(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependency]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_SingletonToSingletonSlice(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_SingletonToScoped(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependency]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.Nil(t, describer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Singleton] *di.testStructWithDependency =(invalid)=> [Scoped] di.testServiceInterface")
}

func TestNewDefaultDescriber_SingletonToScopedSlice(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.Nil(t, describer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Singleton] *di.testStructWithDependencySlice =(invalid)=> [Scoped] di.testServiceInterface")
}

func TestNewDefaultDescriber_SingletonToTransient(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependency]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_SingletonToTransientSlice(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToSingleton(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependency]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToSingletonSlice(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToScoped(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependency]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToScopedSlice(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToTransient(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependency]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_ScopedToTransientSlice(t *testing.T) {
	descriptor1, _ := NewScopedStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToSingleton(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependency]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToSingletonSlice(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewSingletonStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToScoped(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependency]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToScopedSlice(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToTransient(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependency]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_TransientToTransientSlice(t *testing.T) {
	descriptor1, _ := NewTransientStructPtr[testStructWithDependencySlice]()
	descriptor2, _ := NewTransientStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2}

	describer, err := newDefaultDescriber(descriptors)
	assert.NoError(t, err)
	assert.NotNil(t, describer)
}

func TestNewDefaultDescriber_SingletonToTransientToScoped(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithOtherDependency]()
	descriptor2, _ := NewTransientStructPtr[testStructWithDependency]()
	descriptor3, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2, descriptor3}

	describer, err := newDefaultDescriber(descriptors)
	assert.Nil(t, describer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), 
		"[Singleton] *di.testStructWithOtherDependency ==> [Transient] *di.testStructWithDependency =(invalid)=> [Scoped] di.testServiceInterface")
}

func TestNewDefaultDescriber_SingletonToTransientToScopedSlice(t *testing.T) {
	descriptor1, _ := NewSingletonStructPtr[testStructWithOtherDependencySlice]()
	descriptor2, _ := NewTransientStructPtr[testStructWithDependencySlice]()
	descriptor3, _ := NewScopedStruct[testServiceInterface, testServiceStruct]()
	descriptors := []ServiceDescriptor{descriptor1, descriptor2, descriptor3}

	describer, err := newDefaultDescriber(descriptors)
	assert.Nil(t, describer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(),
		"[Singleton] *di.testStructWithOtherDependencySlice ==> [Transient] *di.testStructWithDependencySlice =(invalid)=> [Scoped] di.testServiceInterface")
}
