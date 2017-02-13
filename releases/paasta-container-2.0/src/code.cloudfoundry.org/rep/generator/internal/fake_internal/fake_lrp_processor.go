// This file was generated by counterfeiter
package fake_internal

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/rep/generator/internal"
)

type FakeLRPProcessor struct {
	ProcessStub        func(lager.Logger, executor.Container)
	processMutex       sync.RWMutex
	processArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLRPProcessor) Process(arg1 lager.Logger, arg2 executor.Container) {
	fake.processMutex.Lock()
	fake.processArgsForCall = append(fake.processArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("Process", []interface{}{arg1, arg2})
	fake.processMutex.Unlock()
	if fake.ProcessStub != nil {
		fake.ProcessStub(arg1, arg2)
	}
}

func (fake *FakeLRPProcessor) ProcessCallCount() int {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return len(fake.processArgsForCall)
}

func (fake *FakeLRPProcessor) ProcessArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return fake.processArgsForCall[i].arg1, fake.processArgsForCall[i].arg2
}

func (fake *FakeLRPProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLRPProcessor) recordInvocation(key string, args []interface{}) {
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

var _ internal.LRPProcessor = new(FakeLRPProcessor)
