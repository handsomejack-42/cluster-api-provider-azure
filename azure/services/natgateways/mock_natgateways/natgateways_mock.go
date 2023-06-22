/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../natgateways.go

// Package mock_natgateways is a generated GoMock package.
package mock_natgateways

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockNatGatewayScope is a mock of NatGatewayScope interface.
type MockNatGatewayScope struct {
	ctrl     *gomock.Controller
	recorder *MockNatGatewayScopeMockRecorder
}

// MockNatGatewayScopeMockRecorder is the mock recorder for MockNatGatewayScope.
type MockNatGatewayScopeMockRecorder struct {
	mock *MockNatGatewayScope
}

// NewMockNatGatewayScope creates a new mock instance.
func NewMockNatGatewayScope(ctrl *gomock.Controller) *MockNatGatewayScope {
	mock := &MockNatGatewayScope{ctrl: ctrl}
	mock.recorder = &MockNatGatewayScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatGatewayScope) EXPECT() *MockNatGatewayScopeMockRecorder {
	return m.recorder
}

// APIServerLB mocks base method.
func (m *MockNatGatewayScope) APIServerLB() *v1beta1.LoadBalancerSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLB")
	ret0, _ := ret[0].(*v1beta1.LoadBalancerSpec)
	return ret0
}

// APIServerLB indicates an expected call of APIServerLB.
func (mr *MockNatGatewayScopeMockRecorder) APIServerLB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLB", reflect.TypeOf((*MockNatGatewayScope)(nil).APIServerLB))
}

// APIServerLBName mocks base method.
func (m *MockNatGatewayScope) APIServerLBName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBName")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBName indicates an expected call of APIServerLBName.
func (mr *MockNatGatewayScopeMockRecorder) APIServerLBName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBName", reflect.TypeOf((*MockNatGatewayScope)(nil).APIServerLBName))
}

// APIServerLBPoolName mocks base method.
func (m *MockNatGatewayScope) APIServerLBPoolName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBPoolName")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBPoolName indicates an expected call of APIServerLBPoolName.
func (mr *MockNatGatewayScopeMockRecorder) APIServerLBPoolName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBPoolName", reflect.TypeOf((*MockNatGatewayScope)(nil).APIServerLBPoolName))
}

// AdditionalTags mocks base method.
func (m *MockNatGatewayScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockNatGatewayScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockNatGatewayScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockNatGatewayScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockNatGatewayScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockNatGatewayScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockNatGatewayScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockNatGatewayScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockNatGatewayScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockNatGatewayScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockNatGatewayScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockNatGatewayScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockNatGatewayScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockNatGatewayScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockNatGatewayScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockNatGatewayScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockNatGatewayScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockNatGatewayScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockNatGatewayScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockNatGatewayScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockNatGatewayScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockNatGatewayScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockNatGatewayScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockNatGatewayScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockNatGatewayScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockNatGatewayScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockNatGatewayScope)(nil).ClusterName))
}

// ControlPlaneRouteTable mocks base method.
func (m *MockNatGatewayScope) ControlPlaneRouteTable() v1beta1.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneRouteTable")
	ret0, _ := ret[0].(v1beta1.RouteTable)
	return ret0
}

// ControlPlaneRouteTable indicates an expected call of ControlPlaneRouteTable.
func (mr *MockNatGatewayScopeMockRecorder) ControlPlaneRouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneRouteTable", reflect.TypeOf((*MockNatGatewayScope)(nil).ControlPlaneRouteTable))
}

// ControlPlaneSubnet mocks base method.
func (m *MockNatGatewayScope) ControlPlaneSubnet() v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(v1beta1.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockNatGatewayScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockNatGatewayScope)(nil).ControlPlaneSubnet))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockNatGatewayScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockNatGatewayScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockNatGatewayScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// ExtendedLocation mocks base method.
func (m *MockNatGatewayScope) ExtendedLocation() *v1beta1.ExtendedLocationSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocation")
	ret0, _ := ret[0].(*v1beta1.ExtendedLocationSpec)
	return ret0
}

// ExtendedLocation indicates an expected call of ExtendedLocation.
func (mr *MockNatGatewayScopeMockRecorder) ExtendedLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocation", reflect.TypeOf((*MockNatGatewayScope)(nil).ExtendedLocation))
}

// ExtendedLocationName mocks base method.
func (m *MockNatGatewayScope) ExtendedLocationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationName indicates an expected call of ExtendedLocationName.
func (mr *MockNatGatewayScopeMockRecorder) ExtendedLocationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationName", reflect.TypeOf((*MockNatGatewayScope)(nil).ExtendedLocationName))
}

// ExtendedLocationType mocks base method.
func (m *MockNatGatewayScope) ExtendedLocationType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationType indicates an expected call of ExtendedLocationType.
func (mr *MockNatGatewayScopeMockRecorder) ExtendedLocationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationType", reflect.TypeOf((*MockNatGatewayScope)(nil).ExtendedLocationType))
}

// FailureDomains mocks base method.
func (m *MockNatGatewayScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockNatGatewayScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockNatGatewayScope)(nil).FailureDomains))
}

// GetLongRunningOperationState mocks base method.
func (m *MockNatGatewayScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockNatGatewayScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockNatGatewayScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// GetPrivateDNSZoneName mocks base method.
func (m *MockNatGatewayScope) GetPrivateDNSZoneName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDNSZoneName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrivateDNSZoneName indicates an expected call of GetPrivateDNSZoneName.
func (mr *MockNatGatewayScopeMockRecorder) GetPrivateDNSZoneName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDNSZoneName", reflect.TypeOf((*MockNatGatewayScope)(nil).GetPrivateDNSZoneName))
}

