// This file was generated by counterfeiter
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
	"code.cloudfoundry.org/guardian/kawasaki/subnets"
	"code.cloudfoundry.org/lager"
)

type FakeSpecParser struct {
	ParseStub        func(log lager.Logger, spec string) (subnets.SubnetSelector, subnets.IPSelector, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		log  lager.Logger
		spec string
	}
	parseReturns struct {
		result1 subnets.SubnetSelector
		result2 subnets.IPSelector
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpecParser) Parse(log lager.Logger, spec string) (subnets.SubnetSelector, subnets.IPSelector, error) {
	fake.parseMutex.Lock()
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		log  lager.Logger
		spec string
	}{log, spec})
	fake.recordInvocation("Parse", []interface{}{log, spec})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(log, spec)
	} else {
		return fake.parseReturns.result1, fake.parseReturns.result2, fake.parseReturns.result3
	}
}

func (fake *FakeSpecParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *FakeSpecParser) ParseArgsForCall(i int) (lager.Logger, string) {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].log, fake.parseArgsForCall[i].spec
}

func (fake *FakeSpecParser) ParseReturns(result1 subnets.SubnetSelector, result2 subnets.IPSelector, result3 error) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 subnets.SubnetSelector
		result2 subnets.IPSelector
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpecParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpecParser) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.SpecParser = new(FakeSpecParser)
