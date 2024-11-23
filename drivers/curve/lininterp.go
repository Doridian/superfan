package curve

type LinearInterpolatedCurveDriver struct {
	Thresholds []*FixedThreshold
}

func (d *LinearInterpolatedCurveDriver) Init() error {
	return nil
}

func (d *LinearInterpolatedCurveDriver) Close() error {
	return nil
}

func (d *LinearInterpolatedCurveDriver) GetFanSpeedFor(temperature float64) (float64, error) {
	for i, t := range d.Thresholds {
		if temperature <= t.Temperature {
			if i == 0 {
				return t.Speed, nil
			}
			prev := d.Thresholds[i-1]
			// Linear interpolation
			return prev.Speed + (t.Speed-prev.Speed)*(temperature-prev.Temperature)/(t.Temperature-prev.Temperature), nil
		}
	}
	return 1.00, nil
}
