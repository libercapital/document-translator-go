package getnetextrato

import (
	"errors"

	"github.com/libercapital/document-translator-go/internal/parser"
)

const kindPosition = 0

type RegisterType string

const (
	TipoRegistroHeader                RegisterType = "0"
	TipoRegistroResumoTransacional    RegisterType = "1"
	TipoRegistroAnaliticoTransacional RegisterType = "2"
	TipoRegistroAjusteFinanceiro      RegisterType = "3"
	TipoRegistroResumoFinanceiro      RegisterType = "5"
	TipoRegistroDetalheFinanceiro     RegisterType = "6"
	TipoRegistroTrailer               RegisterType = "9"
)

func Kind(line string) (RegisterType, error) {
	kind := line[kindPosition : kindPosition+1]

	switch kind {
	case "0":
		return TipoRegistroHeader, nil
	case "1":
		return TipoRegistroResumoTransacional, nil
	case "2":
		return TipoRegistroAnaliticoTransacional, nil
	case "3":
		return TipoRegistroAjusteFinanceiro, nil
	case "5":
		return TipoRegistroResumoFinanceiro, nil
	case "6":
		return TipoRegistroDetalheFinanceiro, nil
	case "9":
		return TipoRegistroTrailer, nil
	default:
		return RegisterType(""), errors.New("invalid register type")
	}
}

func ParseHeader(line string) (Header, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Header)
		},
	)

	if err != nil {
		return Header{}, err
	}

	return data.(Header), nil
}

func ParseResumoTransacional(line string) (ResumoTransacional, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(ResumoTransacional)
		},
	)

	if err != nil {
		return ResumoTransacional{}, err
	}

	return data.(ResumoTransacional), nil
}

func ParseAnaliticoTransacional(line string) (AnaliticoTransacional, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(AnaliticoTransacional)
		},
	)

	if err != nil {
		return AnaliticoTransacional{}, err
	}

	return data.(AnaliticoTransacional), nil
}

func ParseAjusteFinanceiro(line string) (AjusteFinanceiro, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(AjusteFinanceiro)
		},
	)

	if err != nil {
		return AjusteFinanceiro{}, err
	}

	return data.(AjusteFinanceiro), nil
}

func ParseResumoFinanceiro(line string) (ResumoFinanceiro, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(ResumoFinanceiro)
		},
	)

	if err != nil {
		return ResumoFinanceiro{}, err
	}

	return data.(ResumoFinanceiro), nil
}

func ParseDetalheFinanceiro(line string) (DetalheFinanceiro, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(DetalheFinanceiro)
		},
	)

	if err != nil {
		return DetalheFinanceiro{}, err
	}

	return data.(DetalheFinanceiro), nil
}

func ParseTrailer(line string) (Trailer, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Trailer)
		},
	)

	if err != nil {
		return Trailer{}, err
	}

	return data.(Trailer), nil
}
