// Automatically generated by MockGen. DO NOT EDIT!
// Source: vm.go

package services

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/pkg/api/v1"
	v10 "kubevirt.io/kubevirt/pkg/api/v1"
)

// Mock of VMService interface
type MockVMService struct {
	ctrl     *gomock.Controller
	recorder *_MockVMServiceRecorder
}

// Recorder for MockVMService (not exported)
type _MockVMServiceRecorder struct {
	mock *MockVMService
}

func NewMockVMService(ctrl *gomock.Controller) *MockVMService {
	mock := &MockVMService{ctrl: ctrl}
	mock.recorder = &_MockVMServiceRecorder{mock}
	return mock
}

func (_m *MockVMService) EXPECT() *_MockVMServiceRecorder {
	return _m.recorder
}

func (_m *MockVMService) StartVMPod(_param0 *v10.VM) error {
	ret := _m.ctrl.Call(_m, "StartVMPod", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) StartVMPod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartVMPod", arg0)
}

func (_m *MockVMService) DeleteVMPod(_param0 *v10.VM) error {
	ret := _m.ctrl.Call(_m, "DeleteVMPod", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) DeleteVMPod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVMPod", arg0)
}

func (_m *MockVMService) GetRunningVMPods(_param0 *v10.VM) (*v1.PodList, error) {
	ret := _m.ctrl.Call(_m, "GetRunningVMPods", _param0)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMServiceRecorder) GetRunningVMPods(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRunningVMPods", arg0)
}

func (_m *MockVMService) DeleteMigrationTargetPods(_param0 *v10.Migration) error {
	ret := _m.ctrl.Call(_m, "DeleteMigrationTargetPods", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) DeleteMigrationTargetPods(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteMigrationTargetPods", arg0)
}

func (_m *MockVMService) GetRunningMigrationPods(_param0 *v10.Migration) (*v1.PodList, error) {
	ret := _m.ctrl.Call(_m, "GetRunningMigrationPods", _param0)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMServiceRecorder) GetRunningMigrationPods(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRunningMigrationPods", arg0)
}

func (_m *MockVMService) CreateMigrationTargetPod(migration *v10.Migration, vm *v10.VM) error {
	ret := _m.ctrl.Call(_m, "CreateMigrationTargetPod", migration, vm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) CreateMigrationTargetPod(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateMigrationTargetPod", arg0, arg1)
}

func (_m *MockVMService) UpdateMigration(migration *v10.Migration) error {
	ret := _m.ctrl.Call(_m, "UpdateMigration", migration)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) UpdateMigration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMigration", arg0)
}

func (_m *MockVMService) FetchVM(vmName string) (*v10.VM, bool, error) {
	ret := _m.ctrl.Call(_m, "FetchVM", vmName)
	ret0, _ := ret[0].(*v10.VM)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVMServiceRecorder) FetchVM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchVM", arg0)
}

func (_m *MockVMService) FetchMigration(migrationName string) (*v10.Migration, bool, error) {
	ret := _m.ctrl.Call(_m, "FetchMigration", migrationName)
	ret0, _ := ret[0].(*v10.Migration)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVMServiceRecorder) FetchMigration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchMigration", arg0)
}

func (_m *MockVMService) StartMigration(migration *v10.Migration, vm *v10.VM, sourceNode *v1.Node, targetNode *v1.Node, targetPod *v1.Pod) error {
	ret := _m.ctrl.Call(_m, "StartMigration", migration, vm, sourceNode, targetNode, targetPod)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMServiceRecorder) StartMigration(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartMigration", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockVMService) GetMigrationJob(migration *v10.Migration) (*v1.Pod, bool, error) {
	ret := _m.ctrl.Call(_m, "GetMigrationJob", migration)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVMServiceRecorder) GetMigrationJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMigrationJob", arg0)
}

func (_m *MockVMService) PutVm(vm *v10.VM) (*v10.VM, error) {
	ret := _m.ctrl.Call(_m, "PutVm", vm)
	ret0, _ := ret[0].(*v10.VM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMServiceRecorder) PutVm(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutVm", arg0)
}