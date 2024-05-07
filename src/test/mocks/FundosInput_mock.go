// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/port/input/FundosInput.go
//
// Generated by this command:
//
//	mockgen -source=src/application/port/input/FundosInput.go -destination=src/test/mocks/FundosInput_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "api-fundos-investimentos/application/domain"
	resterrors "api-fundos-investimentos/configuration/resterrors"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFundosDomainService is a mock of FundosDomainService interface.
type MockFundosDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockFundosDomainServiceMockRecorder
}

// MockFundosDomainServiceMockRecorder is the mock recorder for MockFundosDomainService.
type MockFundosDomainServiceMockRecorder struct {
	mock *MockFundosDomainService
}

// NewMockFundosDomainService creates a new mock instance.
func NewMockFundosDomainService(ctrl *gomock.Controller) *MockFundosDomainService {
	mock := &MockFundosDomainService{ctrl: ctrl}
	mock.recorder = &MockFundosDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFundosDomainService) EXPECT() *MockFundosDomainServiceMockRecorder {
	return m.recorder
}

// CreateBalanceteService mocks base method.
func (m *MockFundosDomainService) CreateBalanceteService(arg0 []domain.BalanceteDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBalanceteService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateBalanceteService indicates an expected call of CreateBalanceteService.
func (mr *MockFundosDomainServiceMockRecorder) CreateBalanceteService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBalanceteService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateBalanceteService), arg0)
}

// CreateCdaAgroCreditoPrivadoService mocks base method.
func (m *MockFundosDomainService) CreateCdaAgroCreditoPrivadoService(arg0 []domain.CdaAgroCreditoPrivadoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaAgroCreditoPrivadoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaAgroCreditoPrivadoService indicates an expected call of CreateCdaAgroCreditoPrivadoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaAgroCreditoPrivadoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaAgroCreditoPrivadoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaAgroCreditoPrivadoService), arg0)
}

// CreateCdaConfidencialService mocks base method.
func (m *MockFundosDomainService) CreateCdaConfidencialService(arg0 []domain.CdaConfidencialDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaConfidencialService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaConfidencialService indicates an expected call of CreateCdaConfidencialService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaConfidencialService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaConfidencialService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaConfidencialService), arg0)
}

// CreateCdaDemaisAtivosNaoCodificadosService mocks base method.
func (m *MockFundosDomainService) CreateCdaDemaisAtivosNaoCodificadosService(arg0 []domain.CdaDemaisAtivosNaoCodificadosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaDemaisAtivosNaoCodificadosService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaDemaisAtivosNaoCodificadosService indicates an expected call of CreateCdaDemaisAtivosNaoCodificadosService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaDemaisAtivosNaoCodificadosService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaDemaisAtivosNaoCodificadosService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaDemaisAtivosNaoCodificadosService), arg0)
}

// CreateCdaDemaisAtivosService mocks base method.
func (m *MockFundosDomainService) CreateCdaDemaisAtivosService(arg0 []domain.CdaDemaisAtivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaDemaisAtivosService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaDemaisAtivosService indicates an expected call of CreateCdaDemaisAtivosService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaDemaisAtivosService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaDemaisAtivosService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaDemaisAtivosService), arg0)
}

// CreateCdaDepositoAPrazoOutrosAtivosService mocks base method.
func (m *MockFundosDomainService) CreateCdaDepositoAPrazoOutrosAtivosService(arg0 []domain.CdaDepositoAPrazoOutrosAtivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaDepositoAPrazoOutrosAtivosService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaDepositoAPrazoOutrosAtivosService indicates an expected call of CreateCdaDepositoAPrazoOutrosAtivosService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaDepositoAPrazoOutrosAtivosService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaDepositoAPrazoOutrosAtivosService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaDepositoAPrazoOutrosAtivosService), arg0)
}

// CreateCdaFiimConfidencialidade mocks base method.
func (m *MockFundosDomainService) CreateCdaFiimConfidencialidade(arg0 []domain.CdaFiimConfidencialDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaFiimConfidencialidade", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaFiimConfidencialidade indicates an expected call of CreateCdaFiimConfidencialidade.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaFiimConfidencialidade(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaFiimConfidencialidade", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaFiimConfidencialidade), arg0)
}

// CreateCdaFiimService mocks base method.
func (m *MockFundosDomainService) CreateCdaFiimService(arg0 []domain.CdaFiimDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaFiimService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaFiimService indicates an expected call of CreateCdaFiimService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaFiimService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaFiimService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaFiimService), arg0)
}

