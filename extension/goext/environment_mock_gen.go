// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudwan/gohan/extension/goext (interfaces: ICore,ILogger,ISchemas,ISync,IDatabase,ITransaction,IHTTP,IAuth,IConfig,IUtil)

// Package goext is a generated GoMock package.
package goext

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockICore is a mock of ICore interface
type MockICore struct {
	ctrl     *gomock.Controller
	recorder *MockICoreMockRecorder
}

// MockICoreMockRecorder is the mock recorder for MockICore
type MockICoreMockRecorder struct {
	mock *MockICore
}

// NewMockICore creates a new mock instance
func NewMockICore(ctrl *gomock.Controller) *MockICore {
	mock := &MockICore{ctrl: ctrl}
	mock.recorder = &MockICoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICore) EXPECT() *MockICoreMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method
func (m *MockICore) HandleEvent(arg0 string, arg1 Context) error {
	ret := m.ctrl.Call(m, "HandleEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockICoreMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockICore)(nil).HandleEvent), arg0, arg1)
}

// RegisterEventHandler mocks base method
func (m *MockICore) RegisterEventHandler(arg0 string, arg1 Handler, arg2 int) {
	m.ctrl.Call(m, "RegisterEventHandler", arg0, arg1, arg2)
}

// RegisterEventHandler indicates an expected call of RegisterEventHandler
func (mr *MockICoreMockRecorder) RegisterEventHandler(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEventHandler", reflect.TypeOf((*MockICore)(nil).RegisterEventHandler), arg0, arg1, arg2)
}

// RegisterSchemaEventHandler mocks base method
func (m *MockICore) RegisterSchemaEventHandler(arg0, arg1 string, arg2 SchemaHandler, arg3 int) {
	m.ctrl.Call(m, "RegisterSchemaEventHandler", arg0, arg1, arg2, arg3)
}

// RegisterSchemaEventHandler indicates an expected call of RegisterSchemaEventHandler
func (mr *MockICoreMockRecorder) RegisterSchemaEventHandler(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSchemaEventHandler", reflect.TypeOf((*MockICore)(nil).RegisterSchemaEventHandler), arg0, arg1, arg2, arg3)
}

// TriggerEvent mocks base method
func (m *MockICore) TriggerEvent(arg0 string, arg1 Context) error {
	ret := m.ctrl.Call(m, "TriggerEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TriggerEvent indicates an expected call of TriggerEvent
func (mr *MockICoreMockRecorder) TriggerEvent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerEvent", reflect.TypeOf((*MockICore)(nil).TriggerEvent), arg0, arg1)
}

// MockILogger is a mock of ILogger interface
type MockILogger struct {
	ctrl     *gomock.Controller
	recorder *MockILoggerMockRecorder
}

// MockILoggerMockRecorder is the mock recorder for MockILogger
type MockILoggerMockRecorder struct {
	mock *MockILogger
}

// NewMockILogger creates a new mock instance
func NewMockILogger(ctrl *gomock.Controller) *MockILogger {
	mock := &MockILogger{ctrl: ctrl}
	mock.recorder = &MockILoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockILogger) EXPECT() *MockILoggerMockRecorder {
	return m.recorder
}

// Critical mocks base method
func (m *MockILogger) Critical(arg0 string) {
	m.ctrl.Call(m, "Critical", arg0)
}

// Critical indicates an expected call of Critical
func (mr *MockILoggerMockRecorder) Critical(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Critical", reflect.TypeOf((*MockILogger)(nil).Critical), arg0)
}

// Criticalf mocks base method
func (m *MockILogger) Criticalf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Criticalf", varargs...)
}

// Criticalf indicates an expected call of Criticalf
func (mr *MockILoggerMockRecorder) Criticalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Criticalf", reflect.TypeOf((*MockILogger)(nil).Criticalf), varargs...)
}

// Debug mocks base method
func (m *MockILogger) Debug(arg0 string) {
	m.ctrl.Call(m, "Debug", arg0)
}

// Debug indicates an expected call of Debug
func (mr *MockILoggerMockRecorder) Debug(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockILogger)(nil).Debug), arg0)
}

// Debugf mocks base method
func (m *MockILogger) Debugf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockILoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockILogger)(nil).Debugf), varargs...)
}

