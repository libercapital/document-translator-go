package brf240

import (
	"github.com/libercapital/document-translator-go/internal/parser"
)

const (
	kindPosition        = 7
	segmentPosition     = 13
	instructionPosition = 15
	instructionLength   = 2
)

var objectKind = map[string]interface{}{
	"0": new(BillingFileHeader),
	"1": new(BillingBatchHeader),
	"5": new(BillingBatchTrailer),
	"9": new(BillingFileTrailer),
}

var parseObjectFunc = func(line string) interface{} {

	kind := line[kindPosition : kindPosition+1]

	if parseObject, ok := objectKind[kind]; ok {
		return parseObject
	}

	var segment = line[segmentPosition : segmentPosition+1]

	var instruction = line[instructionPosition : instructionPosition+instructionLength]

	if segment == "A" && instruction == "12" {
		return new(BillingSegmentAReceipt)
	}

	return new(BillingSegmentA)
}

func Parse(line string) (interface{}, error) {
	return parser.LineTo(
		line,
		parseObjectFunc,
	)
}
