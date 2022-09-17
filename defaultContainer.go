package di

import (
	"errors"
	"reflect"
	"sync"
)

var (
	ErrServiceContainerDisposed = errors.New("service container has been disposed")
	ErrMissingServiceDescriber    = errors.New("missing service describer")
)

type defaultContainer struct {
	mutex     sync.Mutex
	describer ServiceDescriber
	parent    ServiceProvider
	data      []*descriptorData
}

type descriptorData struct {
	descriptor ServiceDescriptor
	instances  []ServiceInstance
}

var _ ServiceContainer = (*defaultContainer)(nil)
var _ ServiceProvider = (*defaultContainer)(nil)

// Provider implements ServiceContainer
func (scope *defaultContainer) Provider() ServiceProvider {
	return scope
}

// Dispose implements ServiceContainer
func (scope *defaultContainer) Dispose() {
	scope.mutex.Lock()
	defer scope.mutex.Unlock()
	if !scope.IsDisposed() {
		for _, data := range scope.data {
			for _, instance := range data.instances {
				instance.Disposable.Dispose()
			}
		}
		scope.data = nil
	}
}

// IsDisposed implements ServiceContainer
func (scope *defaultContainer) IsDisposed() bool {
	return scope.data == nil
}

// IsScoped implements ServiceContainer
func (scope *defaultContainer) IsScoped() bool {
	return scope.parent != nil
}

// GetService implements ServiceProvider
func (scope *defaultContainer) GetService(serviceType reflect.Type) (any, error) {
	if scope.IsDisposed() {
		return nil, ErrServiceContainerDisposed
	}

	scope.mutex.Lock()
	defer scope.mutex.Unlock()

	data := scope.findDescriptorData(serviceType)

	if data == nil {

	}

	return nil, nil
}

// GetServiceInfo implements ServiceProvider
func (scope *defaultContainer) GetServiceInfo(serviceType reflect.Type) ServiceInfo {
	if scope.IsDisposed() {
		return newNotFoundServiceInfo(serviceType)
	}

	descriptor := scope.describer.GetServiceDescriptor(serviceType)

	if descriptor == nil {
		return newNotFoundServiceInfo(serviceType)
	}

	scope.mutex.Lock()
	defer scope.mutex.Unlock()

	data := scope.findDescriptorData(serviceType)
	isInstantiated := data != nil

	return newServiceInfo(serviceType, isInstantiated, descriptor.Lifetime())
}

func (scope *defaultContainer) findDescriptorData(serviceType reflect.Type) *descriptorData {
	for _, data := range scope.data {
		if data.descriptor.ServiceType() == serviceType {
			return data
		}
	}

	return nil
}

func newDefaultContainer(
	describer ServiceDescriber,
	descriptors []ServiceDescriptor,
	parent ServiceProvider,
) (*defaultContainer, error) {
	if describer == nil {
		return nil, ErrMissingServiceDescriber
	}

	return &defaultContainer{
		data:   nil,
		parent: parent,
	}, nil
}
