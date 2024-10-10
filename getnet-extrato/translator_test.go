package getnet_extrato

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHeader(t *testing.T) {
	cnabLine := "01609202407132916092024CEADM1001013903        10440482000154GETNET S.A.         000002516GSSANT. V.10.1 400 BYTES                                                                                                                                                                                                                                                                                                "

	parsed, err := ParseHeader(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "0", parsed.TipoRegistro)
	assert.Equal(t, "16092024", parsed.DataCriacaoArquivo.Format("02012006"))
	assert.Equal(t, "071329", parsed.HoraCriacaoArquivo)
	assert.Equal(t, "16092024", parsed.DataMovimento.Format("02012006"))
	assert.Equal(t, "CEADM100", parsed.VersaoArquivo)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimento)
	assert.Equal(t, "10440482000154", parsed.CNPJAdquirente)
	assert.Equal(t, "GETNET S.A.", parsed.NomeAdquirente)
	assert.Equal(t, "000002516", parsed.NumeroSequencial)
	assert.Equal(t, "GS", parsed.CodigoAdquirente)
	assert.Equal(t, "SANT. V.10.1 400 BYTES", parsed.VersaoLayout)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseResumoTransacional(t *testing.T) {
	cnabLine := "11013903        SRMAN051190405030920241609202403300368900000000000000000001000000000000000014890000000014890000000000000000000000000000000000000000000014890000000000000LQ01011013903        00000000000000000000000000000000000000000000000005000511904050001000000000000000000000000000986 -CC000000000001300287672024040902510656452007001                                                                   "

	parsed, err := ParseResumoTransacional(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "1", parsed.TipoRegistro)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoComercial)
	assert.Equal(t, "SR", parsed.CodigoProduto)
	assert.Equal(t, "MAN", parsed.FormaCaptura)
	assert.Equal(t, "051190405", parsed.NumeroRV)
	assert.Equal(t, "03092024", parsed.DataRV.Format("02012006"))
	assert.Equal(t, "16092024", parsed.DataPagamentoRV.Format("02012006"))
	assert.Equal(t, "033", parsed.Banco)
	assert.Equal(t, "003689", parsed.Agencia)
	assert.Equal(t, "00000000000", parsed.Zeros1)
	assert.Equal(t, 1, parsed.NumeroCVsAceitos)
	assert.Equal(t, 0, parsed.NumeroCVsRejeitados)
	assert.Equal(t, 0.0, parsed.ValorBruto)
	assert.Equal(t, 0.0, parsed.ValorLiquido)
	assert.Equal(t, 0.00, parsed.ValorTarifa)
	assert.Equal(t, 0.00, parsed.ValorTaxaDesconto)
	assert.Equal(t, 0.00, parsed.ValorRejeitado)
	assert.Equal(t, 0.00, parsed.ValorCredito)
	assert.Equal(t, "000000000000", parsed.Zeros2)
	assert.Equal(t, "LQ", parsed.IndicadorTipoPagamento)
	assert.Equal(t, 1, parsed.NumeroParcelaRV)
	assert.Equal(t, 1, parsed.QuantidadeParcelasRV)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoCentralizador)
	assert.Equal(t, "000000000000000", parsed.Zeros3)
	assert.Equal(t, "01010001", parsed.DataVencimentoOriginal.Format("02012006"))
	assert.Equal(t, "000000000000", parsed.Zeros4)
	assert.Equal(t, "000000000000", parsed.Zeros5)
	assert.Equal(t, "005000511904050001", parsed.NumeroControleOperacao)
	assert.Equal(t, 0.00, parsed.ValorLiquidoCobranca)
	assert.Equal(t, "000000000000000", parsed.Zeros6)
	assert.Equal(t, "986", parsed.Moeda)
	assert.Equal(t, "", parsed.IdentificadorBaixaCobranca)
	assert.Equal(t, "CC", parsed.TipoContaPagamento)
	assert.Equal(t, "00000000000130028767", parsed.ContaCorrente)
	assert.Equal(t, "2024040902510656452007001", parsed.ChaveUR)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseAnaliticoTransacional(t *testing.T) {
	cnabLine := "21013903        00052459700000200009816082024104422650921******1796   0000024142390000000000000000000000000601000000402374160920240000385282TEFC1013903        T4502939986N+   000000008775                                                                                                                                                                                                                     "

	parsed, err := ParseAnaliticoTransacional(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "2", parsed.TipoRegistro)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoComercial)
	assert.Equal(t, "000524597", parsed.NumeroRV)
	assert.Equal(t, "000002000098", parsed.NSUAdquirente)
	assert.Equal(t, "16082024", parsed.DataTransacao.Format("02012006"))
	assert.Equal(t, "104422", parsed.HoraTransacao)
	assert.Equal(t, "650921******1796", parsed.NumeroCartao)
	assert.Equal(t, 0.00, parsed.ValorTransacao)
	assert.Equal(t, 0.00, parsed.ValorSaque)
	assert.Equal(t, 0.00, parsed.ValorTaxaEmbarque)
	assert.Equal(t, 6, parsed.NumeroParcelas)
	assert.Equal(t, 1, parsed.NumeroParcelaRelacaoCV)
	assert.Equal(t, 0.00, parsed.ValorParcela)
	assert.Equal(t, "16092024", parsed.DataPagamento.Format("02012006"))
	assert.Equal(t, "0000385282", parsed.CodigoAutorizacao)
	assert.Equal(t, "TEF", parsed.FormaCaptura)
	assert.Equal(t, "C", parsed.StatusTransacao)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoCentralizador)
	assert.Equal(t, "T4502939", parsed.CodigoTerminal)
	assert.Equal(t, "986", parsed.Moeda)
	assert.Equal(t, "N", parsed.OrigemEmissorCartao)
	assert.Equal(t, "+", parsed.SinalTransacao)
	assert.Equal(t, "", parsed.CarteiraDigital)
	assert.Equal(t, 0.00, parsed.ValorComissaoVenda)
	assert.Equal(t, "", parsed.IdentificadorTipoProximoConteudo)
	assert.Equal(t, "", parsed.ConteudoDinamico)
	assert.Equal(t, "", parsed.IdentificadorTipoProximoConteudo2)
	assert.Equal(t, "", parsed.ConteudoDinamico2)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseAjusteFinanceiro(t *testing.T) {
	cnabLine := "31013903        0511904050309202416092024240903003689439      -00000001489002000000000                  00000000000000000000000000000LQ                98600000000000003Aluguel-                                                                                                                                                                                                                                "

	parsed, err := ParseAjusteFinanceiro(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "3", parsed.TipoRegistro)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoComercial)
	assert.Equal(t, "051190405", parsed.NumeroRV)
	assert.Equal(t, "03092024", parsed.DataRV.Format("02012006"))
	assert.Equal(t, "16092024", parsed.DataPagamentoRV.Format("02012006"))
	assert.Equal(t, "240903003689439", parsed.IdentificadorAjuste)
	assert.Equal(t, "", parsed.Brancos)
	assert.Equal(t, "-", parsed.SinalValorAjuste)
	assert.Equal(t, 0.00, parsed.ValorAjuste)
	assert.Equal(t, "02", parsed.MotivoAjuste)
	assert.Equal(t, "01010001", parsed.DataCarta.Format("02012006"))
	assert.Equal(t, "0", parsed.NumeroCartao)
	assert.Equal(t, "000000000", parsed.NumeroRVOriginal)
	assert.Equal(t, "000000000000", parsed.NSUAdquirente)
	assert.Equal(t, "01010001", parsed.DataTransacaoOriginal.Format("02012006"))
	assert.Equal(t, "LQ", parsed.IndicadorTipoPagamento)
	assert.Equal(t, "", parsed.NumeroTerminalOriginal)
	assert.Equal(t, "", parsed.DataPagamentoOriginal)
	assert.Equal(t, "986", parsed.Moeda)
	assert.Equal(t, 0.00, parsed.ValorComissaoVendaCancelada)
	assert.Equal(t, "03", parsed.IdentificadorProximoConteudo)
	assert.Equal(t, "Aluguel-", parsed.ConteudoDinamico)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseResumoFinanceiro(t *testing.T) {
	cnabLine := "51013903        120920241609202400000000000000000000PG00000000000000000036381900000000000000000036381900000000000CC03300368900000000000130028767   LIF 000000000000000000110656452007001  000000000                    000000001013903                         EC2024160900810656452007001                                                                                                                      "

	parsed, err := ParseResumoFinanceiro(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "5", parsed.TipoRegistro)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoComercial)
	assert.Equal(t, "12092024", parsed.DataOperacao.Format("02012006"))
	assert.Equal(t, "16092024", parsed.DataCreditoOperacao.Format("02012006"))
	assert.Equal(t, "00000000000000000000", parsed.NumeroOperacao)
	assert.Equal(t, "PG", parsed.TipoOperacao)
	assert.Equal(t, "000000000000", parsed.Zeros1)
	assert.Equal(t, 0.00, parsed.ValorBrutoOperacaoAdquirencia)
	assert.Equal(t, 0.00, parsed.ValorCustoOperacao)
	assert.Equal(t, 0.00, parsed.ValorLiquidoOperacao)
	assert.Equal(t, 0.0000000, parsed.TaxaMensalOperacao)
	assert.Equal(t, "CC", parsed.TipoContaEstabelecimento)
	assert.Equal(t, "033", parsed.Banco)
	assert.Equal(t, "003689", parsed.Agencia)
	assert.Equal(t, "00000000000130028767", parsed.ContaCorrente)
	assert.Equal(t, "", parsed.CanalOperacao)
	assert.Equal(t, "L", parsed.TipoMovimento)
	assert.Equal(t, "IF", parsed.TipoParticipante)
	assert.Equal(t, "000000000000000000", parsed.Zeros2)
	assert.Equal(t, "1", parsed.TipoDocumentoParticipante)
	assert.Equal(t, "10656452007001", parsed.CNPJCPFParticipante)
	assert.Equal(t, "", parsed.TipoContaEstabelecimentoParticipante)
	assert.Equal(t, "000", parsed.BancoDomicilioBancarioParticipante)
	assert.Equal(t, "000000", parsed.AgenciaDomicilioBancarioParticipante)
	assert.Equal(t, "", parsed.ContaDomicilioBancarioParticipante)
	assert.Equal(t, "000000001013903", parsed.CodigoEstabelecimentoCentralizador)
	assert.Equal(t, "", parsed.RazaoSocialParticipante)
	assert.Equal(t, "EC", parsed.CodigoArranjoPagamento)
	assert.Equal(t, "2024160900810656452007001", parsed.ChaveUR)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseDetalheFinanceiro(t *testing.T) {
	cnabLine := "61013903        1609202400000000000000000000LQ000000000000000000SR05082024000000000000000000000000000000000000000000014890  00000000000000000000000000000DIF 000000000000000000 00000000000000  000000000                    000000001013903                         2024050802510656452007001                                                                                                                  "

	parsed, err := ParseDetalheFinanceiro(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "6", parsed.TipoRegistro)
	assert.Equal(t, "1013903", parsed.CodigoEstabelecimentoComercial)
	assert.Equal(t, "16092024", parsed.DataOperacao.Format("02012006"))
	assert.Equal(t, "00000000000000000000", parsed.NumeroOperacao)
	assert.Equal(t, "LQ", parsed.TipoOperacao)
	assert.Equal(t, "000000000000000000", parsed.Zeros1)
	assert.Equal(t, "SR", parsed.CodigoProduto)
	assert.Equal(t, "05082024", parsed.DataVencimentoUR.Format("02012006"))
	assert.Equal(t, "000000000000", parsed.Zeros2)
	assert.Equal(t, 0.00, parsed.ValorLiquidoURNegociada)
	assert.Equal(t, 0.00, parsed.ValorCustoURNegociada)
	assert.Equal(t, 0.00, parsed.ValorBrutoURNegociada)
	assert.Equal(t, "", parsed.TipoContaEstabelecimento)
	assert.Equal(t, "000", parsed.Zeros3)
	assert.Equal(t, "000000", parsed.Zeros4)
	assert.Equal(t, "00000000000000000000", parsed.Zeros5)
	assert.Equal(t, "D", parsed.TipoMovimento)
	assert.Equal(t, "IF", parsed.TipoParticipante)
	assert.Equal(t, "000000000000000000", parsed.Zeros6)
	assert.Equal(t, "", parsed.TipoDocumentoParticipante)
	assert.Equal(t, "00000000000000", parsed.CNPJCPFParticipante)
	assert.Equal(t, "", parsed.Espaco1)
	assert.Equal(t, "000", parsed.Zeros7)
	assert.Equal(t, "000000", parsed.Zeros8)
	assert.Equal(t, "", parsed.Espaco2)
	assert.Equal(t, "000000001013903", parsed.CodigoEstabelecimentoCentralizador)
	assert.Equal(t, "", parsed.RazaoSocialParticipante)
	assert.Equal(t, "2024050802510656452007001", parsed.ChaveUR)
	assert.Equal(t, "", parsed.Reservado)
}

