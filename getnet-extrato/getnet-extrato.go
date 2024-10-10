package getnet_extrato

import (
	"time"

	"github.com/libercapital/document-translator-go/internal/writer"
)

type Header struct {
	TipoRegistro          string    `translator:"part:0..0"`                      // Tipo de Registro Fixo 0               001..001 A(001)
	DataCriacaoArquivo    time.Time `translator:"part:1..8;timeParse:02012006"`   // Data de criação do arquivo            002..009 N(008)
	HoraCriacaoArquivo    string    `translator:"part:9..14"`                     // Hora de criação do arquivo            010..015 N(006)
	DataMovimento         time.Time `translator:"part:15..22;timeParse:02012006"` // Data de referência do movimento       016..023 N(008)
	VersaoArquivo         string    `translator:"part:23..30"`                    // Arquivo e Versão                      024..031 A(008)
	CodigoEstabelecimento string    `translator:"part:31..45"`                    // Código do Estabelecimento             032..046 A(015)
	CNPJAdquirente        string    `translator:"part:46..59"`                    // CNPJ do adquirente                    047..060 N(014)
	NomeAdquirente        string    `translator:"part:60..79"`                    // Nome do adquirente                    061..080 A(020)
	NumeroSequencial      string    `translator:"part:80..88"`                    // Número sequencial da remessa          081..089 N(009)
	CodigoAdquirente      string    `translator:"part:89..90"`                    // Código do adquirente                  090..091 A(002)
	VersaoLayout          string    `translator:"part:91..115"`                   // Versão do Layout                      092..116 A(025)
	Reservado             string    `translator:"part:116..399"`                  // Reservado para uso futuro             117..400 A(284)
}

func (i Header) String() (string, error) {
	return writer.Marshal(i, 400)
}

type ResumoTransacional struct {
	TipoRegistro                       string    `translator:"part:0..0"`                        // Tipo de Registro Fixo 1               001..001 A(001)
	CodigoEstabelecimentoComercial     string    `translator:"part:1..15"`                       // Código do Estabelecimento Comercial   002..016 A(015)
	CodigoProduto                      string    `translator:"part:16..17"`                      // Código do Produto                     017..018 A(002)
	FormaCaptura                       string    `translator:"part:18..20"`                      // Forma de Captura                      019..021 A(003)
	NumeroRV                           string    `translator:"part:21..29"`                      // Número do Resumo de Vendas            022..030 N(009)
	DataRV                             time.Time `translator:"part:30..37;timeParse:02012006"`   // Data do Resumo de Venda             031..038 N(008)
	DataPagamentoRV                    time.Time `translator:"part:38..45;timeParse:02012006"`   // Data de crédito do RV               039..046 N(008)
	Banco                              string    `translator:"part:46..48"`                      // Código do Banco                       047..049 N(003)
	Agencia                            string    `translator:"part:49..54"`                      // Código da Agência                     050..055 N(006)
	Zeros1                             string    `translator:"part:55..65"`                      // Zeros                                 056..066 N(011)
	NumeroCVsAceitos                   int       `translator:"part:66..74"`                      // Número de CVs Aceitos                 067..075 N(009)
	NumeroCVsRejeitados                int       `translator:"part:75..83"`                      // Número de CVs Rejeitados              076..084 N(009)
	ValorBruto                         float64   `translator:"part:84..95;precision:2"`          // Valor Bruto                           085..096 N(012)
	ValorLiquido                       float64   `translator:"part:96..107;precision:2"`         // Valor Líquido                         097..108 N(012)
	ValorTarifa                        float64   `translator:"part:108..119;precision:2"`        // Valor da Tarifa                       109..120 N(012)
	ValorTaxaDesconto                  float64   `translator:"part:120..131;precision:2"`        // Valor da Taxa de Desconto             121..132 N(012)
	ValorRejeitado                     float64   `translator:"part:132..143;precision:2"`        // Valor Rejeitado                       133..144 N(012)
	ValorCredito                       float64   `translator:"part:144..155;precision:2"`        // Valor Crédito                         145..156 N(012)
	Zeros2                             string    `translator:"part:156..167"`                    // Zeros                                 157..168 N(012)
	IndicadorTipoPagamento             string    `translator:"part:168..169"`                    // Indicador do Tipo de Pagamento        169..170 A(002)
	NumeroParcelaRV                    int       `translator:"part:170..171"`                    // Número da Parcela do RV               171..172 N(002)
	QuantidadeParcelasRV               int       `translator:"part:172..173"`                    // Quantidade de Parcelas do RV          173..174 N(002)
	CodigoEstabelecimentoCentralizador string    `translator:"part:174..188"`                    // Código do Estabelecimento Centralizador 175..189 A(015)
	Zeros3                             string    `translator:"part:189..203"`                    // Zeros                                 190..204 N(015)
	DataVencimentoOriginal             time.Time `translator:"part:204..211;timeParse:02012006"` // Data do Vencimento Original         205..212 N(008)
	Zeros4                             string    `translator:"part:212..223"`                    // Zeros                                 213..224 N(012)
	Zeros5                             string    `translator:"part:224..235"`                    // Zeros                                 225..236 N(012)
	NumeroControleOperacao             string    `translator:"part:236..253"`                    // Número de Controle da Operação        237..254 N(018)
	ValorLiquidoCobranca               float64   `translator:"part:254..265;precision:2"`        // Valor Líquido da Cobrança             255..266 N(012)
	Zeros6                             string    `translator:"part:266..280"`                    // Zeros                                 267..281 N(015)
	Moeda                              string    `translator:"part:281..283"`                    // Moeda da Transação                    282..284 N(003)
	IdentificadorBaixaCobranca         string    `translator:"part:284..284"`                    // Identificador de Baixa de Cobrança    285..285 A(001)
	SinalTransacao                     string    `translator:"part:285..285"`                    // Sinal da Transação                    286..286 A(001)
	TipoContaPagamento                 string    `translator:"part:286..287"`                    // Tipo de Conta para Pagamento          287..288 A(002)
	ContaCorrente                      string    `translator:"part:288..307"`                    // Número da Conta Corrente              289..308 N(020)
	ChaveUR                            string    `translator:"part:308..332"`                    // Chave da Unidade de Recebíveis        309..333 A(025)
	Reservado                          string    `translator:"part:358..399"`                    // Reservado para uso futuro             360..400 A(041)
}

