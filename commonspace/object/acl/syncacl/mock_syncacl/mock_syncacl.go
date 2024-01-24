// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/object/acl/syncacl (interfaces: SyncAcl,SyncClient,RequestFactory,AclSyncProtocol)
//
// Generated by this command:
//
//	mockgen -destination mock_syncacl/mock_syncacl.go github.com/anyproto/any-sync/commonspace/object/acl/syncacl SyncAcl,SyncClient,RequestFactory,AclSyncProtocol
//

// Package mock_syncacl is a generated GoMock package.
package mock_syncacl

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	list "github.com/anyproto/any-sync/commonspace/object/acl/list"
	headupdater "github.com/anyproto/any-sync/commonspace/object/acl/syncacl/headupdater"
	spacesyncproto "github.com/anyproto/any-sync/commonspace/spacesyncproto"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	crypto "github.com/anyproto/any-sync/util/crypto"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncAcl is a mock of SyncAcl interface.
type MockSyncAcl struct {
	ctrl     *gomock.Controller
	recorder *MockSyncAclMockRecorder
}

// MockSyncAclMockRecorder is the mock recorder for MockSyncAcl.
type MockSyncAclMockRecorder struct {
	mock *MockSyncAcl
}

// NewMockSyncAcl creates a new mock instance.
func NewMockSyncAcl(ctrl *gomock.Controller) *MockSyncAcl {
	mock := &MockSyncAcl{ctrl: ctrl}
	mock.recorder = &MockSyncAclMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncAcl) EXPECT() *MockSyncAclMockRecorder {
	return m.recorder
}

// AclState mocks base method.
func (m *MockSyncAcl) AclState() *list.AclState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclState")
	ret0, _ := ret[0].(*list.AclState)
	return ret0
}

// AclState indicates an expected call of AclState.
func (mr *MockSyncAclMockRecorder) AclState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclState", reflect.TypeOf((*MockSyncAcl)(nil).AclState))
}

// AddRawRecord mocks base method.
func (m *MockSyncAcl) AddRawRecord(arg0 *consensusproto.RawRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawRecord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawRecord indicates an expected call of AddRawRecord.
func (mr *MockSyncAclMockRecorder) AddRawRecord(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawRecord", reflect.TypeOf((*MockSyncAcl)(nil).AddRawRecord), arg0)
}

// AddRawRecords mocks base method.
func (m *MockSyncAcl) AddRawRecords(arg0 []*consensusproto.RawRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawRecords", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawRecords indicates an expected call of AddRawRecords.
func (mr *MockSyncAclMockRecorder) AddRawRecords(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawRecords", reflect.TypeOf((*MockSyncAcl)(nil).AddRawRecords), arg0)
}

// Close mocks base method.
func (m *MockSyncAcl) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncAclMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncAcl)(nil).Close), arg0)
}

// Get mocks base method.
func (m *MockSyncAcl) Get(arg0 string) (*list.AclRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*list.AclRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSyncAclMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSyncAcl)(nil).Get), arg0)
}

// GetIndex mocks base method.
func (m *MockSyncAcl) GetIndex(arg0 int) (*list.AclRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndex", arg0)
	ret0, _ := ret[0].(*list.AclRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndex indicates an expected call of GetIndex.
func (mr *MockSyncAclMockRecorder) GetIndex(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndex", reflect.TypeOf((*MockSyncAcl)(nil).GetIndex), arg0)
}

// HandleMessage mocks base method.
func (m *MockSyncAcl) HandleMessage(arg0 context.Context, arg1 string, arg2 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockSyncAclMockRecorder) HandleMessage(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockSyncAcl)(nil).HandleMessage), arg0, arg1, arg2)
}

// HandleRequest mocks base method.
func (m *MockSyncAcl) HandleRequest(arg0 context.Context, arg1 string, arg2 *spacesyncproto.ObjectSyncMessage) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleRequest indicates an expected call of HandleRequest.
func (mr *MockSyncAclMockRecorder) HandleRequest(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRequest", reflect.TypeOf((*MockSyncAcl)(nil).HandleRequest), arg0, arg1, arg2)
}

// HasHead mocks base method.
func (m *MockSyncAcl) HasHead(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasHead", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasHead indicates an expected call of HasHead.
func (mr *MockSyncAclMockRecorder) HasHead(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasHead", reflect.TypeOf((*MockSyncAcl)(nil).HasHead), arg0)
}

// Head mocks base method.
func (m *MockSyncAcl) Head() *list.AclRecord {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(*list.AclRecord)
	return ret0
}

// Head indicates an expected call of Head.
func (mr *MockSyncAclMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockSyncAcl)(nil).Head))
}

