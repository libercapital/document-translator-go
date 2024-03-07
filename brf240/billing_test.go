package brf240_test

import (
	"testing"
	"time"

	"github.com/libercapital/document-translator-go/brf240"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestBillingFileHeaderParse(t *testing.T) {
	file_header := "35300000         272493216000147003320500085000000650189370000005361516 BRF S/A                       Banco Santander                         20306201921310000589206006250                                                                     "

	expected_header := brf240.BillingFileHeader{
		BankCode:         "353",
		BatchNumber:      0,
		RegistryKind:     0,
		KindBuyer:        2,
		BuyerDocument:    "72493216000147",
		ContractNumber:   "00332050008500000065",
		Agency:           "1893",
		AgencyCd:         "7",
		Account:          "536151",
		AccountCd:        "6",
		CheckDigit:       "",
		BuyerName:        "BRF S/A",
		BankName:         "Banco Santander",
		FileKind:         2,
		FileDate:         time.Date(2019, time.June, 3, 0, 0, 0, 0, time.UTC).Round(0),
		FileTime:         time.Date(0, time.January, 1, 21, 31, 0, 0, time.UTC).Round(0),
		FileDateTime:     time.Date(2019, time.June, 3, 21, 31, 0, 0, time.UTC).Round(0),
		SequentialNumber: 5892,
		LayoutVersion:    "060",
		RecordDensity:    6250,
		BankReserved:     "",
		BuyerReserved:    "",
	}
	parsed, err := brf240.Parse(file_header)

	assert.NoError(t, err)
	assert.Equal(t, expected_header, parsed)
}

func TestBillingBatchHeaderParse(t *testing.T) {
	file_batch_header := "35300011C2003060 272493216000147003320500085000000650189370000005361516 BRF S/A                       TITULO DISPONIVEL PARA NEGOCIACAO       0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000          "

	expected_batch_header := brf240.BillingBatchHeader{
		BankCode:           "353",
		BatchNumber:        1,
		RegistryKind:       1,
		OperationKind:      "C",
		ServiceKind:        20,
		ReleaseKind:        3,
		BatchLayoutVersion: "060",
		KindBuyer:          2,
		BuyerDocument:      "72493216000147",
		ContractNumber:     "00332050008500000065",
		Agency:             "1893",
		AgencyCd:           "7",
		Account:            "536151",
		AccountCd:          "6",
		CheckDigit:         "",
		BuyerName:          "BRF S/A",
		GenericMessage:     "TITULO DISPONIVEL PARA NEGOCIACAO",
		AddressStreet:      "000000000000000000000000000000",
		AddressNumber:      0,
		AddressComplement:  "000000000000000",
		AddressCity:        "00000000000000000000",
		AddressZipCode:     0,
		AddressState:       "00",
		PaymentMethod:      "00",
		Occurrence:         "",
	}
	parsed, err := brf240.Parse(file_batch_header)

	assert.NoError(t, err)
	assert.Equal(t, expected_batch_header, parsed)
}

func TestBillingBillingSegmentAParse(t *testing.T) {
	file_segment_a := "3530001300001A000FORNECEDOR 1                        2           3648853400015600033000003808      130023471       000014000-1-00103062019030620190000000000000009523570000000000000000000000000000000000000033PS250051005512682019001          "
	expected_payment_value, _ := decimal.NewFromString("9523.57")
	expected_discount_value, _ := decimal.NewFromString("0.00")
	expected_financing_value, _ := decimal.NewFromString("0.00")
	expected_discount_rate, _ := decimal.NewFromString("0.0000")

	expected_segment_a := brf240.BillingSegmentA{
		BankCode:              "353",
		BatchNumber:           1,
		RegistryKind:          3,
		BatchSequentialNumber: 1,
		SegmentKind:           "A",
		ActionKind:            0,
		ActionInstructionKind: 0,
		VendorName:            "FORNECEDOR 1",
		DocumentKind:          2,
		FinancingDate:         "",
		Document:              "36488534000156",
		VendorBankCode:        "033",
		VendorAgency:          "3808",
		VendorAgencyCd:        "",
		VendorAccount:         "13002347",
		VendorAccountCd:       "1",
		PaymentNumber:         "000014000-1-001",
		IssueDate:             time.Date(2019, time.June, 3, 0, 0, 0, 0, time.UTC).Round(0),
		DueDate:               time.Date(2019, time.June, 3, 0, 0, 0, 0, time.UTC).Round(0),
		PaymentValue:          expected_payment_value,
		DiscountValue:         expected_discount_value,
		FinancingValue:        expected_financing_value,
		DiscountRate:          expected_discount_rate,
		ReferenceNumberPrefix: "000033PS",
		ReferenceNumber:       "250051005512682019001",
		Occurrence:            "",
	}
	parsed, err := brf240.Parse(file_segment_a)

	assert.NoError(t, err)
	assert.Equal(t, expected_segment_a, parsed)
}

func TestBillingSegmentAReceiptParse(t *testing.T) {
	file_segment_a := "BRF0001300001A112G10 TRANSPORTES LTDA                2           0756916100049200000000000000         732807     000085997-001-001180920200000000000000000000000000000000000000000000000000059454850255425BDBRFSA25005102375484201900112          "

	expected_segment_a := brf240.BillingSegmentAReceipt{
		BankCode:              "BRF",
		BatchNumber:           1,
		RegistryKind:          3,
		BatchSequentialNumber: 1,
		SegmentKind:           "A",
		ActionKind:            1,
		ActionInstructionKind: 12,
		VendorName:            "G10 TRANSPORTES LTDA",
		DocumentKind:          2,
		FinancingDate:         "",
		Document:              "07569161000492",
		VendorBankCode:        "000",
		VendorAgency:          "",
		VendorAgencyCd:        "",
		VendorAccount:         "73280",
		VendorAccountCd:       "7",
		PaymentNumber:         "000085997-001-001",
		LiquidationDate:       time.Date(2020, time.September, 18, 0, 0, 0, 0, time.UTC).Round(0),
		ProtocolNumber:        "0000000000000000000000000000000000000000000000000059454850255425BD",
		ReferenceNumber:       "250051023754842019001",
		ReferenceNumberPrefix: "BRFSA",
		Ocurrence:             "12",
	}

	parsed, err := brf240.Parse(file_segment_a)
	t.Log(err)

	assert.NoError(t, err)
	assert.Equal(t, expected_segment_a, parsed)
}

func TestBillingBillingSegmentY52Parse(t *testing.T) {
	file_segment_y52 := "2370001300001Y 00520000000000120510000000004431823009202341230900766315001035570200000120511001205197                                                                                                                                          "

	expected_segment_y52 := brf240.BillingSegmentY52{
		BankCode:              "237",
		BatchNumber:           1,
		RegistryKind:          3,
		BatchSequentialNumber: 1,
		SegmentKind:           "Y",
		ActionInstructionKind: 0,
		OptionalRegistryId:    "52",
		FiscalDocumentNumber1: "12051",
		FiscalDocumentValue1:  "443182",
		FiscalDocumentDate1:   "30092023",
		FiscalDocumentKey1:    "41230900766315001035570200000120511001205197",
		FiscalDocumentNumber2: "",
		FiscalDocumentValue2:  "",
		FiscalDocumentDate2:   "",
		FiscalDocumentKey2:    "",
	}
	parsed, err := brf240.Parse(file_segment_y52)

	assert.NoError(t, err)
	assert.Equal(t, expected_segment_y52, parsed)
}

func TestBillingBatchTrailerParse(t *testing.T) {
	file_trailer := "35300015         035193000000040420170731000000000000000000                                                                                                                                                                                     "
	expected_value_amount, _ := decimal.NewFromString("404201707.31")

	expected_file_trailer := brf240.BillingBatchTrailer{
		BankCode:                "353",
		BatchNumber:             1,
		RegistryKind:            5,
		QuantityRegistries:      35193,
		ValueAmount:             expected_value_amount,
		CurrencyQuantity:        0,
		DebitNotificationNumber: "",
		Occurrence:              "",
	}
	parsed, err := brf240.Parse(file_trailer)

	assert.NoError(t, err)
	assert.Equal(t, expected_file_trailer, parsed)
}
func TestBillingFileTrailerParse(t *testing.T) {
	file_trailer := "35399999         000001035195                                                                                                                                                                                                                   "

	expected_file_trailer := brf240.BillingFileTrailer{
		BankCode:             "353",
		BatchNumber:          9999,
		RegistryKind:         9,
		BatchesQuantity:      1,
		FileRegistryQuantity: 35195,
	}
	parsed, err := brf240.Parse(file_trailer)

	assert.NoError(t, err)
	assert.Equal(t, expected_file_trailer, parsed)
}