// HashKey mocks base method.
func (m *MockNatGatewayScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockNatGatewayScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockNatGatewayScope)(nil).HashKey))
}

// IsAPIServerPrivate mocks base method.
func (m *MockNatGatewayScope) IsAPIServerPrivate() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAPIServerPrivate")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAPIServerPrivate indicates an expected call of IsAPIServerPrivate.
func (mr *MockNatGatewayScopeMockRecorder) IsAPIServerPrivate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAPIServerPrivate", reflect.TypeOf((*MockNatGatewayScope)(nil).IsAPIServerPrivate))
}

// IsIPv6Enabled mocks base method.
func (m *MockNatGatewayScope) IsIPv6Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIPv6Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIPv6Enabled indicates an expected call of IsIPv6Enabled.
func (mr *MockNatGatewayScopeMockRecorder) IsIPv6Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIPv6Enabled", reflect.TypeOf((*MockNatGatewayScope)(nil).IsIPv6Enabled))
}

// IsVnetManaged mocks base method.
func (m *MockNatGatewayScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockNatGatewayScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockNatGatewayScope)(nil).IsVnetManaged))
}

// Location mocks base method.
func (m *MockNatGatewayScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockNatGatewayScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockNatGatewayScope)(nil).Location))
}

// NatGatewaySpecs mocks base method.
func (m *MockNatGatewayScope) NatGatewaySpecs() []azure.ResourceSpecGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NatGatewaySpecs")
	ret0, _ := ret[0].([]azure.ResourceSpecGetter)
	return ret0
}

// NatGatewaySpecs indicates an expected call of NatGatewaySpecs.
func (mr *MockNatGatewayScopeMockRecorder) NatGatewaySpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatGatewaySpecs", reflect.TypeOf((*MockNatGatewayScope)(nil).NatGatewaySpecs))
}

// NodeSubnets mocks base method.
func (m *MockNatGatewayScope) NodeSubnets() []v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnets")
	ret0, _ := ret[0].([]v1beta1.SubnetSpec)
	return ret0
}

// NodeSubnets indicates an expected call of NodeSubnets.
func (mr *MockNatGatewayScopeMockRecorder) NodeSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnets", reflect.TypeOf((*MockNatGatewayScope)(nil).NodeSubnets))
}

// OutboundLBName mocks base method.
func (m *MockNatGatewayScope) OutboundLBName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundLBName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundLBName indicates an expected call of OutboundLBName.
func (mr *MockNatGatewayScopeMockRecorder) OutboundLBName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundLBName", reflect.TypeOf((*MockNatGatewayScope)(nil).OutboundLBName), arg0)
}

// OutboundPoolName mocks base method.
func (m *MockNatGatewayScope) OutboundPoolName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundPoolName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundPoolName indicates an expected call of OutboundPoolName.
func (mr *MockNatGatewayScopeMockRecorder) OutboundPoolName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundPoolName", reflect.TypeOf((*MockNatGatewayScope)(nil).OutboundPoolName), arg0)
}

// ResourceGroup mocks base method.
func (m *MockNatGatewayScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockNatGatewayScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockNatGatewayScope)(nil).ResourceGroup))
}

// SetLongRunningOperationState mocks base method.
func (m *MockNatGatewayScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockNatGatewayScopeMockRecorder) SetLongRunningOperationState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockNatGatewayScope)(nil).SetLongRunningOperationState), arg0)
}

// SetNatGatewayIDInSubnets mocks base method.
func (m *MockNatGatewayScope) SetNatGatewayIDInSubnets(natGatewayName, natGatewayID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNatGatewayIDInSubnets", natGatewayName, natGatewayID)
}

// SetNatGatewayIDInSubnets indicates an expected call of SetNatGatewayIDInSubnets.
func (mr *MockNatGatewayScopeMockRecorder) SetNatGatewayIDInSubnets(natGatewayName, natGatewayID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNatGatewayIDInSubnets", reflect.TypeOf((*MockNatGatewayScope)(nil).SetNatGatewayIDInSubnets), natGatewayName, natGatewayID)
}

// SetSubnet mocks base method.
func (m *MockNatGatewayScope) SetSubnet(arg0 v1beta1.SubnetSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnet", arg0)
}

// SetSubnet indicates an expected call of SetSubnet.
func (mr *MockNatGatewayScopeMockRecorder) SetSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnet", reflect.TypeOf((*MockNatGatewayScope)(nil).SetSubnet), arg0)
}

// Subnet mocks base method.
func (m *MockNatGatewayScope) Subnet(arg0 string) v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnet", arg0)
	ret0, _ := ret[0].(v1beta1.SubnetSpec)
	return ret0
}

// Subnet indicates an expected call of Subnet.
func (mr *MockNatGatewayScopeMockRecorder) Subnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnet", reflect.TypeOf((*MockNatGatewayScope)(nil).Subnet), arg0)
}

// Subnets mocks base method.
func (m *MockNatGatewayScope) Subnets() v1beta1.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1beta1.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockNatGatewayScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockNatGatewayScope)(nil).Subnets))
}

// SubscriptionID mocks base method.
func (m *MockNatGatewayScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockNatGatewayScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockNatGatewayScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockNatGatewayScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockNatGatewayScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockNatGatewayScope)(nil).TenantID))
}

// UpdateDeleteStatus mocks base method.
func (m *MockNatGatewayScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockNatGatewayScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockNatGatewayScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockNatGatewayScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockNatGatewayScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockNatGatewayScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockNatGatewayScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockNatGatewayScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockNatGatewayScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}

// Vnet mocks base method.
func (m *MockNatGatewayScope) Vnet() *v1beta1.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1beta1.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockNatGatewayScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockNatGatewayScope)(nil).Vnet))
}
