package bradescorating

import (
	"github.com/libercapital/document-translator-go/internal/writer"
	"github.com/shopspring/decimal"
)

type Rating struct {
	DocumentNumber string          `translator:"part:0..14"`
	Amount         decimal.Decimal `translator:"part:15..29;precision:2"`
}

func (c Rating) String() (string, error) {
	return writer.Marshal(c, 30)
}
