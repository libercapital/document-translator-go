package bradesco600

import (
	"github.com/libercapital/document-translator-go/internal/parser"
)

var parseObjectFunc = func(line string) interface{} {
	return new(CreditAssessment)
}

func Parse(line string) (interface{}, error) {
	return parser.LineTo(
		line,
		parseObjectFunc,
	)
}
