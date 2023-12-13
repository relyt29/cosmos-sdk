// Code generated by MockGen. DO NOT EDIT.
// Source: x/distribution/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	types "cosmossdk.io/x/staking/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types0.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx context.Context, name string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(arg0 context.Context, arg1 types0.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", arg0, arg1)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BlockedAddr mocks base method.
func (m *MockBankKeeper) BlockedAddr(addr types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockedAddr", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockedAddr indicates an expected call of BlockedAddr.
func (mr *MockBankKeeperMockRecorder) BlockedAddr(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockedAddr", reflect.TypeOf((*MockBankKeeper)(nil).BlockedAddr), addr)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx context.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types0.AccAddress, recipientModule string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx context.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockPoolKeeper is a mock of PoolKeeper interface.
type MockPoolKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPoolKeeperMockRecorder
}

// MockPoolKeeperMockRecorder is the mock recorder for MockPoolKeeper.
type MockPoolKeeperMockRecorder struct {
	mock *MockPoolKeeper
}

// NewMockPoolKeeper creates a new mock instance.
func NewMockPoolKeeper(ctrl *gomock.Controller) *MockPoolKeeper {
	mock := &MockPoolKeeper{ctrl: ctrl}
	mock.recorder = &MockPoolKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolKeeper) EXPECT() *MockPoolKeeperMockRecorder {
	return m.recorder
}

// DistributeFromCommunityPool mocks base method.
func (m *MockPoolKeeper) DistributeFromCommunityPool(ctx context.Context, amount types0.Coins, receiveAddr types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DistributeFromCommunityPool", ctx, amount, receiveAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeFromCommunityPool indicates an expected call of DistributeFromCommunityPool.
func (mr *MockPoolKeeperMockRecorder) DistributeFromCommunityPool(ctx, amount, receiveAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeFromCommunityPool", reflect.TypeOf((*MockPoolKeeper)(nil).DistributeFromCommunityPool), ctx, amount, receiveAddr)
}

// FundCommunityPool mocks base method.
func (m *MockPoolKeeper) FundCommunityPool(ctx context.Context, amount types0.Coins, sender types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FundCommunityPool", ctx, amount, sender)
	ret0, _ := ret[0].(error)
	return ret0
}

// FundCommunityPool indicates an expected call of FundCommunityPool.
func (mr *MockPoolKeeperMockRecorder) FundCommunityPool(ctx, amount, sender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FundCommunityPool", reflect.TypeOf((*MockPoolKeeper)(nil).FundCommunityPool), ctx, amount, sender)
}

// GetCommunityPool mocks base method.
func (m *MockPoolKeeper) GetCommunityPool(ctx context.Context) (types0.Coins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommunityPool", ctx)
	ret0, _ := ret[0].(types0.Coins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommunityPool indicates an expected call of GetCommunityPool.
func (mr *MockPoolKeeperMockRecorder) GetCommunityPool(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommunityPool", reflect.TypeOf((*MockPoolKeeper)(nil).GetCommunityPool), ctx)
}

// SetToDistribute mocks base method.
func (m *MockPoolKeeper) SetToDistribute(ctx context.Context, amount types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetToDistribute", ctx, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetToDistribute indicates an expected call of SetToDistribute.
func (mr *MockPoolKeeperMockRecorder) SetToDistribute(ctx, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToDistribute", reflect.TypeOf((*MockPoolKeeper)(nil).SetToDistribute), ctx, amount)
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BondDenom mocks base method.
func (m *MockStakingKeeper) BondDenom(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BondDenom indicates an expected call of BondDenom.
func (mr *MockStakingKeeperMockRecorder) BondDenom(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), ctx)
}

// ConsensusAddressCodec mocks base method.
func (m *MockStakingKeeper) ConsensusAddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsensusAddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// ConsensusAddressCodec indicates an expected call of ConsensusAddressCodec.
func (mr *MockStakingKeeperMockRecorder) ConsensusAddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsensusAddressCodec", reflect.TypeOf((*MockStakingKeeper)(nil).ConsensusAddressCodec))
}

// Delegation mocks base method.
func (m *MockStakingKeeper) Delegation(arg0 context.Context, arg1 types0.AccAddress, arg2 types0.ValAddress) (types0.DelegationI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegation", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.DelegationI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delegation indicates an expected call of Delegation.
func (mr *MockStakingKeeperMockRecorder) Delegation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegation", reflect.TypeOf((*MockStakingKeeper)(nil).Delegation), arg0, arg1, arg2)
}

// GetAllDelegatorDelegations mocks base method.
func (m *MockStakingKeeper) GetAllDelegatorDelegations(ctx context.Context, delegator types0.AccAddress) ([]types.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDelegatorDelegations", ctx, delegator)
	ret0, _ := ret[0].([]types.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDelegatorDelegations indicates an expected call of GetAllDelegatorDelegations.
func (mr *MockStakingKeeperMockRecorder) GetAllDelegatorDelegations(ctx, delegator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDelegatorDelegations", reflect.TypeOf((*MockStakingKeeper)(nil).GetAllDelegatorDelegations), ctx, delegator)
}

// GetAllSDKDelegations mocks base method.
func (m *MockStakingKeeper) GetAllSDKDelegations(ctx context.Context) ([]types.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSDKDelegations", ctx)
	ret0, _ := ret[0].([]types.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSDKDelegations indicates an expected call of GetAllSDKDelegations.
func (mr *MockStakingKeeperMockRecorder) GetAllSDKDelegations(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSDKDelegations", reflect.TypeOf((*MockStakingKeeper)(nil).GetAllSDKDelegations), ctx)
}

// GetAllValidators mocks base method.
func (m *MockStakingKeeper) GetAllValidators(ctx context.Context) ([]types.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllValidators", ctx)
	ret0, _ := ret[0].([]types.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllValidators indicates an expected call of GetAllValidators.
func (mr *MockStakingKeeperMockRecorder) GetAllValidators(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllValidators", reflect.TypeOf((*MockStakingKeeper)(nil).GetAllValidators), ctx)
}

// IterateDelegations mocks base method.
func (m *MockStakingKeeper) IterateDelegations(ctx context.Context, delegator types0.AccAddress, fn func(int64, types0.DelegationI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateDelegations", ctx, delegator, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateDelegations indicates an expected call of IterateDelegations.
func (mr *MockStakingKeeperMockRecorder) IterateDelegations(ctx, delegator, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateDelegations", reflect.TypeOf((*MockStakingKeeper)(nil).IterateDelegations), ctx, delegator, fn)
}

// IterateValidators mocks base method.
func (m *MockStakingKeeper) IterateValidators(arg0 context.Context, arg1 func(int64, types0.ValidatorI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateValidators", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateValidators indicates an expected call of IterateValidators.
func (mr *MockStakingKeeperMockRecorder) IterateValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateValidators", reflect.TypeOf((*MockStakingKeeper)(nil).IterateValidators), arg0, arg1)
}

// Validator mocks base method.
func (m *MockStakingKeeper) Validator(arg0 context.Context, arg1 types0.ValAddress) (types0.ValidatorI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator", arg0, arg1)
	ret0, _ := ret[0].(types0.ValidatorI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validator indicates an expected call of Validator.
func (mr *MockStakingKeeperMockRecorder) Validator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockStakingKeeper)(nil).Validator), arg0, arg1)
}

// ValidatorAddressCodec mocks base method.
func (m *MockStakingKeeper) ValidatorAddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorAddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// ValidatorAddressCodec indicates an expected call of ValidatorAddressCodec.
func (mr *MockStakingKeeperMockRecorder) ValidatorAddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorAddressCodec", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorAddressCodec))
}

// ValidatorByConsAddr mocks base method.
func (m *MockStakingKeeper) ValidatorByConsAddr(arg0 context.Context, arg1 types0.ConsAddress) (types0.ValidatorI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types0.ValidatorI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockStakingKeeperMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockStakingHooks is a mock of StakingHooks interface.
type MockStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockStakingHooksMockRecorder
}

// MockStakingHooksMockRecorder is the mock recorder for MockStakingHooks.
type MockStakingHooksMockRecorder struct {
	mock *MockStakingHooks
}

// NewMockStakingHooks creates a new mock instance.
func NewMockStakingHooks(ctrl *gomock.Controller) *MockStakingHooks {
	mock := &MockStakingHooks{ctrl: ctrl}
	mock.recorder = &MockStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingHooks) EXPECT() *MockStakingHooksMockRecorder {
	return m.recorder
}

// AfterDelegationModified mocks base method.
func (m *MockStakingHooks) AfterDelegationModified(ctx context.Context, delAddr types0.AccAddress, valAddr types0.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterDelegationModified", ctx, delAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterDelegationModified indicates an expected call of AfterDelegationModified.
func (mr *MockStakingHooksMockRecorder) AfterDelegationModified(ctx, delAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterDelegationModified", reflect.TypeOf((*MockStakingHooks)(nil).AfterDelegationModified), ctx, delAddr, valAddr)
}

// AfterValidatorCreated mocks base method.
func (m *MockStakingHooks) AfterValidatorCreated(ctx context.Context, valAddr types0.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorCreated", ctx, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorCreated indicates an expected call of AfterValidatorCreated.
func (mr *MockStakingHooksMockRecorder) AfterValidatorCreated(ctx, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorCreated", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorCreated), ctx, valAddr)
}
