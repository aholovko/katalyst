package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"katalyst/fizzbuzz"
)

func TestFizzbuzz(t *testing.T) {
	/*
		t.Run("should convert 1 to 1", func(t *testing.T) {
			require.Equal(t, "1", fizzbuzz.Convert(1))
		})

		t.Run("should convert 2 to 2", func(t *testing.T) {
			require.Equal(t, "2", fizzbuzz.Convert(2))
		})

		t.Run("should convert 4 to 4", func(t *testing.T) {
			require.Equal(t, "4", fizzbuzz.Convert(4))
		})

		// it's important to keep duplication as low as possible in tests, so three tests above were converted to
		// parameterized test during refactoring step (https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
	*/

	t.Run("should convert number to FizzBuzz string", func(t *testing.T) {
		tests := []struct {
			arg int
			exp string
		}{
			{1, "1"},
			{2, "2"},
			{4, "4"},
			{3, "Fizz"},
			{6, "Fizz"},
			{5, "Buzz"},
			{10, "Buzz"},
			{15, "FizzBuzz"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("fizzbuzz(%d)", tt.arg), func(t *testing.T) {
				require.Equal(t, tt.exp, fizzbuzz.Convert(tt.arg))
			})
		}
	})

	/*
		t.Run("should convert 3 to Fizz", func(t *testing.T) {
			require.Equal(t, "Fizz", fizzbuzz.Convert(3))
		})

		t.Run("should convert 6 to Fizz", func(t *testing.T) {
			require.Equal(t, "Fizz", fizzbuzz.Convert(6))
		})

		t.Run("should convert 5 to Buzz", func(t *testing.T) {
			require.Equal(t, "Buzz", fizzbuzz.Convert(5))
		})

		t.Run("should convert 10 to Buzz", func(t *testing.T) {
			require.Equal(t, "Buzz", fizzbuzz.Convert(10))
		})
	*/
}
