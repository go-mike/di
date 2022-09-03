package di

import "reflect"

type descriptor struct {
	serviceType reflect.Type
	lifetime    Lifetime
	factory     ServiceFactory
}

// ServiceDescriptor implementation for descriptor

// Factory implements ServiceDescriptor.Factory to return the service factory.
func (desc *descriptor) Factory() ServiceFactory {
	return desc.factory
}

// Lifetime implements ServiceDescriptor.Lifetime to return the service lifetime.
func (desc *descriptor) Lifetime() Lifetime {
	return desc.lifetime
}

// ServiceType implements ServiceDescriptor.ServiceType to return the service type.
func (desc *descriptor) ServiceType() reflect.Type {
	return desc.serviceType
}

// Factories for descriptor as a ServiceDescriptor

// NewDescriptorForType creates a new service descriptor.
// parameters:
// 	serviceType - the service type
// 	lifetime - the service lifetime
// 	factory - the service factory
// returns:
// 	the new service descriptor
func NewDescriptorForType(serviceType reflect.Type, lifetime Lifetime, factory ServiceFactory) ServiceDescriptor {
	return &descriptor{
		serviceType: serviceType,
		lifetime:    lifetime,
		factory:     factory,
	}
}

// NewDescriptor creates a new service descriptor for the given service type.
func NewDescriptor[T any](lifetime Lifetime, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(
		reflect.TypeOf((*T)(nil)).Elem(),
		lifetime, factory)
}


// NewSingletonServiceFactoryForType creates a new singleton service descriptor for the given service type.
func NewSingletonServiceFactoryForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Singleton, factory)
}

// NewSingletonServiceFactory creates a new singleton service descriptor for the given service type.
func NewSingletonServiceFactory[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Singleton, factory)
}

// NewSingletonFactoryForType creates a new singleton service descriptor for the given service type.
func NewSingletonFactoryForType(
	serviceType reflect.Type,
	factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewSingletonServiceFactoryForType(serviceType, NewFactory(factoryFunc))
}

// NewSingletonFactory creates a new singleton service descriptor for the given service type.
func NewSingletonFactory[T any](factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewSingletonServiceFactory[T](NewFactory(factoryFunc))
}

// NewSingletonStructForType creates a new singleton service descriptor for the given service type.
func NewSingletonStructForType(
	serviceType reflect.Type, structType reflect.Type) (ServiceDescriptor, error) {
	factory, err := NewStructFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return NewSingletonServiceFactoryForType(serviceType, factory), nil
}

// NewSingletonStruct creates a new singleton service descriptor for the given service type.
func NewSingletonStruct[T any, Impl any]() (ServiceDescriptor, error) {
	factory, err := NewStructFactory[Impl]()
	if err != nil {
		return nil, err
	}
	return NewSingletonServiceFactory[T](factory), nil
}


// NewScopedServiceFactoryForType creates a new singleton service descriptor for the given service type.
func NewScopedServiceFactoryForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Scoped, factory)
}

// NewScopedServiceFactory creates a new singleton service descriptor for the given service type.
func NewScopedServiceFactory[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Scoped, factory)
}

// NewScopedFactoryForType creates a new singleton service descriptor for the given service type.
func NewScopedFactoryForType(
	serviceType reflect.Type,
	factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewScopedServiceFactoryForType(serviceType, NewFactory(factoryFunc))
}

// NewScopedFactory creates a new singleton service descriptor for the given service type.
func NewScopedFactory[T any](factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewScopedServiceFactory[T](NewFactory(factoryFunc))
}

// NewScopedStructForType creates a new singleton service descriptor for the given service type.
func NewScopedStructForType(
	serviceType reflect.Type, structType reflect.Type) (ServiceDescriptor, error) {
	factory, err := NewStructFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return NewScopedServiceFactoryForType(serviceType, factory), nil
}

// NewScopedStruct creates a new singleton service descriptor for the given service type.
func NewScopedStruct[T any, Impl any]() (ServiceDescriptor, error) {
	factory, err := NewStructFactory[Impl]()
	if err != nil {
		return nil, err
	}
	return NewScopedServiceFactory[T](factory), nil
}


// NewTransientServiceFactoryForType creates a new singleton service descriptor for the given service type.
func NewTransientServiceFactoryForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Transient, factory)
}

// NewTransientServiceFactory creates a new singleton service descriptor for the given service type.
func NewTransientServiceFactory[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Transient, factory)
}

// NewTransientFactoryForType creates a new singleton service descriptor for the given service type.
func NewTransientFactoryForType(
	serviceType reflect.Type,
	factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewTransientServiceFactoryForType(serviceType, NewFactory(factoryFunc))
}

// NewTransientFactory creates a new singleton service descriptor for the given service type.
func NewTransientFactory[T any](factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewTransientServiceFactory[T](NewFactory(factoryFunc))
}

// NewTransientStructForType creates a new singleton service descriptor for the given service type.
func NewTransientStructForType(
	serviceType reflect.Type, structType reflect.Type) (ServiceDescriptor, error) {
	factory, err := NewStructFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return NewTransientServiceFactoryForType(serviceType, factory), nil
}

// NewTransientStruct creates a new singleton service descriptor for the given service type.
func NewTransientStruct[T any, Impl any]() (ServiceDescriptor, error) {
	factory, err := NewStructFactory[Impl]()
	if err != nil {
		return nil, err
	}
	return NewTransientServiceFactory[T](factory), nil
}

// NewInstanceForType creates a new singleton service descriptor for the given service instance.
func NewInstanceForType(serviceType reflect.Type, instance any) (ServiceDescriptor, error) {
	factory, err := newInstanceFactory(instance)
	if err != nil {
		return nil, err
	}
	return NewSingletonServiceFactoryForType(serviceType, factory), nil
}

// NewInstance creates a new singleton service descriptor for the given service instance.
func NewInstance[T any](instance T) (ServiceDescriptor, error) {
	factory, err := newInstanceFactory(instance)
	if err != nil {
		return nil, err
	}
	return NewSingletonServiceFactory[T](factory), nil
}
