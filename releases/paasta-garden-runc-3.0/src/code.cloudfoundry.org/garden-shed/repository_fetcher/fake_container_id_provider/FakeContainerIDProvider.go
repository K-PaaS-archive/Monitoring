// This file was generated by counterfeiter
package fake_container_id_provider

import (
	"sync"

	"code.cloudfoundry.org/garden-shed/layercake"
	"code.cloudfoundry.org/garden-shed/repository_fetcher"
)

type FakeContainerIDProvider struct {
	ProvideIDStub        func(path string) layercake.ID
	provideIDMutex       sync.RWMutex
	provideIDArgsForCall []struct {
		path string
	}
	provideIDReturns struct {
		result1 layercake.ID
	}
	provideIDReturnsOnCall map[int]struct {
		result1 layercake.ID
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerIDProvider) ProvideID(path string) layercake.ID {
	fake.provideIDMutex.Lock()
	ret, specificReturn := fake.provideIDReturnsOnCall[len(fake.provideIDArgsForCall)]
	fake.provideIDArgsForCall = append(fake.provideIDArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ProvideID", []interface{}{path})
	fake.provideIDMutex.Unlock()
	if fake.ProvideIDStub != nil {
		return fake.ProvideIDStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.provideIDReturns.result1
}

func (fake *FakeContainerIDProvider) ProvideIDCallCount() int {
	fake.provideIDMutex.RLock()
	defer fake.provideIDMutex.RUnlock()
	return len(fake.provideIDArgsForCall)
}

func (fake *FakeContainerIDProvider) ProvideIDArgsForCall(i int) string {
	fake.provideIDMutex.RLock()
	defer fake.provideIDMutex.RUnlock()
	return fake.provideIDArgsForCall[i].path
}

func (fake *FakeContainerIDProvider) ProvideIDReturns(result1 layercake.ID) {
	fake.ProvideIDStub = nil
	fake.provideIDReturns = struct {
		result1 layercake.ID
	}{result1}
}

func (fake *FakeContainerIDProvider) ProvideIDReturnsOnCall(i int, result1 layercake.ID) {
	fake.ProvideIDStub = nil
	if fake.provideIDReturnsOnCall == nil {
		fake.provideIDReturnsOnCall = make(map[int]struct {
			result1 layercake.ID
		})
	}
	fake.provideIDReturnsOnCall[i] = struct {
		result1 layercake.ID
	}{result1}
}

func (fake *FakeContainerIDProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provideIDMutex.RLock()
	defer fake.provideIDMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeContainerIDProvider) recordInvocation(key string, args []interface{}) {
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

var _ repository_fetcher.ContainerIDProvider = new(FakeContainerIDProvider)
