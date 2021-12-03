package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			// neither a multiple of 3 nor a multiple of 5, case1
			name:     "normal, case1",
			input:    uint64(1),
			expected: "1",
		},
		{
			// neither a multiple of 3 nor a multiple of 5, case2
			name:     "normal, case2",
			input:    uint64(11),
			expected: "11",
		},
		{
			// a multiple of 3, case1
			name:     "multiple of 3, case1",
			input:    uint64(3),
			expected: "Fizz",
		},
		{
			// a multiple of 3, case2
			name:     "multiple of 3, case2",
			input:    uint64(33),
			expected: "Fizz",
		},
		{
			// a multiple of 5, case1
			name:     "multiple of 5, case1",
			input:    uint64(5),
			expected: "Buzz",
		},
		{
			// a multiple of 5, case2
			name:     "multiple of 5, case2",
			input:    uint64(50),
			expected: "Buzz",
		},
		{
			// a multiple of 15, case1
			name:     "multiple of 15, case1",
			input:    uint64(15),
			expected: "FizzBuzz",
		},
		{
			// a multiple of 15, case2
			name:     "multiple of 15, case2",
			input:    uint64(150),
			expected: "FizzBuzz",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != Fizzbuzz(tc.input) {
				t.Errorf("Fizzbuzz(%v) = %v but expected %v\n", tc.input, Fizzbuzz(tc.input), tc.expected)
			}
		})
	}
}
