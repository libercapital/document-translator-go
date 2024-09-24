package brf240

import (
	"time"

	"github.com/libercapital/document-translator-go/internal/writer"
	"github.com/shopspring/decimal"
)

type BillingReturnFileHeader struct {
	BankCode         string    `translator:"part:0..2"`                        //Código do Banco                          001..003   9(003)
	BatchNumber      int       `translator:"part:3..6"`                        //Lote de Serviço                          004..007   9(004)
	RegistryKind     int       `translator:"part:7..7"`                        //Tipo de Registro                         008..008   9(001)
	KindBuyer        int       `translator:"part:17..17"`                      //Tipo de Inscrição da Empresa             018..018   9(001)
	BuyerDocument    string    `translator:"part:18..31"`                      //Número Inscrição da Empresa              019..032   9(014)
	ContractNumber   string    `translator:"part:32..51"`                      //Código do Convenio no Banco              033..052   X(020)
	Agency           int       `translator:"part:52..56"`                      //Agência Mantenedora da Conta             053..057   9(005)
	AgencyCd         string    `translator:"part:57..57"`                      //Dígito Verificador da Agência            058..058   X(001)
	Account          int       `translator:"part:58..69"`                      //Número da Conta Corrente                 059..070   9(012)
	AccountCd        string    `translator:"part:70..70"`                      //Dígito Verificador da Conta              071..071   X(001)
	CheckDigit       string    `translator:"part:71..71"`                      //Dígito Verificador da Agência / Conta    072..072   X(001)
	BuyerName        string    `translator:"part:72..101"`                     //Nome da Empresa                          073..102   X(030)
	BankName         string    `translator:"part:102..131"`                    //Nome do Banco                            103..132   X(030)
	FileKind         int       `translator:"part:142..142"`                    //Código Remessa / Retorno                 143..143   9(001)
	FileDate         time.Time `translator:"part:143..150;timeParse:02012006"` //Data da Geração do Arquivo               144..151   9(008)
	FileTime         time.Time `translator:"part:151..156;timeParse:150405"`   //Hora da Geração do Arquivo               152..157   9(006)
	SequentialNumber int       `translator:"part:157..162"`                    //Número Seqüencial do Arquivo             158..163   9(006)
	LayoutVersion    string    `translator:"part:163..165"`                    //Número da Versão do Layout               164..166   9(003)
	RecordDensity    int       `translator:"part:166..170"`                    //Densidade de Gravação Arquivo            167..171   9(005)
	BankReserved     string    `translator:"part:171..190"`                    //Uso Reservado do Banco                   172..191   X(020)
	BuyerReserved    string    `translator:"part:191..210"`                    //Uso Reservado da Empresa                 192..211   X(020)
}

// TODO criar tag default nas estruturas
func (b BillingReturnFileHeader) New() BillingReturnFileHeader {
	b.RegistryKind = 0
	b.FileKind = 2
	b.LayoutVersion = "060"
	b.RecordDensity = 6250

	return b
}

func (b BillingReturnFileHeader) String() (string, error) {
	return writer.Marshal(b, 240)
}