func (i ResumoTransacional) String() (string, error) {
	return writer.Marshal(i, 400)
}

type AnaliticoTransacional struct {
	TipoRegistro                       string    `translator:"part:0..0"`                        // Tipo de Registro Fixo 2               001..001 A(001)
	CodigoEstabelecimentoComercial     string    `translator:"part:1..15"`                       // Código do Estabelecimento Comercial   002..016 A(015)
	NumeroRV                           string    `translator:"part:16..24"`                      // Número do Resumo de Vendas            017..025 N(009)
	NSUAdquirente                      string    `translator:"part:25..36"`                      // NSU do adquirente                     026..037 N(012)
	DataTransacao                      time.Time `translator:"part:37..44;timeParse:02012006"`   // Data da Transação                     038..045 N(008)
	HoraTransacao                      string    `translator:"part:45..50"`                      // Hora da Transação                     046..051 N(006)
	NumeroCartao                       string    `translator:"part:51..69"`                      // Número do Cartão (truncado)           052..070 A(019)
	ValorTransacao                     float64   `translator:"part:70..81;precision:2"`          // Valor da Transação                    071..082 N(012)
	ValorSaque                         float64   `translator:"part:82..93;precision:2"`          // Valor do Saque                        083..094 N(012)
	ValorTaxaEmbarque                  float64   `translator:"part:94..105;precision:2"`         // Valor da Taxa de Embarque             095..106 N(012)
	NumeroParcelas                     int       `translator:"part:106..107"`                    // Número de Parcelas                    107..108 N(002)
	NumeroParcelaRelacaoCV             int       `translator:"part:108..109"`                    // Número da Parcela Relacionada ao CV   109..110 N(002)
	ValorParcela                       float64   `translator:"part:110..121;precision:2"`        // Valor da Parcela                      111..122 N(012)
	DataPagamento                      time.Time `translator:"part:122..129;timeParse:02012006"` // Data do Pagamento                 123..130 N(008)
	CodigoAutorizacao                  string    `translator:"part:130..139"`                    // Código de Autorização                 131..140 A(010)
	FormaCaptura                       string    `translator:"part:140..142"`                    // Forma de Captura                      141..143 A(003)
	StatusTransacao                    string    `translator:"part:143..143"`                    // Status da Transação                   144..144 A(001)
	CodigoEstabelecimentoCentralizador string    `translator:"part:144..158"`                    // Código do Estabelecimento Centralizador 145..159 A(015)
	CodigoTerminal                     string    `translator:"part:159..166"`                    // Código do Terminal                    160..167 A(008)
	Moeda                              string    `translator:"part:167..169"`                    // Moeda da Transação                    168..170 N(003)
	OrigemEmissorCartao                string    `translator:"part:170..170"`                    // Origem do Emissor do Cartão           171..171 A(001)
	SinalTransacao                     string    `translator:"part:171..171"`                    // Sinal da Transação                    172..172 A(001)
	CarteiraDigital                    string    `translator:"part:172..174"`                    // Carteira Digital                      173..175 A(003)
	ValorComissaoVenda                 float64   `translator:"part:175..186;precision:2"`        // Valor Comissão da Venda               176..187 N(012)
	IdentificadorTipoProximoConteudo   string    `translator:"part:187..188"`                    // Identificador de Tipo do Próximo Conteúdo 188..189 A(002)
	ConteudoDinamico                   string    `translator:"part:189..306"`                    // Conteúdo Dinâmico                     190..307 A(118)
	IdentificadorTipoProximoConteudo2  string    `translator:"part:307..308"`                    // Identificador de Tipo do Próximo Conteúdo 308..309 A(002)
	ConteudoDinamico2                  string    `translator:"part:309..358"`                    // Conteúdo Dinâmico 2                   310..359 A(050)
	Reservado                          string    `translator:"part:359..399"`                    // Reservado para uso futuro             360..400 A(041)
}