// Error mocks base method
func (m *MockILogger) Error(arg0 string) {
	m.ctrl.Call(m, "Error", arg0)
}

// Error indicates an expected call of Error
func (mr *MockILoggerMockRecorder) Error(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockILogger)(nil).Error), arg0)
}

// Errorf mocks base method
func (m *MockILogger) Errorf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockILoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockILogger)(nil).Errorf), varargs...)
}

// Info mocks base method
func (m *MockILogger) Info(arg0 string) {
	m.ctrl.Call(m, "Info", arg0)
}

// Info indicates an expected call of Info
func (mr *MockILoggerMockRecorder) Info(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockILogger)(nil).Info), arg0)
}

// Infof mocks base method
func (m *MockILogger) Infof(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockILoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockILogger)(nil).Infof), varargs...)
}

// Notice mocks base method
func (m *MockILogger) Notice(arg0 string) {
	m.ctrl.Call(m, "Notice", arg0)
}

// Notice indicates an expected call of Notice
func (mr *MockILoggerMockRecorder) Notice(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notice", reflect.TypeOf((*MockILogger)(nil).Notice), arg0)
}

// Noticef mocks base method
func (m *MockILogger) Noticef(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Noticef", varargs...)
}

// Noticef indicates an expected call of Noticef
func (mr *MockILoggerMockRecorder) Noticef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Noticef", reflect.TypeOf((*MockILogger)(nil).Noticef), varargs...)
}

// Warning mocks base method
func (m *MockILogger) Warning(arg0 string) {
	m.ctrl.Call(m, "Warning", arg0)
}

// Warning indicates an expected call of Warning
func (mr *MockILoggerMockRecorder) Warning(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockILogger)(nil).Warning), arg0)
}

// Warningf mocks base method
func (m *MockILogger) Warningf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockILoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockILogger)(nil).Warningf), varargs...)
}

// MockISchemas is a mock of ISchemas interface
type MockISchemas struct {
	ctrl     *gomock.Controller
	recorder *MockISchemasMockRecorder
}

// MockISchemasMockRecorder is the mock recorder for MockISchemas
type MockISchemasMockRecorder struct {
	mock *MockISchemas
}

// NewMockISchemas creates a new mock instance
func NewMockISchemas(ctrl *gomock.Controller) *MockISchemas {
	mock := &MockISchemas{ctrl: ctrl}
	mock.recorder = &MockISchemasMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISchemas) EXPECT() *MockISchemasMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockISchemas) Find(arg0 string) ISchema {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(ISchema)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockISchemasMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockISchemas)(nil).Find), arg0)
}

// List mocks base method
func (m *MockISchemas) List() []ISchema {
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]ISchema)
	return ret0
}

// List indicates an expected call of List
func (mr *MockISchemasMockRecorder) List() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockISchemas)(nil).List))
}

// Relations mocks base method
func (m *MockISchemas) Relations(arg0 string) []SchemaRelationInfo {
	ret := m.ctrl.Call(m, "Relations", arg0)
	ret0, _ := ret[0].([]SchemaRelationInfo)
	return ret0
}

// Relations indicates an expected call of Relations
func (mr *MockISchemasMockRecorder) Relations(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations", reflect.TypeOf((*MockISchemas)(nil).Relations), arg0)
}

// MockISync is a mock of ISync interface
type MockISync struct {
	ctrl     *gomock.Controller
	recorder *MockISyncMockRecorder
}

// MockISyncMockRecorder is the mock recorder for MockISync
type MockISyncMockRecorder struct {
	mock *MockISync
}

// NewMockISync creates a new mock instance
func NewMockISync(ctrl *gomock.Controller) *MockISync {
	mock := &MockISync{ctrl: ctrl}
	mock.recorder = &MockISyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISync) EXPECT() *MockISyncMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockISync) Delete(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockISyncMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISync)(nil).Delete), arg0, arg1)
}

// Fetch mocks base method
func (m *MockISync) Fetch(arg0 string) (*Node, error) {
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(*Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockISyncMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockISync)(nil).Fetch), arg0)
}

