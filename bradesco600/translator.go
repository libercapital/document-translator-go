package bradesco600

import (
	"github.com/libercapital/document-translator-go/internal/parser"
)

func Parse(line string) (CreditAssessment, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(CreditAssessment)
		},
	)

	if err != nil {
		return CreditAssessment{}, err
	}

	return data.(CreditAssessment), nil
}