// CreateCdaFundosInvestimentosService mocks base method.
func (m *MockFundosDomainService) CreateCdaFundosInvestimentosService(arg0 []domain.CdaFundosInvestimentosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaFundosInvestimentosService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaFundosInvestimentosService indicates an expected call of CreateCdaFundosInvestimentosService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaFundosInvestimentosService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaFundosInvestimentosService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaFundosInvestimentosService), arg0)
}

// CreateCdaInvestimentosExteriorService mocks base method.
func (m *MockFundosDomainService) CreateCdaInvestimentosExteriorService(arg0 []domain.CdaInvestimentosExteriorDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaInvestimentosExteriorService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaInvestimentosExteriorService indicates an expected call of CreateCdaInvestimentosExteriorService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaInvestimentosExteriorService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaInvestimentosExteriorService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaInvestimentosExteriorService), arg0)
}

// CreateCdaPatrominioLiquido mocks base method.
func (m *MockFundosDomainService) CreateCdaPatrominioLiquido(arg0 []domain.CdaPatrimonioLiquidoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaPatrominioLiquido", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaPatrominioLiquido indicates an expected call of CreateCdaPatrominioLiquido.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaPatrominioLiquido(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaPatrominioLiquido", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaPatrominioLiquido), arg0)
}

// CreateCdaSelicService mocks base method.
func (m *MockFundosDomainService) CreateCdaSelicService(arg0 []domain.CdaSelicDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaSelicService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaSelicService indicates an expected call of CreateCdaSelicService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaSelicService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaSelicService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaSelicService), arg0)
}

// CreateCdaSwapService mocks base method.
func (m *MockFundosDomainService) CreateCdaSwapService(arg0 []domain.CdaSwapDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCdaSwapService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateCdaSwapService indicates an expected call of CreateCdaSwapService.
func (mr *MockFundosDomainServiceMockRecorder) CreateCdaSwapService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCdaSwapService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateCdaSwapService), arg0)
}

// CreateExtratoService mocks base method.
func (m *MockFundosDomainService) CreateExtratoService(arg0 []domain.ExtratoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExtratoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateExtratoService indicates an expected call of CreateExtratoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateExtratoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExtratoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateExtratoService), arg0)
}

// CreateFundosService mocks base method.
func (m *MockFundosDomainService) CreateFundosService(arg0 []domain.FundosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFundosService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateFundosService indicates an expected call of CreateFundosService.
func (mr *MockFundosDomainServiceMockRecorder) CreateFundosService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFundosService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateFundosService), arg0)
}

// CreateInformacaoComplementarCotistaService mocks base method.
func (m *MockFundosDomainService) CreateInformacaoComplementarCotistaService(arg0 []domain.InformacoesCotistaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInformacaoComplementarCotistaService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateInformacaoComplementarCotistaService indicates an expected call of CreateInformacaoComplementarCotistaService.
func (mr *MockFundosDomainServiceMockRecorder) CreateInformacaoComplementarCotistaService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInformacaoComplementarCotistaService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateInformacaoComplementarCotistaService), arg0)
}

// CreateInformacaoComplementarDivulgacaoService mocks base method.
func (m *MockFundosDomainService) CreateInformacaoComplementarDivulgacaoService(arg0 []domain.InformacoesDivulgacaoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInformacaoComplementarDivulgacaoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateInformacaoComplementarDivulgacaoService indicates an expected call of CreateInformacaoComplementarDivulgacaoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateInformacaoComplementarDivulgacaoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInformacaoComplementarDivulgacaoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateInformacaoComplementarDivulgacaoService), arg0)
}

// CreateInformacaoComplementarFundoService mocks base method.
func (m *MockFundosDomainService) CreateInformacaoComplementarFundoService(arg0 []domain.InformacoesFundoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInformacaoComplementarFundoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateInformacaoComplementarFundoService indicates an expected call of CreateInformacaoComplementarFundoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateInformacaoComplementarFundoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInformacaoComplementarFundoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateInformacaoComplementarFundoService), arg0)
}

// CreateInformacaoComplementarServicoPrestadoService mocks base method.
func (m *MockFundosDomainService) CreateInformacaoComplementarServicoPrestadoService(arg0 []domain.ServicoPrestadoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInformacaoComplementarServicoPrestadoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateInformacaoComplementarServicoPrestadoService indicates an expected call of CreateInformacaoComplementarServicoPrestadoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateInformacaoComplementarServicoPrestadoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInformacaoComplementarServicoPrestadoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateInformacaoComplementarServicoPrestadoService), arg0)
}

