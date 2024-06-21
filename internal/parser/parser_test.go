package parser

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	documenttranslator "github.com/libercapital/document-translator-go"
	"github.com/libercapital/document-translator-go/internal/utils"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_checkDeliminatorSize(t *testing.T) {
	type TestStruct struct {
		Field1 string `translator:"part:0..2"`
		Field2 string `translator:"part:3..5"`
		Field3 string `translator:"part:6..8"`
	}

	tests := []struct {
		name        string
		line        string
		parseOpt    ParseOpt
		valueOf     reflect.Value
		expectedErr error
	}{
		{
			name: "should validated without error",
			line: "123456789",
			parseOpt: ParseOpt{
				Params: []ParseParams{
					{Deliminator: []int{0, 2}},
					{Deliminator: []int{3, 5}},
					{Deliminator: []int{6, 8}},
				},
			},
			valueOf:     reflect.ValueOf(TestStruct{}),
			expectedErr: nil,
		},
		{
			name: "should throw a documenttranslator.ErrParseShorterThenDeliminator error",
			line: "123456",
			parseOpt: ParseOpt{
				Params: []ParseParams{
					{Deliminator: []int{0, 2}},
					{Deliminator: []int{3, 5}},
					{Deliminator: []int{6, 8}},
				},
			},
			valueOf:     reflect.ValueOf(TestStruct{}),
			expectedErr: documenttranslator.ErrParseShorterThenDeliminator,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDeliminatorSize(tt.line, tt.parseOpt, tt.valueOf)

			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func Test_validateKindAndSegment(t *testing.T) {
	type TestStruct struct {
		Field1 int    `translator:"kind:1"`
		Field2 string `translator:"segment:segment_value"`
	}

	type TestStructErrKindInt struct {
		Field1 int `translator:"kind:1"`
	}

	type TestStructErrKindString struct {
		Field1 string `translator:"kind:1"`
	}

	type TestStructErrSegmentInt struct {
		Field1 int `translator:"kind:1"`
		Field2 int `translator:"segment:segment_value"`
	}

	type TestStructErrSegmentInconsistencyKindInt struct {
		Field1 int    `translator:"kind:1"`
		Field2 string `translator:"segment:segment_value"`
	}

	type TestStructErrSegmentInconsistencyKindString struct {
		Field1 string `translator:"kind:1"`
		Field2 string `translator:"segment:segment_value"`
	}

	tests := []struct {
		name        string
		parseOpt    ParseOpt
		valueOf     reflect.Value
		expectedErr error
	}{
		{
			name: "should validated without error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(1),
					Value:      "segment_value",
				},
			},
			valueOf: reflect.ValueOf(TestStruct{Field1: 1, Field2: "segment_value"}),
		},
		{
			name: "should throw a documenttranslator.ErrKindInconsistency error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "2",
				},
			},
			valueOf:     reflect.ValueOf(TestStructErrKindInt{Field1: 1}),
			expectedErr: documenttranslator.ErrKindInconsistency,
		},
		{
			name: "should throw a documenttranslator.ErrKindInconsistency error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "2",
				},
			},
			valueOf:     reflect.ValueOf(TestStructErrKindString{Field1: "1"}),
			expectedErr: documenttranslator.ErrKindInconsistency,
		},
		{
			name: "should throw a documenttranslator.ErrSegmentMustBeString error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(1),
					Value:      "segment_value",
				},
			},
			valueOf:     reflect.ValueOf(TestStructErrSegmentInt{Field1: 1}),
			expectedErr: documenttranslator.ErrSegmentMustBeString,
		},
		{
			name: "should throw string a documenttranslator.ErrSegmentInconsistency error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(1),
					Value:      "segment_value",
				},
			},
			valueOf:     reflect.ValueOf(TestStructErrSegmentInconsistencyKindInt{Field1: 1}),
			expectedErr: documenttranslator.ErrSegmentInconsistency,
		},
		{
			name: "should throw string a documenttranslator.ErrSegmentInconsistency error",
			parseOpt: ParseOpt{
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(0),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(1),
					Value:      "segment_value",
				},
			},
			valueOf:     reflect.ValueOf(TestStructErrSegmentInconsistencyKindString{Field1: "1"}),
			expectedErr: documenttranslator.ErrSegmentInconsistency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateKindAndSegment(tt.parseOpt, tt.valueOf)

			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func Test_extractTags(t *testing.T) {
	type TestStruct struct {
		Field1 int    `translator:"part:0..5"`
		Field2 string `translator:"timeParse:02012006150405"`
		Field5 int    `translator:"precision:4"`
		Field6 string `translator:"prefixFrom:PS,PA,SP,SA,EN,DM"`
		Field7 int    `translator:"splitAfter:PS,PA,SP,SA,EN,DM"`
		Field3 int    `translator:"kind:1"`
		Field4 string `translator:"segment:segment_value"`
	}

	type TestStructErrorPart struct {
		Field1 int `translator:"part:0..B"`
	}
	type TestStructErrorPrecision struct {
		Field1 int `translator:"precision:C"`
	}

	type args struct {
		structTagged reflect.Type
	}
	tests := []struct {
		name         string
		args         args
		wantParseOpt ParseOpt
		wantErr      bool
		wantErrType  error
	}{
		{
			name: "should parse without error",
			args: args{
				structTagged: reflect.TypeOf(TestStruct{}),
			},
			wantParseOpt: ParseOpt{
				Params: []ParseParams{
					{
						Deliminator: []int{0, 5},
					},
					{
						TimeParse: "02012006150405",
					},
					{
						Precision: 4,
					},
					{
						PrefixFrom: []string{"PS", "PA", "SP", "SA", "EN", "DM"},
					},
					{
						SplitAfter: []string{"PS", "PA", "SP", "SA", "EN", "DM"},
					},
					{},
					{},
				},
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(5),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(6),
					Value:      "segment_value",
				},
			},
		},
		{
			name: "should throw an error on field deliminator",
			args: args{
				structTagged: reflect.TypeOf(TestStructErrorPart{}),
			},
			wantParseOpt: ParseOpt{
				Params: []ParseParams{
					{
						Deliminator: []int{0, 5},
					},
					{},
					{},
					{},
					{},
					{},
					{},
				},
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(5),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(6),
					Value:      "segment_value",
				},
			},
			wantErr:     true,
			wantErrType: strconv.ErrSyntax,
		},
		{
			name: "should throw an error on field precision",
			args: args{
				structTagged: reflect.TypeOf(TestStructErrorPrecision{}),
			},
			wantParseOpt: ParseOpt{
				Params: []ParseParams{
					{},
					{},
					{
						Precision: 4,
					},
					{},
					{},
					{},
					{},
				},
				Kind: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(5),
					Value:      "1",
				},
				Segment: struct {
					FieldIndex *int
					Value      string
				}{
					FieldIndex: utils.PtrAny(6),
					Value:      "segment_value",
				},
			},
			wantErr:     true,
			wantErrType: strconv.ErrSyntax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParseOpt, err := extractTags(tt.args.structTagged)

			if assert.ErrorIs(t, err, tt.wantErrType) {
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantParseOpt, gotParseOpt)
		})
	}
}

