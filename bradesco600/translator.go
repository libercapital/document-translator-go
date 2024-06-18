package bradesco600

import (
	"github.com/libercapital/document-translator-go/internal/parser"
)

// const (
//
//	kindPosition             = 7
//	segmentPosition          = 13
//	instructionPosition      = 15
//	optionalRegistryPosition = 17
//	instructionLength        = 2
//
// )

var parseObjectFunc = func(line string) interface{} {
	//
	// if parseObject, ok := objectKind[kind]; ok {
	// 	return parseObject
	// }
	//
	// var segment = line[segmentPosition : segmentPosition+1]
	//
	// var instruction = line[instructionPosition : instructionPosition+instructionLength]
	//
	// if segment == "A" && instruction == "12" {
	// 	return new(BillingSegmentAReceipt)
	// }
	//
	// var optionalRegistry = line[optionalRegistryPosition : optionalRegistryPosition+instructionLength]
	//
	// if segment == "Y" && optionalRegistry == "52" {
	// 	return new(BillingSegmentY52)
	// }
	//
	// return new(BillingSegmentA)
	return new(CreditAssessment)
}

func Parse(line string) (interface{}, error) {
	return parser.LineTo(
		line,
		parseObjectFunc,
	)
}
