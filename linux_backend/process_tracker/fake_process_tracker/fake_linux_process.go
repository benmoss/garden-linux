// This file was generated by counterfeiter
package fake_process_tracker

import (
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/cloudfoundry-incubator/warden-linux/linux_backend/process_tracker"

	"sync"
)

type FakeLinuxProcess struct {
	IDStub        func() uint32
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns struct {
		result1 uint32
	}
	WaitStub        func() (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns struct {
		result1 int
		result2 error
	}
	SetTTYStub        func(warden.TTYSpec) error
	setTTYMutex       sync.RWMutex
	setTTYArgsForCall []struct {
		arg1 warden.TTYSpec
	}
	setTTYReturns struct {
		result1 error
	}
	WithTTYStub        func() bool
	withTTYMutex       sync.RWMutex
	withTTYArgsForCall []struct{}
	withTTYReturns struct {
		result1 bool
	}
}

func (fake *FakeLinuxProcess) ID() uint32 {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeLinuxProcess) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeLinuxProcess) IDReturns(result1 uint32) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeLinuxProcess) Wait() (int, error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1, fake.waitReturns.result2
	}
}

func (fake *FakeLinuxProcess) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeLinuxProcess) WaitReturns(result1 int, result2 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLinuxProcess) SetTTY(arg1 warden.TTYSpec) error {
	fake.setTTYMutex.Lock()
	defer fake.setTTYMutex.Unlock()
	fake.setTTYArgsForCall = append(fake.setTTYArgsForCall, struct {
		arg1 warden.TTYSpec
	}{arg1})
	if fake.SetTTYStub != nil {
		return fake.SetTTYStub(arg1)
	} else {
		return fake.setTTYReturns.result1
	}
}

func (fake *FakeLinuxProcess) SetTTYCallCount() int {
	fake.setTTYMutex.RLock()
	defer fake.setTTYMutex.RUnlock()
	return len(fake.setTTYArgsForCall)
}

func (fake *FakeLinuxProcess) SetTTYArgsForCall(i int) warden.TTYSpec {
	fake.setTTYMutex.RLock()
	defer fake.setTTYMutex.RUnlock()
	return fake.setTTYArgsForCall[i].arg1
}

func (fake *FakeLinuxProcess) SetTTYReturns(result1 error) {
	fake.SetTTYStub = nil
	fake.setTTYReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLinuxProcess) WithTTY() bool {
	fake.withTTYMutex.Lock()
	defer fake.withTTYMutex.Unlock()
	fake.withTTYArgsForCall = append(fake.withTTYArgsForCall, struct{}{})
	if fake.WithTTYStub != nil {
		return fake.WithTTYStub()
	} else {
		return fake.withTTYReturns.result1
	}
}

func (fake *FakeLinuxProcess) WithTTYCallCount() int {
	fake.withTTYMutex.RLock()
	defer fake.withTTYMutex.RUnlock()
	return len(fake.withTTYArgsForCall)
}

func (fake *FakeLinuxProcess) WithTTYReturns(result1 bool) {
	fake.WithTTYStub = nil
	fake.withTTYReturns = struct {
		result1 bool
	}{result1}
}

var _ process_tracker.LinuxProcess = new(FakeLinuxProcess)