func TestParseTrailer(t *testing.T) {
	cnabLine := "9000000047                                                                                                                                                                                                                                                                                                                                                                                                      "

	parsed, err := ParseTrailer(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "9", parsed.TipoRegistro)
	assert.Equal(t, 47, parsed.QuantidadeRegistros)
	assert.Equal(t, "", parsed.Reservado)
}

func TestKindT0(t *testing.T) {
	cnabLine := "01609202407132916092024CEADM1001013903        10440482000154GETNET S.A.         000002516GSSANT. V.10.1 400 BYTES                                                                                                                                                                                                                                                                                                "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroHeader, kind)
}

func TestKindT1(t *testing.T) {
	cnabLine := "11013903        SRMAN051190405030920241609202403300368900000000000000000001000000000000000014890000000014890000000000000000000000000000000000000000000014890000000000000LQ01011013903        00000000000000000000000000000000000000000000000005000511904050001000000000000000000000000000986 -CC000000000001300287672024040902510656452007001                                                                   "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroResumoTransacional, kind)
}

func TestKindT2(t *testing.T) {
	cnabLine := "21013903        00052459700000200009816082024104422650921******1796   0000024142390000000000000000000000000601000000402374160920240000385282TEFC1013903        T4502939986N+   000000008775                                                                                                                                                                                                                     "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroAnaliticoTransacional, kind)
}