// Id mocks base method.
func (m *MockSyncAcl) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSyncAclMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSyncAcl)(nil).Id))
}

// Init mocks base method.
func (m *MockSyncAcl) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSyncAclMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSyncAcl)(nil).Init), arg0)
}

// IsAfter mocks base method.
func (m *MockSyncAcl) IsAfter(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAfter", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAfter indicates an expected call of IsAfter.
func (mr *MockSyncAclMockRecorder) IsAfter(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAfter", reflect.TypeOf((*MockSyncAcl)(nil).IsAfter), arg0, arg1)
}

// Iterate mocks base method.
func (m *MockSyncAcl) Iterate(arg0 func(*list.AclRecord) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Iterate", arg0)
}

// Iterate indicates an expected call of Iterate.
func (mr *MockSyncAclMockRecorder) Iterate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockSyncAcl)(nil).Iterate), arg0)
}

// IterateFrom mocks base method.
func (m *MockSyncAcl) IterateFrom(arg0 string, arg1 func(*list.AclRecord) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateFrom", arg0, arg1)
}

// IterateFrom indicates an expected call of IterateFrom.
func (mr *MockSyncAclMockRecorder) IterateFrom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateFrom", reflect.TypeOf((*MockSyncAcl)(nil).IterateFrom), arg0, arg1)
}

// KeyStorage mocks base method.
func (m *MockSyncAcl) KeyStorage() crypto.KeyStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyStorage")
	ret0, _ := ret[0].(crypto.KeyStorage)
	return ret0
}

// KeyStorage indicates an expected call of KeyStorage.
func (mr *MockSyncAclMockRecorder) KeyStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyStorage", reflect.TypeOf((*MockSyncAcl)(nil).KeyStorage))
}

// Lock mocks base method.
func (m *MockSyncAcl) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockSyncAclMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockSyncAcl)(nil).Lock))
}

// Name mocks base method.
func (m *MockSyncAcl) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSyncAclMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSyncAcl)(nil).Name))
}

// RLock mocks base method.
func (m *MockSyncAcl) RLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RLock")
}

// RLock indicates an expected call of RLock.
func (mr *MockSyncAclMockRecorder) RLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RLock", reflect.TypeOf((*MockSyncAcl)(nil).RLock))
}

// RUnlock mocks base method.
func (m *MockSyncAcl) RUnlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RUnlock")
}

// RUnlock indicates an expected call of RUnlock.
func (mr *MockSyncAclMockRecorder) RUnlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RUnlock", reflect.TypeOf((*MockSyncAcl)(nil).RUnlock))
}

// RecordBuilder mocks base method.
func (m *MockSyncAcl) RecordBuilder() list.AclRecordBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordBuilder")
	ret0, _ := ret[0].(list.AclRecordBuilder)
	return ret0
}

// RecordBuilder indicates an expected call of RecordBuilder.
func (mr *MockSyncAclMockRecorder) RecordBuilder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordBuilder", reflect.TypeOf((*MockSyncAcl)(nil).RecordBuilder))
}

// Records mocks base method.
func (m *MockSyncAcl) Records() []*list.AclRecord {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Records")
	ret0, _ := ret[0].([]*list.AclRecord)
	return ret0
}

// Records indicates an expected call of Records.
func (mr *MockSyncAclMockRecorder) Records() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Records", reflect.TypeOf((*MockSyncAcl)(nil).Records))
}

// RecordsAfter mocks base method.
func (m *MockSyncAcl) RecordsAfter(arg0 context.Context, arg1 string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordsAfter", arg0, arg1)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordsAfter indicates an expected call of RecordsAfter.
func (mr *MockSyncAclMockRecorder) RecordsAfter(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordsAfter", reflect.TypeOf((*MockSyncAcl)(nil).RecordsAfter), arg0, arg1)
}

// RecordsBefore mocks base method.
func (m *MockSyncAcl) RecordsBefore(arg0 context.Context, arg1 string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordsBefore", arg0, arg1)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordsBefore indicates an expected call of RecordsBefore.
func (mr *MockSyncAclMockRecorder) RecordsBefore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordsBefore", reflect.TypeOf((*MockSyncAcl)(nil).RecordsBefore), arg0, arg1)
}

// Root mocks base method.
func (m *MockSyncAcl) Root() *consensusproto.RawRecordWithId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	return ret0
}

// Root indicates an expected call of Root.
func (mr *MockSyncAclMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockSyncAcl)(nil).Root))
}

// Run mocks base method.
func (m *MockSyncAcl) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockSyncAclMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSyncAcl)(nil).Run), arg0)
}

