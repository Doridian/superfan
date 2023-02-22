package thermal

type Driver interface {
	Init() error
	Close() error
	GetTemperature() (float64, error)
}
