package control

import (
	"errors"
	"time"

	"github.com/gofrs/flock"
)

var lockFile = flock.New("/var/lock/ipmi.lock")
var LockTimeout = time.Second * 10
var LockRetryInterval = time.Millisecond * 10

func acquireLock() (err error) {
	var locked bool
	lockWaitEnd := time.Now().Add(LockTimeout)
	for time.Now().Before(lockWaitEnd) && !locked {
		locked, err = lockFile.TryLock()
		if err != nil || locked {
			return
		}
		time.Sleep(LockRetryInterval)
	}
	return errors.New("could not acquire IPMI lock")
}

func RunLocked[T any](f func() (T, error)) (T, error) {
	if err := acquireLock(); err != nil {
		var zero T
		return zero, err
	}
	defer lockFile.Unlock()
	return f()
}

func RunLockedVoid(f func() error) error {
	if err := acquireLock(); err != nil {
		return err
	}
	defer lockFile.Unlock()
	return f()
}
