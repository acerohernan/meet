// Code generated by counterfeiter. DO NOT EDIT.
package rtcfakes

import (
	"context"
	"net/http"
	"sync"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/rtc"
)

type FakeRTCManager struct {
	CreateRoomStub        func(context.Context, string) error
	createRoomMutex       sync.RWMutex
	createRoomArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	createRoomReturns struct {
		result1 error
	}
	createRoomReturnsOnCall map[int]struct {
		result1 error
	}
	GetRoomStub        func(context.Context, string) (*rtc.Room, error)
	getRoomMutex       sync.RWMutex
	getRoomArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getRoomReturns struct {
		result1 *rtc.Room
		result2 error
	}
	getRoomReturnsOnCall map[int]struct {
		result1 *rtc.Room
		result2 error
	}
	ServeHTTPStub        func(http.ResponseWriter, *http.Request)
	serveHTTPMutex       sync.RWMutex
	serveHTTPArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	StartParticipantSignalStub        func(string, *auth.Grants) (*core.SignalResponse, error)
	startParticipantSignalMutex       sync.RWMutex
	startParticipantSignalArgsForCall []struct {
		arg1 string
		arg2 *auth.Grants
	}
	startParticipantSignalReturns struct {
		result1 *core.SignalResponse
		result2 error
	}
	startParticipantSignalReturnsOnCall map[int]struct {
		result1 *core.SignalResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRTCManager) CreateRoom(arg1 context.Context, arg2 string) error {
	fake.createRoomMutex.Lock()
	ret, specificReturn := fake.createRoomReturnsOnCall[len(fake.createRoomArgsForCall)]
	fake.createRoomArgsForCall = append(fake.createRoomArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateRoomStub
	fakeReturns := fake.createRoomReturns
	fake.recordInvocation("CreateRoom", []interface{}{arg1, arg2})
	fake.createRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRTCManager) CreateRoomCallCount() int {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	return len(fake.createRoomArgsForCall)
}

func (fake *FakeRTCManager) CreateRoomCalls(stub func(context.Context, string) error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = stub
}

func (fake *FakeRTCManager) CreateRoomArgsForCall(i int) (context.Context, string) {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	argsForCall := fake.createRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRTCManager) CreateRoomReturns(result1 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	fake.createRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRTCManager) CreateRoomReturnsOnCall(i int, result1 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	if fake.createRoomReturnsOnCall == nil {
		fake.createRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRTCManager) GetRoom(arg1 context.Context, arg2 string) (*rtc.Room, error) {
	fake.getRoomMutex.Lock()
	ret, specificReturn := fake.getRoomReturnsOnCall[len(fake.getRoomArgsForCall)]
	fake.getRoomArgsForCall = append(fake.getRoomArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetRoomStub
	fakeReturns := fake.getRoomReturns
	fake.recordInvocation("GetRoom", []interface{}{arg1, arg2})
	fake.getRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRTCManager) GetRoomCallCount() int {
	fake.getRoomMutex.RLock()
	defer fake.getRoomMutex.RUnlock()
	return len(fake.getRoomArgsForCall)
}

func (fake *FakeRTCManager) GetRoomCalls(stub func(context.Context, string) (*rtc.Room, error)) {
	fake.getRoomMutex.Lock()
	defer fake.getRoomMutex.Unlock()
	fake.GetRoomStub = stub
}

func (fake *FakeRTCManager) GetRoomArgsForCall(i int) (context.Context, string) {
	fake.getRoomMutex.RLock()
	defer fake.getRoomMutex.RUnlock()
	argsForCall := fake.getRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRTCManager) GetRoomReturns(result1 *rtc.Room, result2 error) {
	fake.getRoomMutex.Lock()
	defer fake.getRoomMutex.Unlock()
	fake.GetRoomStub = nil
	fake.getRoomReturns = struct {
		result1 *rtc.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRTCManager) GetRoomReturnsOnCall(i int, result1 *rtc.Room, result2 error) {
	fake.getRoomMutex.Lock()
	defer fake.getRoomMutex.Unlock()
	fake.GetRoomStub = nil
	if fake.getRoomReturnsOnCall == nil {
		fake.getRoomReturnsOnCall = make(map[int]struct {
			result1 *rtc.Room
			result2 error
		})
	}
	fake.getRoomReturnsOnCall[i] = struct {
		result1 *rtc.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRTCManager) ServeHTTP(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.serveHTTPMutex.Lock()
	fake.serveHTTPArgsForCall = append(fake.serveHTTPArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	stub := fake.ServeHTTPStub
	fake.recordInvocation("ServeHTTP", []interface{}{arg1, arg2})
	fake.serveHTTPMutex.Unlock()
	if stub != nil {
		fake.ServeHTTPStub(arg1, arg2)
	}
}

func (fake *FakeRTCManager) ServeHTTPCallCount() int {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return len(fake.serveHTTPArgsForCall)
}

func (fake *FakeRTCManager) ServeHTTPCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.serveHTTPMutex.Lock()
	defer fake.serveHTTPMutex.Unlock()
	fake.ServeHTTPStub = stub
}

func (fake *FakeRTCManager) ServeHTTPArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	argsForCall := fake.serveHTTPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRTCManager) StartParticipantSignal(arg1 string, arg2 *auth.Grants) (*core.SignalResponse, error) {
	fake.startParticipantSignalMutex.Lock()
	ret, specificReturn := fake.startParticipantSignalReturnsOnCall[len(fake.startParticipantSignalArgsForCall)]
	fake.startParticipantSignalArgsForCall = append(fake.startParticipantSignalArgsForCall, struct {
		arg1 string
		arg2 *auth.Grants
	}{arg1, arg2})
	stub := fake.StartParticipantSignalStub
	fakeReturns := fake.startParticipantSignalReturns
	fake.recordInvocation("StartParticipantSignal", []interface{}{arg1, arg2})
	fake.startParticipantSignalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRTCManager) StartParticipantSignalCallCount() int {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	return len(fake.startParticipantSignalArgsForCall)
}

func (fake *FakeRTCManager) StartParticipantSignalCalls(stub func(string, *auth.Grants) (*core.SignalResponse, error)) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = stub
}

func (fake *FakeRTCManager) StartParticipantSignalArgsForCall(i int) (string, *auth.Grants) {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	argsForCall := fake.startParticipantSignalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRTCManager) StartParticipantSignalReturns(result1 *core.SignalResponse, result2 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	fake.startParticipantSignalReturns = struct {
		result1 *core.SignalResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRTCManager) StartParticipantSignalReturnsOnCall(i int, result1 *core.SignalResponse, result2 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	if fake.startParticipantSignalReturnsOnCall == nil {
		fake.startParticipantSignalReturnsOnCall = make(map[int]struct {
			result1 *core.SignalResponse
			result2 error
		})
	}
	fake.startParticipantSignalReturnsOnCall[i] = struct {
		result1 *core.SignalResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRTCManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	fake.getRoomMutex.RLock()
	defer fake.getRoomMutex.RUnlock()
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRTCManager) recordInvocation(key string, args []interface{}) {
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

var _ rtc.RTCManager = new(FakeRTCManager)
