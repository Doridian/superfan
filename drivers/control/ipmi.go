package control

import (
	"github.com/u-root/u-root/pkg/ipmi"
)

type IPMIDriver struct {
	dev *ipmi.IPMI
}

func (d *IPMIDriver) Init() error {
	dev, err := ipmi.Open(0)
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
