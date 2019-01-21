
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/bccsp/idemix/handlers"
)

type User struct {
	NewKeyStub        func() (handlers.Big, error)
	newKeyMutex       sync.RWMutex
	newKeyArgsForCall []struct{}
	newKeyReturns     struct {
		result1 handlers.Big
		result2 error
	}
	newKeyReturnsOnCall map[int]struct {
		result1 handlers.Big
		result2 error
	}
	NewKeyFromBytesStub        func(raw []byte) (handlers.Big, error)
	newKeyFromBytesMutex       sync.RWMutex
	newKeyFromBytesArgsForCall []struct {
		raw []byte
	}
	newKeyFromBytesReturns struct {
		result1 handlers.Big
		result2 error
	}
	newKeyFromBytesReturnsOnCall map[int]struct {
		result1 handlers.Big
		result2 error
	}
	MakeNymStub        func(sk handlers.Big, key handlers.IssuerPublicKey) (handlers.Ecp, handlers.Big, error)
	makeNymMutex       sync.RWMutex
	makeNymArgsForCall []struct {
		sk  handlers.Big
		key handlers.IssuerPublicKey
	}
	makeNymReturns struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}
	makeNymReturnsOnCall map[int]struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}
	NewPublicNymFromBytesStub        func(raw []byte) (handlers.Ecp, error)
	newPublicNymFromBytesMutex       sync.RWMutex
	newPublicNymFromBytesArgsForCall []struct {
		raw []byte
	}
	newPublicNymFromBytesReturns struct {
		result1 handlers.Ecp
		result2 error
	}
	newPublicNymFromBytesReturnsOnCall map[int]struct {
		result1 handlers.Ecp
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *User) NewKey() (handlers.Big, error) {
	fake.newKeyMutex.Lock()
	ret, specificReturn := fake.newKeyReturnsOnCall[len(fake.newKeyArgsForCall)]
	fake.newKeyArgsForCall = append(fake.newKeyArgsForCall, struct{}{})
	fake.recordInvocation("NewKey", []interface{}{})
	fake.newKeyMutex.Unlock()
	if fake.NewKeyStub != nil {
		return fake.NewKeyStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newKeyReturns.result1, fake.newKeyReturns.result2
}

func (fake *User) NewKeyCallCount() int {
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	return len(fake.newKeyArgsForCall)
}

func (fake *User) NewKeyReturns(result1 handlers.Big, result2 error) {
	fake.NewKeyStub = nil
	fake.newKeyReturns = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyReturnsOnCall(i int, result1 handlers.Big, result2 error) {
	fake.NewKeyStub = nil
	if fake.newKeyReturnsOnCall == nil {
		fake.newKeyReturnsOnCall = make(map[int]struct {
			result1 handlers.Big
			result2 error
		})
	}
	fake.newKeyReturnsOnCall[i] = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyFromBytes(raw []byte) (handlers.Big, error) {
	var rawCopy []byte
	if raw != nil {
		rawCopy = make([]byte, len(raw))
		copy(rawCopy, raw)
	}
	fake.newKeyFromBytesMutex.Lock()
	ret, specificReturn := fake.newKeyFromBytesReturnsOnCall[len(fake.newKeyFromBytesArgsForCall)]
	fake.newKeyFromBytesArgsForCall = append(fake.newKeyFromBytesArgsForCall, struct {
		raw []byte
	}{rawCopy})
	fake.recordInvocation("NewKeyFromBytes", []interface{}{rawCopy})
	fake.newKeyFromBytesMutex.Unlock()
	if fake.NewKeyFromBytesStub != nil {
		return fake.NewKeyFromBytesStub(raw)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newKeyFromBytesReturns.result1, fake.newKeyFromBytesReturns.result2
}

func (fake *User) NewKeyFromBytesCallCount() int {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	return len(fake.newKeyFromBytesArgsForCall)
}

func (fake *User) NewKeyFromBytesArgsForCall(i int) []byte {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	return fake.newKeyFromBytesArgsForCall[i].raw
}

func (fake *User) NewKeyFromBytesReturns(result1 handlers.Big, result2 error) {
	fake.NewKeyFromBytesStub = nil
	fake.newKeyFromBytesReturns = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyFromBytesReturnsOnCall(i int, result1 handlers.Big, result2 error) {
	fake.NewKeyFromBytesStub = nil
	if fake.newKeyFromBytesReturnsOnCall == nil {
		fake.newKeyFromBytesReturnsOnCall = make(map[int]struct {
			result1 handlers.Big
			result2 error
		})
	}
	fake.newKeyFromBytesReturnsOnCall[i] = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) MakeNym(sk handlers.Big, key handlers.IssuerPublicKey) (handlers.Ecp, handlers.Big, error) {
	fake.makeNymMutex.Lock()
	ret, specificReturn := fake.makeNymReturnsOnCall[len(fake.makeNymArgsForCall)]
	fake.makeNymArgsForCall = append(fake.makeNymArgsForCall, struct {
		sk  handlers.Big
		key handlers.IssuerPublicKey
	}{sk, key})
	fake.recordInvocation("MakeNym", []interface{}{sk, key})
	fake.makeNymMutex.Unlock()
	if fake.MakeNymStub != nil {
		return fake.MakeNymStub(sk, key)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.makeNymReturns.result1, fake.makeNymReturns.result2, fake.makeNymReturns.result3
}

func (fake *User) MakeNymCallCount() int {
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	return len(fake.makeNymArgsForCall)
}

func (fake *User) MakeNymArgsForCall(i int) (handlers.Big, handlers.IssuerPublicKey) {
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	return fake.makeNymArgsForCall[i].sk, fake.makeNymArgsForCall[i].key
}

func (fake *User) MakeNymReturns(result1 handlers.Ecp, result2 handlers.Big, result3 error) {
	fake.MakeNymStub = nil
	fake.makeNymReturns = struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}{result1, result2, result3}
}

func (fake *User) MakeNymReturnsOnCall(i int, result1 handlers.Ecp, result2 handlers.Big, result3 error) {
	fake.MakeNymStub = nil
	if fake.makeNymReturnsOnCall == nil {
		fake.makeNymReturnsOnCall = make(map[int]struct {
			result1 handlers.Ecp
			result2 handlers.Big
			result3 error
		})
	}
	fake.makeNymReturnsOnCall[i] = struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}{result1, result2, result3}
}

func (fake *User) NewPublicNymFromBytes(raw []byte) (handlers.Ecp, error) {
	var rawCopy []byte
	if raw != nil {
		rawCopy = make([]byte, len(raw))
		copy(rawCopy, raw)
	}
	fake.newPublicNymFromBytesMutex.Lock()
	ret, specificReturn := fake.newPublicNymFromBytesReturnsOnCall[len(fake.newPublicNymFromBytesArgsForCall)]
	fake.newPublicNymFromBytesArgsForCall = append(fake.newPublicNymFromBytesArgsForCall, struct {
		raw []byte
	}{rawCopy})
	fake.recordInvocation("NewPublicNymFromBytes", []interface{}{rawCopy})
	fake.newPublicNymFromBytesMutex.Unlock()
	if fake.NewPublicNymFromBytesStub != nil {
		return fake.NewPublicNymFromBytesStub(raw)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newPublicNymFromBytesReturns.result1, fake.newPublicNymFromBytesReturns.result2
}

func (fake *User) NewPublicNymFromBytesCallCount() int {
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	return len(fake.newPublicNymFromBytesArgsForCall)
}

func (fake *User) NewPublicNymFromBytesArgsForCall(i int) []byte {
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	return fake.newPublicNymFromBytesArgsForCall[i].raw
}

func (fake *User) NewPublicNymFromBytesReturns(result1 handlers.Ecp, result2 error) {
	fake.NewPublicNymFromBytesStub = nil
	fake.newPublicNymFromBytesReturns = struct {
		result1 handlers.Ecp
		result2 error
	}{result1, result2}
}

func (fake *User) NewPublicNymFromBytesReturnsOnCall(i int, result1 handlers.Ecp, result2 error) {
	fake.NewPublicNymFromBytesStub = nil
	if fake.newPublicNymFromBytesReturnsOnCall == nil {
		fake.newPublicNymFromBytesReturnsOnCall = make(map[int]struct {
			result1 handlers.Ecp
			result2 error
		})
	}
	fake.newPublicNymFromBytesReturnsOnCall[i] = struct {
		result1 handlers.Ecp
		result2 error
	}{result1, result2}
}

func (fake *User) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *User) recordInvocation(key string, args []interface{}) {
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

var _ handlers.User = new(User)