func (i AnaliticoTransacional) String() (string, error) {
	return writer.Marshal(i, 400)
}

type AjusteFinanceiro struct {
	TipoRegistro                   string    `translator:"part:0..0"`                        // Tipo de Registro Fixo 3               001..001 A(001)
	CodigoEstabelecimentoComercial string    `translator:"part:1..15"`                       // Código do Estabelecimento Origem     002..016 A(015)
	NumeroRV                       string    `translator:"part:16..24"`                      // Número do RV ajustado                017..025 N(009)
	DataRV                         time.Time `translator:"part:25..32;timeParse:02012006"`   // Data do RV                           026..033 N(008)
	DataPagamentoRV                time.Time `translator:"part:33..40;timeParse:02012006"`   // Data de crédito do Resumo de Venda   034..041 N(008)
	IdentificadorAjuste            string    `translator:"part:41..60"`                      // Identificador do Ajuste              042..061 N(020)
	Brancos                        string    `translator:"part:61..61"`                      // Preenchido com brancos               062..062 A(001)
	SinalValorAjuste               string    `translator:"part:62..62"`                      // Sinal do valor do ajuste             063..063 A(001)
	ValorAjuste                    float64   `translator:"part:63..74;precision:2"`          // Valor do Ajuste                      064..075 N(012)
	MotivoAjuste                   string    `translator:"part:75..76"`                      // Motivo do Ajuste                     076..077 A(002)
	DataCarta                      time.Time `translator:"part:77..84;timeParse:02012006"`   // Data da Carta de Cancelamento        078..085 N(008)
	NumeroCartao                   string    `translator:"part:85..103"`                     // Número do Cartão (truncado)          086..104 A(019)
	NumeroRVOriginal               string    `translator:"part:104..112"`                    // Número do RV Original                105..113 N(009)
	NSUAdquirente                  string    `translator:"part:113..124"`                    // NSU do Adquirente                    114..125 N(012)
	DataTransacaoOriginal          time.Time `translator:"part:125..132;timeParse:02012006"` // Data da Transação Original        126..133 N(008)
	IndicadorTipoPagamento         string    `translator:"part:133..134"`                    // Indicador do Tipo de Pagamento       134..135 A(002)
	NumeroTerminalOriginal         string    `translator:"part:135..142"`                    // Número do Terminal Original          136..143 A(008)
	DataPagamentoOriginal          string    `translator:"part:143..150"`                    // Data do Pagamento Original       144..151 N(008)
	Moeda                          string    `translator:"part:151..153"`                    // Moeda da Transação                   152..154 N(003)
	ValorComissaoVendaCancelada    float64   `translator:"part:154..165;precision:2"`        // Valor da Comissão da Venda Cancelada 155..166 N(012)
	IdentificadorProximoConteudo   string    `translator:"part:166..167"`                    // Identificador do Próximo Conteúdo    167..168 A(002)
	ConteudoDinamico               string    `translator:"part:168..285"`                    // Conteúdo Dinâmico                    169..286 A(118)
	Reservado                      string    `translator:"part:286..399"`                    // Reservado para uso futuro            287..400 A(114)
}

