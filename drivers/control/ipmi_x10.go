package control

import (
	"fmt"
)

type X10IPMIDriver struct {
	IPMIDriver
}

func (d *X10IPMIDriver) Init() error {
	err := d.IPMIDriver.Init()
	if err != nil {
		return err
	}
	// Set fan mode FULL
	_, err = d.dev.RawCmd([]byte{0x30, 0x45, 0x01, 0x01})
	if err != nil {
		return err
	}
	return nil
}

func (d *X10IPMIDriver) SetFanSpeed(speed float64) error {
	speedByte := byte(speed * 100)
	for i := byte(0); i < 2; i++ {
		_, err := d.dev.RawCmd([]byte{0x30, 0x70, 0x66, 0x01, i, speedByte})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *X10IPMIDriver) GetFanSpeed() (float64, error) {
	resp, err := d.dev.RawCmd([]byte{0x30, 0x70, 0x66, 0x00, 0x00})
	if err != nil {
		return 0, err
	}
	if len(resp) != 2 || resp[0] != 0 {
		return 0, fmt.Errorf("unexpected response %v", resp)
	}
	return float64(resp[1]) / 100, err
}

func (d *X10IPMIDriver) Close() error {
	// Set fan mode OPTIMAL
	_, err := d.dev.RawCmd([]byte{0x30, 0x45, 0x01, 0x02})
	if err != nil {
		return err
	}
	return d.IPMIDriver.Close()
}
