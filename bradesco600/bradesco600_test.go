package bradesco600

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditAssessment(t *testing.T) {
	segment := "1501202400000000000117304ELETROZEMA S/A                          20264047310001961                   150120241402202400000000000000000000000083-REAL                0000000000011740400000000000000000000000000001174040000000000000000000000000000117404000000000001174040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000                                          0000124798515.01.2024001                                                                                                            "

	parsed, err := Parse(segment)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// TODO: test is incomplete. Fields don't match perfectly.
	assert.Equal(t, "15012024", parsed.BaseDate.Format("02012006"))
	assert.Equal(t, "ELETROZEMA S/A", parsed.CustomerName)
	assert.Equal(t, "083-REAL", parsed.Indexer)
}
