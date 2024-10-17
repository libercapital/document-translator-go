package bradesco80

import (
	"time"

	"github.com/libercapital/document-translator-go/internal/writer"
	"github.com/shopspring/decimal"
)

type ContractSettlementHeader struct {
	TipoRegistro  int       `translator:"part:0..0"`
	DataMovimento time.Time `translator:"part:1..9;timeParse:02012006"`
	Nome          string    `translator:"part:10..49"`
	EmpresaOrigem int       `translator:"part:50..57"`
	Filler        string    `translator:"part:58..79"`
}

func (c ContractSettlementHeader) String() (string, error) {
	return writer.Marshal(c, 80)
}

type ContractSettlementRegister struct {
	TipoRegistro                    int             `translator:"part:0..0"`
	SistemaOrigem                   int             `translator:"part:1..4"`
	CodigoConvenio                  int             `translator:"part:5..13"`
	ContratoOrigem                  int             `translator:"part:14..22"`
	TipoPagamento                   int             `translator:"part:23..23"`
	DataVencimentoParcela           time.Time       `translator:"part:24..31;timeParse:02012006"`
	Produto                         string          `translator:"part:32..34"`
	Familia                         string          `translator:"part:35..35"`
	Contrato                        string          `translator:"part:36..44"`
	ValorPagameto                   decimal.Decimal `translator:"part:45..61;precision:2"`
	NumeroParcela                   int             `translator:"part:62..64"`
	ADebitarNaConta                 int             `translator:"part:65..65"`
	Filler                          int             `translator:"part:66..78"`
	IdentificadorRecompraLiquidacao string          `translator:"part:79..79"`
}

func (c ContractSettlementRegister) String() (string, error) {
	return writer.Marshal(c, 80)
}

type ContractSettlementTrailer struct {
	TipoRegistro        int    `translator:"part:0..0"`
	QuantidadeRegistros int    `translator:"part:1..7"`
	Filler              string `translator:"part:8..79"`
}

func (c ContractSettlementTrailer) String() (string, error) {
	return writer.Marshal(c, 80)
}
