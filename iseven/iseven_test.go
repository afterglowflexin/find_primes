package iseven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isEven_Test(t *testing.T) {
	testCases := []struct {
		name   string
		number int
		isEven bool
	}{
		{
			name:   "even",
			number: 1,
			isEven: true,
		},
		{
			name:   "even",
			number: -345,
			isEven: true,
		},
		{
			name:   "even",
			number: 184297,
			isEven: true,
		},
		{
			name:   "odd",
			number: 1891356,
			isEven: false,
		},
		{
			name:   "odd",
			number: -489412,
			isEven: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isEven {
				assert.True(t, isEven(tc.number))
			} else {
				assert.False(t, isEven(tc.number))
			}
		})
	}
}
