package othello

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
	Decode(s string) Square
}

type Encoder interface {
	Encode(s Square) string
}

// Decode implements the Decoder interface.
func (AlgebraicNotation) Decode(s string) Square {
	return 37
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