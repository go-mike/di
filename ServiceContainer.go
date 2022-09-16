package di

type ServiceContainer interface {
	Provider() ServiceProvider
	IsScoped() bool
	Dispose()
	IsDisposed() bool
}
