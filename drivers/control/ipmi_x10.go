package control

type X10IPMIDriver struct {
	IPMIDriver
}

func (d *X10IPMIDriver) Init() error {
	err := d.IPMIDriver.Init()
	if err != nil {
		return err
	}
	// Set fan mode FULL
	return nil
}

func (d *X10IPMIDriver) Close() error {
	// Set fan mode OPTIMAL
	return d.IPMIDriver.Close()
}
