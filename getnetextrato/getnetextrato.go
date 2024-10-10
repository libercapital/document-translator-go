package getnetextrato

import (
	"time"

	"github.com/libercapital/document-translator-go/internal/writer"
	"github.com/shopspring/decimal"
)

type Header struct {
	TipoRegistro          string    `translator:"part:0..0"`                      // 001..001 A(001)
	DataCriacaoArquivo    time.Time `translator:"part:1..8;timeParse:02012006"`   // 002..009 N(008)
	HoraCriacaoArquivo    string    `translator:"part:9..14"`                     // 010..015 N(006)
	DataMovimento         time.Time `translator:"part:15..22;timeParse:02012006"` // 016..023 N(008)
	VersaoArquivo         string    `translator:"part:23..30"`                    // 024..031 A(008)
	CodigoEstabelecimento string    `translator:"part:31..45"`                    // 032..046 A(015)
	CNPJAdquirente        string    `translator:"part:46..59"`                    // 047..060 N(014)
	NomeAdquirente        string    `translator:"part:60..79"`                    // 061..080 A(020)
	NumeroSequencial      string    `translator:"part:80..88"`                    // 081..089 N(009)
	CodigoAdquirente      string    `translator:"part:89..90"`                    // 090..091 A(002)
	VersaoLayout          string    `translator:"part:91..115"`                   // 092..116 A(025)
	Reservado             string    `translator:"part:116..399"`                  // 117..400 A(284)
}

func (i Header) String() (string, error) {
	return writer.Marshal(i, 400)
}

type ResumoTransacional struct {
	TipoRegistro                       string          `translator:"part:0..0"`                        // 001..001 A(001)
	CodigoEstabelecimentoComercial     string          `translator:"part:1..15"`                       // 002..016 A(015)
	CodigoProduto                      string          `translator:"part:16..17"`                      // 017..018 A(002)
	FormaCaptura                       string          `translator:"part:18..20"`                      // 019..021 A(003)
	NumeroRV                           string          `translator:"part:21..29"`                      // 022..030 N(009)
	DataRV                             time.Time       `translator:"part:30..37;timeParse:02012006"`   // 031..038 N(008)
	DataPagamentoRV                    time.Time       `translator:"part:38..45;timeParse:02012006"`   // 039..046 N(008)
	Banco                              string          `translator:"part:46..48"`                      // 047..049 N(003)
	Agencia                            string          `translator:"part:49..54"`                      // 050..055 N(006)
	Zeros1                             string          `translator:"part:55..65"`                      // 056..066 N(011)
	NumeroCVsAceitos                   int             `translator:"part:66..74"`                      // 067..075 N(009)
	NumeroCVsRejeitados                int             `translator:"part:75..83"`                      // 076..084 N(009)
	ValorBruto                         decimal.Decimal `translator:"part:84..95;precision:2"`          // 085..096 N(012)
	ValorLiquido                       decimal.Decimal `translator:"part:96..107;precision:2"`         // 097..108 N(012)
	ValorTarifa                        decimal.Decimal `translator:"part:108..119;precision:2"`        // 109..120 N(012)
	ValorTaxaDesconto                  decimal.Decimal `translator:"part:120..131;precision:2"`        // 121..132 N(012)
	ValorRejeitado                     decimal.Decimal `translator:"part:132..143;precision:2"`        // 133..144 N(012)
	ValorCredito                       decimal.Decimal `translator:"part:144..155;precision:2"`        // 145..156 N(012)
	Zeros2                             string          `translator:"part:156..167"`                    // 157..168 N(012)
	IndicadorTipoPagamento             string          `translator:"part:168..169"`                    // 169..170 A(002)
	NumeroParcelaRV                    int             `translator:"part:170..171"`                    // 171..172 N(002)
	QuantidadeParcelasRV               int             `translator:"part:172..173"`                    // 173..174 N(002)
	CodigoEstabelecimentoCentralizador string          `translator:"part:174..188"`                    // 175..189 A(015)
	Zeros3                             string          `translator:"part:189..203"`                    // 190..204 N(015)
	DataVencimentoOriginal             time.Time       `translator:"part:204..211;timeParse:02012006"` // 205..212 N(008)
	Zeros4                             string          `translator:"part:212..223"`                    // 213..224 N(012)
	Zeros5                             string          `translator:"part:224..235"`                    // 225..236 N(012)
	NumeroControleOperacao             string          `translator:"part:236..253"`                    // 237..254 N(018)
	ValorLiquidoCobranca               decimal.Decimal `translator:"part:254..265;precision:2"`        // 255..266 N(012)
	Zeros6                             string          `translator:"part:266..280"`                    // 267..281 N(015)
	Moeda                              string          `translator:"part:281..283"`                    // 282..284 N(003)
	IdentificadorBaixaCobranca         string          `translator:"part:284..284"`                    // 285..285 A(001)
	SinalTransacao                     string          `translator:"part:285..285"`                    // 286..286 A(001)
	TipoContaPagamento                 string          `translator:"part:286..287"`                    // 287..288 A(002)
	ContaCorrente                      string          `translator:"part:288..307"`                    // 289..308 N(020)
	ChaveUR                            string          `translator:"part:308..332"`                    // 309..333 A(025)
	Reservado                          string          `translator:"part:358..399"`                    // 360..400 A(041)
}

