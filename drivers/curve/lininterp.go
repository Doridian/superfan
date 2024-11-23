package curve

type LinearInterpolated struct {
	Points []*Point
}

func (d *LinearInterpolated) Init() error {
	return nil
}

func (d *LinearInterpolated) Close() error {
	return nil
}

func (d *LinearInterpolated) GetFanSpeedFor(temperature float64) (float64, error) {
	for i, t := range d.Points {
		if temperature <= t.Temperature {
			if i == 0 {
				return t.Speed, nil
			}
			prev := d.Points[i-1]
			// Linear interpolation
			return prev.Speed + (t.Speed-prev.Speed)*(temperature-prev.Temperature)/(t.Temperature-prev.Temperature), nil
		}
	}
	return 1.00, nil
}