type BillingReturnBatchHeader struct {
	BankCode           string `translator:"part:0..2"`     //Código do Banco                          001..003   9(003)
	BatchNumber        int    `translator:"part:3..6"`     //Lote de Serviço                          004..007   9(004)
	RegistryKind       int    `translator:"part:7..7"`     //Tipo de Registro                         008..008   9(001)
	OperationKind      string `translator:"part:8..8"`     //Tipo da Operação                         009..009   X(001)
	ServiceKind        int    `translator:"part:9..10"`    //Tipo de Serviço                          010..011   9(002)
	ReleaseKind        int    `translator:"part:11..12"`   //Forma de Lançamento                      012..013   9(002)
	BatchLayoutVersion int    `translator:"part:13..15"`   //Número da Versão do Lote                 014..016   9(003)
	KindBuyer          int    `translator:"part:17..17"`   //Tipo de Inscrição da Empresa             018..018   9(001)
	BuyerDocument      int    `translator:"part:18..31"`   //Número de Inscrição da Empresa           019..032   9(014)
	ContractNumber     string `translator:"part:32..51"`   //Código do Convenio no Banco              033..052   X(020)
	Agency             int    `translator:"part:52..56"`   //Agência Mantenedora da Conta             053..057   9(005)
	AgencyCd           string `translator:"part:57..57"`   //Dígito Verificador da Agência            058..058   X(001)
	Account            int    `translator:"part:58..69"`   //Número da Conta Corrente                 059..070   9(012)
	AccountCd          string `translator:"part:70..70"`   //Dígito Verificador da Conta              071..071   X(001)
	CheckDigit         string `translator:"part:71..71"`   //Dígito Verificador da Agência/Conta      072..072   X(001)
	BuyerName          string `translator:"part:72..101"`  //Nome da Empresa                          073..102   X(030)
	GenericMessage     string `translator:"part:102..141"` //Informação 1 - Mensagem                  103..142   X(040)
	AddressStreet      string `translator:"part:142..171"` //Endereço                                 143..172   X(030)
	AddressNumber      int    `translator:"part:172..176"` //Número                                   173..177   9(005)
	AddressComplement  string `translator:"part:177..191"` //Complemento do Endereço                  178..192   X(015)
	AddressCity        string `translator:"part:192..211"` //Cidade                                   193..212   X(020)
	AddressZipCode     int    `translator:"part:212..219"` //CEP                                      213..217   9(005)
	AddressState       string `translator:"part:220..221"` //UF                                       221..222   X(002)
	Filler             string `translator:"part:222..229"` //Filler                                   223..230   X(008)
	Occurrence         string `translator:"part:230..239"` //Ocorrências para o Retorno               231..240   X(010)
}

// TODO criar tag default nas estruturas
func (b BillingReturnBatchHeader) New() BillingReturnBatchHeader {
	b.RegistryKind = 1
	b.OperationKind = "C"
	b.ServiceKind = 20
	b.ReleaseKind = 3
	b.BatchLayoutVersion = 60

	return b
}

func (b BillingReturnBatchHeader) String() (string, error) {
	return writer.Marshal(b, 240)
}

type BillingReturnSegmentA struct {
	BankCode              string          `translator:"part:0..2"`                        //Código do Banco                         001..003   9(003)
	BatchNumber           int             `translator:"part:3..6"`                        //Lote de Serviço                         004..007   9(004)
	RegistryKind          int             `translator:"part:7..7"`                        //Tipo de Registro                        008..008   9(001)
	BatchSequentialNumber int             `translator:"part:8..12"`                       //Número Seqüencial do Registro no Lote   009..013   9(005)
	SegmentKind           string          `translator:"part:13..13"`                      //Código Segmento do Registro Detalhe     014..014   X(001)
	ActionKind            int             `translator:"part:14..14"`                      //Tipo de Movimento                       015..015   9(001)
	ActionInstructionKind int             `translator:"part:15..16"`                      //Código da Instrução para Movimento      016..017   9(002)
	VendorName            string          `translator:"part:17..52"`                      //Nome do Fornecedor                      018..053   X(036)
	DocumentKind          int             `translator:"part:53..53"`                      //Se CNPJ = "2". Se CPF = "1"             054..054   9(001)
	FinancingDate         time.Time       `translator:"part:54..61;timeParse:02012006"`   //Data de financiamento                   055..062   X(008)
	Document              string          `translator:"part:62..78"`                      //CPNJ ou CPF                             063..079   9(017)
	VendorBankCode        int             `translator:"part:79..83"`                      //Número do Banco Fornecedor              080..084   9(005)
	VendorAgency          int             `translator:"part:84..92"`                      //Agência do Banco Fornecedor             085..093   9(009)
	VendorAgencyCd        string          `translator:"part:93..93"`                      //Dígito da Agência                       094..094   X(001)
	VendorAccount         string          `translator:"part:94..106;align:right"`         //Conta Bancária                          095..107   9(013)
	VendorAccountCd       string          `translator:"part:107..107"`                    //Dígito Verificador da Conta             108..108   X(001)
	PaymentNumber         string          `translator:"part:108..129"`                    //Número da Nota Fiscal/Fatura            109..130   X(022)
	IssueDate             time.Time       `translator:"part:130..137;timeParse:02012006"` //Data de emissão do documento            131..138   X(008)
	DueDate               time.Time       `translator:"part:138..145;timeParse:02012006"` //Data do vencimento                      139..146   X(008)
	PaymentValue          decimal.Decimal `translator:"part:146..166;precision:2"`        //Valor do título                         147..167   9(021)(2)
	DiscountValue         decimal.Decimal `translator:"part:167..180;precision:2"`        //Valor do desconto                       168..181   9(012)(2)
	FinancingValue        decimal.Decimal `translator:"part:181..194;precision:2"`        //Valor liquido                           182..195   9(012)(2)
	DiscountRate          decimal.Decimal `translator:"part:195..200;precision:4"`        //Taxa de adiantamento                    196..201   9(2)(4)
	ReferenceNumber       string          `translator:"part:201..229"`                    //Numero de referência                    202..230   X(029)
	Occurrence            string          `translator:"part:230..239"`                    //Status da Partida/Código de ocorrência  231..240   X(010)
}

