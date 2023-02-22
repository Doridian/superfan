package control

type X9IPMIDriver struct {
	IPMIDriver
}

func (d *X9IPMIDriver) Init() error {
	err := d.IPMIDriver.Init()
	if err != nil {
		return err
	}
	// Set fan mode FULL
	// Set fan bank 0 override mode
	// Set fan bank 1 override mode
	return nil
}

func (d *X9IPMIDriver) Close() error {
	// Set fan mode STANDARD
	return d.IPMIDriver.Close()
}
