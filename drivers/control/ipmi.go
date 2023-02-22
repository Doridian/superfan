package control

import (
	"github.com/u-root/u-root/pkg/ipmi"
)

type IPMIDriver struct {
	dev         *ipmi.IPMI
	DeviceIndex int
}

func (d *IPMIDriver) Init() error {
	dev, err := ipmi.Open(d.DeviceIndex)
	if err != nil {
		return err
	}
	d.dev = dev
	return nil
}

func (d *IPMIDriver) Close() error {
	if d.dev != nil {
		return d.dev.Close()
	}
	return nil
}
