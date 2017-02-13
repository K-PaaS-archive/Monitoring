// This file was generated by counterfeiter
package auctioneerfakes

import (
	"sync"

	"code.cloudfoundry.org/auctioneer"
)

type FakeClient struct {
	RequestLRPAuctionsStub        func(lrpStart []*auctioneer.LRPStartRequest) error
	requestLRPAuctionsMutex       sync.RWMutex
	requestLRPAuctionsArgsForCall []struct {
		lrpStart []*auctioneer.LRPStartRequest
	}
	requestLRPAuctionsReturns struct {
		result1 error
	}
	RequestTaskAuctionsStub        func(tasks []*auctioneer.TaskStartRequest) error
	requestTaskAuctionsMutex       sync.RWMutex
	requestTaskAuctionsArgsForCall []struct {
		tasks []*auctioneer.TaskStartRequest
	}
	requestTaskAuctionsReturns struct {
		result1 error
	}
}

func (fake *FakeClient) RequestLRPAuctions(lrpStart []*auctioneer.LRPStartRequest) error {
	fake.requestLRPAuctionsMutex.Lock()
	fake.requestLRPAuctionsArgsForCall = append(fake.requestLRPAuctionsArgsForCall, struct {
		lrpStart []*auctioneer.LRPStartRequest
	}{lrpStart})
	fake.requestLRPAuctionsMutex.Unlock()
	if fake.RequestLRPAuctionsStub != nil {
		return fake.RequestLRPAuctionsStub(lrpStart)
	} else {
		return fake.requestLRPAuctionsReturns.result1
	}
}

func (fake *FakeClient) RequestLRPAuctionsCallCount() int {
	fake.requestLRPAuctionsMutex.RLock()
	defer fake.requestLRPAuctionsMutex.RUnlock()
	return len(fake.requestLRPAuctionsArgsForCall)
}

func (fake *FakeClient) RequestLRPAuctionsArgsForCall(i int) []*auctioneer.LRPStartRequest {
	fake.requestLRPAuctionsMutex.RLock()
	defer fake.requestLRPAuctionsMutex.RUnlock()
	return fake.requestLRPAuctionsArgsForCall[i].lrpStart
}

func (fake *FakeClient) RequestLRPAuctionsReturns(result1 error) {
	fake.RequestLRPAuctionsStub = nil
	fake.requestLRPAuctionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RequestTaskAuctions(tasks []*auctioneer.TaskStartRequest) error {
	fake.requestTaskAuctionsMutex.Lock()
	fake.requestTaskAuctionsArgsForCall = append(fake.requestTaskAuctionsArgsForCall, struct {
		tasks []*auctioneer.TaskStartRequest
	}{tasks})
	fake.requestTaskAuctionsMutex.Unlock()
	if fake.RequestTaskAuctionsStub != nil {
		return fake.RequestTaskAuctionsStub(tasks)
	} else {
		return fake.requestTaskAuctionsReturns.result1
	}
}

func (fake *FakeClient) RequestTaskAuctionsCallCount() int {
	fake.requestTaskAuctionsMutex.RLock()
	defer fake.requestTaskAuctionsMutex.RUnlock()
	return len(fake.requestTaskAuctionsArgsForCall)
}

func (fake *FakeClient) RequestTaskAuctionsArgsForCall(i int) []*auctioneer.TaskStartRequest {
	fake.requestTaskAuctionsMutex.RLock()
	defer fake.requestTaskAuctionsMutex.RUnlock()
	return fake.requestTaskAuctionsArgsForCall[i].tasks
}

func (fake *FakeClient) RequestTaskAuctionsReturns(result1 error) {
	fake.RequestTaskAuctionsStub = nil
	fake.requestTaskAuctionsReturns = struct {
		result1 error
	}{result1}
}

var _ auctioneer.Client = new(FakeClient)