// SetAclUpdater mocks base method.
func (m *MockSyncAcl) SetAclUpdater(arg0 headupdater.AclUpdater) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAclUpdater", arg0)
}

// SetAclUpdater indicates an expected call of SetAclUpdater.
func (mr *MockSyncAclMockRecorder) SetAclUpdater(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAclUpdater", reflect.TypeOf((*MockSyncAcl)(nil).SetAclUpdater), arg0)
}

// SetHeadUpdater mocks base method.
func (m *MockSyncAcl) SetHeadUpdater(arg0 headupdater.HeadUpdater) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeadUpdater", arg0)
}

// SetHeadUpdater indicates an expected call of SetHeadUpdater.
func (mr *MockSyncAclMockRecorder) SetHeadUpdater(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeadUpdater", reflect.TypeOf((*MockSyncAcl)(nil).SetHeadUpdater), arg0)
}

// SyncWithPeer mocks base method.
func (m *MockSyncAcl) SyncWithPeer(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncWithPeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncWithPeer indicates an expected call of SyncWithPeer.
func (mr *MockSyncAclMockRecorder) SyncWithPeer(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncWithPeer", reflect.TypeOf((*MockSyncAcl)(nil).SyncWithPeer), arg0, arg1)
}

// Unlock mocks base method.
func (m *MockSyncAcl) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockSyncAclMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockSyncAcl)(nil).Unlock))
}

// ValidateRawRecord mocks base method.
func (m *MockSyncAcl) ValidateRawRecord(arg0 *consensusproto.RawRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRawRecord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRawRecord indicates an expected call of ValidateRawRecord.
func (mr *MockSyncAclMockRecorder) ValidateRawRecord(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRawRecord", reflect.TypeOf((*MockSyncAcl)(nil).ValidateRawRecord), arg0)
}

// MockSyncClient is a mock of SyncClient interface.
type MockSyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockSyncClientMockRecorder
}

// MockSyncClientMockRecorder is the mock recorder for MockSyncClient.
type MockSyncClientMockRecorder struct {
	mock *MockSyncClient
}

// NewMockSyncClient creates a new mock instance.
func NewMockSyncClient(ctrl *gomock.Controller) *MockSyncClient {
	mock := &MockSyncClient{ctrl: ctrl}
	mock.recorder = &MockSyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncClient) EXPECT() *MockSyncClientMockRecorder {
	return m.recorder
}

// Broadcast mocks base method.
func (m *MockSyncClient) Broadcast(arg0 *consensusproto.LogSyncMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Broadcast", arg0)
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockSyncClientMockRecorder) Broadcast(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockSyncClient)(nil).Broadcast), arg0)
}

// CreateFullSyncRequest mocks base method.
func (m *MockSyncClient) CreateFullSyncRequest(arg0 list.AclList, arg1 string) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncRequest", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncRequest indicates an expected call of CreateFullSyncRequest.
func (mr *MockSyncClientMockRecorder) CreateFullSyncRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncRequest", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncRequest), arg0, arg1)
}

// CreateFullSyncResponse mocks base method.
func (m *MockSyncClient) CreateFullSyncResponse(arg0 list.AclList, arg1 string) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncResponse", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncResponse indicates an expected call of CreateFullSyncResponse.
func (mr *MockSyncClientMockRecorder) CreateFullSyncResponse(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncResponse", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncResponse), arg0, arg1)
}

// CreateHeadUpdate mocks base method.
func (m *MockSyncClient) CreateHeadUpdate(arg0 list.AclList, arg1 []*consensusproto.RawRecordWithId) *consensusproto.LogSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHeadUpdate", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	return ret0
}

// CreateHeadUpdate indicates an expected call of CreateHeadUpdate.
func (mr *MockSyncClientMockRecorder) CreateHeadUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHeadUpdate", reflect.TypeOf((*MockSyncClient)(nil).CreateHeadUpdate), arg0, arg1)
}

// QueueRequest mocks base method.
func (m *MockSyncClient) QueueRequest(arg0 string, arg1 *consensusproto.LogSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueRequest indicates an expected call of QueueRequest.
func (mr *MockSyncClientMockRecorder) QueueRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueRequest", reflect.TypeOf((*MockSyncClient)(nil).QueueRequest), arg0, arg1)
}

// SendRequest mocks base method.
func (m *MockSyncClient) SendRequest(arg0 context.Context, arg1 string, arg2 *consensusproto.LogSyncMessage) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRequest indicates an expected call of SendRequest.
func (mr *MockSyncClientMockRecorder) SendRequest(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockSyncClient)(nil).SendRequest), arg0, arg1, arg2)
}

