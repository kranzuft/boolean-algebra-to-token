// Package types defines types for conditional grammar
// Defines the rules governing types in stoc, and are the core building block of the condition grammar.
// stoc is heavily influenced by type-theory.
// The types are pivotal to the lexer package
package types

type TokenType string

// TokenInfo defines a configuration of a TokenType.
type TokenInfo struct {
	// desc the description of the type.
	desc string
	// key the keyword for the type, used by the lexer.
	key []rune
	// size the length of the keyword.
	size int
}

// TokenType are a set of non-binding descriptors for the syntax of stoc conditions
const (
	EOL     TokenType = "END_OF_LINE"
	UNKNOWN TokenType = "UNKNOWN" // ignore first value
	AND     TokenType = "AND"
	OR      TokenType = "OR"
	ANDNOT  TokenType = "ANDNOT"
	ORNOT   TokenType = "ORNOT"
	NOT     TokenType = "NOT"
	EXP     TokenType = "EXPRESSION"
	LBR     TokenType = "LEFT_BRACKET"
	RBR     TokenType = "RIGHT_BRACKET"
	TRUE    TokenType = "TRUE"
	DQUOTE  TokenType = "DOUBLE_INVERTED_COMMA"
	SQUOTE  TokenType = "SINGLE_INVERTED_COMMA"
)

// IsLeftAssociative returns true if TokenType matches a left-to-right associative type.
func IsLeftAssociative(tok TokenType) bool {
	return tok == AND || tok == OR
}

// IsOp returns true if TokenType matches any operator.
func IsOp(tok TokenType) bool {
	return tok == AND || tok == OR || IsComplexOp(tok)
}

// IsComplexOp returns true if TokenType matches a complex operator
// a complex operator is a conjunction (types.AND) or disjunction (types.OR) that also applies an inversion (types.NOT).
func IsComplexOp(tok TokenType) bool {
	return tok == ANDNOT || tok == ORNOT
}

// CouldBeComplexOp returns true if TokenType matches any operator
// that a unary operator could modify to create a complex type.
func CouldBeComplexOp(tok TokenType) bool {
	return tok == AND || tok == OR
}

func IsWhitespace(r []rune, index int) bool {
	return (len(r)-index) > 0 && (r[index] == ' ' || r[index] == '\n' || r[index] == '\t')
}
