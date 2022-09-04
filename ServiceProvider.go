package di

import "reflect"

type ServiceProvider interface {
	GetService(serviceType reflect.Type) (any, error)
	GetServiceInfo(serviceType reflect.Type) ServiceInfo
}

type ServiceInfo struct {
	ServiceType    reflect.Type
	IsInstantiated bool
	Lifetime       Lifetime
}

func newNotFoundServiceInfo(serviceType reflect.Type) ServiceInfo {
	return ServiceInfo{
		ServiceType:    serviceType,
		IsInstantiated: false,
		Lifetime:       UnknownLifetime,
	}
}

func newServiceInfo(serviceType reflect.Type, instantiated bool, lifetime Lifetime) ServiceInfo {
	return ServiceInfo{
		ServiceType:    serviceType,
		IsInstantiated: instantiated,
		Lifetime:       lifetime,
	}
}