func (i ResumoTransacional) String() (string, error) {
	return writer.Marshal(i, 400)
}

type AnaliticoTransacional struct {
	TipoRegistro                       string          `translator:"part:0..0"`                        // 001..001 A(001)
	CodigoEstabelecimentoComercial     string          `translator:"part:1..15"`                       // 002..016 A(015)
	NumeroRV                           string          `translator:"part:16..24"`                      // 017..025 N(009)
	NSUAdquirente                      string          `translator:"part:25..36"`                      // 026..037 N(012)
	DataTransacao                      time.Time       `translator:"part:37..44;timeParse:02012006"`   // 038..045 N(008)
	HoraTransacao                      string          `translator:"part:45..50"`                      // 046..051 N(006)
	NumeroCartao                       string          `translator:"part:51..69"`                      // 052..070 A(019)
	ValorTransacao                     decimal.Decimal `translator:"part:70..81;precision:2"`          // 071..082 N(012)
	ValorSaque                         decimal.Decimal `translator:"part:82..93;precision:2"`          // 083..094 N(012)
	ValorTaxaEmbarque                  decimal.Decimal `translator:"part:94..105;precision:2"`         // 095..106 N(012)
	NumeroParcelas                     int             `translator:"part:106..107"`                    // 107..108 N(002)
	NumeroParcelaRelacaoCV             int             `translator:"part:108..109"`                    // 109..110 N(002)
	ValorParcela                       decimal.Decimal `translator:"part:110..121;precision:2"`        // 111..122 N(012)
	DataPagamento                      time.Time       `translator:"part:122..129;timeParse:02012006"` // 123..130 N(008)
	CodigoAutorizacao                  string          `translator:"part:130..139"`                    // 131..140 A(010)
	FormaCaptura                       string          `translator:"part:140..142"`                    // 141..143 A(003)
	StatusTransacao                    string          `translator:"part:143..143"`                    // 144..144 A(001)
	CodigoEstabelecimentoCentralizador string          `translator:"part:144..158"`                    // 145..159 A(015)
	CodigoTerminal                     string          `translator:"part:159..166"`                    // 160..167 A(008)
	Moeda                              string          `translator:"part:167..169"`                    // 168..170 N(003)
	OrigemEmissorCartao                string          `translator:"part:170..170"`                    // 171..171 A(001)
	SinalTransacao                     string          `translator:"part:171..171"`                    // 172..172 A(001)
	CarteiraDigital                    string          `translator:"part:172..174"`                    // 173..175 A(003)
	ValorComissaoVenda                 decimal.Decimal `translator:"part:175..186;precision:2"`        // 176..187 N(012)
	IdentificadorTipoProximoConteudo   string          `translator:"part:187..188"`                    // 188..189 A(002)
	ConteudoDinamico                   string          `translator:"part:189..306"`                    // 190..307 A(118)
	IdentificadorTipoProximoConteudo2  string          `translator:"part:307..308"`                    // 308..309 A(002)
	ConteudoDinamico2                  string          `translator:"part:309..358"`                    // 310..359 A(050)
	Reservado                          string          `translator:"part:359..399"`                    // 360..400 A(041)
}

func (i AnaliticoTransacional) String() (string, error) {
	return writer.Marshal(i, 400)
}

