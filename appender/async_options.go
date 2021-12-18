package appender

import (
	"errors"
	"time"
	"zap_ing/appender/appendercore"
)

type AsyncOption interface {
	apply(*Async) error
}

type asyncOptionsFunc func(*Async) error

func (f asyncOptionsFunc) apply(a *Async) error {
	return f(a)
}

func AsyncMaxQueueLength(length uint32) AsyncOption {
	return asyncOptionsFunc(func(a *Async) error {
		a.maxQueueLength = length
		return nil
	})
}

// AsyncOnQueueNearlyFullForwardTo
// fallback is wrapped in a Synchronizing appender.
func AsyncOnQueueNearlyFullForwardTo(fallback appendercore.Appender) AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		if fallback == nil {
			return errors.New("fallback must not be nil")
		}
		async.fallback = appendercore.NewSynchronizing(fallback)
		return nil
	})
}

func AsyncOnQueueNearlyFullDropMessages() AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		async.fallback = NewDiscard()
		return nil
	})
}

func AsyncQueueMinFreePercent(minFreePercent float32) AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		if minFreePercent < 0 || minFreePercent >= 1 {
			return errors.New("minFreePercent must be between 0 and 1")
		}
		async.calculateDropThresholdFn = func(a *Async) (uint32, error) {
			threshold := float32(async.maxQueueLength) * minFreePercent
			return uint32(threshold), nil
		}
		return nil
	})
}

func AsyncQueueMinFreeItems(minFree uint32) AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		async.calculateDropThresholdFn = func(a *Async) (uint32, error) {
			if a.maxQueueLength < minFree {
				return 0, errors.New("minFree must less than the max queue size")
			}
			return minFree, nil
		}
		return nil
	})
}

func AsyncQueueMonitorPeriod(period time.Duration) AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		if period <= time.Duration(0) {
			return errors.New("period must be positive")
		}
		async.monitorPeriod = period
		return nil
	})
}

func AsyncSyncTimeout(timeout time.Duration) AsyncOption {
	return asyncOptionsFunc(func(async *Async) error {
		if timeout <= time.Duration(0) {
			return errors.New("timeout must be positive")
		}
		async.syncTimeout = timeout
		return nil
	})
}