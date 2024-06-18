package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	documenttranslator "github.com/libercapital/document-translator-go"
	"github.com/libercapital/document-translator-go/internal/wraperrors"
	"github.com/shopspring/decimal"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type ParseObjectFunction func(line string) interface{}

type ParseOpt struct {
	Params []ParseParams // Params contains the parsing parameters for each field.
	Kind   struct {      // Kind represents the kind field options.
		FieldIndex *int   // FieldIndex holds the field index for the kind field.
		Value      string // Value is the expected value for the kind field.
	}
	Segment struct { // Segment represents the segment field options.
		FieldIndex *int   // FieldIndex holds the field index for the segment field.
		Value      string // Value is the expected value for the segment field.
	}
}

type ParseParams struct {
	Deliminator   []int    // Deliminator holds the delimiter indices for the field.
	TimeParse     string   // TimeParse specifies the time parsing format for the field.
	Precision     int      // Precision specifies the decimal precision for the field.
	PrefixFrom    []string // PrefixFrom specifies the prefix strings to extract from the field.
	SplitAfter    []string // SplitAfter specifies the suffix strings to split the field after.
	ClearZeroLeft string   // ClearZeroLeft specifies the parser to clear all zeros to the left of string.
	LastDigits    int      // LastDigits specifies the parser to extract only the N digits at the end of string.
}

// LineTo parses a line of text and returns the parsed struct corresponding to the kind value.
//
// Parameters:
// - line: The line of text to parse.
// - kindPosition: The position of the kind field in the line.
// - parseTo: A map of kind values to target structs.
//
// Returns:
// - structParsed: The parsed struct corresponding to the kind value.
// - err: An error if there was an issue during parsing or if the kind value is not found in the parseTo map.
//
// Example:
//
//	structParsed, err := LineTo(line, kindPosition, parseTo)
//	if err != nil {
//	    // Handle the error
//	}
func LineTo(line string, parseObjectFunc ParseObjectFunction) (structParsed interface{}, err error) {
	line, err = removeAccents(sanitize(line))

	if err != nil {
		return
	}

	parseObject := parseObjectFunc(line)

	typeOf := reflect.TypeOf(parseObject).Elem()
	valueOf := reflect.ValueOf(parseObject).Elem()

	parseOpt, err := extractTags(typeOf)

	if err != nil {
		return structParsed, err
	}

	if err = checkDeliminatorSize(line, parseOpt, valueOf); err != nil {
		return
	}

	if err = parseLine(line, parseOpt, valueOf, typeOf); err != nil {
		return
	}

	if err = validateKindAndSegment(parseOpt, valueOf); err != nil {
		return
	}

	return reflect.ValueOf(parseObject).Elem().Interface(), nil
}

// parseLine parses a line of input using the provided parse options and sets the corresponding values
// in the target struct using reflection.
//
// The line parameter represents the input line to be parsed.
//
// The parseOpt parameter contains the parsing options that define how to extract and set values from the input line.
//
// The valueOf parameter is a reflect.Value representing the target struct in which the parsed values will be set.
//
// The typeOf parameter is a reflect.Type representing the type of the target struct.
//
// The function iterates over the parse options and sets the values in the target struct based on the specified rules.
// It returns an error if any error occurs during the value setting process.
// The error is wrapped with additional context including the struct type and field name where the error occurred.
//
// Example usage:
//
//	var document Document
//	err := parseLine(inputLine, parseOpt, reflect.ValueOf(&document).Elem(), reflect.TypeOf(document))
//	if err != nil {
//	  fmt.Println("Error parsing line:", err)
//	}
func parseLine(line string, parseOpt ParseOpt, valueOf reflect.Value, typeOf reflect.Type) (err error) {
	for index, param := range parseOpt.Params {
		if err = setValues(line, valueOf.Field(index), param); err != nil {
			return wraperrors.NewErrWrap(err, fmt.Errorf("at struct %s and field %s", valueOf.Type(), typeOf.Field(index).Name))
		}
	}

	return
}

