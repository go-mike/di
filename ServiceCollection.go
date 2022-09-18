package di

import "reflect"

// ServiceCollection is a collection of services, describing a dependency graph of services.
type ServiceCollection interface {
	ListDescriptors() []ServiceDescriptor

	FindDescriptors(predicate func(ServiceDescriptor) bool) []ServiceDescriptor
	FindFirstDescriptor(predicate func(ServiceDescriptor) bool) ServiceDescriptor

	FindDescriptorsForType(serviceType reflect.Type) []ServiceDescriptor
	FindFirstDescriptorForType(serviceType reflect.Type) ServiceDescriptor

	Add(descriptor ServiceDescriptor) ServiceCollection
	AddRange(descriptors ...ServiceDescriptor) ServiceCollection

	UpdateDescriptors(
		predicate func(ServiceDescriptor) bool,
		replace func([]ServiceDescriptor) []ServiceDescriptor,
	) ServiceCollection

	TryAdd(descriptor ServiceDescriptor) ServiceCollection
	TryAddRange(descriptors ...ServiceDescriptor) ServiceCollection

	// TODO: Decorate

	Build() (ServiceContainer, error)
}