// CreateInformacaoDiariaService mocks base method.
func (m *MockFundosDomainService) CreateInformacaoDiariaService(arg0 []domain.InformacaoDiariaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInformacaoDiariaService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateInformacaoDiariaService indicates an expected call of CreateInformacaoDiariaService.
func (mr *MockFundosDomainServiceMockRecorder) CreateInformacaoDiariaService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInformacaoDiariaService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateInformacaoDiariaService), arg0)
}

// CreateLaminaCarteiraService mocks base method.
func (m *MockFundosDomainService) CreateLaminaCarteiraService(arg0 []domain.LaminaCarteiraDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaminaCarteiraService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateLaminaCarteiraService indicates an expected call of CreateLaminaCarteiraService.
func (mr *MockFundosDomainServiceMockRecorder) CreateLaminaCarteiraService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaminaCarteiraService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateLaminaCarteiraService), arg0)
}

// CreateLaminaRentabilidadeAnoService mocks base method.
func (m *MockFundosDomainService) CreateLaminaRentabilidadeAnoService(arg0 []domain.LaminaRentabilidadeAnoDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaminaRentabilidadeAnoService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateLaminaRentabilidadeAnoService indicates an expected call of CreateLaminaRentabilidadeAnoService.
func (mr *MockFundosDomainServiceMockRecorder) CreateLaminaRentabilidadeAnoService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaminaRentabilidadeAnoService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateLaminaRentabilidadeAnoService), arg0)
}

// CreateLaminaRentabilidadeMesService mocks base method.
func (m *MockFundosDomainService) CreateLaminaRentabilidadeMesService(arg0 []domain.LaminaRentabilidadeMesDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaminaRentabilidadeMesService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateLaminaRentabilidadeMesService indicates an expected call of CreateLaminaRentabilidadeMesService.
func (mr *MockFundosDomainServiceMockRecorder) CreateLaminaRentabilidadeMesService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaminaRentabilidadeMesService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateLaminaRentabilidadeMesService), arg0)
}

// CreateLaminaService mocks base method.
func (m *MockFundosDomainService) CreateLaminaService(arg0 []domain.LaminaDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaminaService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreateLaminaService indicates an expected call of CreateLaminaService.
func (mr *MockFundosDomainServiceMockRecorder) CreateLaminaService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaminaService", reflect.TypeOf((*MockFundosDomainService)(nil).CreateLaminaService), arg0)
}

// CreatePerfilMensalService mocks base method.
func (m *MockFundosDomainService) CreatePerfilMensalService(arg0 []domain.PerfilMensalDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerfilMensalService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// CreatePerfilMensalService indicates an expected call of CreatePerfilMensalService.
func (mr *MockFundosDomainServiceMockRecorder) CreatePerfilMensalService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerfilMensalService", reflect.TypeOf((*MockFundosDomainService)(nil).CreatePerfilMensalService), arg0)
}

// DownloadArquivosCVMService mocks base method.
func (m *MockFundosDomainService) DownloadArquivosCVMService(arg0 domain.ArquivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadArquivosCVMService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// DownloadArquivosCVMService indicates an expected call of DownloadArquivosCVMService.
func (mr *MockFundosDomainServiceMockRecorder) DownloadArquivosCVMService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArquivosCVMService", reflect.TypeOf((*MockFundosDomainService)(nil).DownloadArquivosCVMService), arg0)
}

// ProcessarArquivosCVMService mocks base method.
func (m *MockFundosDomainService) ProcessarArquivosCVMService(arg0 domain.ArquivosDomain) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessarArquivosCVMService", arg0)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// ProcessarArquivosCVMService indicates an expected call of ProcessarArquivosCVMService.
func (mr *MockFundosDomainServiceMockRecorder) ProcessarArquivosCVMService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessarArquivosCVMService", reflect.TypeOf((*MockFundosDomainService)(nil).ProcessarArquivosCVMService), arg0)
}

// QueueFundosSincronizarService mocks base method.
func (m *MockFundosDomainService) QueueFundosSincronizarService(arg0 string, arg1 bool) *resterrors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueFundosSincronizarService", arg0, arg1)
	ret0, _ := ret[0].(*resterrors.RestErr)
	return ret0
}

// QueueFundosSincronizarService indicates an expected call of QueueFundosSincronizarService.
func (mr *MockFundosDomainServiceMockRecorder) QueueFundosSincronizarService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueFundosSincronizarService", reflect.TypeOf((*MockFundosDomainService)(nil).QueueFundosSincronizarService), arg0, arg1)
}