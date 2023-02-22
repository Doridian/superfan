package curve

type CurveDriver interface {
	Init() error
	Close() error
	GetFanSpeedFor(temperature float64) (float64, error)
}
