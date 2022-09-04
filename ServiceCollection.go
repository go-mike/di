package di

import (
	"reflect"

	"golang.org/x/exp/slices"
)

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

	Build() (ServiceScope, error)
}

// ServiceCollection is a collection of services, describing a dependency graph of services.
type serviceCollection struct {
	descriptors []ServiceDescriptor
}

// NewServiceCollection creates a new ServiceCollection.
func NewServiceCollection() ServiceCollection {
	return &serviceCollection{
		descriptors: []ServiceDescriptor{},
	}
}

func (services *serviceCollection) ListDescriptors() []ServiceDescriptor {
	return slices.Clone(services.descriptors)
}

func (services *serviceCollection) FindDescriptors(predicate func(ServiceDescriptor) bool) []ServiceDescriptor {
	return filterSlice(services.descriptors, predicate)
}

func (services *serviceCollection) FindFirstDescriptor(predicate func(ServiceDescriptor) bool) ServiceDescriptor {
	found := findSlice(services.descriptors, predicate)
	if found == nil {
		return nil
	}
	return *found
}

func (services *serviceCollection) FindDescriptorsForType(serviceType reflect.Type) []ServiceDescriptor {
	return services.FindDescriptors(func(descriptor ServiceDescriptor) bool {
		return descriptor.ServiceType() == serviceType
	})
}

func (services *serviceCollection) FindFirstDescriptorForType(serviceType reflect.Type) ServiceDescriptor {
	return services.FindFirstDescriptor(func(descriptor ServiceDescriptor) bool {
		return descriptor.ServiceType() == serviceType
	})
}

func (services *serviceCollection) AddRange(descriptors ...ServiceDescriptor) ServiceCollection {
	services.descriptors = append(services.descriptors, descriptors...)
	return services
}

func (services *serviceCollection) Add(descriptor ServiceDescriptor) ServiceCollection {
	services.descriptors = append(services.descriptors, descriptor)
	return services
}

func (services *serviceCollection) UpdateDescriptors(
	predicate func(ServiceDescriptor) bool,
	replace func([]ServiceDescriptor) []ServiceDescriptor,
) ServiceCollection {
	toReplace, toRemain := partitionSlice(services.descriptors, predicate)
	replaced := replace(toReplace)
	services.descriptors = append(toRemain, replaced...)
	return services
}

func (services *serviceCollection) TryAdd(descriptor ServiceDescriptor) ServiceCollection {
	serviceType := descriptor.ServiceType()
	return services.UpdateDescriptors(
		func(descriptor ServiceDescriptor) bool {
			return descriptor.ServiceType() == serviceType
		},
		func(toReplace []ServiceDescriptor) []ServiceDescriptor {
			if len(toReplace) == 0 {
				return []ServiceDescriptor{descriptor}
			} else {
				return toReplace
			}
		})
}

func (services *serviceCollection) TryAddRange(descriptors ...ServiceDescriptor) ServiceCollection {
	for _, descriptor := range descriptors {
		services.TryAdd(descriptor)
	}
	return services
}

func (services *serviceCollection) Build() (ServiceScope, error) {
	return newSingletonScope(slices.Clone(services.descriptors))
}