func Test_parseLine(t *testing.T) {
	type TestStruct struct {
		Field1 int             `translator:"part:0..2;"`
		Field2 time.Time       `translator:"part:3..16;timeParse:02012006150405"`
		Field3 decimal.Decimal `translator:"part:17..25;precision:2"`
		Field4 string          `translator:"part:26..35;prefixFrom:PS,PA,SP,SA,EN,DM"`
		Field5 string          `translator:"part:26..35;splitAfter:PS,PA,SP,SA,EN,DM"`
		Field6 int             `translator:"part:36..36;kind:1"`
		Field7 string          `translator:"part:37..39;segment:abc"`
	}

	type TestStructIntError struct {
		Field1 int `translator:"part:0..2;"`
	}

	type TestStructTimeError struct {
		Field1 time.Time `translator:"part:0..13;timeParse:02012006150405"`
	}

	type TestStructDecimalError struct {
		Field1 decimal.Decimal `translator:"part:0..5;precision:2"`
	}

	type args struct {
		line string
		v    interface{}
	}
	tests := []struct {
		name        string
		args        args
		expected    interface{}
		wantErr     bool
		expectedErr error
	}{
		{
			name: "should parse the line without error",
			args: args{
				line: "012020120061504050047864547t8946PS891abc",
				v:    &TestStruct{},
			},
			expected: TestStruct{
				Field1: 12,
				Field2: func() time.Time {
					time, _ := time.Parse("02012006150405", "02012006150405")
					return time
				}(),
				Field3: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("47864.54")
					return dec
				}(),
				Field4: "7t8946PS",
				Field5: "89",
				Field6: 1,
				Field7: "abc",
			},
		},
		{
			name: "should parse the line with strconv error",
			args: args{
				line: "ABC",
				v:    &TestStructIntError{},
			},
			wantErr:     true,
			expectedErr: strconv.ErrSyntax,
		},
		{
			name: "should parse the line with time error",
			args: args{
				line: "45012006150405",
				v:    &TestStructTimeError{},
			},
			wantErr: true,
		},
		{
			name: "should parse the line with decimal error",
			args: args{
				line: "A12345",
				v:    &TestStructDecimalError{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeOf := reflect.TypeOf(tt.args.v).Elem()
			valueOf := reflect.ValueOf(tt.args.v).Elem()

			params, err := extractTags(typeOf)

			if !tt.wantErr && !assert.NoError(t, err) {
				t.Fail()
			}

			err = parseLine(tt.args.line, params, valueOf, typeOf)

			if !tt.wantErr && !assert.NoError(t, err) {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && (tt.expectedErr == nil || assert.ErrorIs(t, err, tt.expectedErr)) {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.expected, reflect.ValueOf(tt.args.v).Elem().Interface())
		})
	}
}

func TestLineTo(t *testing.T) {
	type TestStruct struct {
		Field1 int             `translator:"part:0..2;"`
		Field2 time.Time       `translator:"part:3..16;timeParse:02012006150405"`
		Field3 decimal.Decimal `translator:"part:17..25;precision:2"`
		Field4 string          `translator:"part:26..35;prefixFrom:PS,PA,SP,SA,EN,DM"`
		Field5 string          `translator:"part:26..35;splitAfter:PS,PA,SP,SA,EN,DM"`
		Field6 int             `translator:"part:36..36;kind:1"`
		Field7 string          `translator:"part:37..39;segment:abc"`
	}

	type args struct {
		line            string
		kindPosition    int
		parseObjectFunc ParseObjectFunction
	}
	tests := []struct {
		name             string
		args             args
		wantStructParsed interface{}
		wantErr          bool
	}{
		{
			args: args{
				line:         "012020120061504050047864547t8946PS891abc",
				kindPosition: 0,
				parseObjectFunc: func(line string) interface{} {
					return new(TestStruct)
				},
			},
			wantStructParsed: TestStruct{
				Field1: 12,
				Field2: func() time.Time {
					time, _ := time.Parse("02012006150405", "02012006150405")
					return time
				}(),
				Field3: func() decimal.Decimal {
					dec, _ := decimal.NewFromString("47864.54")
					return dec
				}(),
				Field4: "7t8946PS",
				Field5: "89",
				Field6: 1,
				Field7: "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStructParsed, err := LineTo(tt.args.line, tt.args.parseObjectFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStructParsed, tt.wantStructParsed) {
				t.Errorf("LineTo() = %v, want %v", gotStructParsed, tt.wantStructParsed)
			}
		})
	}
}

func TestUnicodeSpecialCharacter(t *testing.T) {
	type TestStruct struct {
		Field1 string `translator:"part:0..0;"`
		Field2 string `translator:"part:1..1;kind:1;"`
	}

	type args struct {
		line            string
		kindPosition    int
		parseObjectFunc ParseObjectFunction
	}
	tests := []struct {
		name             string
		args             args
		wantStructParsed interface{}
		wantErr          bool
	}{
		{
			args: args{
				line:         "�1",
				kindPosition: 0,
				parseObjectFunc: func(line string) interface{} {
					return new(TestStruct)
				},
			},
			wantStructParsed: TestStruct{
				Field1: "",
				Field2: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStructParsed, err := LineTo(tt.args.line, tt.args.parseObjectFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStructParsed, tt.wantStructParsed) {
				t.Errorf("LineTo() = %v, want %v", gotStructParsed, tt.wantStructParsed)
			}
		})
	}
}
