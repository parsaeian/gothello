package othello

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("length of input variable doesn't equal with 2", func(t *testing.T) {
		notation := AlgebraicNotation{}
		_, err := notation.Decode("f11")

		var got = ""
		if err != nil {
			got = err.Error()
		}
		want := fmt.Sprintf(DecodeLengthErrorMessage, 3)
		assertError(t, got, want)
	})

	t.Run("wrong column format", func(t *testing.T) {
		notation := AlgebraicNotation{}
		_, err := notation.Decode("x1")

		var got = ""
		if err != nil {
			got = err.Error()
		}
		want := fmt.Sprintf(DecodeColumnErrorMessage, "x")
		assertError(t, got, want)
	})

	t.Run("wrong column format", func(t *testing.T) {
		notation := AlgebraicNotation{}
		_, err := notation.Decode("c9")

		var got = ""
		if err != nil {
			got = err.Error()
		}
		want := fmt.Sprintf(DecodeRowErrorMessage, "9")
		assertError(t, got, want)
	})
}

func assertError(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct error, got %v, want %v", got, want)
	}
}
