// Code generated by counterfeiter. DO NOT EDIT.
package signalingfakes

import (
	"sync"

	"github.com/ryanrolds/club/pkg/signaling"
)

type FakePeerConnection struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ReadMessageStub        func() (int, []byte, error)
	readMessageMutex       sync.RWMutex
	readMessageArgsForCall []struct {
	}
	readMessageReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	readMessageReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	WriteJSONStub        func(interface{}) error
	writeJSONMutex       sync.RWMutex
	writeJSONArgsForCall []struct {
		arg1 interface{}
	}
	writeJSONReturns struct {
		result1 error
	}
	writeJSONReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePeerConnection) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakePeerConnection) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePeerConnection) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakePeerConnection) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) ReadMessage() (int, []byte, error) {
	fake.readMessageMutex.Lock()
	ret, specificReturn := fake.readMessageReturnsOnCall[len(fake.readMessageArgsForCall)]
	fake.readMessageArgsForCall = append(fake.readMessageArgsForCall, struct {
	}{})
	fake.recordInvocation("ReadMessage", []interface{}{})
	fake.readMessageMutex.Unlock()
	if fake.ReadMessageStub != nil {
		return fake.ReadMessageStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.readMessageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePeerConnection) ReadMessageCallCount() int {
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	return len(fake.readMessageArgsForCall)
}

func (fake *FakePeerConnection) ReadMessageCalls(stub func() (int, []byte, error)) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = stub
}

func (fake *FakePeerConnection) ReadMessageReturns(result1 int, result2 []byte, result3 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	fake.readMessageReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePeerConnection) ReadMessageReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	if fake.readMessageReturnsOnCall == nil {
		fake.readMessageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.readMessageReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePeerConnection) WriteJSON(arg1 interface{}) error {
	fake.writeJSONMutex.Lock()
	ret, specificReturn := fake.writeJSONReturnsOnCall[len(fake.writeJSONArgsForCall)]
	fake.writeJSONArgsForCall = append(fake.writeJSONArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("WriteJSON", []interface{}{arg1})
	fake.writeJSONMutex.Unlock()
	if fake.WriteJSONStub != nil {
		return fake.WriteJSONStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeJSONReturns
	return fakeReturns.result1
}

func (fake *FakePeerConnection) WriteJSONCallCount() int {
	fake.writeJSONMutex.RLock()
	defer fake.writeJSONMutex.RUnlock()
	return len(fake.writeJSONArgsForCall)
}

func (fake *FakePeerConnection) WriteJSONCalls(stub func(interface{}) error) {
	fake.writeJSONMutex.Lock()
	defer fake.writeJSONMutex.Unlock()
	fake.WriteJSONStub = stub
}

func (fake *FakePeerConnection) WriteJSONArgsForCall(i int) interface{} {
	fake.writeJSONMutex.RLock()
	defer fake.writeJSONMutex.RUnlock()
	argsForCall := fake.writeJSONArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) WriteJSONReturns(result1 error) {
	fake.writeJSONMutex.Lock()
	defer fake.writeJSONMutex.Unlock()
	fake.WriteJSONStub = nil
	fake.writeJSONReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) WriteJSONReturnsOnCall(i int, result1 error) {
	fake.writeJSONMutex.Lock()
	defer fake.writeJSONMutex.Unlock()
	fake.WriteJSONStub = nil
	if fake.writeJSONReturnsOnCall == nil {
		fake.writeJSONReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeJSONReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	fake.writeJSONMutex.RLock()
	defer fake.writeJSONMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePeerConnection) recordInvocation(key string, args []interface{}) {
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

var _ signaling.PeerConnection = new(FakePeerConnection)