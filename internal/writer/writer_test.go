package writer

import (
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_serializerOpts_String(t *testing.T) {
	var tests = []struct {
		name          string
		serializerOpt serializerOpt
		want          string
	}{
		{
			name: "test serialize simple serializerOpt",
			serializerOpt: serializerOpt{
				Length: 11,
				Params: []serializerParams{
					{
						FillType:    FillString,
						Deliminator: []int{0, 4},
						Value:       "Value",
					},
					{
						FillType:    FillString,
						Deliminator: []int{5, 9},
						Value:       "Value",
					},
				},
			},
			want: "VALUEVALUE ",
		},
		{
			name: "test serialize multiple params serializerOpt",
			serializerOpt: serializerOpt{
				Length: 20,
				Params: []serializerParams{
					{
						FillType:    FillString,
						Deliminator: []int{15, 19},
						Value:       "TESTEEEEEEEE",
					},
					{
						FillType:    FillNumber,
						Deliminator: []int{0, 10},
						Value:       "100",
					},
					{
						FillType:    FillNumber,
						Deliminator: []int{11, 14},
						Value:       "200",
					},
				},
			},
			want: "000000001000200TESTE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializedString := tt.serializerOpt.String()
			assert.Equal(t, tt.want, serializedString)
		})
	}
}

func Test_structToString(t *testing.T) {
	date, _ := time.Parse("02012006", "10102023")
	var tests = []struct {
		name    string
		value   interface{}
		length  int
		wantErr error
		want    string
	}{
		{
			name: "successful simple struct to string",
			value: struct {
				BankCode int    `translator:"part:0..10"`
				BankName string `translator:"part:11..19"`
			}{
				BankCode: 269,
				BankName: "HSBC",
			},
			length:  20,
			wantErr: nil,
			want:    "00000000269HSBC     ",
		},
		{
			name: "successful complex struct to string",
			value: struct {
				Segment          string          `translator:"part:0..1;segment:T"`
				Kind             string          `translator:"part:2..2;kind:1"`
				BankCode         int             `translator:"part:3..10"`
				BankName         string          `translator:"part:11..20;align:right"`
				Date             time.Time       `translator:"part:21..28;timeParse:02012006"`
				TransactionValue decimal.Decimal `translator:"part:29..39"`
			}{
				Segment:  "T",
				Kind:     "1",
				BankCode: 269,
				BankName: "HSBC",
				Date:     date,
				TransactionValue: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("2.53")
					return dec
				}(),
			},
			length:  40,
			wantErr: nil,
			want:    "T 100000269      HSBC1010202300000000253",
		},
		{
			name: "successful complex struct to string",
			value: struct {
				TransactionValue decimal.Decimal `translator:"part:0..10"`
			}{
				TransactionValue: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("2")
					return dec
				}(),
			},
			length:  11,
			wantErr: nil,
			want:    "00000000200",
		},
		{
			name: "successful struct to string with zero date",
			value: struct {
				ZeroDate time.Time `translator:"part:0..7;timeParse:02012006"`
			}{
				ZeroDate: time.Time{},
			},
			length:  8,
			wantErr: nil,
			want:    "        ", // Esperado formato da data zero
		},
		{
			name: "successful struct to string with value with precision",
			value: struct {
				Rate decimal.Decimal `translator:"part:0..4;precision:4"`
			}{
				Rate: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("0.00644927")
					return dec
				}(),
			},
			length:  5,
			wantErr: nil,
			want:    "00064",
		},
		{
			name: "error when struct to string with wrong precision",
			value: struct {
				Rate decimal.Decimal `translator:"part:0..4;precision:TEST"`
			}{
				Rate: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("0.0064")
					return dec
				}(),
			},
			length:  5,
			wantErr: &strconv.NumError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structSerialized, err := structToString(tt.value, tt.length)
			if err == nil {
				assert.Equal(t, tt.want, structSerialized)
				assert.Nil(t, err)
			}
			if err != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.wantErr, err)
			}
		})
	}
}
