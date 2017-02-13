// This file was generated by counterfeiter
package dadoofakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/dadoo"
)

type FakePidGetter struct {
	PidStub        func(pidFilePath string) (int, error)
	pidMutex       sync.RWMutex
	pidArgsForCall []struct {
		pidFilePath string
	}
	pidReturns struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePidGetter) Pid(pidFilePath string) (int, error) {
	fake.pidMutex.Lock()
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct {
		pidFilePath string
	}{pidFilePath})
	fake.recordInvocation("Pid", []interface{}{pidFilePath})
	fake.pidMutex.Unlock()
	if fake.PidStub != nil {
		return fake.PidStub(pidFilePath)
	} else {
		return fake.pidReturns.result1, fake.pidReturns.result2
	}
}

func (fake *FakePidGetter) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

func (fake *FakePidGetter) PidArgsForCall(i int) string {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return fake.pidArgsForCall[i].pidFilePath
}

func (fake *FakePidGetter) PidReturns(result1 int, result2 error) {
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakePidGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePidGetter) recordInvocation(key string, args []interface{}) {
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

var _ dadoo.PidGetter = new(FakePidGetter)
