package types

// Token is the core element of the lexed grammar.
// It is made up by a type and the phrase that matched the type.
// In this context we call this phrase the expression (Exp),
// as distinct from the expression type in the condition grammar, which is a filter-rule.
type Token struct {
	// Typ The type of the token
	Typ TokenType
	// Exp The phrase that matches the type in a given context
	Exp string
}
