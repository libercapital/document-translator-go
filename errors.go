package documenttranslator

import (
	"errors"
)

var (
	ErrParseShorterThenDeliminator = errors.New("line to parse is shorter then deliminator")
	ErrKindInconsistency           = errors.New("row invalid due kind inconsistency")
	ErrSegmentInconsistency        = errors.New("row invalid due segment inconsistency")
	ErrSegmentMustBeString         = errors.New("segment must be string")
)
