// Code generated by counterfeiter. DO NOT EDIT.
package stopperfakes

import (
	"sync"
	"syscall"

	"code.cloudfoundry.org/guardian/rundmc/stopper"
)

type FakeKiller struct {
	KillStub        func(signal syscall.Signal, pid ...int)
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		signal syscall.Signal
		pid    []int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKiller) Kill(signal syscall.Signal, pid ...int) {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		signal syscall.Signal
		pid    []int
	}{signal, pid})
	fake.recordInvocation("Kill", []interface{}{signal, pid})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		fake.KillStub(signal, pid...)
	}
}

func (fake *FakeKiller) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeKiller) KillArgsForCall(i int) (syscall.Signal, []int) {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return fake.killArgsForCall[i].signal, fake.killArgsForCall[i].pid
}

func (fake *FakeKiller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKiller) recordInvocation(key string, args []interface{}) {
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

var _ stopper.Killer = new(FakeKiller)