type AjusteFinanceiro struct {
	TipoRegistro                   string          `translator:"part:0..0"`                        // 001..001 A(001)
	CodigoEstabelecimentoComercial string          `translator:"part:1..15"`                       // 002..016 A(015)
	NumeroRV                       string          `translator:"part:16..24"`                      // 017..025 N(009)
	DataRV                         time.Time       `translator:"part:25..32;timeParse:02012006"`   // 026..033 N(008)
	DataPagamentoRV                time.Time       `translator:"part:33..40;timeParse:02012006"`   // 034..041 N(008)
	IdentificadorAjuste            string          `translator:"part:41..60"`                      // 042..061 N(020)
	Brancos                        string          `translator:"part:61..61"`                      // 062..062 A(001)
	SinalValorAjuste               string          `translator:"part:62..62"`                      // 063..063 A(001)
	ValorAjuste                    decimal.Decimal `translator:"part:63..74;precision:2"`          // 064..075 N(012)
	MotivoAjuste                   string          `translator:"part:75..76"`                      // 076..077 A(002)
	DataCarta                      time.Time       `translator:"part:77..84;timeParse:02012006"`   // 078..085 N(008)
	NumeroCartao                   string          `translator:"part:85..103"`                     // 086..104 A(019)
	NumeroRVOriginal               string          `translator:"part:104..112"`                    // 105..113 N(009)
	NSUAdquirente                  string          `translator:"part:113..124"`                    // 114..125 N(012)
	DataTransacaoOriginal          time.Time       `translator:"part:125..132;timeParse:02012006"` // 126..133 N(008)
	IndicadorTipoPagamento         string          `translator:"part:133..134"`                    // 134..135 A(002)
	NumeroTerminalOriginal         string          `translator:"part:135..142"`                    // 136..143 A(008)
	DataPagamentoOriginal          string          `translator:"part:143..150"`                    // 144..151 N(008)
	Moeda                          string          `translator:"part:151..153"`                    // 152..154 N(003)
	ValorComissaoVendaCancelada    decimal.Decimal `translator:"part:154..165;precision:2"`        // 155..166 N(012)
	IdentificadorProximoConteudo   string          `translator:"part:166..167"`                    // 167..168 A(002)
	ConteudoDinamico               string          `translator:"part:168..285"`                    // 169..286 A(118)
	Reservado                      string          `translator:"part:286..399"`                    // 287..400 A(114)
}

