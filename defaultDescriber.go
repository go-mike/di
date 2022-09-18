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

type validatedDescriptor struct {
	descriptor            ServiceDescriptor
	validatedForRecurse   bool
	validatedForSingleton bool
	validatedForScoped    bool
}

func validateDescriptors(
	descriptors []ServiceDescriptor,
) error {
	validations := mapSlice(
		descriptors,
		func(descriptor ServiceDescriptor) *validatedDescriptor {
			return &validatedDescriptor{
				descriptor: descriptor,
			}
		},
	)

	messages := validateDescriptorsAux(validations)

	if len(messages) > 0 {
		messages = distinctSlice(messages)
		errs := mapSlice(messages, func(s string) error { return errors.New(s) })
		return &ServiceDependencyError{Errors: errs}
	}

	return nil
}

func validateDescriptorsAux(
	validations []*validatedDescriptor,
) []string {
	var messages []string

	validate := func(recurse bool) {
		for _, validation := range validations {
			requestChain := []ServiceDescriptor{validation.descriptor}
			moreErrors := validateDescriptor(
				validation.descriptor.Lifetime(),
				requestChain,
				validation,
				validations,
				recurse,
			)
			messages = append(messages, moreErrors...)
		}
	}

	validate(false)
	validate(true)

	return messages
}

func validateDescriptor(
	lifetime Lifetime,
	requestChain []ServiceDescriptor,
	validation *validatedDescriptor,
	validations []*validatedDescriptor,
	recurse bool,
) []string {
	if lifetime == Scoped {
		if validation.validatedForScoped && validation.validatedForRecurse == recurse {
			return nil
		}
		validation.validatedForScoped = true
		validation.validatedForRecurse = recurse
	}

	if lifetime != Scoped {
		if validation.validatedForSingleton && validation.validatedForRecurse == recurse {
			return nil
		}
		validation.validatedForSingleton = true
		validation.validatedForRecurse = recurse
	}

	var messages []string

	requirements := validation.descriptor.Factory().Requirements()

	for _, requirement := range requirements {
		if requirement.Kind() == reflect.Slice {
			for _, current := range validations {
				if current.descriptor.ServiceType() == requirement.Elem() {
					// Validate the requirement
					messages = append(
						messages,
						validateRequirement(
							lifetime,
							current.descriptor,
							requestChain)...)
					// Validate the requirement's requirements from the current lifetime
					if recurse {
						messages = append(
							messages,
							validateDescriptor(
								lifetime,
								append(requestChain, current.descriptor),
								current,
								validations,
								recurse)...)
					}
				}
			}
		} else {
			for i := len(validations) - 1; i >= 0; i-- {
				current := validations[i]
				if current.descriptor.ServiceType() == requirement {
					// Validate the requirement
					messages = append(
						messages,
						validateRequirement(
							lifetime,
							current.descriptor,
							requestChain)...)
					// Validate the requirement's requirements from the current lifetime
					if recurse {
						messages = append(
							messages,
							validateDescriptor(
								lifetime,
								append(requestChain, current.descriptor),
								current,
								validations,
								recurse)...)
					}
					return messages
				}
			}

			messages = append(
				messages,
				fmt.Sprintf(
					"service request %s =(not found)=> %s fails",
					requestChainString(requestChain),
					requirement))
		}
	}

	return messages
}

func validateRequirement(
	lifetime Lifetime,
	requirementDescriptor ServiceDescriptor,
	requestChain []ServiceDescriptor,
) []string {
	var messages []string

	if !isValidRequirement(lifetime, requirementDescriptor) {
		messages = append(
			messages,
			fmt.Sprintf(
				"service request %s =(invalid)=> %s fails",
				requestChainString(requestChain),
				requirementDescriptor))
	}

	return messages
}

func isValidRequirement(
	lifetime Lifetime,
	requirementDescriptor ServiceDescriptor,
) bool {
	if lifetime == Singleton && requirementDescriptor.Lifetime() == Scoped {
		return false
	}

	return true
}

func requestChainString(requestChain []ServiceDescriptor) string {
	var sb strings.Builder
	for i, descriptor := range requestChain {
		sb.WriteString(descriptor.String())
		if i < len(requestChain)-1 {
			sb.WriteString(" ==> ")
		}
	}
	return sb.String()
}