// setValues parses a line of text and assigns the parsed value to a reflect.Value based on the provided ParseParams.
//
// Parameters:
// - line: The line of text to parse.
// - v: A reflect.Value representing the target value to assign the parsed result.
// - param: The ParseParams containing parsing options for the value.
//
// Returns:
// - An error if there was an issue parsing the line or assigning the value.
// - nil if the line was successfully parsed and the value was assigned.
//
// Example:
//
//	err := parseLine(line, v, param)
//	if err != nil {
//	    // Handle the error
//	}
func setValues(line string, v reflect.Value, param ParseParams) error {
	value := line[int(param.Deliminator[0]) : param.Deliminator[1]+1]

	switch v.Interface().(type) {
	case int, int32, int64:
		valueInt, err := strconv.Atoi(strings.TrimSpace(value))

		if err != nil {
			return err
		}
		v.SetInt(int64(valueInt))
	case string:
		if len(param.PrefixFrom) > 0 {
			for _, prefixValue := range param.PrefixFrom {
				index := strings.Index(value, prefixValue)
				if index != -1 {
					result := value[:index+len(prefixValue)]
					v.SetString(strings.TrimSpace(result))
					return nil
				}
			}
		}

		if len(param.SplitAfter) > 0 {
			for _, sufixValue := range param.SplitAfter {
				index := strings.Index(value, sufixValue)
				if index != -1 {
					result := value[index+len(sufixValue):]
					v.SetString(strings.TrimSpace(result))
					return nil
				}
			}
		}

		if len(param.ClearZeroLeft) > 0 {
			trimmed := strings.TrimLeft(value, "0")

			v.SetString(strings.TrimSpace(trimmed))
			return nil
		}

		if param.LastDigits > 0 {
			substring := value[len(value)-param.LastDigits:]
			v.SetString(strings.TrimSpace(substring))
			return nil
		}

		value := strings.TrimSpace(value)

		v.SetString(value)
	case time.Time:
		var timeParsed time.Time
		var err error

		if dateValueIsEmpty(value) {
			timeParsed = time.Time{}
		} else {
			timeParsed, err = time.Parse(param.TimeParse, value)

			if err != nil {
				return err
			}

		}

		v.Set(reflect.ValueOf(timeParsed.UTC()))
	case decimal.Decimal:
		value = strings.TrimSpace(value)
		value = value[:len(value)-param.Precision] + "." + value[(len(value))-param.Precision:]

		valueDecimal, err := decimal.NewFromString(value)

		if err != nil {
			return err
		}

		v.Set(reflect.ValueOf(valueDecimal))
	}

	return nil
}

func dateValueIsEmpty(value string) bool {
	return strings.Trim(value, "0") == ""
}

// convertStringToIntSlice converts a slice of string values to a slice of integers.
//
// Parameters:
// - values: A variadic parameter representing the string values to convert.
//
// Returns:
// - ret: A slice of integers obtained from converting the string values.
// - err: An error if there was an issue converting the string values to integers.
//
// Example:
//
//	ret, err := toInt("1", "2", "3")
//	if err != nil {
//	    // Handle the error
//	}
func convertStringToIntSlice(values ...string) (ret []int, err error) {
	for _, value := range values {
		var intValue int

		if intValue, err = strconv.Atoi(value); err != nil {
			return
		}

		ret = append(ret, intValue)
	}

	return
}

// extractTags extracts and parses tags from the fields of a struct with tag annotations.
// It populates a ParseOpt struct based on the extracted tags.
//
// Parameters:
// - structTagged: The reflect.Type of the struct with tag annotations.
//
// Returns:
// - parseOpt: A ParseOpt struct containing the parsed tag information.
// - err: An error if there was an issue extracting or parsing the tags.
//
// Example:
//
//	parseOpt, err := extractTags(structTagged)
//	if err != nil {
//	    // Handle the error
//	}
func extractTags(structTagged reflect.Type) (parseOpt ParseOpt, err error) {
	parseOpt.Params = make([]ParseParams, structTagged.NumField())

	for i := 0; i < structTagged.NumField(); i++ {
		f := structTagged.Field(i)

		translatorTags := strings.Split(f.Tag.Get("translator"), ";")

		for _, rule := range translatorTags {
			parts := strings.Split(rule, ";")
			key := parts[0]

			if key == "" {
				continue
			}

			resultGroup := strings.Split(key, ":")

			switch resultGroup[0] {
			case "part":
				convertToInt, err := convertStringToIntSlice(strings.Split(resultGroup[1], "..")...)

				if err != nil {
					return parseOpt, err
				}

				parseOpt.Params[i].Deliminator = convertToInt
			case "clearZeroLeft":
				parseOpt.Params[i].ClearZeroLeft = resultGroup[0]
			case "lastDigits":
				nDigits, err := strconv.Atoi(resultGroup[1])
				if err != nil {
					return parseOpt, err
				}
				parseOpt.Params[i].LastDigits = nDigits
			case "timeParse":
				parseOpt.Params[i].TimeParse = resultGroup[1]
			case "kind":
				parseOpt.Kind.FieldIndex = new(int)
				*parseOpt.Kind.FieldIndex = i
				parseOpt.Kind.Value = resultGroup[1]
			case "segment":
				parseOpt.Segment.FieldIndex = new(int)
				*parseOpt.Segment.FieldIndex = i
				parseOpt.Segment.Value = resultGroup[1]
			case "precision":
				precision, err := strconv.Atoi(resultGroup[1])
				if err != nil {
					return parseOpt, err
				}
				parseOpt.Params[i].Precision = precision
			case "prefixFrom":
				parseOpt.Params[i].PrefixFrom = strings.Split(resultGroup[1], ",")
			case "splitAfter":
				parseOpt.Params[i].SplitAfter = strings.Split(resultGroup[1], ",")
			}
		}
	}

	return
}

