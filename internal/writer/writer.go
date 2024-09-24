package writer

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/libercapital/document-translator-go/internal/utils"
	"github.com/shopspring/decimal"
)

type fillType string

const (
	FillString fillType = "STRING"
	FillNumber fillType = "NUMBER"
)

type serializerOpt struct {
	Length int
	Params []serializerParams // Params contains the serializer parameters for each field.
}

type serializerParams struct {
	FillType    fillType
	Deliminator []int
	TimeParse   string
	Value       string
	Align       string
	Precision   int // Precision specifies the decimal precision for the field.
}

func (s *serializerOpt) String() string {
	var line = utils.EmptyArray(s.Length)
	for _, param := range s.Params {
		length := (param.Deliminator[1] - param.Deliminator[0] + 1)
		var data = fillValue(param.Value, length, param.FillType, param.Align)
		copy(line[param.Deliminator[0]:], data[:length])
	}
	return string(line)
}

func extractValues(structValue reflect.Value, opt *serializerOpt) error {

	for i := 0; i < structValue.NumField(); i++ {
		opt.Params[i].Value = getValue(&opt.Params[i], structValue.Field(i))
	}
	return nil

}

func extractTags(structTagged reflect.Type) (serializerOpt serializerOpt, err error) {
	serializerOpt.Params = make([]serializerParams, structTagged.NumField())

	for i := 0; i < structTagged.NumField(); i++ {
		f := structTagged.Field(i)

		translatorTags := strings.Split(f.Tag.Get("translator"), ";")

		for _, rule := range translatorTags {
			parts := strings.Split(rule, ";")
			defaultPrecision := true

			for _, part := range parts {
				keyValuePair := strings.Split(part, ":")
				key := keyValuePair[0]
				value := keyValuePair[1]

				switch key {
				case "part":
					convertToInt, err := utils.ConvertStringToIntSlice(strings.Split(value, "..")...)

					if err != nil {
						return serializerOpt, err
					}

					serializerOpt.Params[i].Deliminator = convertToInt
				case "timeParse":
					serializerOpt.Params[i].TimeParse = value
				case "align":
					serializerOpt.Params[i].Align = value
				case "precision":
					precision, err := strconv.Atoi(keyValuePair[1])
					if err != nil {
						return serializerOpt, err
					}
					serializerOpt.Params[i].Precision = precision
					defaultPrecision = false
				}
			}

			if defaultPrecision {
				serializerOpt.Params[i].Precision = 2
			}
		}
	}

	return
}

func structToString(value interface{}, length int) (string, error) {

	serializerOpts, err := extractTags(reflect.TypeOf(value))
	if err != nil {
		return "", err
	}
	serializerOpts.Length = length

	if err := extractValues(reflect.ValueOf(value), &serializerOpts); err != nil {
		return "", err
	}

	return serializerOpts.String(), nil

}

func getValue(param *serializerParams, structValue reflect.Value) string {

	switch structValue.Interface().(type) {
	case int, int32, int64:
		param.FillType = FillNumber
		return strconv.FormatInt(structValue.Int(), 10)

	case string:
		param.FillType = FillString
		return structValue.String()

	case time.Time:
		param.FillType = FillString
		timeValue := structValue.Interface().(time.Time)
		if timeValue.IsZero() {
			return ""
		}
		return timeValue.Format(param.TimeParse)

	case decimal.Decimal:
		param.FillType = FillNumber
		decimal := structValue.Interface().(decimal.Decimal)
		return strings.ReplaceAll(decimal.StringFixed(int32(param.Precision)), ".", "")
	}

	return ""
}

func fillValue(value string, length int, fillType fillType, align string) []byte {
	value = strings.ToUpper(value)
	if len(value) > length {
		value = value[:length]
	}

	paddingLength := length - len(value)
	var padding string

	if fillType == FillNumber {
		padding = strings.Repeat("0", paddingLength)
		return []byte(padding + value) // Números: preenchidos com zeros à esquerda
	}

	padding = strings.Repeat(" ", paddingLength)
	if align == "right" {
		return []byte(padding + value)
	}

	return []byte(value + padding)
}

func Marshal(value interface{}, length int) (string, error) {
	return structToString(value, length)
}
