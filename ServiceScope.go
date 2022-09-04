package di

type ServiceScope interface {
	Provider() ServiceProvider
	Dispose()
}
