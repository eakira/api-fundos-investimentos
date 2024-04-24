// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/port/output/FundosOutput.go
//
// Generated by this command:
//
//	mockgen -source=src/application/port/output/FundosOutput.go -destination=src/test/mocks/FundosOutput_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	response "api-fundos-investimentos/adapter/output/model/response"
	domain "api-fundos-investimentos/application/domain"
	resterrors "api-fundos-investimentos/configuration/resterrors"
	reflect "reflect"
	sync "sync"

	gomock "go.uber.org/mock/gomock"
)

// MockFundosPort is a mock of FundosPort interface.
type MockFundosPort struct {
	ctrl     *gomock.Controller
	recorder *MockFundosPortMockRecorder
}

// MockFundosPortMockRecorder is the mock recorder for MockFundosPort.
type MockFundosPortMockRecorder struct {
	mock *MockFundosPort
}

// NewMockFundosPort creates a new mock instance.
func NewMockFundosPort(ctrl *gomock.Controller) *MockFundosPort {
	mock := &MockFundosPort{ctrl: ctrl}
	mock.recorder = &MockFundosPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFundosPort) EXPECT() *MockFundosPortMockRecorder {
	return m.recorder
}

// CreateArquivosRepository mocks base method.
func (m *MockFundosPort) CreateArquivosRepository(arg0 domain.ArquivosDomain) (*domain.ArquivosDomain, *resterrors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArquivosRepository", arg0)
	ret0, _ := ret[0].(*domain.ArquivosDomain)
	ret1, _ := ret[1].(*resterrors.RestErr)
	return ret0, ret1
}

// CreateArquivosRepository indicates an expected call of CreateArquivosRepository.
func (mr *MockFundosPortMockRecorder) CreateArquivosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArquivosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateArquivosRepository), arg0)
}

// CreateFundosRepository mocks base method.
func (m *MockFundosPort) CreateFundosRepository(arg0 domain.FundosDomain) (*domain.FundosDomain, *resterrors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFundosRepository", arg0)
	ret0, _ := ret[0].(*domain.FundosDomain)
	ret1, _ := ret[1].(*resterrors.RestErr)
	return ret0, ret1
}

// CreateFundosRepository indicates an expected call of CreateFundosRepository.
func (mr *MockFundosPortMockRecorder) CreateFundosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFundosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateFundosRepository), arg0)
}

// CreateManyBalecenteRepository mocks base method.
func (m *MockFundosPort) CreateManyBalecenteRepository(arg0 []domain.BalanceteDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyBalecenteRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyBalecenteRepository indicates an expected call of CreateManyBalecenteRepository.
func (mr *MockFundosPortMockRecorder) CreateManyBalecenteRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyBalecenteRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyBalecenteRepository), arg0)
}

// CreateManyCdaAgroCreditoPrivadoRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaAgroCreditoPrivadoRepository(arg0 []domain.CdaAgroCreditoPrivadoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaAgroCreditoPrivadoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaAgroCreditoPrivadoRepository indicates an expected call of CreateManyCdaAgroCreditoPrivadoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaAgroCreditoPrivadoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaAgroCreditoPrivadoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaAgroCreditoPrivadoRepository), arg0)
}

// CreateManyCdaConfidencialRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaConfidencialRepository(arg0 []domain.CdaConfidencialDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaConfidencialRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaConfidencialRepository indicates an expected call of CreateManyCdaConfidencialRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaConfidencialRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaConfidencialRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaConfidencialRepository), arg0)
}

// CreateManyCdaDemaisAtivosNaoCodificadosRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaDemaisAtivosNaoCodificadosRepository(arg0 []domain.CdaDemaisAtivosNaoCodificadosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaDemaisAtivosNaoCodificadosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaDemaisAtivosNaoCodificadosRepository indicates an expected call of CreateManyCdaDemaisAtivosNaoCodificadosRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaDemaisAtivosNaoCodificadosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaDemaisAtivosNaoCodificadosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaDemaisAtivosNaoCodificadosRepository), arg0)
}

// CreateManyCdaDemaisAtivosRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaDemaisAtivosRepository(arg0 []domain.CdaDemaisAtivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaDemaisAtivosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaDemaisAtivosRepository indicates an expected call of CreateManyCdaDemaisAtivosRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaDemaisAtivosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaDemaisAtivosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaDemaisAtivosRepository), arg0)
}

// CreateManyCdaDepositoAPrazoOutrosAtivosRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaDepositoAPrazoOutrosAtivosRepository(arg0 []domain.CdaDepositoAPrazoOutrosAtivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaDepositoAPrazoOutrosAtivosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaDepositoAPrazoOutrosAtivosRepository indicates an expected call of CreateManyCdaDepositoAPrazoOutrosAtivosRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaDepositoAPrazoOutrosAtivosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaDepositoAPrazoOutrosAtivosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaDepositoAPrazoOutrosAtivosRepository), arg0)
}

// CreateManyCdaFiimConfidencialRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaFiimConfidencialRepository(arg0 []domain.CdaFiimConfidencialDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaFiimConfidencialRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaFiimConfidencialRepository indicates an expected call of CreateManyCdaFiimConfidencialRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaFiimConfidencialRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaFiimConfidencialRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaFiimConfidencialRepository), arg0)
}

// CreateManyCdaFiimRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaFiimRepository(arg0 []domain.CdaFiimDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaFiimRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaFiimRepository indicates an expected call of CreateManyCdaFiimRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaFiimRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaFiimRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaFiimRepository), arg0)
}

// CreateManyCdaFundosInvestimentosRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaFundosInvestimentosRepository(arg0 []domain.CdaFundosInvestimentosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaFundosInvestimentosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaFundosInvestimentosRepository indicates an expected call of CreateManyCdaFundosInvestimentosRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaFundosInvestimentosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaFundosInvestimentosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaFundosInvestimentosRepository), arg0)
}

// CreateManyCdaInvestimentosExteriorRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaInvestimentosExteriorRepository(arg0 []domain.CdaInvestimentosExteriorDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaInvestimentosExteriorRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaInvestimentosExteriorRepository indicates an expected call of CreateManyCdaInvestimentosExteriorRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaInvestimentosExteriorRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaInvestimentosExteriorRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaInvestimentosExteriorRepository), arg0)
}

// CreateManyCdaPatrimonioLiquidoRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaPatrimonioLiquidoRepository(arg0 []domain.CdaPatrimonioLiquidoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaPatrimonioLiquidoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaPatrimonioLiquidoRepository indicates an expected call of CreateManyCdaPatrimonioLiquidoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaPatrimonioLiquidoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaPatrimonioLiquidoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaPatrimonioLiquidoRepository), arg0)
}

// CreateManyCdaSelicRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaSelicRepository(arg0 []domain.CdaSelicDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaSelicRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaSelicRepository indicates an expected call of CreateManyCdaSelicRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaSelicRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaSelicRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaSelicRepository), arg0)
}

// CreateManyCdaSwapRepository mocks base method.
func (m *MockFundosPort) CreateManyCdaSwapRepository(arg0 []domain.CdaSwapDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyCdaSwapRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyCdaSwapRepository indicates an expected call of CreateManyCdaSwapRepository.
func (mr *MockFundosPortMockRecorder) CreateManyCdaSwapRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyCdaSwapRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyCdaSwapRepository), arg0)
}

// CreateManyExtratoRepository mocks base method.
func (m *MockFundosPort) CreateManyExtratoRepository(arg0 []domain.ExtratoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyExtratoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyExtratoRepository indicates an expected call of CreateManyExtratoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyExtratoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyExtratoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyExtratoRepository), arg0)
}

// CreateManyFundosRepository mocks base method.
func (m *MockFundosPort) CreateManyFundosRepository(arg0 []domain.FundosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyFundosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyFundosRepository indicates an expected call of CreateManyFundosRepository.
func (mr *MockFundosPortMockRecorder) CreateManyFundosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyFundosRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyFundosRepository), arg0)
}