// Watch mocks base method
func (m *MockISync) Watch(arg0 context.Context, arg1 string, arg2 time.Duration, arg3 int64) ([]*Event, error) {
	ret := m.ctrl.Call(m, "Watch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockISyncMockRecorder) Watch(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockISync)(nil).Watch), arg0, arg1, arg2, arg3)
}

// MockIDatabase is a mock of IDatabase interface
type MockIDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase
type MockIDatabaseMockRecorder struct {
	mock *MockIDatabase
}

// NewMockIDatabase creates a new mock instance
func NewMockIDatabase(ctrl *gomock.Controller) *MockIDatabase {
	mock := &MockIDatabase{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDatabase) EXPECT() *MockIDatabaseMockRecorder {
	return m.recorder
}

// Begin mocks base method
func (m *MockIDatabase) Begin() (ITransaction, error) {
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(ITransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin
func (mr *MockIDatabaseMockRecorder) Begin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockIDatabase)(nil).Begin))
}

// BeginTx mocks base method
func (m *MockIDatabase) BeginTx(arg0 Context, arg1 *TxOptions) (ITransaction, error) {
	ret := m.ctrl.Call(m, "BeginTx", arg0, arg1)
	ret0, _ := ret[0].(ITransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx
func (mr *MockIDatabaseMockRecorder) BeginTx(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockIDatabase)(nil).BeginTx), arg0, arg1)
}

// Options mocks base method
func (m *MockIDatabase) Options() DbOptions {
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(DbOptions)
	return ret0
}

// Options indicates an expected call of Options
func (mr *MockIDatabaseMockRecorder) Options() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockIDatabase)(nil).Options))
}

// MockITransaction is a mock of ITransaction interface
type MockITransaction struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionMockRecorder
}

// MockITransactionMockRecorder is the mock recorder for MockITransaction
type MockITransactionMockRecorder struct {
	mock *MockITransaction
}

// NewMockITransaction creates a new mock instance
func NewMockITransaction(ctrl *gomock.Controller) *MockITransaction {
	mock := &MockITransaction{ctrl: ctrl}
	mock.recorder = &MockITransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITransaction) EXPECT() *MockITransactionMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockITransaction) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockITransactionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockITransaction)(nil).Close))
}

// Closed mocks base method
func (m *MockITransaction) Closed() bool {
	ret := m.ctrl.Call(m, "Closed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Closed indicates an expected call of Closed
func (mr *MockITransactionMockRecorder) Closed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Closed", reflect.TypeOf((*MockITransaction)(nil).Closed))
}

// Commit mocks base method
func (m *MockITransaction) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockITransactionMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockITransaction)(nil).Commit))
}

// Create mocks base method
func (m *MockITransaction) Create(arg0 context.Context, arg1 ISchema, arg2 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockITransactionMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITransaction)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockITransaction) Delete(arg0 context.Context, arg1 ISchema, arg2 interface{}) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockITransactionMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITransaction)(nil).Delete), arg0, arg1, arg2)
}

// Exec mocks base method
func (m *MockITransaction) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec
func (mr *MockITransactionMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockITransaction)(nil).Exec), varargs...)
}

// Fetch mocks base method
func (m *MockITransaction) Fetch(arg0 context.Context, arg1 ISchema, arg2 Filter) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockITransactionMockRecorder) Fetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockITransaction)(nil).Fetch), arg0, arg1, arg2)
}

// GetIsolationLevel mocks base method
func (m *MockITransaction) GetIsolationLevel() Type {
	ret := m.ctrl.Call(m, "GetIsolationLevel")
	ret0, _ := ret[0].(Type)
	return ret0
}

// GetIsolationLevel indicates an expected call of GetIsolationLevel
func (mr *MockITransactionMockRecorder) GetIsolationLevel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsolationLevel", reflect.TypeOf((*MockITransaction)(nil).GetIsolationLevel))
}

// List mocks base method
func (m *MockITransaction) List(arg0 context.Context, arg1 ISchema, arg2 Filter, arg3 *ListOptions, arg4 *Paginator) ([]map[string]interface{}, uint64, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockITransactionMockRecorder) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockITransaction)(nil).List), arg0, arg1, arg2, arg3, arg4)
}

// LockFetch mocks base method
func (m *MockITransaction) LockFetch(arg0 context.Context, arg1 ISchema, arg2 Filter, arg3 LockPolicy) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "LockFetch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockFetch indicates an expected call of LockFetch
func (mr *MockITransactionMockRecorder) LockFetch(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockFetch", reflect.TypeOf((*MockITransaction)(nil).LockFetch), arg0, arg1, arg2, arg3)
}

