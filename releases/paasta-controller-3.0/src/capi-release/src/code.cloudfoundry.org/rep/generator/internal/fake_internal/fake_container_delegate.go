// This file was generated by counterfeiter
package fake_internal

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/rep/generator/internal"
)

type FakeContainerDelegate struct {
	GetContainerStub        func(logger lager.Logger, guid string) (executor.Container, bool)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	getContainerReturns struct {
		result1 executor.Container
		result2 bool
	}
	RunContainerStub        func(logger lager.Logger, req *executor.RunRequest) bool
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		logger lager.Logger
		req    *executor.RunRequest
	}
	runContainerReturns struct {
		result1 bool
	}
	StopContainerStub        func(logger lager.Logger, guid string) bool
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	stopContainerReturns struct {
		result1 bool
	}
	DeleteContainerStub        func(logger lager.Logger, guid string) bool
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	deleteContainerReturns struct {
		result1 bool
	}
	FetchContainerResultFileStub        func(logger lager.Logger, guid string, filename string) (string, error)
	fetchContainerResultFileMutex       sync.RWMutex
	fetchContainerResultFileArgsForCall []struct {
		logger   lager.Logger
		guid     string
		filename string
	}
	fetchContainerResultFileReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerDelegate) GetContainer(logger lager.Logger, guid string) (executor.Container, bool) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.recordInvocation("GetContainer", []interface{}{logger, guid})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(logger, guid)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2
	}
}

func (fake *FakeContainerDelegate) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeContainerDelegate) GetContainerArgsForCall(i int) (lager.Logger, string) {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].logger, fake.getContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) GetContainerReturns(result1 executor.Container, result2 bool) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 executor.Container
		result2 bool
	}{result1, result2}
}

func (fake *FakeContainerDelegate) RunContainer(logger lager.Logger, req *executor.RunRequest) bool {
	fake.runContainerMutex.Lock()
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		logger lager.Logger
		req    *executor.RunRequest
	}{logger, req})
	fake.recordInvocation("RunContainer", []interface{}{logger, req})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(logger, req)
	} else {
		return fake.runContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeContainerDelegate) RunContainerArgsForCall(i int) (lager.Logger, *executor.RunRequest) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].logger, fake.runContainerArgsForCall[i].req
}

func (fake *FakeContainerDelegate) RunContainerReturns(result1 bool) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) StopContainer(logger lager.Logger, guid string) bool {
	fake.stopContainerMutex.Lock()
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.recordInvocation("StopContainer", []interface{}{logger, guid})
	fake.stopContainerMutex.Unlock()
	if fake.StopContainerStub != nil {
		return fake.StopContainerStub(logger, guid)
	} else {
		return fake.stopContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeContainerDelegate) StopContainerArgsForCall(i int) (lager.Logger, string) {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return fake.stopContainerArgsForCall[i].logger, fake.stopContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) StopContainerReturns(result1 bool) {
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) DeleteContainer(logger lager.Logger, guid string) bool {
	fake.deleteContainerMutex.Lock()
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.recordInvocation("DeleteContainer", []interface{}{logger, guid})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(logger, guid)
	} else {
		return fake.deleteContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeContainerDelegate) DeleteContainerArgsForCall(i int) (lager.Logger, string) {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].logger, fake.deleteContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) DeleteContainerReturns(result1 bool) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) FetchContainerResultFile(logger lager.Logger, guid string, filename string) (string, error) {
	fake.fetchContainerResultFileMutex.Lock()
	fake.fetchContainerResultFileArgsForCall = append(fake.fetchContainerResultFileArgsForCall, struct {
		logger   lager.Logger
		guid     string
		filename string
	}{logger, guid, filename})
	fake.recordInvocation("FetchContainerResultFile", []interface{}{logger, guid, filename})
	fake.fetchContainerResultFileMutex.Unlock()
	if fake.FetchContainerResultFileStub != nil {
		return fake.FetchContainerResultFileStub(logger, guid, filename)
	} else {
		return fake.fetchContainerResultFileReturns.result1, fake.fetchContainerResultFileReturns.result2
	}
}

func (fake *FakeContainerDelegate) FetchContainerResultFileCallCount() int {
	fake.fetchContainerResultFileMutex.RLock()
	defer fake.fetchContainerResultFileMutex.RUnlock()
	return len(fake.fetchContainerResultFileArgsForCall)
}

func (fake *FakeContainerDelegate) FetchContainerResultFileArgsForCall(i int) (lager.Logger, string, string) {
	fake.fetchContainerResultFileMutex.RLock()
	defer fake.fetchContainerResultFileMutex.RUnlock()
	return fake.fetchContainerResultFileArgsForCall[i].logger, fake.fetchContainerResultFileArgsForCall[i].guid, fake.fetchContainerResultFileArgsForCall[i].filename
}

func (fake *FakeContainerDelegate) FetchContainerResultFileReturns(result1 string, result2 error) {
	fake.FetchContainerResultFileStub = nil
	fake.fetchContainerResultFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	fake.fetchContainerResultFileMutex.RLock()
	defer fake.fetchContainerResultFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeContainerDelegate) recordInvocation(key string, args []interface{}) {
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

var _ internal.ContainerDelegate = new(FakeContainerDelegate)