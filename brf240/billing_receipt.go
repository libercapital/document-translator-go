package brf240

import "time"

type BillingSegmentAReceipt struct {
	BankCode              string    `translator:"part:0..2"`                                  //Código do Banco                         001..003   9(003)
	BatchNumber           int       `translator:"part:3..6"`                                  //Lote de Serviço                         004..007   9(004)
	RegistryKind          int       `translator:"part:7..7;kind:3"`                           //Tipo de Registro                        008..008   9(001)
	BatchSequentialNumber int       `translator:"part:8..12"`                                 //Número Seqüencial do Registro no Lote   009..013   9(005)
	SegmentKind           string    `translator:"part:13..13;segment:A"`                      //Código Segmento do Registro Detalhe     014..014   X(001)
	ActionKind            int       `translator:"part:14..14"`                                //Tipo de Movimento                       015..015   9(001)
	ActionInstructionKind int       `translator:"part:15..16"`                                //Código da Instrução para Movimento      016..017   9(002)
	VendorName            string    `translator:"part:17..52"`                                //Nome do Fornecedor                      018..053   X(036)
	DocumentKind          int       `translator:"part:53..53"`                                //Se CNPJ = "2". Se CPF = "1"             054..054   9(001)
	FinancingDate         string    `translator:"part:54..61"`                                //Data de financiamento                   055..062   X(008)
	Document              string    `translator:"part:62..78"`                                //CPNJ ou CPF                             063..079   9(017)
	VendorBankCode        string    `translator:"part:79..83;lastDigits:3"`                   //Número do Banco Fornecedor              080..084   9(005)
	VendorAgency          string    `translator:"part:84..92;clearZeroLeft"`                  //Agência do Banco Fornecedor             085..093   9(009)
	VendorAgencyCd        string    `translator:"part:93..93"`                                //Dígito da Agência                       094..094   X(001)
	VendorAccount         string    `translator:"part:94..106;clearZeroLeft"`                 //Conta Bancária                          095..107   9(013)
	VendorAccountCd       string    `translator:"part:107..107"`                              //Dígito Verificador da Conta             108..108   X(001)
	PaymentNumber         string    `translator:"part:108..129"`                              //Número da Nota Fiscal/Fatura            109..130   X(022)
	LiquidationDate       time.Time `translator:"part:130..137;timeParse:02012006"`           //Data da Liquidação                      131..138   X(008)
	ProtocolNumber        string    `translator:"part:138..203"`                              //Protocolo Bancário                      139..202   X(063)
	ReferenceNumberPrefix string    `translator:"part:204..229;prefixFrom:PS,PA,SP,SA,EN,DM"` //Numero de referência                    203..230   X(026)
	ReferenceNumber       string    `translator:"part:204..229;splitAfter:PS,PA,SP,SA,EN,DM"` //Numero de referência                    203..230   X(026)
	Ocurrence             string    `translator:"part:230..231"`                              //Status da Partida/Código de ocorrência  231..232   X(002)
}
