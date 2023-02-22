package control

type X9IPMIDriver struct {
	lastSetSpeed byte
	IPMIDriver
}

func (d *X9IPMIDriver) Init() error {
	err := d.IPMIDriver.Init()
	if err != nil {
		return err
	}
	// Set fan mode FULL
	_, err = d.dev.RawCmd([]byte{0x30, 0x45, 0x01, 0x01})
	if err != nil {
		return err
	}
	// Set fan bank 0-3 override mode
	for i := byte(0); i < 4; i++ {
		_, err = d.dev.RawCmd([]byte{0x30, 0x91, 0x5A, 0x03, i, 0x00})
		if err != nil {
			return err
		}
	}
	return err
}

func (d *X9IPMIDriver) SetFanSpeed(speed float64) error {
	speedByte := byte(speed * 255)
	for i := byte(0); i < 2; i++ {
		_, err := d.dev.RawCmd([]byte{0x30, 0x91, 0x5A, 0x03, i, speedByte})
		if err != nil {
			return err
		}
	}
	d.lastSetSpeed = speedByte
	return nil
}

func (d *X9IPMIDriver) GetFanSpeed() (float64, error) {
	return float64(d.lastSetSpeed) / 255, nil
}

func (d *X9IPMIDriver) Close() error {
	// Set fan mode STANDARD
	_, err := d.dev.RawCmd([]byte{0x30, 0x45, 0x01, 0x00})
	if err != nil {
		return err
	}
	return d.IPMIDriver.Close()
}