// CreateManyInformacaoComplementarCotistaRepository mocks base method.
func (m *MockFundosPort) CreateManyInformacaoComplementarCotistaRepository(arg0 []domain.InformacoesCotistaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyInformacaoComplementarCotistaRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyInformacaoComplementarCotistaRepository indicates an expected call of CreateManyInformacaoComplementarCotistaRepository.
func (mr *MockFundosPortMockRecorder) CreateManyInformacaoComplementarCotistaRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyInformacaoComplementarCotistaRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyInformacaoComplementarCotistaRepository), arg0)
}

// CreateManyInformacaoComplementarDivulgacaoRepository mocks base method.
func (m *MockFundosPort) CreateManyInformacaoComplementarDivulgacaoRepository(arg0 []domain.InformacoesDivulgacaoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyInformacaoComplementarDivulgacaoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyInformacaoComplementarDivulgacaoRepository indicates an expected call of CreateManyInformacaoComplementarDivulgacaoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyInformacaoComplementarDivulgacaoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyInformacaoComplementarDivulgacaoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyInformacaoComplementarDivulgacaoRepository), arg0)
}

// CreateManyInformacaoComplementarFundoRepository mocks base method.
func (m *MockFundosPort) CreateManyInformacaoComplementarFundoRepository(arg0 []domain.InformacoesFundoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyInformacaoComplementarFundoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyInformacaoComplementarFundoRepository indicates an expected call of CreateManyInformacaoComplementarFundoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyInformacaoComplementarFundoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyInformacaoComplementarFundoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyInformacaoComplementarFundoRepository), arg0)
}

// CreateManyInformacaoComplementarServicoPrestadoRepository mocks base method.
func (m *MockFundosPort) CreateManyInformacaoComplementarServicoPrestadoRepository(arg0 []domain.ServicoPrestadoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyInformacaoComplementarServicoPrestadoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyInformacaoComplementarServicoPrestadoRepository indicates an expected call of CreateManyInformacaoComplementarServicoPrestadoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyInformacaoComplementarServicoPrestadoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyInformacaoComplementarServicoPrestadoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyInformacaoComplementarServicoPrestadoRepository), arg0)
}

// CreateManyInformacaoDiariaRepository mocks base method.
func (m *MockFundosPort) CreateManyInformacaoDiariaRepository(arg0 []domain.InformacaoDiariaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyInformacaoDiariaRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyInformacaoDiariaRepository indicates an expected call of CreateManyInformacaoDiariaRepository.
func (mr *MockFundosPortMockRecorder) CreateManyInformacaoDiariaRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyInformacaoDiariaRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyInformacaoDiariaRepository), arg0)
}

// CreateManyLaminaCarteiraRepository mocks base method.
func (m *MockFundosPort) CreateManyLaminaCarteiraRepository(arg0 []domain.LaminaCarteiraDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyLaminaCarteiraRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyLaminaCarteiraRepository indicates an expected call of CreateManyLaminaCarteiraRepository.
func (mr *MockFundosPortMockRecorder) CreateManyLaminaCarteiraRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyLaminaCarteiraRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyLaminaCarteiraRepository), arg0)
}

// CreateManyLaminaRentabilidadeAnoRepository mocks base method.
func (m *MockFundosPort) CreateManyLaminaRentabilidadeAnoRepository(arg0 []domain.LaminaRentabilidadeAnoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyLaminaRentabilidadeAnoRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyLaminaRentabilidadeAnoRepository indicates an expected call of CreateManyLaminaRentabilidadeAnoRepository.
func (mr *MockFundosPortMockRecorder) CreateManyLaminaRentabilidadeAnoRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyLaminaRentabilidadeAnoRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyLaminaRentabilidadeAnoRepository), arg0)
}

// CreateManyLaminaRentabilidadeMesRepository mocks base method.
func (m *MockFundosPort) CreateManyLaminaRentabilidadeMesRepository(arg0 []domain.LaminaRentabilidadeMesDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyLaminaRentabilidadeMesRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyLaminaRentabilidadeMesRepository indicates an expected call of CreateManyLaminaRentabilidadeMesRepository.
func (mr *MockFundosPortMockRecorder) CreateManyLaminaRentabilidadeMesRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyLaminaRentabilidadeMesRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyLaminaRentabilidadeMesRepository), arg0)
}