// TODO criar tag default nas estruturas
func (b BillingReturnSegmentA) New() BillingReturnSegmentA {
	b.RegistryKind = 3
	b.SegmentKind = "A"
	b.ActionKind = 0
	b.ActionInstructionKind = 0

	return b
}

func (b BillingReturnSegmentA) String() (string, error) {
	return writer.Marshal(b, 240)
}

type BillingReturnBatchTrailer struct {
	BankCode                string          `translator:"part:0..2"`               //Código do Banco                      001..003   9(003)
	BatchNumber             int             `translator:"part:3..6"`               //Lote de Serviço                      004..007   9(004
	RegistryKind            int             `translator:"part:7..7"`               //Tipo de Registro                     008..008   9(001)
	QuantityRegistries      int             `translator:"part:17..22"`             //Quantidade de Registros do Lote      018..023   9(006)
	ValueAmount             decimal.Decimal `translator:"part:23..40;precision:2"` //Somatória dos Valores                024..041   9(016)V2
	CurrencyQuantity        decimal.Decimal `translator:"part:41..58;precision:2"` //Somatória Quantidade Moeda           042..059   9(013)V5
	DebitNotificationNumber int             `translator:"part:59..64"`             //Número Aviso de Débito               060..065   9(006)
	Occurrence              string          `translator:"part:230..239"`           //Ocorrências para o Retorno           231..240   X(010)
}

// TODO criar tag default nas estruturas
func (b BillingReturnBatchTrailer) New() BillingReturnBatchTrailer {
	b.RegistryKind = 5

	return b
}

func (b BillingReturnBatchTrailer) String() (string, error) {
	return writer.Marshal(b, 240)
}

type BillingReturnFileTrailer struct {
	BankCode             string `translator:"part:0..2"`   //Código do Banco                        001..003   9(003)
	BatchNumber          int    `translator:"part:3..6"`   //Lote de Serviço                        004..007   9(004)
	RegistryKind         int    `translator:"part:7..7"`   //Tipo de Registro                       008..008   9(001)
	BatchesQuantity      int    `translator:"part:17..22"` //Quantidade de lotes do arquivo         018..023   9(006)
	FileRegistryQuantity int    `translator:"part:23..28"` //Quantidade de registros no arquivo     024..029   9(006)
}

// TODO criar tag default nas estruturas
func (b BillingReturnFileTrailer) New() BillingReturnFileTrailer {
	b.BatchNumber = 9999
	b.RegistryKind = 9

	return b
}

func (b BillingReturnFileTrailer) String() (string, error) {
	return writer.Marshal(b, 240)
}
