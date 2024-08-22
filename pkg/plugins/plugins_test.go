package plugins

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_stringToSliceHookFunc(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected []string
	}{
		{
			desc:     "without separator",
			data:     "║24║abc",
			expected: []string{"abc"},
		},
		{
			desc:     "with the file separator",
			data:     "║24║a║b║c",
			expected: []string{"a", "b", "c"},
		},
		{
			desc:     "with the label separator",
			data:     "a,b,c",
			expected: []string{"a", "b", "c"},
		},
		{
			desc:     "with the file separator and values with commas",
			data:     "║24║a,z║b,w║c,x,y",
			expected: []string{"a,z", "b,w", "c,x,y"},
		},
		{
			desc:     "with the file separator and empty item",
			data:     "║24║║a║z",
			expected: []string{"", "a", "z"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			values, err := stringToSliceHookFunc(reflect.String, reflect.Slice, test.data)
			require.NoError(t, err)

			assert.EqualValues(t, test.expected, values)
		})
	}
}
