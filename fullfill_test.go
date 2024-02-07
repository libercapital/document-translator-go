package documenttranslator

import "testing"

func TestFullFill(t *testing.T) {
	tests := []struct {
		fillSize int
		fillWith string
		position FullFillPosition
		data     string
		expected string
	}{
		{
			fillSize: 3,
			fillWith: "*",
			position: FullFillLeft,
			data:     "Hello",
			expected: "***Hello",
		},
		{
			fillSize: 4,
			fillWith: "-",
			position: FullFillRight,
			data:     "World",
			expected: "World----",
		},
		{
			fillSize: 2,
			fillWith: "+",
			position: FullFillLeft,
			data:     "Test",
			expected: "++Test",
		},
		{
			fillSize: 0,
			fillWith: "@",
			position: FullFillRight,
			data:     "NoFill",
			expected: "NoFill",
		},
	}

	for _, test := range tests {
		result := FullFill(test.fillSize, test.fillWith, test.position, test.data)
		if result != test.expected {
			t.Errorf("FullFill(%d, %q, %v, %q) = %q; expected %q", test.fillSize, test.fillWith, test.position, test.data, result, test.expected)
		}
	}
}