// LockList mocks base method
func (m *MockITransaction) LockList(arg0 context.Context, arg1 ISchema, arg2 Filter, arg3 *ListOptions, arg4 *Paginator, arg5 LockPolicy) ([]map[string]interface{}, uint64, error) {
	ret := m.ctrl.Call(m, "LockList", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LockList indicates an expected call of LockList
func (mr *MockITransactionMockRecorder) LockList(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockList", reflect.TypeOf((*MockITransaction)(nil).LockList), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Query mocks base method
func (m *MockITransaction) Query(arg0 context.Context, arg1 ISchema, arg2 string, arg3 []interface{}) ([]map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockITransactionMockRecorder) Query(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockITransaction)(nil).Query), arg0, arg1, arg2, arg3)
}

// RawTransaction mocks base method
func (m *MockITransaction) RawTransaction() interface{} {
	ret := m.ctrl.Call(m, "RawTransaction")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// RawTransaction indicates an expected call of RawTransaction
func (mr *MockITransactionMockRecorder) RawTransaction() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawTransaction", reflect.TypeOf((*MockITransaction)(nil).RawTransaction))
}

// StateFetch mocks base method
func (m *MockITransaction) StateFetch(arg0 context.Context, arg1 ISchema, arg2 Filter) (ResourceState, error) {
	ret := m.ctrl.Call(m, "StateFetch", arg0, arg1, arg2)
	ret0, _ := ret[0].(ResourceState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateFetch indicates an expected call of StateFetch
func (mr *MockITransactionMockRecorder) StateFetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateFetch", reflect.TypeOf((*MockITransaction)(nil).StateFetch), arg0, arg1, arg2)
}

// StateUpdate mocks base method
func (m *MockITransaction) StateUpdate(arg0 context.Context, arg1 ISchema, arg2 map[string]interface{}, arg3 *ResourceState) error {
	ret := m.ctrl.Call(m, "StateUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StateUpdate indicates an expected call of StateUpdate
func (mr *MockITransactionMockRecorder) StateUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateUpdate", reflect.TypeOf((*MockITransaction)(nil).StateUpdate), arg0, arg1, arg2, arg3)
}

// Update mocks base method
func (m *MockITransaction) Update(arg0 context.Context, arg1 ISchema, arg2 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockITransactionMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITransaction)(nil).Update), arg0, arg1, arg2)
}

// MockIHTTP is a mock of IHTTP interface
type MockIHTTP struct {
	ctrl     *gomock.Controller
	recorder *MockIHTTPMockRecorder
}

// MockIHTTPMockRecorder is the mock recorder for MockIHTTP
type MockIHTTPMockRecorder struct {
	mock *MockIHTTP
}

