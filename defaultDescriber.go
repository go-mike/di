package di

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
	return searchDescriptor(describer.descriptors, serviceType)
}

func (describer *defaultDescriber) GetServiceDescriptors(serviceType reflect.Type) []ServiceDescriptor {
	return searchDescriptors(describer.descriptors, serviceType)
}

func searchDescriptor(descriptors []ServiceDescriptor, serviceType reflect.Type) ServiceDescriptor {
	for i := len(descriptors) - 1; i >= 0; i-- {
		descriptor := descriptors[i]
		if descriptor.ServiceType() == serviceType {
			return descriptor
		}
	}

	return nil
}

func searchDescriptors(descriptors []ServiceDescriptor, serviceType reflect.Type) []ServiceDescriptor {
	var result []ServiceDescriptor
	for _, descriptor := range descriptors {
		if descriptor.ServiceType() == serviceType {
			result = append(result, descriptor)
		}
	}

	return result
}

// newDefaultDescriber creates a new default service describer
func newDefaultDescriber(descriptors []ServiceDescriptor) (*defaultDescriber, error) {
	if err := validateDescriptors(descriptors); err != nil {
		return nil, err
	}

	return &defaultDescriber{
		descriptors: cloneSlice(descriptors),
	}, nil
}

type ServiceDependencyError struct {
	Errors []error
}

var _ error = (*ServiceDependencyError)(nil)

func (err *ServiceDependencyError) Error() string {
	var sb strings.Builder
	for i, e := range err.Errors {
		sb.WriteString(e.Error())
		if i < len(err.Errors)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func validateDescriptors(descriptors []ServiceDescriptor) error {
	var es []error

	for _, descriptor := range descriptors {
		requirements := descriptor.Factory().Requirements()
		for _, requirement := range requirements {
			reqErrors := validateRequirement(
				descriptor.Factory().DisplayName(),
				descriptor.Lifetime(),
				requirement,
				descriptors)
			es = append(es, reqErrors...)
		}
	}

	if len(es) > 0 {
		return &ServiceDependencyError{Errors: es}
	}

	return nil
}

func validateRequirement(
	service string,
	lifetime Lifetime,
	requirement reflect.Type,
	descriptors []ServiceDescriptor,
) []error {
	var es []error

	if requirement.Kind() == reflect.Slice {
		dependencies := searchDescriptors(descriptors, requirement.Elem())

		if len(dependencies) == 0 {
			return nil
		}

		for _, dependency := range dependencies {
			reqErrors := validateRequirementDescriptor(
				service, lifetime, dependency, descriptors)
			es = append(es, reqErrors...)
		}
	} else {

		dependency := searchDescriptor(descriptors, requirement)

		if dependency == nil {
			es = append(es, fmt.Errorf(
				"service %s requires %s but it is not found",
				service, requirement))
		} else {
			reqErrors := validateRequirementDescriptor(
				service, lifetime, dependency, descriptors)
			es = append(es, reqErrors...)
		}
	}

	return es
}

func validateRequirementDescriptor(
	service string,
	lifetime Lifetime,
	requirementDescriptor ServiceDescriptor,
	descriptors []ServiceDescriptor,
) []error {
	var es []error

	if lifetime == Singleton {
		if requirementDescriptor.Lifetime() == Scoped {
			es = append(es, fmt.Errorf(
				"singleton service %s cannot request scoped service %s",
				service, requirementDescriptor.ServiceType()))
		}
	}

	return es
}
