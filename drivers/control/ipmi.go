package control

import (
	"github.com/gofrs/flock"
	"github.com/u-root/u-root/pkg/ipmi"
)

var lockFile = flock.New("/var/lock/ipmi.lock")

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