// CreateManyLaminaRepository mocks base method.
func (m *MockFundosPort) CreateManyLaminaRepository(arg0 []domain.LaminaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyLaminaRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyLaminaRepository indicates an expected call of CreateManyLaminaRepository.
func (mr *MockFundosPortMockRecorder) CreateManyLaminaRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyLaminaRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyLaminaRepository), arg0)
}

// CreateManyPerfilMensalRepository mocks base method.
func (m *MockFundosPort) CreateManyPerfilMensalRepository(arg0 []domain.PerfilMensalDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManyPerfilMensalRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateManyPerfilMensalRepository indicates an expected call of CreateManyPerfilMensalRepository.
func (mr *MockFundosPortMockRecorder) CreateManyPerfilMensalRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManyPerfilMensalRepository", reflect.TypeOf((*MockFundosPort)(nil).CreateManyPerfilMensalRepository), arg0)
}

// UpdateArquivosRepository mocks base method.
func (m *MockFundosPort) UpdateArquivosRepository(arg0 domain.ArquivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArquivosRepository", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// UpdateArquivosRepository indicates an expected call of UpdateArquivosRepository.
func (mr *MockFundosPortMockRecorder) UpdateArquivosRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArquivosRepository", reflect.TypeOf((*MockFundosPort)(nil).UpdateArquivosRepository), arg0)
}

// MockFundosExternoPort is a mock of FundosExternoPort interface.
type MockFundosExternoPort struct {
	ctrl     *gomock.Controller
	recorder *MockFundosExternoPortMockRecorder
}

// MockFundosExternoPortMockRecorder is the mock recorder for MockFundosExternoPort.
type MockFundosExternoPortMockRecorder struct {
	mock *MockFundosExternoPort
}

// NewMockFundosExternoPort creates a new mock instance.
func NewMockFundosExternoPort(ctrl *gomock.Controller) *MockFundosExternoPort {
	mock := &MockFundosExternoPort{ctrl: ctrl}
	mock.recorder = &MockFundosExternoPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFundosExternoPort) EXPECT() *MockFundosExternoPortMockRecorder {
	return m.recorder
}

// DownloadArquivosCVMPort mocks base method.
func (m *MockFundosExternoPort) DownloadArquivosCVMPort(arg0 string, arg1 bool) ([]string, *resterrors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadArquivosCVMPort", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*resterrors.RestErr)
	return ret0, ret1
}

// DownloadArquivosCVMPort indicates an expected call of DownloadArquivosCVMPort.
func (mr *MockFundosExternoPortMockRecorder) DownloadArquivosCVMPort(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArquivosCVMPort", reflect.TypeOf((*MockFundosExternoPort)(nil).DownloadArquivosCVMPort), arg0, arg1)
}

// MockFundosQueuePort is a mock of FundosQueuePort interface.
type MockFundosQueuePort struct {
	ctrl     *gomock.Controller
	recorder *MockFundosQueuePortMockRecorder
}

// MockFundosQueuePortMockRecorder is the mock recorder for MockFundosQueuePort.
type MockFundosQueuePortMockRecorder struct {
	mock *MockFundosQueuePort
}

// NewMockFundosQueuePort creates a new mock instance.
func NewMockFundosQueuePort(ctrl *gomock.Controller) *MockFundosQueuePort {
	mock := &MockFundosQueuePort{ctrl: ctrl}
	mock.recorder = &MockFundosQueuePortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFundosQueuePort) EXPECT() *MockFundosQueuePortMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockFundosQueuePort) Produce(arg0 response.FundosQueueResponse) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockFundosQueuePortMockRecorder) Produce(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockFundosQueuePort)(nil).Produce), arg0)
}

// ProduceLote mocks base method.
func (m *MockFundosQueuePort) ProduceLote(arg0 chan response.FundosQueueResponse, arg1 *sync.WaitGroup) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceLote", arg0, arg1)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// ProduceLote indicates an expected call of ProduceLote.
func (mr *MockFundosQueuePortMockRecorder) ProduceLote(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceLote", reflect.TypeOf((*MockFundosQueuePort)(nil).ProduceLote), arg0, arg1)
}