// validateKindAndSegment validates the consistency of the kind and segment fields within a struct.
// It performs different validations based on the kind of the field.
//
// Parameters:
// - parseOpt: The ParseOpt struct containing the kind and segment information.
// - valueOf: A reflect.Value representing the struct type.
//
// Returns:
// - An error if the kind or segment fields are inconsistent with the provided options.
// - nil if the kind and segment fields are consistent.
//
// Example:
//
//	err := validateKindAndSegment(parseOpt, valueOf)
//	if err != nil {
//	    // Handle the error
//	}
func validateKindAndSegment(parseOpt ParseOpt, valueOf reflect.Value) (err error) {
	// Some documents don't have a kind.
	if parseOpt.Kind.Value == "" {
		return nil
	}

	var segmentField reflect.Value
	kindField := valueOf.Field(*parseOpt.Kind.FieldIndex)

	if parseOpt.Segment.FieldIndex != nil {
		segmentField = valueOf.Field(*parseOpt.Segment.FieldIndex)
	}

	switch kindField.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		parseKindValue := -1

		if parseOpt.Kind.Value != "" {
			parseKindValue, err = strconv.Atoi(parseOpt.Kind.Value)
			if err != nil {
				return
			}
		}

		if kindField.Int() != int64(parseKindValue) {
			return documenttranslator.ErrKindInconsistency
		}

		if parseOpt.Segment.FieldIndex != nil {
			if segmentField.Kind() == reflect.Int ||
				segmentField.Kind() == reflect.Int32 ||
				segmentField.Kind() == reflect.Int64 {
				return documenttranslator.ErrSegmentMustBeString
			}

			if segmentField.String() != parseOpt.Segment.Value {
				return documenttranslator.ErrSegmentInconsistency
			}
		}

	case reflect.String:
		if kindField.String() != parseOpt.Kind.Value {
			return documenttranslator.ErrKindInconsistency
		}

		if parseOpt.Segment.FieldIndex != nil && segmentField.String() != parseOpt.Segment.Value {
			return documenttranslator.ErrSegmentInconsistency
		}
	}

	return nil
}

// checkDeliminatorSize checks if the highest delimiter value in the ParseOpt struct is greater than the length of the given line.
// If the highest delimiter value is greater than the line length, it returns an error indicating that the line is shorter than the delimiter.
//
// Parameters:
// - line: The input line of text to be checked.
// - parseOpt: The ParseOpt struct containing delimiter information.
// - valueOf: A reflect.Value representing the struct type.
//
// Returns:
// - An error if the highest delimiter value is greater than the line length.
// - nil if the highest delimiter value is not greater than the line length.
//
// Example:
//
//	err := checkDeliminatorSize(line, parseOpt, valueOf)
//	if err != nil {
//	    // Handle the error
//	}
func checkDeliminatorSize(line string, parseOpt ParseOpt, valueOf reflect.Value) error {
	var highestDeliminator int

	for _, opt := range parseOpt.Params {
		if opt.Deliminator[1] > highestDeliminator {
			highestDeliminator = opt.Deliminator[1]
		}
	}

	if highestDeliminator > len(line) {
		return wraperrors.NewErrWrap(documenttranslator.ErrParseShorterThenDeliminator, fmt.Errorf("at struct %s", valueOf.Type()))
	}

	return nil
}

func removeAccents(raw string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	r, _, err := transform.String(t, raw)

	return r, err
}

func sanitize(raw string) string {
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return ' '
		}
		return r
	}, raw)
}
