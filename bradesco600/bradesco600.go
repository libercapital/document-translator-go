package bradesco600

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/libercapital/document-translator-go/internal/writer"
)

type CreditAssessment struct {
	BaseDate                    time.Time       `translator:"part:0..7;timeParse:02012006"`       // Data base                             001..008 9(008)
	ContractNumber              string          `translator:"part:8..24"`                         // Número do contrato                    009..025 9(017)
	CustomerName                string          `translator:"part:25..64"`                        // Nome do cliente                       026..065 X(040)
	PersonType                  string          `translator:"part:65..65"`                        // Tipo de pessoa                        066..066 X(001)
	DocumentNumber              string          `translator:"part:66..80"`                        // CNPJ / CPF                            067..081 9(015)
	AssessmentType              string          `translator:"part:81..100"`                       // Modalidade                            082..101 X(020)
	ContractStartDate           time.Time       `translator:"part:101..108;timeParse:02012006"`   // Data Início Contrato                  102..109 9(008)
	ContractEndDate             time.Time       `translator:"part:109..116;timeParse:02012006"`   // Data Fim Contrato                     110..117 9(008)
	PaidInstallments            uint            `translator:"part:117..120"`                      // Qtde de parcelas pagas                118..121 9(004)
	OverdueInstallments         uint            `translator:"part:121..124"`                      // Qtde de parcelas vencidas             122..125 9(004)
	QtyInstallments             uint            `translator:"part:125..128"`                      // Quantidade total de parcelas          122..125 9(004)
	AnualContractFee            decimal.Decimal `translator:"part:129..139;precision:7"`          // Taxa ao ano                           130..140 9(011)(7)
	Indexer                     string          `translator:"part:140..163"`                      // Indexador                             141..164 X(024)
	InstallmentPrice            decimal.Decimal `translator:"part:164..180;precision:2"`          // Valor da parcela                      165..181 9(017)(2)
	ContractPrice               decimal.Decimal `translator:"part:181..197;precision:2"`          // Valor principal                       182..198 9(017)(2)
	GuaranteePrice              decimal.Decimal `translator:"part:198..214;precision:2"`          // Valor da garantia                     199..215 9(017)(2)
	InitialContractPrice        decimal.Decimal `translator:"part:215..231;precision:2"`          // Valor de entrada                      216..232 9(017)(2)
	DuePrice                    decimal.Decimal `translator:"part:232..248;precision:2"`          // Saldo a vencer total                  233..249 9(017)(2)
	DuePriceNext15To30Days      decimal.Decimal `translator:"part:249..265;precision:2"`          // Saldo a vencer entre 15 e 30 dias     250..266 9(017)(2)
	DuePriceNext31To60Days      decimal.Decimal `translator:"part:266..282;precision:2"`          // Saldo a vencer entre 31 e 60 dias     267..283 9(017)(2)
	DuePriceNext61To90Days      decimal.Decimal `translator:"part:283..299;precision:2"`          // Saldo a vencer entre 61 e 90 dias     267..283 9(017)(2)
	DuePriceNext91To120Days     decimal.Decimal `translator:"part:300..316;precision:2"`          // Saldo a vencer entre 91 e 120 dias    301..317 9(017)(2)
	DuePriceNext121To150Days    decimal.Decimal `translator:"part:317..333;precision:2"`          // Saldo a vencer entre 121 e 150 dias   318..334 9(017)(2)
	DuePriceNext151To180Days    decimal.Decimal `translator:"part:334..350;precision:2"`          // Saldo a vencer entre 151 e 180 dias   335..351 9(017)(2)
	DuePriceNext181To360Days    decimal.Decimal `translator:"part:351..367;precision:2"`          // Saldo a vencer entre 181 e 360 dias   352..368 9(017)(2)
	DuePriceOver360Days         decimal.Decimal `translator:"part:368..384;precision:2"`          // Saldo a vencer acima de 360 dias      369..385 9(017)(2)
	TotalOverDuePrice           decimal.Decimal `translator:"part:385..401;precision:2"`          // Saldo vencido total                   386..402 9(017)(2)
	FirstOverDueInstallmentDate time.Time       `translator:"part:402..409;timeParse:02012006"`   // Data da primeira parcela vencida      403..410 9(008)
	OverDuePrice15To30Days      decimal.Decimal `translator:"part:410..426;precision:2"`          // Saldo vencido entre 15 e 30 dias      411..427 9(017)(2)
	OverDuePrice31To60Days      decimal.Decimal `translator:"part:427..443;precision:2"`          // Saldo vencido entre 31 e 60 dias      428..444 9(017)(2)
	OverDuePrice61To90Days      decimal.Decimal `translator:"part:444..460;precision:2"`          // Saldo vencido entre 61 e 90 dias      445..461 9(017)(2)
	OverDuePrice91To120Days     decimal.Decimal `translator:"part:461..477;precision:2"`          // Saldo vencido entre 91 e 120 dias     462..478 9(017)(2)
	OverDuePrice121To150Days    decimal.Decimal `translator:"part:478..494;precision:2"`          // Saldo vencido entre 121 e 150 dias    479..495 9(017)(2)
	OverDuePrice151To180Days    decimal.Decimal `translator:"part:495..511;precision:2"`          // Saldo vencido entre 151 e 180 dias    496..512 9(017)(2)
	OverDuePrice181To360Days    decimal.Decimal `translator:"part:512..528;precision:2"`          // Saldo vencido entre 181 e 360 dias    513..529 9(017)(2)
	OverDuePriceOver360Days     decimal.Decimal `translator:"part:529..545;precision:2"`          // Saldo vencido há mais de 360 dias     530..546 9(017)(2)
	OperationRating             string          `translator:"part:546..547"`                      // Rating da operação                    547..548 X(002)
	VehicleBrand                string          `translator:"part:548..567"`                      // Marca do veículo                      549..568 X(020)
	VehicleModel                string          `translator:"part:568..587"`                      // Modelo do veículo                     569..588 X(020)
	VehicleYear                 uint            `translator:"part:588..591"`                      // Ano do veículo                        589..592 9(004)
	SystemSource                uint            `translator:"part:592..598"`                      // Sistema de origem                     593..599 9(007)
	AquisitionDate              time.Time       `translator:"part:599..608;timeParse:02.01.2006"` // Data de aquisição                     600..609 X(010)
	AquisitionCode              uint            `translator:"part:609..611"`                      // Código de aquisição                   610..612 9(003)
	Return                      string          `translator:"part:612..712"`                      // Retorno                               613..713 X(100)
}

func (c CreditAssessment) String() (string, error) {
	return writer.Marshal(c, 713)
}
