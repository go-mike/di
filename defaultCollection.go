package di

import (
	"reflect"

	"golang.org/x/exp/slices"
)

// ServiceCollection is a collection of services, describing a dependency graph of services.
type defaultCollection struct {
	descriptors []ServiceDescriptor
}

// NewServiceCollection creates a new ServiceCollection.
func NewServiceCollection() ServiceCollection {
	return &defaultCollection{
		descriptors: []ServiceDescriptor{},
	}
}

func (services *defaultCollection) ListDescriptors() []ServiceDescriptor {
	return slices.Clone(services.descriptors)
}

func (services *defaultCollection) FindDescriptors(predicate func(ServiceDescriptor) bool) []ServiceDescriptor {
	return filterSlice(services.descriptors, predicate)
}

func (services *defaultCollection) FindFirstDescriptor(predicate func(ServiceDescriptor) bool) ServiceDescriptor {
	found := findSlice(services.descriptors, predicate)
	if found == nil {
		return nil
	}
	return *found
}

func (services *defaultCollection) FindDescriptorsForType(serviceType reflect.Type) []ServiceDescriptor {
	return services.FindDescriptors(func(descriptor ServiceDescriptor) bool {
		return descriptor.ServiceType() == serviceType
	})
}

func (services *defaultCollection) FindFirstDescriptorForType(serviceType reflect.Type) ServiceDescriptor {
	return services.FindFirstDescriptor(func(descriptor ServiceDescriptor) bool {
		return descriptor.ServiceType() == serviceType
	})
}

func (services *defaultCollection) AddRange(descriptors ...ServiceDescriptor) ServiceCollection {
	services.descriptors = append(services.descriptors, descriptors...)
	return services
}

func (services *defaultCollection) Add(descriptor ServiceDescriptor) ServiceCollection {
	services.descriptors = append(services.descriptors, descriptor)
	return services
}

func (services *defaultCollection) UpdateDescriptors(
	predicate func(ServiceDescriptor) bool,
	replace func([]ServiceDescriptor) []ServiceDescriptor,
) ServiceCollection {
	toReplace, toRemain := partitionSlice(services.descriptors, predicate)
	replaced := replace(toReplace)
	services.descriptors = append(toRemain, replaced...)
	return services
}

func (services *defaultCollection) TryAdd(descriptor ServiceDescriptor) ServiceCollection {
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

func (services *defaultCollection) TryAddRange(descriptors ...ServiceDescriptor) ServiceCollection {
	for _, descriptor := range descriptors {
		services.TryAdd(descriptor)
	}
	return services
}

func (services *defaultCollection) Build() (ServiceContainer, error) {
	describer, err := newDefaultDescriber(services.descriptors)
	if err != nil {
		return nil, err
	}
	return newDefaultContainer(
		describer,
		slices.Clone(services.descriptors),
		nil,
	)
}