// NewMockIHTTP creates a new mock instance
func NewMockIHTTP(ctrl *gomock.Controller) *MockIHTTP {
	mock := &MockIHTTP{ctrl: ctrl}
	mock.recorder = &MockIHTTPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHTTP) EXPECT() *MockIHTTPMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockIHTTP) Request(arg0 context.Context, arg1, arg2 string, arg3 map[string]interface{}, arg4 interface{}, arg5 bool) (*Response, error) {
	ret := m.ctrl.Call(m, "Request", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request
func (mr *MockIHTTPMockRecorder) Request(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockIHTTP)(nil).Request), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RequestRaw mocks base method
func (m *MockIHTTP) RequestRaw(arg0 context.Context, arg1, arg2 string, arg3 map[string]string, arg4 string) (*Response, error) {
	ret := m.ctrl.Call(m, "RequestRaw", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestRaw indicates an expected call of RequestRaw
func (mr *MockIHTTPMockRecorder) RequestRaw(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestRaw", reflect.TypeOf((*MockIHTTP)(nil).RequestRaw), arg0, arg1, arg2, arg3, arg4)
}

// MockIAuth is a mock of IAuth interface
type MockIAuth struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthMockRecorder
}

// MockIAuthMockRecorder is the mock recorder for MockIAuth
type MockIAuthMockRecorder struct {
	mock *MockIAuth
}

// NewMockIAuth creates a new mock instance
func NewMockIAuth(ctrl *gomock.Controller) *MockIAuth {
	mock := &MockIAuth{ctrl: ctrl}
	mock.recorder = &MockIAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAuth) EXPECT() *MockIAuthMockRecorder {
	return m.recorder
}

// GetTenantName mocks base method
func (m *MockIAuth) GetTenantName(arg0 Context) string {
	ret := m.ctrl.Call(m, "GetTenantName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTenantName indicates an expected call of GetTenantName
func (mr *MockIAuthMockRecorder) GetTenantName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantName", reflect.TypeOf((*MockIAuth)(nil).GetTenantName), arg0)
}

// HasRole mocks base method
func (m *MockIAuth) HasRole(arg0 Context, arg1 string) bool {
	ret := m.ctrl.Call(m, "HasRole", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasRole indicates an expected call of HasRole
func (mr *MockIAuthMockRecorder) HasRole(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasRole", reflect.TypeOf((*MockIAuth)(nil).HasRole), arg0, arg1)
}

// IsAdmin mocks base method
func (m *MockIAuth) IsAdmin(arg0 Context) bool {
	ret := m.ctrl.Call(m, "IsAdmin", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdmin indicates an expected call of IsAdmin
func (mr *MockIAuthMockRecorder) IsAdmin(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdmin", reflect.TypeOf((*MockIAuth)(nil).IsAdmin), arg0)
}

// MockIConfig is a mock of IConfig interface
type MockIConfig struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigMockRecorder
}

// MockIConfigMockRecorder is the mock recorder for MockIConfig
type MockIConfigMockRecorder struct {
	mock *MockIConfig
}

// NewMockIConfig creates a new mock instance
func NewMockIConfig(ctrl *gomock.Controller) *MockIConfig {
	mock := &MockIConfig{ctrl: ctrl}
	mock.recorder = &MockIConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConfig) EXPECT() *MockIConfigMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockIConfig) Config(arg0 string, arg1 interface{}) interface{} {
	ret := m.ctrl.Call(m, "Config", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockIConfigMockRecorder) Config(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockIConfig)(nil).Config), arg0, arg1)
}

// MockIUtil is a mock of IUtil interface
type MockIUtil struct {
	ctrl     *gomock.Controller
	recorder *MockIUtilMockRecorder
}

// MockIUtilMockRecorder is the mock recorder for MockIUtil
type MockIUtilMockRecorder struct {
	mock *MockIUtil
}

// NewMockIUtil creates a new mock instance
func NewMockIUtil(ctrl *gomock.Controller) *MockIUtil {
	mock := &MockIUtil{ctrl: ctrl}
	mock.recorder = &MockIUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUtil) EXPECT() *MockIUtilMockRecorder {
	return m.recorder
}

// GetTransaction mocks base method
func (m *MockIUtil) GetTransaction(arg0 Context) (ITransaction, bool) {
	ret := m.ctrl.Call(m, "GetTransaction", arg0)
	ret0, _ := ret[0].(ITransaction)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockIUtilMockRecorder) GetTransaction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockIUtil)(nil).GetTransaction), arg0)
}

// NewUUID mocks base method
func (m *MockIUtil) NewUUID() string {
	ret := m.ctrl.Call(m, "NewUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// NewUUID indicates an expected call of NewUUID
func (mr *MockIUtilMockRecorder) NewUUID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUID", reflect.TypeOf((*MockIUtil)(nil).NewUUID))
}

// ResourceFromMapForType mocks base method
func (m *MockIUtil) ResourceFromMapForType(arg0 map[string]interface{}, arg1 interface{}) (Resource, error) {
	ret := m.ctrl.Call(m, "ResourceFromMapForType", arg0, arg1)
	ret0, _ := ret[0].(Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceFromMapForType indicates an expected call of ResourceFromMapForType
func (mr *MockIUtilMockRecorder) ResourceFromMapForType(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceFromMapForType", reflect.TypeOf((*MockIUtil)(nil).ResourceFromMapForType), arg0, arg1)
}

// ResourceToMap mocks base method
func (m *MockIUtil) ResourceToMap(arg0 interface{}) map[string]interface{} {
	ret := m.ctrl.Call(m, "ResourceToMap", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// ResourceToMap indicates an expected call of ResourceToMap
func (mr *MockIUtilMockRecorder) ResourceToMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceToMap", reflect.TypeOf((*MockIUtil)(nil).ResourceToMap), arg0)
}