func (i AjusteFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type ResumoFinanceiro struct {
	TipoRegistro                         string    `translator:"part:0..0"`                      // Tipo de Registro Fixo 5               001..001 A(001)
	CodigoEstabelecimentoComercial       string    `translator:"part:1..15"`                     // Código do Estabelecimento Comercial   002..016 A(015)
	DataOperacao                         time.Time `translator:"part:16..23;timeParse:02012006"` // Data da Operação                    017..024 N(008)
	DataCreditoOperacao                  time.Time `translator:"part:24..31;timeParse:02012006"` // Data do Crédito da Operação         025..032 N(008)
	NumeroOperacao                       string    `translator:"part:32..51"`                    // Número da Operação                  033..052 A(020)
	TipoOperacao                         string    `translator:"part:52..53"`                    // Tipo de Operação                    053..054 A(002)
	Zeros1                               string    `translator:"part:54..65"`                    // Zeros                               055..066 N(012)
	ValorBrutoOperacaoAdquirencia        float64   `translator:"part:66..77;precision:2"`        // Valor Bruto da Operação - Adquirência 067..078 N(012)
	ValorCustoOperacao                   float64   `translator:"part:78..89;precision:2"`        // Valor do Custo da Operação          079..090 N(012)
	ValorLiquidoOperacao                 float64   `translator:"part:90..101;precision:2"`       // Valor Líquido da Operação           091..102 N(012)
	TaxaMensalOperacao                   float64   `translator:"part:102..112;precision:7"`      // Taxa Mensal da Operação (%)         103..113 N(011)
	TipoContaEstabelecimento             string    `translator:"part:113..114"`                  // Tipo de Conta do Estabelecimento    114..115 A(002)
	Banco                                string    `translator:"part:115..117"`                  // Código do Banco                     116..118 N(003)
	Agencia                              string    `translator:"part:118..123"`                  // Código da Agência                   119..124 N(006)
	ContaCorrente                        string    `translator:"part:124..143"`                  // Número da Conta Corrente            125..144 A(020)
	CanalOperacao                        string    `translator:"part:144..146"`                  // Canal de Operação                   145..147 A(003)
	TipoMovimento                        string    `translator:"part:147..147"`                  // Tipo de Movimento                   148..148 A(001)
	TipoParticipante                     string    `translator:"part:148..150"`                  // Tipo de Participante                149..151 A(003)
	Zeros2                               string    `translator:"part:151..168"`                  // Zeros                               152..169 N(018)
	TipoDocumentoParticipante            string    `translator:"part:169..169"`                  // Tipo de Documento do Participante   170..170 A(001)
	CNPJCPFParticipante                  string    `translator:"part:170..183"`                  // CNPJ/CPF do Participante            171..184 N(014)
	TipoContaEstabelecimentoParticipante string    `translator:"part:184..185"`                  // Tipo da Conta do Estabelecimento Participante  185..186 A(002)
	BancoDomicilioBancarioParticipante   string    `translator:"part:186..188"`                  // Banco do Domicílio Bancário Participante 187..189 N(003)
	AgenciaDomicilioBancarioParticipante string    `translator:"part:189..194"`                  // Agência do Domicílio Bancário Participante 190..195 N(006)
	ContaDomicilioBancarioParticipante   string    `translator:"part:195..214"`                  // Conta do Domicílio Bancário Participante 196..215 A(020)
	CodigoEstabelecimentoCentralizador   string    `translator:"part:215..229"`                  // Código do Estabelecimento Comercial Centralizador 216..230 A(015)
	RazaoSocialParticipante              string    `translator:"part:230..254"`                  // Razão Social do Participante        231..255 A(025)
	CodigoArranjoPagamento               string    `translator:"part:255..256"`                  // Código do Arranjo de Pagamento      256..257 A(002)
	ChaveUR                              string    `translator:"part:257..281"`                  // Chave da UR                         258..282 A(025)
	Reservado                            string    `translator:"part:282..399"`                  // Reservado para uso futuro           283..400 A(118)
}

func (i ResumoFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type DetalheFinanceiro struct {
	TipoRegistro                       string    `translator:"part:0..0"`                      // Tipo de Registro Fixo 6               001..001 A(001)
	CodigoEstabelecimentoComercial     string    `translator:"part:1..15"`                     // Código do Estabelecimento Origem     002..016 A(015)
	DataOperacao                       time.Time `translator:"part:16..23;timeParse:02012006"` // Data da Operação                     017..024 N(008)
	NumeroOperacao                     string    `translator:"part:24..43"`                    // Número da operação                   025..044 A(020)
	TipoOperacao                       string    `translator:"part:44..45"`                    // Tipo de operação                     045..046 A(002)
	Zeros1                             string    `translator:"part:46..63"`                    // Zeros                                047..064 N(018)
	CodigoProduto                      string    `translator:"part:64..65"`                    // Código do Produto                    065..066 A(002)
	DataVencimentoUR                   time.Time `translator:"part:66..73;timeParse:02012006"` // Data de vencimento da UR            067..074 N(008)
	Zeros2                             string    `translator:"part:74..85"`                    // Zeros                                075..086 N(012)
	ValorLiquidoURNegociada            float64   `translator:"part:86..97;precision:2"`        // Valor Líquido da UR Negociada        087..098 N(012)
	ValorCustoURNegociada              float64   `translator:"part:98..109;precision:2"`       // Valor do Custo da UR Negociada       099..110 N(012)
	ValorBrutoURNegociada              float64   `translator:"part:110..121;precision:2"`      // Valor Bruto da UR Negociada          111..122 N(012)
	TipoContaEstabelecimento           string    `translator:"part:122..123"`                  // Tipo de conta do estabelecimento    123..124 A(002)
	Zeros3                             string    `translator:"part:124..126"`                  // Zeros                                125..127 N(003)
	Zeros4                             string    `translator:"part:127..132"`                  // Zeros                                128..133 N(006)
	Zeros5                             string    `translator:"part:133..152"`                  // Zeros                                134..153 A(020)
	TipoMovimento                      string    `translator:"part:153..153"`                  // Tipo de movimento                    154..154 A(001)
	TipoParticipante                   string    `translator:"part:154..156"`                  // Tipo de participante                 155..157 A(003)
	Zeros6                             string    `translator:"part:157..174"`                  // Zeros                                158..175 N(018)
	TipoDocumentoParticipante          string    `translator:"part:175..175"`                  // Tipo de documento do participante    176..176 A(001)
	CNPJCPFParticipante                string    `translator:"part:176..189"`                  // CNPJ/CPF do participante             177..190 N(014)
	Espaco1                            string    `translator:"part:190..191"`                  // Espaço                               191..192 A(002)
	Zeros7                             string    `translator:"part:192..194"`                  // Zeros                                193..195 N(003)
	Zeros8                             string    `translator:"part:195..200"`                  // Zeros                                196..201 N(006)
	Espaco2                            string    `translator:"part:201..220"`                  // Espaço                               202..221 A(020)
	CodigoEstabelecimentoCentralizador string    `translator:"part:221..235"`                  // Código do Estabelecimento Centralizador 222..236 A(015)
	RazaoSocialParticipante            string    `translator:"part:236..260"`                  // Razão Social do Participante         237..261 A(025)
	ChaveUR                            string    `translator:"part:261..285"`                  // Chave da Unidade de Recebíveis       262..286 A(025)
	Reservado                          string    `translator:"part:286..399"`                  // Reservado para uso futuro            287..400 A(114)
}

func (i DetalheFinanceiro) String() (string, error) {
	return writer.Marshal(i, 400)
}

type Trailer struct {
	TipoRegistro        string `translator:"part:0..0"`    // Tipo de Registro Fixo 9               001..001 A(001)
	QuantidadeRegistros int    `translator:"part:1..9"`    // Quantidade total de registros        002..010 N(009)
	Reservado           string `translator:"part:10..399"` // Reservado para uso futuro            011..400 A(390)
}

func (i Trailer) String() (string, error) {
	return writer.Marshal(i, 400)
}
