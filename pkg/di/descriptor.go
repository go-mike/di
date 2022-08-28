package di

import "reflect"

type descriptor struct {
	serviceType reflect.Type
	lifetime    Lifetime
	factory     ServiceFactory
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
	return NewDescriptorForType(reflect.TypeOf((*T)(nil)).Elem(), lifetime, factory)
}

// NewDescriptorForType creates a new service descriptor.
func NewDescriptorForTypeExplicitFactory(
	serviceType reflect.Type, 
	lifetime Lifetime, 
	factoryFunc SimpleServiceFactoryFunc,
	requirements []reflect.Type, 
	displayName string) ServiceDescriptor {
	return NewDescriptorForType(
		serviceType, lifetime,
		NewExplicitFactory(factoryFunc, requirements, displayName))
}


// NewSingletonForType creates a new singleton service descriptor for the given service type.
func NewSingletonForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Singleton, factory)
}

// NewSingleton creates a new singleton service descriptor for the given service type.
func NewSingleton[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Singleton, factory)
}

// NewSingletonFactoryForType creates a new singleton service descriptor for the given service type.
func NewSingletonFactoryForType(
	serviceType reflect.Type,
	factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewSingletonForType(serviceType, NewFactory(factoryFunc))
}

// NewSingletonFactory creates a new singleton service descriptor for the given service type.
func NewSingletonFactory[T any](factoryFunc SimpleServiceFactoryFunc) ServiceDescriptor {
	return NewSingleton[T](NewFactory(factoryFunc))
}

// NewSingletonStructForType creates a new singleton service descriptor for the given service type.
func NewSingletonStructForType(
	serviceType reflect.Type, structType reflect.Type) (ServiceDescriptor, error) {
	factory, err := NewStructFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return NewSingletonForType(serviceType, factory), nil
}

// NewSingletonStruct creates a new singleton service descriptor for the given service type.
func NewSingletonStruct[T any, Impl any]() (ServiceDescriptor, error) {
	factory, err := NewStructFactory[Impl]()
	if err != nil {
		return nil, err
	}
	return NewSingleton[T](factory), nil
}


// NewScopedForType creates a new scoped service descriptor for the given service type.
func NewScopedForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Scoped, factory)
}

// NewScoped creates a new scoped service descriptor for the given service type.
func NewScoped[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Scoped, factory)
}


// NewTransientForType creates a new transient service descriptor for the given service type.
func NewTransientForType(serviceType reflect.Type, factory ServiceFactory) ServiceDescriptor {
	return NewDescriptorForType(serviceType, Transient, factory)
}

// NewTransient creates a new transient service descriptor for the given service type.
func NewTransient[T any](factory ServiceFactory) ServiceDescriptor {
	return NewDescriptor[T](Transient, factory)
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
