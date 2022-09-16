package di

import (
	"errors"
	"reflect"
)

var (
	ErrServiceNotFound = errors.New("service not found")
)

type defaultDescriber struct {
	descriptors []ServiceDescriptor
}

var _ ServiceDescriber = (*defaultDescriber)(nil)

// GetServiceDef implements ServiceDescriber
func (describer *defaultDescriber) GetServiceDescriptor(serviceType reflect.Type) ServiceDescriptor {
	// Look from back to front to prioritize the last added services?
	for _, descriptor := range describer.descriptors {
		if descriptor.ServiceType() == serviceType {
			return descriptor
		}
	}

	return nil
}

// newDefaultDescriber creates a new default service describer
func newDefaultDescriber(descriptors []ServiceDescriptor) (*defaultDescriber, error) {
	if err := findValidationErrors(descriptors); err != nil {
		return nil, err
	}

	return &defaultDescriber{

	}, nil
}

func findValidationErrors(descriptors []ServiceDescriptor) error {
	// TODO: Validate
	return nil
}
