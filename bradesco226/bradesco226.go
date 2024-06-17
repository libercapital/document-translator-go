package bradesco226

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/libercapital/document-translator-go/internal/writer"
)

type Header struct {
	RegisterType           string          `translator:"part:0..0;kind:1"`               // Tipo de Registro Fixo 1                            001..001 9(001)
	ContractNumber         string          `translator:"part:1..9"`                      // Numero do contrato original                        002..010 9(009)
	MovementDate           time.Time       `translator:"part:10..17;timeParse:02012006"` // Data do movimento                                  011..018 9(008)
	ContractQuantity       decimal.Decimal `translator:"part:18..26"`                    // Quantidade de contratos                            019..027 9(009)
	BorrowerQuantity       decimal.Decimal `translator:"part:27..35"`                    // Quantidade de devedores                            028..036 9(009)
	InstallmentQuantity    decimal.Decimal `translator:"part:36..44"`                    // Quantidade de parcelas                             037..045 9(009)
	SourceCompanyCode      string          `translator:"part:45..51"`                    // Codigo da empresa origem                           046..052 9(007)
	RetroactiveAccountCode string          `translator:"part:52..52"`                    // Codigo da conta retroativa                         053..053 9(001)
}

func (h Header) String() (string, error) {
	return writer.Marshal(h, 226)
}

type Contract struct {
	RegisterType                  string          `translator:"part:0..0;kind:2"`               // Tipo de Registro Fixo 2                  001..001 9(001)
	ContractNumber                string          `translator:"part:1..9"`                      // Numero do contrato                       002..010 9(009)
	AgreementProduct              string          `translator:"part:10..12"`                    // Contrato Convênio - Produto              011..013 9(003)
	AgreementContractFamily       string          `translator:"part:13..13"`                    // Contrato Convênio - Familia              014..014 9(001)
	AgreementContractNumber       string          `translator:"part:14..22"`                    // Contrato Convênio - Numero               015..023 9(009)
	BorrowerDocumentNumber        string          `translator:"part:23..31"`                    // Tomador - Numero do Documento            024..032 9(009)
	BorrowerFilial                string          `translator:"part:32..35"`                    // Tomador - Filial                         033..036 9(004)
	BorrowerDocumentControl       string          `translator:"part:36..37"`                    // Tomador - Controle do Documento          037..038 9(002)
	Date                          time.Time       `translator:"part:38..45;timeParse:02012006"` // Data do Contrato                         039..046 9(008)
	ContractValue                 decimal.Decimal `translator:"part:46..62;precision:2"`        // Valor do Contrato                        047..063 9(015)(2)
	ProductValue                  decimal.Decimal `translator:"part:63..79;precision:2"`        // Valor do Bem                             064..080 9(015)(2)
	InterestTax                   decimal.Decimal `translator:"part:80..92;precision:7"`        // Taxa de Juros                            081..093 9(006)(7)
	FinancingObject               string          `translator:"part:93..97"`                    // Objeto do Financiamento                  094..098 9(005)
	Installments                  string          `translator:"part:98..100"`                   // Quantidade de Parcelas                   099..101 9(003)
	PaymentMethod                 string          `translator:"part:101..103"`                  // Código do meio de pagamento              102..104 9(003)
	PaymentBranch                 string          `translator:"part:104..108"`                  // Agência de Pagamento                     105..109 9(005)
	PaymentAccountNumber          string          `translator:"part:109..115"`                  // Conta de Pagamento                       110..116 9(007)
	IOF                           decimal.Decimal `translator:"part:116..132;precision:2"`      // IOF                                      117..133 9(015)(2)
	TAC                           decimal.Decimal `translator:"part:133..149;precision:2"`      // TAC                                      134..150 9(015)(2)
	Modality                      string          `translator:"part:150..150"`                  // Código da Modalidade                     151..151 9(001)
	EconomicIndex                 string          `translator:"part:151..155"`                  // Indice Econômico                         152..156 9(005)
	CreditApprovalBranch          string          `translator:"part:156..160"`                  // Agência de Liberação de Crédito          157..161 9(005)
	CreditApprovalAccountNumber   string          `translator:"part:161..167"`                  // Conta de Liberação de Crédito            162..168 9(007)
	CreditApprovalMethod          string          `translator:"part:168..168"`                  // Código do Meio de Liberação de Crédito   169..169 9(001)
	FloatRate                     string          `translator:"part:169..170"`                  // Float (Zeros)                            170..171 9(002)
	VehicleInsuranceBranch        string          `translator:"part:171..175"`                  // Agência de crédito do seguro             172..176 9(005)
	VehicleInsuranceAccountNumber string          `translator:"part:176..182"`                  // Conta de crédito do seguro               177..183 9(007)
	VehicleInsuranceAmount        decimal.Decimal `translator:"part:183..199;precision:2"`      // Valor do seguro do veículo               184..200 9(015)(2)
	Commission                    decimal.Decimal `translator:"part:200..216;precision:2"`      // Comissão                                 201..217 9(015)(2)
	BillOfExchangeNumber          string          `translator:"part:217..225"`                  // Número da duplicata                      218..226 X(009)
}

