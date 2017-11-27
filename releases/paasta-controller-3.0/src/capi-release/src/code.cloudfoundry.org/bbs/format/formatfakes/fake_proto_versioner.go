// This file was generated by counterfeiter
package formatfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/format"
)

type FakeProtoVersioner struct {
	ResetStub         func()
	resetMutex        sync.RWMutex
	resetArgsForCall  []struct{}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	ProtoMessageStub        func()
	protoMessageMutex       sync.RWMutex
	protoMessageArgsForCall []struct{}
	VersionStub             func() format.Version
	versionMutex            sync.RWMutex
	versionArgsForCall      []struct{}
	versionReturns          struct {
		result1 format.Version
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct{}
	validateReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProtoVersioner) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct{}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *FakeProtoVersioner) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *FakeProtoVersioner) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	} else {
		return fake.stringReturns.result1
	}
}

func (fake *FakeProtoVersioner) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeProtoVersioner) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeProtoVersioner) ProtoMessage() {
	fake.protoMessageMutex.Lock()
	fake.protoMessageArgsForCall = append(fake.protoMessageArgsForCall, struct{}{})
	fake.recordInvocation("ProtoMessage", []interface{}{})
	fake.protoMessageMutex.Unlock()
	if fake.ProtoMessageStub != nil {
		fake.ProtoMessageStub()
	}
}

func (fake *FakeProtoVersioner) ProtoMessageCallCount() int {
	fake.protoMessageMutex.RLock()
	defer fake.protoMessageMutex.RUnlock()
	return len(fake.protoMessageArgsForCall)
}

func (fake *FakeProtoVersioner) Version() format.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeProtoVersioner) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeProtoVersioner) VersionReturns(result1 format.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 format.Version
	}{result1}
}

func (fake *FakeProtoVersioner) Validate() error {
	fake.validateMutex.Lock()
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct{}{})
	fake.recordInvocation("Validate", []interface{}{})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub()
	} else {
		return fake.validateReturns.result1
	}
}

func (fake *FakeProtoVersioner) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeProtoVersioner) ValidateReturns(result1 error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProtoVersioner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.protoMessageMutex.RLock()
	defer fake.protoMessageMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProtoVersioner) recordInvocation(key string, args []interface{}) {
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

var _ format.ProtoVersioner = new(FakeProtoVersioner)