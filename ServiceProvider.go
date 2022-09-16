package di

import "reflect"

type ServiceProvider interface {
	GetService(serviceType reflect.Type) (any, error)
	GetServiceInfo(serviceType reflect.Type) ServiceInfo
}

type ServiceInfo struct {
	ServiceType    reflect.Type
	Lifetime       Lifetime
	IsInstantiated bool
}

func (info ServiceInfo) IsNotFound() bool {
	return info.Lifetime == UnknownLifetime
}

func newNotFoundServiceInfo(serviceType reflect.Type) ServiceInfo {
	return ServiceInfo{
		ServiceType:    serviceType,
		Lifetime:       UnknownLifetime,
		IsInstantiated: false,
	}
}

func newServiceInfo(serviceType reflect.Type, instantiated bool, lifetime Lifetime) ServiceInfo {
	return ServiceInfo{
		ServiceType:    serviceType,
		Lifetime:       lifetime,
		IsInstantiated: instantiated,
	}
}
