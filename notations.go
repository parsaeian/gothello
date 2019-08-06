package othello

import "fmt"

const (
	asciiCodeOfLowercaseA = 97
	asciiCodeOf1          = 49
	DecodeLengthErrorMessage = "length of input variable does'nt equal with 2, length: %d"
	DecodeColumnErrorMessage = "wrong column format (between a-h): %s"
	DecodeRowErrorMessage = "wrong column format (between 1-8): %s"
)

// Notation is the interface implemented by objects that can
// encode and decode moves.
type Notation interface {
	Encoder
	Decoder
}

// AlgebraicNotation is the official othello notation.
// Examples: f5
type AlgebraicNotation struct{}

type Decoder interface {
	Decode(s string) (Square, error)
}

type Encoder interface {
	Encode(s Square) string
}

// Decode implements the Decoder interface.
func (AlgebraicNotation) Decode(s string) (Square, error) {
	if len(s) != 2 {
		return Square(0), fmt.Errorf(DecodeLengthErrorMessage, len(s))
	}
	columnChar := s[0]
	ascii := int(columnChar)
	columnInt := ascii - asciiCodeOfLowercaseA
	if columnInt < 0 || columnInt > 7 {
		return Square(0), fmt.Errorf(DecodeColumnErrorMessage, string(columnChar))
	}

	rowChar := s[1]
	rowInt := int(rowChar) - asciiCodeOf1
	if rowInt < 0 || rowInt > 7 {
		return Square(0), fmt.Errorf(DecodeRowErrorMessage, string(rowChar))
	}

	return Square(rowInt*8 + columnInt), nil
}

// Encode implements the Encoder interface.
func (AlgebraicNotation) Encode(s Square) string {
	return ""
}

// String implements the fmt.Stringer interface and returns
// the notation's name.
func (AlgebraicNotation) String() string {
	return "Algebraic Notation"
}
