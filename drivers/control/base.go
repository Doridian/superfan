package control

type Driver interface {
	Init() error
	Close() error
	SetFanSpeed(speed float64) error
	GetFanSpeed() (float64, error)
}
