package bradesco80

import (
	"time"

	"github.com/libercapital/document-translator-go/internal/writer"
	"github.com/shopspring/decimal"
)

type ContractSettlementHeader struct {
	TipoRegistro  int       `translator:"part:0..0"`
	DataMovimento time.Time `translator:"part:1..8;timeParse:02012006"`
	Nome          string    `translator:"part:9..48"`
	EmpresaOrigem int       `translator:"part:49..55"`
	Filler        string    `translator:"part:56..79"`
}

func (c ContractSettlementHeader) String() (string, error) {
	return writer.Marshal(c, 80)
}

type ContractSettlementRegister struct {
	TipoRegistro                    string          `translator:"part:0..0"`
	SistemaOrigem                   string          `translator:"part:1..3"`
	CodigoConvenio                  string          `translator:"part:4..12"`
	ContratoOrigem                  string          `translator:"part:13..21"`
	TipoPagamento                   string          `translator:"part:22..22"`
	DataVencimentoParcela           time.Time       `translator:"part:23..30;timeParse:02012006"`
	Produto                         string          `translator:"part:31..33"`
	Familia                         string          `translator:"part:34..34"`
	Contrato                        string          `translator:"part:35..43"`
	ValorPagameto                   decimal.Decimal `translator:"part:44..60;precision:2"`
	NumeroParcela                   string          `translator:"part:61..63"`
	ADebitarNaConta                 string          `translator:"part:64..64"`
	Filler                          string          `translator:"part:65..78"`
	IdentificadorRecompraLiquidacao string          `translator:"part:79..79"`
}

func (c ContractSettlementRegister) String() (string, error) {
	return writer.Marshal(c, 80)
}

type ContractSettlementTrailer struct {
	TipoRegistro        int    `translator:"part:0..0"`
	QuantidadeRegistros int    `translator:"part:1..6"`
	Filler              string `translator:"part:7..79"`
}

func (c ContractSettlementTrailer) String() (string, error) {
	return writer.Marshal(c, 80)
}
