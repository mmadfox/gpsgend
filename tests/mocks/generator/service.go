// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/mmadfox/projects/gpsio/gpsgend/internal/generator/service.go

// Package mock_generator is a generated GoMock package.
package mock_generator

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	go_gpsgen "github.com/mmadfox/go-gpsgen"
	proto "github.com/mmadfox/go-gpsgen/proto"
	generator "github.com/mmadfox/gpsgend/internal/generator"
	types "github.com/mmadfox/gpsgend/internal/types"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddRoutes mocks base method.
func (m *MockService) AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*go_gpsgen.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoutes", ctx, trackerID, newRoutes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoutes indicates an expected call of AddRoutes.
func (mr *MockServiceMockRecorder) AddRoutes(ctx, trackerID, newRoutes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoutes", reflect.TypeOf((*MockService)(nil).AddRoutes), ctx, trackerID, newRoutes)
}

// AddSensor mocks base method.
func (m *MockService) AddSensor(ctx context.Context, trackerID types.ID, sensor *types.Sensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSensor", ctx, trackerID, sensor)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSensor indicates an expected call of AddSensor.
func (mr *MockServiceMockRecorder) AddSensor(ctx, trackerID, sensor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSensor", reflect.TypeOf((*MockService)(nil).AddSensor), ctx, trackerID, sensor)
}

// FindTracker mocks base method.
func (m *MockService) FindTracker(ctx context.Context, trackerID types.ID) (*generator.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTracker", ctx, trackerID)
	ret0, _ := ret[0].(*generator.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTracker indicates an expected call of FindTracker.
func (mr *MockServiceMockRecorder) FindTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTracker", reflect.TypeOf((*MockService)(nil).FindTracker), ctx, trackerID)
}

// MoveToRoute mocks base method.
func (m *MockService) MoveToRoute(ctx context.Context, trackerID types.ID, routeIndex int) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveToRoute", ctx, trackerID, routeIndex)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveToRoute indicates an expected call of MoveToRoute.
func (mr *MockServiceMockRecorder) MoveToRoute(ctx, trackerID, routeIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToRoute", reflect.TypeOf((*MockService)(nil).MoveToRoute), ctx, trackerID, routeIndex)
}

// MoveToRouteByID mocks base method.
func (m *MockService) MoveToRouteByID(ctx context.Context, trackerID, routeID types.ID) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveToRouteByID", ctx, trackerID, routeID)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveToRouteByID indicates an expected call of MoveToRouteByID.
func (mr *MockServiceMockRecorder) MoveToRouteByID(ctx, trackerID, routeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToRouteByID", reflect.TypeOf((*MockService)(nil).MoveToRouteByID), ctx, trackerID, routeID)
}

// MoveToSegment mocks base method.
func (m *MockService) MoveToSegment(ctx context.Context, trackerID types.ID, routeIndex, trackIndex, segmentIndex int) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveToSegment", ctx, trackerID, routeIndex, trackIndex, segmentIndex)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveToSegment indicates an expected call of MoveToSegment.
func (mr *MockServiceMockRecorder) MoveToSegment(ctx, trackerID, routeIndex, trackIndex, segmentIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToSegment", reflect.TypeOf((*MockService)(nil).MoveToSegment), ctx, trackerID, routeIndex, trackIndex, segmentIndex)
}

// MoveToTrack mocks base method.
func (m *MockService) MoveToTrack(ctx context.Context, trackerID types.ID, routeIndex, trackIndex int) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveToTrack", ctx, trackerID, routeIndex, trackIndex)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveToTrack indicates an expected call of MoveToTrack.
func (mr *MockServiceMockRecorder) MoveToTrack(ctx, trackerID, routeIndex, trackIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToTrack", reflect.TypeOf((*MockService)(nil).MoveToTrack), ctx, trackerID, routeIndex, trackIndex)
}

// MoveToTrackByID mocks base method.
func (m *MockService) MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID types.ID) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveToTrackByID", ctx, trackerID, routeID, trackID)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveToTrackByID indicates an expected call of MoveToTrackByID.
func (mr *MockServiceMockRecorder) MoveToTrackByID(ctx, trackerID, routeID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToTrackByID", reflect.TypeOf((*MockService)(nil).MoveToTrackByID), ctx, trackerID, routeID, trackID)
}

// NewTracker mocks base method.
func (m *MockService) NewTracker(ctx context.Context, opts *generator.NewTrackerOptions) (*generator.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTracker", ctx, opts)
	ret0, _ := ret[0].(*generator.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTracker indicates an expected call of NewTracker.
func (mr *MockServiceMockRecorder) NewTracker(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTracker", reflect.TypeOf((*MockService)(nil).NewTracker), ctx, opts)
}

// RemoveRoute mocks base method.
func (m *MockService) RemoveRoute(ctx context.Context, trackerID, routeID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRoute", ctx, trackerID, routeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRoute indicates an expected call of RemoveRoute.
func (mr *MockServiceMockRecorder) RemoveRoute(ctx, trackerID, routeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRoute", reflect.TypeOf((*MockService)(nil).RemoveRoute), ctx, trackerID, routeID)
}

// RemoveSensor mocks base method.
func (m *MockService) RemoveSensor(ctx context.Context, trackerID, sensorID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSensor", ctx, trackerID, sensorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSensor indicates an expected call of RemoveSensor.
func (mr *MockServiceMockRecorder) RemoveSensor(ctx, trackerID, sensorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSensor", reflect.TypeOf((*MockService)(nil).RemoveSensor), ctx, trackerID, sensorID)
}

// RemoveTracker mocks base method.
func (m *MockService) RemoveTracker(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTracker", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTracker indicates an expected call of RemoveTracker.
func (mr *MockServiceMockRecorder) RemoveTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTracker", reflect.TypeOf((*MockService)(nil).RemoveTracker), ctx, trackerID)
}

// ResetNavigator mocks base method.
func (m *MockService) ResetNavigator(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetNavigator", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetNavigator indicates an expected call of ResetNavigator.
func (mr *MockServiceMockRecorder) ResetNavigator(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetNavigator", reflect.TypeOf((*MockService)(nil).ResetNavigator), ctx, trackerID)
}

// ResetRoutes mocks base method.
func (m *MockService) ResetRoutes(ctx context.Context, trackerID types.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetRoutes", ctx, trackerID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetRoutes indicates an expected call of ResetRoutes.
func (mr *MockServiceMockRecorder) ResetRoutes(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetRoutes", reflect.TypeOf((*MockService)(nil).ResetRoutes), ctx, trackerID)
}

// ResumeTracker mocks base method.
func (m *MockService) ResumeTracker(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeTracker", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeTracker indicates an expected call of ResumeTracker.
func (mr *MockServiceMockRecorder) ResumeTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeTracker", reflect.TypeOf((*MockService)(nil).ResumeTracker), ctx, trackerID)
}

// RouteAt mocks base method.
func (m *MockService) RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (*go_gpsgen.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteAt", ctx, trackerID, routeIndex)
	ret0, _ := ret[0].(*go_gpsgen.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteAt indicates an expected call of RouteAt.
func (mr *MockServiceMockRecorder) RouteAt(ctx, trackerID, routeIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteAt", reflect.TypeOf((*MockService)(nil).RouteAt), ctx, trackerID, routeIndex)
}

// RouteByID mocks base method.
func (m *MockService) RouteByID(ctx context.Context, trackerID, routeID types.ID) (*go_gpsgen.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteByID", ctx, trackerID, routeID)
	ret0, _ := ret[0].(*go_gpsgen.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteByID indicates an expected call of RouteByID.
func (mr *MockServiceMockRecorder) RouteByID(ctx, trackerID, routeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteByID", reflect.TypeOf((*MockService)(nil).RouteByID), ctx, trackerID, routeID)
}

// Routes mocks base method.
func (m *MockService) Routes(ctx context.Context, trackerID types.ID) ([]*go_gpsgen.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes", ctx, trackerID)
	ret0, _ := ret[0].([]*go_gpsgen.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Routes indicates an expected call of Routes.
func (mr *MockServiceMockRecorder) Routes(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockService)(nil).Routes), ctx, trackerID)
}

// SearchTrackers mocks base method.
func (m *MockService) SearchTrackers(ctx context.Context, f generator.Filter) (generator.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTrackers", ctx, f)
	ret0, _ := ret[0].(generator.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTrackers indicates an expected call of SearchTrackers.
func (mr *MockServiceMockRecorder) SearchTrackers(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTrackers", reflect.TypeOf((*MockService)(nil).SearchTrackers), ctx, f)
}

// Sensors mocks base method.
func (m *MockService) Sensors(ctx context.Context, trackerID types.ID) ([]*types.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sensors", ctx, trackerID)
	ret0, _ := ret[0].([]*types.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sensors indicates an expected call of Sensors.
func (mr *MockServiceMockRecorder) Sensors(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sensors", reflect.TypeOf((*MockService)(nil).Sensors), ctx, trackerID)
}

// ShutdownTracker mocks base method.
func (m *MockService) ShutdownTracker(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownTracker", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutdownTracker indicates an expected call of ShutdownTracker.
func (mr *MockServiceMockRecorder) ShutdownTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownTracker", reflect.TypeOf((*MockService)(nil).ShutdownTracker), ctx, trackerID)
}

// StartTracker mocks base method.
func (m *MockService) StartTracker(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTracker", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTracker indicates an expected call of StartTracker.
func (mr *MockServiceMockRecorder) StartTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTracker", reflect.TypeOf((*MockService)(nil).StartTracker), ctx, trackerID)
}

// StopTracker mocks base method.
func (m *MockService) StopTracker(ctx context.Context, trackerID types.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopTracker", ctx, trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopTracker indicates an expected call of StopTracker.
func (mr *MockServiceMockRecorder) StopTracker(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopTracker", reflect.TypeOf((*MockService)(nil).StopTracker), ctx, trackerID)
}

// ToNextRoute mocks base method.
func (m *MockService) ToNextRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToNextRoute", ctx, trackerID)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ToNextRoute indicates an expected call of ToNextRoute.
func (mr *MockServiceMockRecorder) ToNextRoute(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToNextRoute", reflect.TypeOf((*MockService)(nil).ToNextRoute), ctx, trackerID)
}

// ToPrevRoute mocks base method.
func (m *MockService) ToPrevRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToPrevRoute", ctx, trackerID)
	ret0, _ := ret[0].(types.Navigator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ToPrevRoute indicates an expected call of ToPrevRoute.
func (mr *MockServiceMockRecorder) ToPrevRoute(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToPrevRoute", reflect.TypeOf((*MockService)(nil).ToPrevRoute), ctx, trackerID)
}

// TrackerState mocks base method.
func (m *MockService) TrackerState(ctx context.Context, trackerID types.ID) (*proto.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackerState", ctx, trackerID)
	ret0, _ := ret[0].(*proto.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrackerState indicates an expected call of TrackerState.
func (mr *MockServiceMockRecorder) TrackerState(ctx, trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackerState", reflect.TypeOf((*MockService)(nil).TrackerState), ctx, trackerID)
}

// UpdateTracker mocks base method.
func (m *MockService) UpdateTracker(ctx context.Context, trackerID types.ID, opts generator.UpdateTrackerOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTracker", ctx, trackerID, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTracker indicates an expected call of UpdateTracker.
func (mr *MockServiceMockRecorder) UpdateTracker(ctx, trackerID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTracker", reflect.TypeOf((*MockService)(nil).UpdateTracker), ctx, trackerID, opts)
}