func (c Contract) String() (string, error) {
	return writer.Marshal(c, 226)
}

type Borrower struct {
	RegisterType           string `translator:"part:0..0;kind:3"`          // Tipo de Registro - Fixo 3                  001..001 9(001)
	ContractNumber         string `translator:"part:1..9"`                 // Numero do contrato                         002..010 9(009)
	PersonType             string `translator:"part:10..10"`               // Tipo de Pessoa - 1 p/ PF 2 p/PJ            011..011 9(001)
	Name                   string `translator:"part:11..70;align:right"`   // Nome do Cliente                            012..071 X(060)
	Address                string `translator:"part:71..110;align:right"`  // Logradouro                                 072..111 X(040)
	AddressNumber          string `translator:"part:111..115;align:right"` // Numero do Logradouro                       112..116 X(005)
	AddressComplement      string `translator:"part:116..125;align:right"` // Complemento do Logradouro                  117..126 X(010)
	Neighborhood           string `translator:"part:126..145;align:right"` // Bairro                                     127..146 X(020)
	ZipCode                string `translator:"part:146..153"`             // CEP                                        147..154 9(008)
	CompanySize            string `translator:"part:154..156"`             // Porte da Empresa                           155..157 9(003)
	LegalStatus            string `translator:"part:157..159"`             // Natureza Jurídica                          158..160 9(003)
	ActivityCode           string `translator:"part:160..164"`             // Código da Atividade                        161..165 9(005)
	Phone                  string `translator:"part:165..176"`             // Telefone                                   166..177 9(012)
	PhoneExtension         string `translator:"part:177..181"`             // Ramal do Telefone                          178..182 9(005)
	OriginalContractNumber string `translator:"part:182..221"`             // Numero do contrato original na C3          183..222 9(040)
}

func (b Borrower) String() (string, error) {
	return writer.Marshal(b, 226)
}

type Installment struct {
	RegisterType               string          `translator:"part:0..0;kind:4"`               // Tipo de Registro Fixo 4                  001..001 9(001)
	ContractNumber             string          `translator:"part:1..9"`                      // Numero do contrato                       002..010 9(009)
	InstallmentNumber          string          `translator:"part:10..12"`                    // Numero da parcela                        011..013 9(003)
	DueDate                    time.Time       `translator:"part:13..20;timeParse:02012006"` // Data de vencimento                       014..021 9(008)
	Amount                     decimal.Decimal `translator:"part:21..37;precision:2"`        // Valor da parcela                         022..038 9(015)(2)
	Number                     string          `translator:"part:38..48"`                    // Nosso número                             039..049 9(011)
	PreDatedCheckBank          string          `translator:"part:49..51"`                    // Banco do cheque pré-datado               050..052 9(003)
	PreDatedCheckBranch        string          `translator:"part:52..56"`                    // Agência do cheque pré-datado             053..057 9(005)
	PreDatedCheckAccountNumber string          `translator:"part:57..67"`                    // Conta do cheque pré-datado               058..068 9(011)
	PreDatedCheckNumber        string          `translator:"part:68..74"`                    // Numero do cheque pré-datado              069..075 9(007)
	DailyDefaultValue          decimal.Decimal `translator:"part:75..87;precision:2"`        // Valor Mora Dia                           076..088 9(011)(2)
	Status                     string          `translator:"part:88..88"`                    // Situação da parcela                      089..089 9(001)
	InstallmentInQuantity      decimal.Decimal `translator:"part:89..105;precision:5"`       // Parcela em quantitdade                   090..106 9(012)(5)
	DailyDefaultInQuantity     decimal.Decimal `translator:"part:106..122;precision:5"`      // Mora Dia em Quantidade                   107..123 9(012)(5)
	ChargePortfolio            string          `translator:"part:210..212"`                  // Carteira de Cobrança                     211..213 9(003)
	AssignorBranch             string          `translator:"part:213..217"`                  // Agência Cedente                          214..218 9(005)
	AssignorAccountNumber      string          `translator:"part:218..224"`                  // Conta Cedente                            219..225 9(007)
	DuplicatedTitles           string          `translator:"part:225..225"`                  // Títulos Duplicados                       226..226 9(001)
}

func (i Installment) String() (string, error) {
	return writer.Marshal(i, 226)
}
