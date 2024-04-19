// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	basepeer "github.com/IBM-Blockchain/fabric-operator/pkg/offering/base/peer"
	v1a "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeploymentManager struct {
	CheckForSecretChangeStub        func(v1.Object, string, func(string, *v1a.Deployment) bool) error
	checkForSecretChangeMutex       sync.RWMutex
	checkForSecretChangeArgsForCall []struct {
		arg1 v1.Object
		arg2 string
		arg3 func(string, *v1a.Deployment) bool
	}
	checkForSecretChangeReturns struct {
		result1 error
	}
	checkForSecretChangeReturnsOnCall map[int]struct {
		result1 error
	}
	CheckStateStub        func(v1.Object) error
	checkStateMutex       sync.RWMutex
	checkStateArgsForCall []struct {
		arg1 v1.Object
	}
	checkStateReturns struct {
		result1 error
	}
	checkStateReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(v1.Object) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 v1.Object
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeploymentStatusStub        func(v1.Object) (v1a.DeploymentStatus, error)
	deploymentStatusMutex       sync.RWMutex
	deploymentStatusArgsForCall []struct {
		arg1 v1.Object
	}
	deploymentStatusReturns struct {
		result1 v1a.DeploymentStatus
		result2 error
	}
	deploymentStatusReturnsOnCall map[int]struct {
		result1 v1a.DeploymentStatus
		result2 error
	}
	ExistsStub        func(v1.Object) bool
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		arg1 v1.Object
	}
	existsReturns struct {
		result1 bool
	}
	existsReturnsOnCall map[int]struct {
		result1 bool
	}
	GetStub        func(v1.Object) (client.Object, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 v1.Object
	}
	getReturns struct {
		result1 client.Object
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 client.Object
		result2 error
	}
	GetNameStub        func(v1.Object) string
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
		arg1 v1.Object
	}
	getNameReturns struct {
		result1 string
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
	}
	GetSchemeStub        func() *runtime.Scheme
	getSchemeMutex       sync.RWMutex
	getSchemeArgsForCall []struct {
	}
	getSchemeReturns struct {
		result1 *runtime.Scheme
	}
	getSchemeReturnsOnCall map[int]struct {
		result1 *runtime.Scheme
	}
	ReconcileStub        func(v1.Object, bool) error
	reconcileMutex       sync.RWMutex
	reconcileArgsForCall []struct {
		arg1 v1.Object
		arg2 bool
	}
	reconcileReturns struct {
		result1 error
	}
	reconcileReturnsOnCall map[int]struct {
		result1 error
	}
	RestoreStateStub        func(v1.Object) error
	restoreStateMutex       sync.RWMutex
	restoreStateArgsForCall []struct {
		arg1 v1.Object
	}
	restoreStateReturns struct {
		result1 error
	}
	restoreStateReturnsOnCall map[int]struct {
		result1 error
	}
	SetCustomNameStub        func(string)
	setCustomNameMutex       sync.RWMutex
	setCustomNameArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeploymentManager) CheckForSecretChange(arg1 v1.Object, arg2 string, arg3 func(string, *v1a.Deployment) bool) error {
	fake.checkForSecretChangeMutex.Lock()
	ret, specificReturn := fake.checkForSecretChangeReturnsOnCall[len(fake.checkForSecretChangeArgsForCall)]
	fake.checkForSecretChangeArgsForCall = append(fake.checkForSecretChangeArgsForCall, struct {
		arg1 v1.Object
		arg2 string
		arg3 func(string, *v1a.Deployment) bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("CheckForSecretChange", []interface{}{arg1, arg2, arg3})
	fake.checkForSecretChangeMutex.Unlock()
	if fake.CheckForSecretChangeStub != nil {
		return fake.CheckForSecretChangeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkForSecretChangeReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) CheckForSecretChangeCallCount() int {
	fake.checkForSecretChangeMutex.RLock()
	defer fake.checkForSecretChangeMutex.RUnlock()
	return len(fake.checkForSecretChangeArgsForCall)
}

func (fake *DeploymentManager) CheckForSecretChangeCalls(stub func(v1.Object, string, func(string, *v1a.Deployment) bool) error) {
	fake.checkForSecretChangeMutex.Lock()
	defer fake.checkForSecretChangeMutex.Unlock()
	fake.CheckForSecretChangeStub = stub
}

func (fake *DeploymentManager) CheckForSecretChangeArgsForCall(i int) (v1.Object, string, func(string, *v1a.Deployment) bool) {
	fake.checkForSecretChangeMutex.RLock()
	defer fake.checkForSecretChangeMutex.RUnlock()
	argsForCall := fake.checkForSecretChangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *DeploymentManager) CheckForSecretChangeReturns(result1 error) {
	fake.checkForSecretChangeMutex.Lock()
	defer fake.checkForSecretChangeMutex.Unlock()
	fake.CheckForSecretChangeStub = nil
	fake.checkForSecretChangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) CheckForSecretChangeReturnsOnCall(i int, result1 error) {
	fake.checkForSecretChangeMutex.Lock()
	defer fake.checkForSecretChangeMutex.Unlock()
	fake.CheckForSecretChangeStub = nil
	if fake.checkForSecretChangeReturnsOnCall == nil {
		fake.checkForSecretChangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkForSecretChangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) CheckState(arg1 v1.Object) error {
	fake.checkStateMutex.Lock()
	ret, specificReturn := fake.checkStateReturnsOnCall[len(fake.checkStateArgsForCall)]
	fake.checkStateArgsForCall = append(fake.checkStateArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("CheckState", []interface{}{arg1})
	fake.checkStateMutex.Unlock()
	if fake.CheckStateStub != nil {
		return fake.CheckStateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkStateReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) CheckStateCallCount() int {
	fake.checkStateMutex.RLock()
	defer fake.checkStateMutex.RUnlock()
	return len(fake.checkStateArgsForCall)
}

func (fake *DeploymentManager) CheckStateCalls(stub func(v1.Object) error) {
	fake.checkStateMutex.Lock()
	defer fake.checkStateMutex.Unlock()
	fake.CheckStateStub = stub
}

func (fake *DeploymentManager) CheckStateArgsForCall(i int) v1.Object {
	fake.checkStateMutex.RLock()
	defer fake.checkStateMutex.RUnlock()
	argsForCall := fake.checkStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) CheckStateReturns(result1 error) {
	fake.checkStateMutex.Lock()
	defer fake.checkStateMutex.Unlock()
	fake.CheckStateStub = nil
	fake.checkStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) CheckStateReturnsOnCall(i int, result1 error) {
	fake.checkStateMutex.Lock()
	defer fake.checkStateMutex.Unlock()
	fake.CheckStateStub = nil
	if fake.checkStateReturnsOnCall == nil {
		fake.checkStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) Delete(arg1 v1.Object) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *DeploymentManager) DeleteCalls(stub func(v1.Object) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *DeploymentManager) DeleteArgsForCall(i int) v1.Object {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) DeploymentStatus(arg1 v1.Object) (v1a.DeploymentStatus, error) {
	fake.deploymentStatusMutex.Lock()
	ret, specificReturn := fake.deploymentStatusReturnsOnCall[len(fake.deploymentStatusArgsForCall)]
	fake.deploymentStatusArgsForCall = append(fake.deploymentStatusArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("DeploymentStatus", []interface{}{arg1})
	fake.deploymentStatusMutex.Unlock()
	if fake.DeploymentStatusStub != nil {
		return fake.DeploymentStatusStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deploymentStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeploymentManager) DeploymentStatusCallCount() int {
	fake.deploymentStatusMutex.RLock()
	defer fake.deploymentStatusMutex.RUnlock()
	return len(fake.deploymentStatusArgsForCall)
}

func (fake *DeploymentManager) DeploymentStatusCalls(stub func(v1.Object) (v1a.DeploymentStatus, error)) {
	fake.deploymentStatusMutex.Lock()
	defer fake.deploymentStatusMutex.Unlock()
	fake.DeploymentStatusStub = stub
}

func (fake *DeploymentManager) DeploymentStatusArgsForCall(i int) v1.Object {
	fake.deploymentStatusMutex.RLock()
	defer fake.deploymentStatusMutex.RUnlock()
	argsForCall := fake.deploymentStatusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) DeploymentStatusReturns(result1 v1a.DeploymentStatus, result2 error) {
	fake.deploymentStatusMutex.Lock()
	defer fake.deploymentStatusMutex.Unlock()
	fake.DeploymentStatusStub = nil
	fake.deploymentStatusReturns = struct {
		result1 v1a.DeploymentStatus
		result2 error
	}{result1, result2}
}

func (fake *DeploymentManager) DeploymentStatusReturnsOnCall(i int, result1 v1a.DeploymentStatus, result2 error) {
	fake.deploymentStatusMutex.Lock()
	defer fake.deploymentStatusMutex.Unlock()
	fake.DeploymentStatusStub = nil
	if fake.deploymentStatusReturnsOnCall == nil {
		fake.deploymentStatusReturnsOnCall = make(map[int]struct {
			result1 v1a.DeploymentStatus
			result2 error
		})
	}
	fake.deploymentStatusReturnsOnCall[i] = struct {
		result1 v1a.DeploymentStatus
		result2 error
	}{result1, result2}
}

func (fake *DeploymentManager) Exists(arg1 v1.Object) bool {
	fake.existsMutex.Lock()
	ret, specificReturn := fake.existsReturnsOnCall[len(fake.existsArgsForCall)]
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("Exists", []interface{}{arg1})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.existsReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *DeploymentManager) ExistsCalls(stub func(v1.Object) bool) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = stub
}

func (fake *DeploymentManager) ExistsArgsForCall(i int) v1.Object {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	argsForCall := fake.existsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) ExistsReturns(result1 bool) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *DeploymentManager) ExistsReturnsOnCall(i int, result1 bool) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = nil
	if fake.existsReturnsOnCall == nil {
		fake.existsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.existsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *DeploymentManager) Get(arg1 v1.Object) (client.Object, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeploymentManager) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *DeploymentManager) GetCalls(stub func(v1.Object) (client.Object, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *DeploymentManager) GetArgsForCall(i int) v1.Object {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) GetReturns(result1 client.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 client.Object
		result2 error
	}{result1, result2}
}

func (fake *DeploymentManager) GetReturnsOnCall(i int, result1 client.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 client.Object
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 client.Object
		result2 error
	}{result1, result2}
}

func (fake *DeploymentManager) GetName(arg1 v1.Object) string {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("GetName", []interface{}{arg1})
	fake.getNameMutex.Unlock()
	if fake.GetNameStub != nil {
		return fake.GetNameStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getNameReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *DeploymentManager) GetNameCalls(stub func(v1.Object) string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *DeploymentManager) GetNameArgsForCall(i int) v1.Object {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	argsForCall := fake.getNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) GetNameReturns(result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *DeploymentManager) GetNameReturnsOnCall(i int, result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *DeploymentManager) GetScheme() *runtime.Scheme {
	fake.getSchemeMutex.Lock()
	ret, specificReturn := fake.getSchemeReturnsOnCall[len(fake.getSchemeArgsForCall)]
	fake.getSchemeArgsForCall = append(fake.getSchemeArgsForCall, struct {
	}{})
	fake.recordInvocation("GetScheme", []interface{}{})
	fake.getSchemeMutex.Unlock()
	if fake.GetSchemeStub != nil {
		return fake.GetSchemeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getSchemeReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) GetSchemeCallCount() int {
	fake.getSchemeMutex.RLock()
	defer fake.getSchemeMutex.RUnlock()
	return len(fake.getSchemeArgsForCall)
}

func (fake *DeploymentManager) GetSchemeCalls(stub func() *runtime.Scheme) {
	fake.getSchemeMutex.Lock()
	defer fake.getSchemeMutex.Unlock()
	fake.GetSchemeStub = stub
}

func (fake *DeploymentManager) GetSchemeReturns(result1 *runtime.Scheme) {
	fake.getSchemeMutex.Lock()
	defer fake.getSchemeMutex.Unlock()
	fake.GetSchemeStub = nil
	fake.getSchemeReturns = struct {
		result1 *runtime.Scheme
	}{result1}
}

func (fake *DeploymentManager) GetSchemeReturnsOnCall(i int, result1 *runtime.Scheme) {
	fake.getSchemeMutex.Lock()
	defer fake.getSchemeMutex.Unlock()
	fake.GetSchemeStub = nil
	if fake.getSchemeReturnsOnCall == nil {
		fake.getSchemeReturnsOnCall = make(map[int]struct {
			result1 *runtime.Scheme
		})
	}
	fake.getSchemeReturnsOnCall[i] = struct {
		result1 *runtime.Scheme
	}{result1}
}

func (fake *DeploymentManager) Reconcile(arg1 v1.Object, arg2 bool) error {
	fake.reconcileMutex.Lock()
	ret, specificReturn := fake.reconcileReturnsOnCall[len(fake.reconcileArgsForCall)]
	fake.reconcileArgsForCall = append(fake.reconcileArgsForCall, struct {
		arg1 v1.Object
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("Reconcile", []interface{}{arg1, arg2})
	fake.reconcileMutex.Unlock()
	if fake.ReconcileStub != nil {
		return fake.ReconcileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.reconcileReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) ReconcileCallCount() int {
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	return len(fake.reconcileArgsForCall)
}

func (fake *DeploymentManager) ReconcileCalls(stub func(v1.Object, bool) error) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = stub
}

func (fake *DeploymentManager) ReconcileArgsForCall(i int) (v1.Object, bool) {
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	argsForCall := fake.reconcileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *DeploymentManager) ReconcileReturns(result1 error) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = nil
	fake.reconcileReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) ReconcileReturnsOnCall(i int, result1 error) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = nil
	if fake.reconcileReturnsOnCall == nil {
		fake.reconcileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reconcileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) RestoreState(arg1 v1.Object) error {
	fake.restoreStateMutex.Lock()
	ret, specificReturn := fake.restoreStateReturnsOnCall[len(fake.restoreStateArgsForCall)]
	fake.restoreStateArgsForCall = append(fake.restoreStateArgsForCall, struct {
		arg1 v1.Object
	}{arg1})
	fake.recordInvocation("RestoreState", []interface{}{arg1})
	fake.restoreStateMutex.Unlock()
	if fake.RestoreStateStub != nil {
		return fake.RestoreStateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.restoreStateReturns
	return fakeReturns.result1
}

func (fake *DeploymentManager) RestoreStateCallCount() int {
	fake.restoreStateMutex.RLock()
	defer fake.restoreStateMutex.RUnlock()
	return len(fake.restoreStateArgsForCall)
}

func (fake *DeploymentManager) RestoreStateCalls(stub func(v1.Object) error) {
	fake.restoreStateMutex.Lock()
	defer fake.restoreStateMutex.Unlock()
	fake.RestoreStateStub = stub
}

func (fake *DeploymentManager) RestoreStateArgsForCall(i int) v1.Object {
	fake.restoreStateMutex.RLock()
	defer fake.restoreStateMutex.RUnlock()
	argsForCall := fake.restoreStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) RestoreStateReturns(result1 error) {
	fake.restoreStateMutex.Lock()
	defer fake.restoreStateMutex.Unlock()
	fake.RestoreStateStub = nil
	fake.restoreStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) RestoreStateReturnsOnCall(i int, result1 error) {
	fake.restoreStateMutex.Lock()
	defer fake.restoreStateMutex.Unlock()
	fake.RestoreStateStub = nil
	if fake.restoreStateReturnsOnCall == nil {
		fake.restoreStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeploymentManager) SetCustomName(arg1 string) {
	fake.setCustomNameMutex.Lock()
	fake.setCustomNameArgsForCall = append(fake.setCustomNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetCustomName", []interface{}{arg1})
	fake.setCustomNameMutex.Unlock()
	if fake.SetCustomNameStub != nil {
		fake.SetCustomNameStub(arg1)
	}
}

func (fake *DeploymentManager) SetCustomNameCallCount() int {
	fake.setCustomNameMutex.RLock()
	defer fake.setCustomNameMutex.RUnlock()
	return len(fake.setCustomNameArgsForCall)
}

func (fake *DeploymentManager) SetCustomNameCalls(stub func(string)) {
	fake.setCustomNameMutex.Lock()
	defer fake.setCustomNameMutex.Unlock()
	fake.SetCustomNameStub = stub
}

func (fake *DeploymentManager) SetCustomNameArgsForCall(i int) string {
	fake.setCustomNameMutex.RLock()
	defer fake.setCustomNameMutex.RUnlock()
	argsForCall := fake.setCustomNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeploymentManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkForSecretChangeMutex.RLock()
	defer fake.checkForSecretChangeMutex.RUnlock()
	fake.checkStateMutex.RLock()
	defer fake.checkStateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deploymentStatusMutex.RLock()
	defer fake.deploymentStatusMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.getSchemeMutex.RLock()
	defer fake.getSchemeMutex.RUnlock()
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	fake.restoreStateMutex.RLock()
	defer fake.restoreStateMutex.RUnlock()
	fake.setCustomNameMutex.RLock()
	defer fake.setCustomNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeploymentManager) recordInvocation(key string, args []interface{}) {
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

var _ basepeer.DeploymentManager = new(DeploymentManager)