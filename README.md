# github.com/go-mike/di

Dependency Injection Abstractions and default implementation

## Glossary

- **Service Interface**: Represents a type that serves as a key to request a service to a service provider, and as a way to interact with the service.
- **Service Implementation**: Represents a type that implements a service interface.
- **Service Lifetime**: Represents the lifetime of a service.
  - *Singleton Lifetime*: The service is instantiated once and shared across all requests.
  - *Scoped Lifetime*: The service is instantiated once per scope.
  - *Transient Lifetime*: The service is instantiated every time it is requested.
- **Disposable**: Represents a type that can be disposed.
- **Service Factory**: Is a function that can create an instance of a service implementation, and a way to dispose of it when not required anymore.
- **Service Descriptor**: Represents a description of a service with a lifetime, a service interface, and a service factory.
- **Service Collection**: Is a mutable collection of service descriptors, which can be used to configure the services available in an application.
- **Service Describer**: Is an immutable collection of service descriptors, used internally by the service provider to resolve services.
- **Service Container**: Is a container of services, which can be used to resolve services, and manage their lifetime.
- **Service Provider**: Is a type to access the services available in an application.

## Table of service possible dependencies

| Can a service | Ask for a service | Answer    |
| ------------- | ----------------- | --------- |
| Singleton     | Singleton         | ✅         |
| Singleton     | Scoped            | ❌         |
| Singleton     | Transient         | ✅         |
| Scoped        | Singleton         | ✅         |
| Scoped        | Scoped            | ✅         |
| Scoped        | Transient         | ✅         |
| Transient     | Singleton         | ✅         |
| Transient     | Scoped            | if scoped |
| Transient     | Transient         | ✅         |

