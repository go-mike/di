# github.com/go-mike/di
Dependency Injection Abstraction and simple implementation

## Table of Service request possibilities

| Can a service | Ask for a service | Answer |
| ------------- | ----------------- | ------ |
| Singleton     | Singleton         | ✅      |
| Singleton     | Scoped            | ❌      |
| Singleton     | Transient         | ✅      |
| Scoped        | Singleton         | ✅      |
| Scoped        | Scoped            | ✅      |
| Scoped        | Transient         | ✅      |
| Transient     | Singleton         | ✅      |
| Transient     | Scoped            | ❌      |
| Transient     | Transient         | ✅      |