func TestKindT3(t *testing.T) {
	cnabLine := "31013903        0511904050309202416092024240903003689439      -00000001489002000000000                  00000000000000000000000000000LQ                98600000000000003Aluguel-                                                                                                                                                                                                                                "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroAjusteFinanceiro, kind)
}

func TestKindT5(t *testing.T) {
	cnabLine := "51013903        120920241609202400000000000000000000PG00000000000000000036381900000000000000000036381900000000000CC033003689000000000000130028767   LIF 000000000000000000110656452007001  000000000                    000000001013903                         EC2024160900810656452007001                                                                                                                      "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroResumoFinanceiro, kind)
}

func TestKindT6(t *testing.T) {
	cnabLine := "61013903        1609202400000000000000000000LQ000000000000000000SR05082024000000000000000000000000000000000000000000014890  00000000000000000000000000000DIF 000000000000000000 00000000000000  000000000                    000000001013903                         2024050802510656452007001                                                                                                                  "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroDetalheFinanceiro, kind)
}

func TestKindT9(t *testing.T) {
	cnabLine := "9000000047                                                                                                                                                                                                                                                                                                                                                                                                      "

	kind, err := Kind(cnabLine)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, TipoRegistroTrailer, kind)
}

func TestKindInvalid(t *testing.T) {
	cnabLine := "X1013903        1609202400000000000000000000LQ000000000000000000SR05082024000000000000000000000000000000000000000000014890  00000000000000000000000000000DIF 000000000000000000 00000000000000  000000000                    000000001013903                         2024050802510656452007001                                                                                                                  "

	kind, err := Kind(cnabLine)
	assert.EqualError(t, err, "invalid register type")
	assert.Equal(t, RegisterType(""), kind)
}
