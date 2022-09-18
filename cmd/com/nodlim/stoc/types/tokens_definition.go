package types

import (
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc/commons"
)

// TokensDefinition A set of definitions for TokenType types.
// The definition is stored in a TokenInfo object.
type TokensDefinition map[TokenType]TokenInfo

// DefineTokenInfo configures a types.TokenInfo object.
// typ is the type of the TokenInfo object.
// key is the keyword for the type in TokenInfo.
// description is the description of the type in TokenInfo.
// The result is added to td TokensDefinition, and td is then returned, allowing chaining of method calls
func (td *TokensDefinition) DefineTokenInfo(typ TokenType, key string, description string) *TokensDefinition {
	(*td)[typ] = TokenInfo{
		desc: description,
		key:  []rune(key),
		size: len([]rune(key)),
	}
	return td
}

// Finalise converts the pointer of the td TokensDefinition into a non-pointer object.
// Primarily a convenience-method that makes defining TokensDefinition cleaner.
// It is used at the end of a chain of DefineTokenInfo method calls,
// so that it can be validly passed to a search function, for instance stoc.SearchStringCustom
func (td *TokensDefinition) Finalise() TokensDefinition {
	return *td
}

// IsLeftBracket returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.LBR at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsLeftBracket(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[LBR].key)
}

// IsLeftBracketI returns the length of the keyword of types.LBR defined by the TokensDefinition
func (td TokensDefinition) IsLeftBracketI() int {
	return td[LBR].size
}

// IsRightBracket returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.RBR at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsRightBracket(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[RBR].key)
}

// IsRightBracketI returns the length of the keyword of types.RBR defined by the TokensDefinition
func (td TokensDefinition) IsRightBracketI() int {
	return td[RBR].size
}

// IsAnd returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.AND at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsAnd(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[AND].key)
}

// IsAndI returns the length of the keyword of types.AND defined by the TokensDefinition
func (td TokensDefinition) IsAndI() int {
	return td[AND].size
}

// IsOr returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.OR at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsOr(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[OR].key)
}

// IsOrI returns the length of the keyword of types.OR defined by the TokensDefinition
func (td TokensDefinition) IsOrI() int {
	return td[OR].size
}

// IsNot returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.NOT at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsNot(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[NOT].key)
}

// IsNotI returns the length of the keyword of types.NOT defined by the TokensDefinition
func (td TokensDefinition) IsNotI() int {
	return td[NOT].size
}

// IsExpRune returns true if the rune array 'r' at numeric index 'index'
// does not match any keyword or whitespace at the start of the remainder of the array.
// All keywords are defined by the TokensDefinition
func (td TokensDefinition) IsExpRune(r []rune, index int) bool {
	return !(IsWhitespace(r, index) || td.IsKeyword(r, index))
}

// IsAssociativeOp returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for either types.OR or types.AND, at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsAssociativeOp(r []rune, index int) bool {
	return td.IsOr(r, index) || td.IsAnd(r, index)
}

// IsKeyword returns true if the rune array 'r' at numeric index 'index'
// matches any keyword at the start of the remainder of the array.
// All keywords are defined by the TokensDefinition
func (td TokensDefinition) IsKeyword(r []rune, index int) bool {
	return td.IsOr(r, index) || td.IsAnd(r, index) || td.IsLeftBracket(r, index) || td.IsRightBracket(r, index) || td.IsNot(r, index)
}

// IsQuote returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for either types.SQUOTE or types.DQUOTE at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsQuote(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[SQUOTE].key, td[DQUOTE].key)
}

// IsSingleInvertedComma returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.SQUOTE at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsSingleInvertedComma(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[SQUOTE].key)
}

// IsDoubleInvertedComma returns true if the rune array 'r' at numeric index 'index'
// follows with the keyword for types.DQUOTE at the start of the remainder of the array.
// The keyword is defined by the TokensDefinition
func (td TokensDefinition) IsDoubleInvertedComma(r []rune, index int) bool {
	return commons.StartsWith(r, index, td[DQUOTE].key)
}

// TokToString gets the TokenInfo description for t TokenType defined by TokensDefinition
func (td TokensDefinition) TokToString(t TokenType) string {
	return td[t].desc
}

// prepareDefaultTokensDefinition defines the default TokensDefinition for the library.
// It is exported by the global DefaultTokensDefinition
// It is used namely by stoc.SearchString
// The code for the method is also the de-facto reference for defining TokensDefinition objects.
func prepareDefaultTokensDefinition() TokensDefinition {
	def := TokensDefinition{}
	return def.
		DefineTokenInfo(AND, "&", "and").
		DefineTokenInfo(OR, "|", "or").
		DefineTokenInfo(NOT, "!", "not").
		DefineTokenInfo(ANDNOT, "&!", "and not").
		DefineTokenInfo(ORNOT, "&!", "or not").
		DefineTokenInfo(TRUE, "True", "true").
		DefineTokenInfo(LBR, "(", "left bracket").
		DefineTokenInfo(RBR, ")", "right bracket").
		DefineTokenInfo(EOL, "\n", "end of line").
		DefineTokenInfo(EXP, "", "expression").
		DefineTokenInfo(DQUOTE, "\"", "double inverted comma").
		DefineTokenInfo(SQUOTE, "'", "single inverted comma").
		Finalise()
}

// DefaultTokensDefinition exports the default tokens definition. See prepareDefaultTokensDefinition for more information.
var DefaultTokensDefinition = prepareDefaultTokensDefinition()
