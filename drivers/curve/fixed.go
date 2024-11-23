package curve

type Fixed struct {
	Points []*Point
}

func (d *Fixed) Init() error {
	return nil
}

func (d *Fixed) Close() error {
	return nil
}

func (d *Fixed) GetFanSpeedFor(temperature float64) (float64, error) {
	for _, t := range d.Points {
		if temperature <= t.Temperature {
			return t.Speed, nil
		}
	}
	return 1.00, nil
}
