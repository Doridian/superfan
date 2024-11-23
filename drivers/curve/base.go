package curve

type Point struct {
	Temperature float64
	Speed       float64
}

type Curve interface {
	Init() error
	Close() error
	GetFanSpeedFor(temperature float64) (float64, error)
}
