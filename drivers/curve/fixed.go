package curve

type FixedThreshold struct {
	Temperature float64
	Speed       float64
}

type FixedCurveDriver struct {
	Thresholds []*FixedThreshold
}

func (d *FixedCurveDriver) Init() error {
	return nil
}

func (d *FixedCurveDriver) Close() error {
	return nil
}

func (d *FixedCurveDriver) GetFanSpeedFor(temperature float64) (float64, error) {
	for _, t := range d.Thresholds {
		if temperature >= t.Temperature {
			return t.Speed, nil
		}
	}
	return 1.00, nil
}
