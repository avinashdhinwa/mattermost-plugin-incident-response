// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-incident-response/server/incident (interfaces: Store)

// Package mock_incident is a generated GoMock package.
package mock_incident

import (
	gomock "github.com/golang/mock/gomock"
	incident "github.com/mattermost/mattermost-plugin-incident-response/server/incident"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateIncident mocks base method
func (m *MockStore) CreateIncident(arg0 *incident.Incident) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncident", arg0)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncident indicates an expected call of CreateIncident
func (mr *MockStoreMockRecorder) CreateIncident(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncident", reflect.TypeOf((*MockStore)(nil).CreateIncident), arg0)
}

// GetAllIncidentMembersCount mocks base method
func (m *MockStore) GetAllIncidentMembersCount(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIncidentMembersCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIncidentMembersCount indicates an expected call of GetAllIncidentMembersCount
func (mr *MockStoreMockRecorder) GetAllIncidentMembersCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIncidentMembersCount", reflect.TypeOf((*MockStore)(nil).GetAllIncidentMembersCount), arg0)
}

// GetCommanders mocks base method
func (m *MockStore) GetCommanders(arg0 incident.HeaderFilterOptions) ([]incident.CommanderInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommanders", arg0)
	ret0, _ := ret[0].([]incident.CommanderInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommanders indicates an expected call of GetCommanders
func (mr *MockStoreMockRecorder) GetCommanders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommanders", reflect.TypeOf((*MockStore)(nil).GetCommanders), arg0)
}

// GetIncident mocks base method
func (m *MockStore) GetIncident(arg0 string) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncident", arg0)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncident indicates an expected call of GetIncident
func (mr *MockStoreMockRecorder) GetIncident(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncident", reflect.TypeOf((*MockStore)(nil).GetIncident), arg0)
}

// GetIncidentIDForChannel mocks base method
func (m *MockStore) GetIncidentIDForChannel(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentIDForChannel", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentIDForChannel indicates an expected call of GetIncidentIDForChannel
func (mr *MockStoreMockRecorder) GetIncidentIDForChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentIDForChannel", reflect.TypeOf((*MockStore)(nil).GetIncidentIDForChannel), arg0)
}

// GetIncidents mocks base method
func (m *MockStore) GetIncidents(arg0 incident.HeaderFilterOptions) (*incident.GetIncidentsResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidents", arg0)
	ret0, _ := ret[0].(*incident.GetIncidentsResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidents indicates an expected call of GetIncidents
func (mr *MockStoreMockRecorder) GetIncidents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidents", reflect.TypeOf((*MockStore)(nil).GetIncidents), arg0)
}

// NukeDB mocks base method
func (m *MockStore) NukeDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NukeDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// NukeDB indicates an expected call of NukeDB
func (mr *MockStoreMockRecorder) NukeDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NukeDB", reflect.TypeOf((*MockStore)(nil).NukeDB))
}

// UpdateIncident mocks base method
func (m *MockStore) UpdateIncident(arg0 *incident.Incident) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIncident", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIncident indicates an expected call of UpdateIncident
func (mr *MockStoreMockRecorder) UpdateIncident(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIncident", reflect.TypeOf((*MockStore)(nil).UpdateIncident), arg0)
}
