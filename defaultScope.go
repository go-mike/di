package di

import (
	"errors"
	"reflect"
)

var ErrServiceScopeDisposed = errors.New("service scope has been disposed")

type defaultScope struct {
	parent   ServiceProvider
	lifetime Lifetime
	data     []*descriptorData
}

type descriptorData struct {
	descriptor ServiceDescriptor
	instances  []ServiceInstance
}

var _ ServiceScope = (*defaultScope)(nil)
var _ ServiceProvider = (*defaultScope)(nil)

// Provider implements ServiceScope
func (scope *defaultScope) Provider() ServiceProvider {
	return scope
}

// Lifetime implements ServiceScope
func (scope *defaultScope) Lifetime() Lifetime {
	return scope.lifetime
}

// Dispose implements ServiceScope
func (scope *defaultScope) Dispose() {
	if !scope.IsDisposed() {
		for _, data := range scope.data {
			for _, instance := range data.instances {
				instance.Disposable.Dispose()
			}
		}
		scope.data = nil
	}
}

// IsDisposed implements ServiceScope
func (scope *defaultScope) IsDisposed() bool {
	return scope.data == nil
}

// GetService implements ServiceProvider
func (scope *defaultScope) GetService(serviceType reflect.Type) (any, error) {
	panic("unimplemented")
}

// GetServiceInfo implements ServiceProvider
func (scope *defaultScope) GetServiceInfo(serviceType reflect.Type) ServiceInfo {
	if scope.IsDisposed() {
		return newNotFoundServiceInfo(serviceType)
	}

	data := scope.findDescriptorData(serviceType)

	if data != nil {
		isInstantiated := false
		lifetime := data.descriptor.Lifetime()

		if lifetime != Transient {
			isInstantiated = len(data.instances) > 0
		}

		return newServiceInfo(serviceType, isInstantiated, lifetime)
	}

	return newNotFoundServiceInfo(serviceType)
}

func (scope *defaultScope) findDescriptorData(serviceType reflect.Type) *descriptorData {
	for _, data := range scope.data {
		if data.descriptor.ServiceType() == serviceType {
			return data
		}
	}

	return nil
}

func newSingletonScope(descriptors []ServiceDescriptor) (*defaultScope, error) {
	err := findValidationErrors(nil, descriptors)
	if err != nil {
		return nil, err
	}

	data := mapSlice(descriptors, func(descriptor ServiceDescriptor) *descriptorData {
		return &descriptorData {
			descriptor: descriptor,
			instances:  nil,
		}
	})

	return &defaultScope{
		lifetime:    Singleton,
		data: data,
	}, nil
}

func findValidationErrors(parent ServiceProvider, descriptors []ServiceDescriptor) error {
	// TODO: Validate
	return nil
}
