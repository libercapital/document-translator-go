package documenttranslator

type FullFillPosition int

const (
	FullFillLeft FullFillPosition = iota
	FullFillRight
)

func FullFill(fillSize int, fillWith string, position FullFillPosition, data string) (ret string) {
	var fillet string

	for a := 0; a < fillSize; a++ {
		fillet += fillWith
	}

	if position == FullFillLeft {
		ret = fillet + data
	} else {
		ret = data + fillet
	}

	return
}
