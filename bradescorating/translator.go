package bradescorating

import "github.com/libercapital/document-translator-go/internal/parser"

func Parse(line string) (interface{}, error) {
	return parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Rating)
		},
	)
}
