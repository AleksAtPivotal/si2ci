// Code generated by counterfeiter. DO NOT EDIT.
package radarfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/radar"
)

type FakeScanner struct {
	RunStub        func(lager.Logger, string) (time.Duration, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	runReturns struct {
		result1 time.Duration
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 time.Duration
		result2 error
	}
	ScanStub        func(lager.Logger, string) error
	scanMutex       sync.RWMutex
	scanArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	scanReturns struct {
		result1 error
	}
	scanReturnsOnCall map[int]struct {
		result1 error
	}
	ScanFromVersionStub        func(lager.Logger, string, atc.Version) error
	scanFromVersionMutex       sync.RWMutex
	scanFromVersionArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 atc.Version
	}
	scanFromVersionReturns struct {
		result1 error
	}
	scanFromVersionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScanner) Run(arg1 lager.Logger, arg2 string) (time.Duration, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runReturns.result1, fake.runReturns.result2
}

func (fake *FakeScanner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeScanner) RunArgsForCall(i int) (lager.Logger, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2
}

func (fake *FakeScanner) RunReturns(result1 time.Duration, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 time.Duration
		result2 error
	}{result1, result2}
}

func (fake *FakeScanner) RunReturnsOnCall(i int, result1 time.Duration, result2 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 time.Duration
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 time.Duration
		result2 error
	}{result1, result2}
}

func (fake *FakeScanner) Scan(arg1 lager.Logger, arg2 string) error {
	fake.scanMutex.Lock()
	ret, specificReturn := fake.scanReturnsOnCall[len(fake.scanArgsForCall)]
	fake.scanArgsForCall = append(fake.scanArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Scan", []interface{}{arg1, arg2})
	fake.scanMutex.Unlock()
	if fake.ScanStub != nil {
		return fake.ScanStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.scanReturns.result1
}

func (fake *FakeScanner) ScanCallCount() int {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return len(fake.scanArgsForCall)
}

func (fake *FakeScanner) ScanArgsForCall(i int) (lager.Logger, string) {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return fake.scanArgsForCall[i].arg1, fake.scanArgsForCall[i].arg2
}

func (fake *FakeScanner) ScanReturns(result1 error) {
	fake.ScanStub = nil
	fake.scanReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) ScanReturnsOnCall(i int, result1 error) {
	fake.ScanStub = nil
	if fake.scanReturnsOnCall == nil {
		fake.scanReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.scanReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) ScanFromVersion(arg1 lager.Logger, arg2 string, arg3 atc.Version) error {
	fake.scanFromVersionMutex.Lock()
	ret, specificReturn := fake.scanFromVersionReturnsOnCall[len(fake.scanFromVersionArgsForCall)]
	fake.scanFromVersionArgsForCall = append(fake.scanFromVersionArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 atc.Version
	}{arg1, arg2, arg3})
	fake.recordInvocation("ScanFromVersion", []interface{}{arg1, arg2, arg3})
	fake.scanFromVersionMutex.Unlock()
	if fake.ScanFromVersionStub != nil {
		return fake.ScanFromVersionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.scanFromVersionReturns.result1
}

func (fake *FakeScanner) ScanFromVersionCallCount() int {
	fake.scanFromVersionMutex.RLock()
	defer fake.scanFromVersionMutex.RUnlock()
	return len(fake.scanFromVersionArgsForCall)
}

func (fake *FakeScanner) ScanFromVersionArgsForCall(i int) (lager.Logger, string, atc.Version) {
	fake.scanFromVersionMutex.RLock()
	defer fake.scanFromVersionMutex.RUnlock()
	return fake.scanFromVersionArgsForCall[i].arg1, fake.scanFromVersionArgsForCall[i].arg2, fake.scanFromVersionArgsForCall[i].arg3
}

func (fake *FakeScanner) ScanFromVersionReturns(result1 error) {
	fake.ScanFromVersionStub = nil
	fake.scanFromVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) ScanFromVersionReturnsOnCall(i int, result1 error) {
	fake.ScanFromVersionStub = nil
	if fake.scanFromVersionReturnsOnCall == nil {
		fake.scanFromVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.scanFromVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	fake.scanFromVersionMutex.RLock()
	defer fake.scanFromVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScanner) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ radar.Scanner = new(FakeScanner)
