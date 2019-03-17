package interfaces

type HealthServiceContract interface {
	HealthStatus() (httpStatusCode int, result interface{})
}

type DatabasePing interface {
	Ping() error
}