// SendUpdate mocks base method.
func (m *MockSyncClient) SendUpdate(arg0 string, arg1 *consensusproto.LogSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendUpdate indicates an expected call of SendUpdate.
func (mr *MockSyncClientMockRecorder) SendUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUpdate", reflect.TypeOf((*MockSyncClient)(nil).SendUpdate), arg0, arg1)
}

// MockRequestFactory is a mock of RequestFactory interface.
type MockRequestFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRequestFactoryMockRecorder
}

// MockRequestFactoryMockRecorder is the mock recorder for MockRequestFactory.
type MockRequestFactoryMockRecorder struct {
	mock *MockRequestFactory
}

// NewMockRequestFactory creates a new mock instance.
func NewMockRequestFactory(ctrl *gomock.Controller) *MockRequestFactory {
	mock := &MockRequestFactory{ctrl: ctrl}
	mock.recorder = &MockRequestFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestFactory) EXPECT() *MockRequestFactoryMockRecorder {
	return m.recorder
}

// CreateFullSyncRequest mocks base method.
func (m *MockRequestFactory) CreateFullSyncRequest(arg0 list.AclList, arg1 string) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncRequest", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncRequest indicates an expected call of CreateFullSyncRequest.
func (mr *MockRequestFactoryMockRecorder) CreateFullSyncRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncRequest", reflect.TypeOf((*MockRequestFactory)(nil).CreateFullSyncRequest), arg0, arg1)
}

// CreateFullSyncResponse mocks base method.
func (m *MockRequestFactory) CreateFullSyncResponse(arg0 list.AclList, arg1 string) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncResponse", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncResponse indicates an expected call of CreateFullSyncResponse.
func (mr *MockRequestFactoryMockRecorder) CreateFullSyncResponse(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncResponse", reflect.TypeOf((*MockRequestFactory)(nil).CreateFullSyncResponse), arg0, arg1)
}

// CreateHeadUpdate mocks base method.
func (m *MockRequestFactory) CreateHeadUpdate(arg0 list.AclList, arg1 []*consensusproto.RawRecordWithId) *consensusproto.LogSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHeadUpdate", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	return ret0
}

// CreateHeadUpdate indicates an expected call of CreateHeadUpdate.
func (mr *MockRequestFactoryMockRecorder) CreateHeadUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHeadUpdate", reflect.TypeOf((*MockRequestFactory)(nil).CreateHeadUpdate), arg0, arg1)
}

// MockAclSyncProtocol is a mock of AclSyncProtocol interface.
type MockAclSyncProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockAclSyncProtocolMockRecorder
}

// MockAclSyncProtocolMockRecorder is the mock recorder for MockAclSyncProtocol.
type MockAclSyncProtocolMockRecorder struct {
	mock *MockAclSyncProtocol
}

// NewMockAclSyncProtocol creates a new mock instance.
func NewMockAclSyncProtocol(ctrl *gomock.Controller) *MockAclSyncProtocol {
	mock := &MockAclSyncProtocol{ctrl: ctrl}
	mock.recorder = &MockAclSyncProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAclSyncProtocol) EXPECT() *MockAclSyncProtocolMockRecorder {
	return m.recorder
}

// FullSyncRequest mocks base method.
func (m *MockAclSyncProtocol) FullSyncRequest(arg0 context.Context, arg1 string, arg2 *consensusproto.LogFullSyncRequest) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullSyncRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FullSyncRequest indicates an expected call of FullSyncRequest.
func (mr *MockAclSyncProtocolMockRecorder) FullSyncRequest(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullSyncRequest", reflect.TypeOf((*MockAclSyncProtocol)(nil).FullSyncRequest), arg0, arg1, arg2)
}

// FullSyncResponse mocks base method.
func (m *MockAclSyncProtocol) FullSyncResponse(arg0 context.Context, arg1 string, arg2 *consensusproto.LogFullSyncResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullSyncResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FullSyncResponse indicates an expected call of FullSyncResponse.
func (mr *MockAclSyncProtocolMockRecorder) FullSyncResponse(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullSyncResponse", reflect.TypeOf((*MockAclSyncProtocol)(nil).FullSyncResponse), arg0, arg1, arg2)
}

// HeadUpdate mocks base method.
func (m *MockAclSyncProtocol) HeadUpdate(arg0 context.Context, arg1 string, arg2 *consensusproto.LogHeadUpdate) (*consensusproto.LogSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*consensusproto.LogSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadUpdate indicates an expected call of HeadUpdate.
func (mr *MockAclSyncProtocolMockRecorder) HeadUpdate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadUpdate", reflect.TypeOf((*MockAclSyncProtocol)(nil).HeadUpdate), arg0, arg1, arg2)
}
