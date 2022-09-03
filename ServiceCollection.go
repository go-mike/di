package di

// // ServiceCollection is a collection of services, describing a dependency graph of services.
// type ServiceCollection interface {
// 	FindDescriptors(predicate func(ServiceDescriptor) bool) []ServiceDescriptor
// 	FindFirstDescriptor(predicate func(ServiceDescriptor) bool) ServiceDescriptor

// 	UpdateDescriptors(
// 		predicate func(ServiceDescriptor) bool,
// 		replace func([]ServiceDescriptor) []ServiceDescriptor,
// 	) ServiceCollection

// 	AddRange(descriptors []ServiceDescriptor) ServiceCollection

// 	Add(descriptor ServiceDescriptor) ServiceCollection
// 	TryAdd(descriptor ServiceDescriptor) ServiceCollection
// 	ReplaceAll(descriptor ServiceDescriptor) ServiceCollection

// 	// TODO: Decorate(descriptor ServiceDescriptor) ServiceCollection

// 	Build() (ServiceProvider, error)
// }

// // ServiceCollection is a collection of services, describing a dependency graph of services.
// type serviceCollection struct {
// 	descriptors []ServiceDescriptor
// }

// // NewServiceCollection creates a new ServiceCollection.
// func NewServiceCollection() ServiceCollection {
// 	return &serviceCollection{
// 		descriptors: []ServiceDescriptor{},
// 	}
// }

// // func (services *ServiceCollection) UpdateServices(
// // 	predicate func(ServiceDescriptor) bool,
// // 	update func([]ServiceDescriptor) []ServiceDescriptor) *ServiceCollection {
// // 	indices := []int{}
// // 	toReplace := []ServiceDescriptor{}
// // 	for i, descriptor := range services.descriptors {
// // 		if predicate(descriptor) {
// // 			indices = append(indices, i)
// // 			toReplace = append(toReplace, descriptor)
// // 		}
// // 	}

// // 	updated := update(toReplace)

// // 	if len(indices) == 0 {
// // 		services.descriptors = append(services.descriptors, updated...)
// // 	} else {
// // 		descriptors := make([]ServiceDescriptor, len(services.descriptors)-len(indices)+len(updated))
// // 		originalPos, indicesPos, updatedPos := 0, 0, 0

// // 		for i := 0; i < len(descriptors); i++ {
// // 			if
// // 		}
// // 	}
// // 	return services
// // }

// func (services *serviceCollection) FindDescriptors(predicate func(ServiceDescriptor) bool) []ServiceDescriptor {
// 	return filterSlice(services.descriptors, predicate)
// }

// func (services *serviceCollection) FindFirstDescriptor(predicate func(ServiceDescriptor) bool) ServiceDescriptor {
// 	found := findSlice(services.descriptors, predicate)
// 	if found == nil {
// 		return nil
// 	}
// 	return *found
// }

// func (services *serviceCollection) UpdateDescriptors(
// 	predicate func(ServiceDescriptor) bool,
// 	replace func([]ServiceDescriptor) []ServiceDescriptor,
// ) ServiceCollection {
// 	toReplace, toRemain := partitionSlice(services.descriptors, predicate)
// 	replaced := replace(toReplace)
// 	services.descriptors = append(toRemain, replaced...)
// 	return services
// }

// func (services *serviceCollection) AddRange(descriptors []ServiceDescriptor) ServiceCollection {
// 	services.descriptors = append(services.descriptors, descriptors...)
// 	return services
// }

// func (services *serviceCollection) Add(descriptor ServiceDescriptor) ServiceCollection {
// 	services.descriptors = append(services.descriptors, descriptor)
// 	return services
// }

// func (services *serviceCollection) TryAdd(descriptor ServiceDescriptor) ServiceCollection {
// 	serviceType := descriptor.ServiceType()
// 	return services.UpdateDescriptors(
// 		func(descriptor ServiceDescriptor) bool { 
// 			return descriptor.ServiceType() == serviceType 
// 		},
// 		func(toReplace []ServiceDescriptor) []ServiceDescriptor {
// 			if len(toReplace) == 0 {
// 				return []ServiceDescriptor{descriptor}
// 			} else {
// 				return toReplace
// 			}
// 		})
// }

// func (services *serviceCollection) ReplaceAll(descriptor ServiceDescriptor) ServiceCollection {
// 	serviceType := descriptor.ServiceType()
// 	return services.UpdateDescriptors(
// 		func(descriptor ServiceDescriptor) bool {
// 			return descriptor.ServiceType() == serviceType
// 		},
// 		func(toReplace []ServiceDescriptor) []ServiceDescriptor {
// 			return []ServiceDescriptor{descriptor}
// 		})
// }

// func (services *serviceCollection) Build() (ServiceProvider, error) {
// 	panic("not implemented")
// }
