package bradesco226

import "github.com/libercapital/document-translator-go/internal/parser"

const kindPosition = 0

type RegisterType string

const (
	RegisterTypeHeader      RegisterType = "1"
	RegisterTypeContract    RegisterType = "2"
	RegisterTypeBorrower    RegisterType = "3"
	RegisterTypeInstallment RegisterType = "4"
)

func Kind(line string) RegisterType {
	kind := line[kindPosition : kindPosition+1]

	switch kind {
	case "1":
		return RegisterTypeHeader
	case "2":
		return RegisterTypeContract
	case "3":
		return RegisterTypeBorrower
	default:
		return RegisterTypeInstallment

	}
}

func ParseContract(line string) (Contract, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Contract)
		},
	)

	if err != nil {
		return Contract{}, err
	}

	return data.(Contract), nil
}

func ParseBorrower(line string) (Borrower, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Borrower)
		},
	)

	if err != nil {
		return Borrower{}, err
	}

	return data.(Borrower), nil
}

func ParseInstallment(line string) (Installment, error) {
	data, err := parser.LineTo(
		line,
		func(line string) interface{} {
			return new(Installment)
		},
	)

	if err != nil {
		return Installment{}, err
	}

	return data.(Installment), nil
}
