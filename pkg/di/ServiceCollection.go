package di

// ServiceCollection is a collection of services, describing a dependency graph of services.
type ServiceCollection struct {
	descriptors []ServiceDescriptor
}

// NewServiceCollection creates a new ServiceCollection.
func NewServiceCollection() *ServiceCollection {
	return &ServiceCollection{
		descriptors: []ServiceDescriptor{},
	}
}

// func (services *ServiceCollection) UpdateServices(
// 	predicate func(ServiceDescriptor) bool,
// 	update func([]ServiceDescriptor) []ServiceDescriptor) *ServiceCollection {
// 	indices := []int{}
// 	toReplace := []ServiceDescriptor{}
// 	for i, descriptor := range services.descriptors {
// 		if predicate(descriptor) {
// 			indices = append(indices, i)
// 			toReplace = append(toReplace, descriptor)
// 		}
// 	}

// 	updated := update(toReplace)

// 	if len(indices) == 0 {
// 		services.descriptors = append(services.descriptors, updated...)
// 	} else {
// 		descriptors := make([]ServiceDescriptor, len(services.descriptors)-len(indices)+len(updated))
// 		originalPos, indicesPos, updatedPos := 0, 0, 0

// 		for i := 0; i < len(descriptors); i++ {
// 			if 
// 		}
// 	}
// 	return services
// }

func (services *ServiceCollection) Add(descriptor ServiceDescriptor) *ServiceCollection {
	services.descriptors = append(services.descriptors, descriptor)
	return services
}

func (services *ServiceCollection) Build() (ServiceProvider, error) {
	panic("not implemented")
}