func (i AjusteFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type ResumoFinanceiro struct {
	TipoRegistro                         string          `translator:"part:0..0"`                      // 001..001 A(001)
	CodigoEstabelecimentoComercial       string          `translator:"part:1..15"`                     // 002..016 A(015)
	DataOperacao                         time.Time       `translator:"part:16..23;timeParse:02012006"` // 017..024 N(008)
	DataCreditoOperacao                  time.Time       `translator:"part:24..31;timeParse:02012006"` // 025..032 N(008)
	NumeroOperacao                       string          `translator:"part:32..51"`                    // 033..052 A(020)
	TipoOperacao                         string          `translator:"part:52..53"`                    // 053..054 A(002)
	Zeros1                               string          `translator:"part:54..65"`                    // 055..066 N(012)
	ValorBrutoOperacaoAdquirencia        decimal.Decimal `translator:"part:66..77;precision:2"`        // 067..078 N(012)
	ValorCustoOperacao                   decimal.Decimal `translator:"part:78..89;precision:2"`        // 079..090 N(012)
	ValorLiquidoOperacao                 decimal.Decimal `translator:"part:90..101;precision:2"`       // 091..102 N(012)
	TaxaMensalOperacao                   decimal.Decimal `translator:"part:102..112;precision:7"`      // 103..113 N(011)
	TipoContaEstabelecimento             string          `translator:"part:113..114"`                  // 114..115 A(002)
	Banco                                string          `translator:"part:115..117"`                  // 116..118 N(003)
	Agencia                              string          `translator:"part:118..123"`                  // 119..124 N(006)
	ContaCorrente                        string          `translator:"part:124..143"`                  // 125..144 A(020)
	CanalOperacao                        string          `translator:"part:144..146"`                  // 145..147 A(003)
	TipoMovimento                        string          `translator:"part:147..147"`                  // 148..148 A(001)
	TipoParticipante                     string          `translator:"part:148..150"`                  // 149..151 A(003)
	Zeros2                               string          `translator:"part:151..168"`                  // 152..169 N(018)
	TipoDocumentoParticipante            string          `translator:"part:169..169"`                  // 170..170 A(001)
	CNPJCPFParticipante                  string          `translator:"part:170..183"`                  // 171..184 N(014)
	TipoContaEstabelecimentoParticipante string          `translator:"part:184..185"`                  // 185..186 A(002)
	BancoDomicilioBancarioParticipante   string          `translator:"part:186..188"`                  // 187..189 N(003)
	AgenciaDomicilioBancarioParticipante string          `translator:"part:189..194"`                  // 190..195 N(006)
	ContaDomicilioBancarioParticipante   string          `translator:"part:195..214"`                  // 196..215 A(020)
	CodigoEstabelecimentoCentralizador   string          `translator:"part:215..229"`                  // 216..230 A(015)
	RazaoSocialParticipante              string          `translator:"part:230..254"`                  // 231..255 A(025)
	CodigoArranjoPagamento               string          `translator:"part:255..256"`                  // 256..257 A(002)
	ChaveUR                              string          `translator:"part:257..281"`                  // 258..282 A(025)
	Reservado                            string          `translator:"part:282..399"`                  // 283..400 A(118)
}

func (i ResumoFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type DetalheFinanceiro struct {
	TipoRegistro                       string          `translator:"part:0..0"`                      // 001..001 A(001)
	CodigoEstabelecimentoComercial     string          `translator:"part:1..15"`                     // 002..016 A(015)
	DataOperacao                       time.Time       `translator:"part:16..23;timeParse:02012006"` // 017..024 N(008)
	NumeroOperacao                     string          `translator:"part:24..43"`                    // 025..044 A(020)
	TipoOperacao                       string          `translator:"part:44..45"`                    // 045..046 A(002)
	Zeros1                             string          `translator:"part:46..63"`                    // 047..064 N(018)
	CodigoProduto                      string          `translator:"part:64..65"`                    // 065..066 A(002)
	DataVencimentoUR                   time.Time       `translator:"part:66..73;timeParse:02012006"` // 067..074 N(008)
	Zeros2                             string          `translator:"part:74..85"`                    // 075..086 N(012)
	ValorLiquidoURNegociada            decimal.Decimal `translator:"part:86..97;precision:2"`        // 087..098 N(012)
	ValorCustoURNegociada              decimal.Decimal `translator:"part:98..109;precision:2"`       // 099..110 N(012)
	ValorBrutoURNegociada              decimal.Decimal `translator:"part:110..121;precision:2"`      // 111..122 N(012)
	TipoContaEstabelecimento           string          `translator:"part:122..123"`                  // 123..124 A(002)
	Zeros3                             string          `translator:"part:124..126"`                  // 125..127 N(003)
	Zeros4                             string          `translator:"part:127..132"`                  // 128..133 N(006)
	Zeros5                             string          `translator:"part:133..152"`                  // 134..153 A(020)
	TipoMovimento                      string          `translator:"part:153..153"`                  // 154..154 A(001)
	TipoParticipante                   string          `translator:"part:154..156"`                  // 155..157 A(003)
	Zeros6                             string          `translator:"part:157..174"`                  // 158..175 N(018)
	TipoDocumentoParticipante          string          `translator:"part:175..175"`                  // 176..176 A(001)
	CNPJCPFParticipante                string          `translator:"part:176..189"`                  // 177..190 N(014)
	Espaco1                            string          `translator:"part:190..191"`                  // 191..192 A(002)
	Zeros7                             string          `translator:"part:192..194"`                  // 193..195 N(003)
	Zeros8                             string          `translator:"part:195..200"`                  // 196..201 N(006)
	Espaco2                            string          `translator:"part:201..220"`                  // 202..221 A(020)
	CodigoEstabelecimentoCentralizador string          `translator:"part:221..235"`                  // 222..236 A(015)
	RazaoSocialParticipante            string          `translator:"part:236..260"`                  // 237..261 A(025)
	ChaveUR                            string          `translator:"part:261..285"`                  // 262..286 A(025)
	Reservado                          string          `translator:"part:286..399"`                  // 287..400 A(114)
}

func (i DetalheFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type Trailer struct {
	TipoRegistro        string `translator:"part:0..0"`    // 001..001 A(001)
	QuantidadeRegistros int    `translator:"part:1..9"`    // 002..010 N(009)
	Reservado           string `translator:"part:10..399"` // 011..400 A(390)
}

func (i Trailer) String() (string, error) {
	return writer.Marshal(i, 400)
}
