package control

import (
	"errors"

	"github.com/gofrs/flock"
)

var lockFile = flock.New("/var/lock/ipmi.lock")

func RunLocked[T any](f func() (T, error)) (T, error) {
	locked, err := lockFile.TryLock()
	if err != nil {
		var zero T
		return zero, err
	}
	if !locked {
		var zero T
		return zero, errors.New("could not acquire IPMI lock")
	}
	defer lockFile.Unlock()
	return f()
}

func RunLockedVoid(f func() error) error {
	locked, err := lockFile.TryLock()
	if err != nil {
		return err
	}
	if !locked {
		return errors.New("could not acquire IPMI lock")
	}
	defer lockFile.Unlock()
	return f()
}
