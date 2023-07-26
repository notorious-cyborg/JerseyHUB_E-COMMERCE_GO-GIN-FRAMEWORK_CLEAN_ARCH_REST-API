// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/helper/interface/helper.go

// Package mockhelper is a generated GoMock package.
package mockhelper

import (
	models "jerseyhub/pkg/utils/models"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// AddImageToS3 mocks base method.
func (m *MockHelper) AddImageToS3(file *multipart.FileHeader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImageToS3", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddImageToS3 indicates an expected call of AddImageToS3.
func (mr *MockHelperMockRecorder) AddImageToS3(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImageToS3", reflect.TypeOf((*MockHelper)(nil).AddImageToS3), file)
}

// GenerateRefferalCode mocks base method.
func (m *MockHelper) GenerateRefferalCode() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefferalCode")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefferalCode indicates an expected call of GenerateRefferalCode.
func (mr *MockHelperMockRecorder) GenerateRefferalCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefferalCode", reflect.TypeOf((*MockHelper)(nil).GenerateRefferalCode))
}

// GenerateTokenAdmin mocks base method.
func (m *MockHelper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokenAdmin", admin)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateTokenAdmin indicates an expected call of GenerateTokenAdmin.
func (mr *MockHelperMockRecorder) GenerateTokenAdmin(admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokenAdmin", reflect.TypeOf((*MockHelper)(nil).GenerateTokenAdmin), admin)
}

// GenerateTokenClients mocks base method.
func (m *MockHelper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokenClients", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokenClients indicates an expected call of GenerateTokenClients.
func (mr *MockHelperMockRecorder) GenerateTokenClients(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokenClients", reflect.TypeOf((*MockHelper)(nil).GenerateTokenClients), user)
}

// PasswordHashing mocks base method.
func (m *MockHelper) PasswordHashing(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordHashing", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordHashing indicates an expected call of PasswordHashing.
func (mr *MockHelperMockRecorder) PasswordHashing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordHashing", reflect.TypeOf((*MockHelper)(nil).PasswordHashing), arg0)
}

// TwilioSendOTP mocks base method.
func (m *MockHelper) TwilioSendOTP(phone, serviceID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwilioSendOTP", phone, serviceID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TwilioSendOTP indicates an expected call of TwilioSendOTP.
func (mr *MockHelperMockRecorder) TwilioSendOTP(phone, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwilioSendOTP", reflect.TypeOf((*MockHelper)(nil).TwilioSendOTP), phone, serviceID)
}

// TwilioSetup mocks base method.
func (m *MockHelper) TwilioSetup(username, password string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TwilioSetup", username, password)
}

// TwilioSetup indicates an expected call of TwilioSetup.
func (mr *MockHelperMockRecorder) TwilioSetup(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwilioSetup", reflect.TypeOf((*MockHelper)(nil).TwilioSetup), username, password)
}

// TwilioVerifyOTP mocks base method.
func (m *MockHelper) TwilioVerifyOTP(serviceID, code, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwilioVerifyOTP", serviceID, code, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// TwilioVerifyOTP indicates an expected call of TwilioVerifyOTP.
func (mr *MockHelperMockRecorder) TwilioVerifyOTP(serviceID, code, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwilioVerifyOTP", reflect.TypeOf((*MockHelper)(nil).TwilioVerifyOTP), serviceID, code, phone)
}