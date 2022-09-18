// Package lexer handles passing tokens from raw text data
//
// # The lexer package is made up of the token lexer and postfix algorithm
//
// Token Lexer
// Primarily consists of functions that process a token, and the return with the given token and a function
// that defines the next step in processing.
//
// Postfix Algorithm
// Consists of two primary functions:
// 1. converts the in-fix tokens array to post-fix
// 2. processes the post-fix tokens against a target text and determines whether the text meets the conditions defined by the tokens
//
// Potential improvement:
// Part of the token lexer design is to minimise repeated decisions. When it comes to determining
// token type, at times we figure out the generic type, which requires checking specific types, and then pass it on to
// a generic function to then figure out the specific type again. For instance the generic type operator is determined
// by figuring out if it matches a specific operator type, then pass to a function to lex any operator. That function then
// has to re-determine the operator type. It might be worth creating simple functions for each type, that call the
// lexOperator function internally with a pre-determined type
package lexer

import (
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/token_error"
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/types"
	"strings"
)

// Definitions interface for lexer type definitions. A valid implementation is types.TokensDefinition.
type Definitions interface {
	IsLeftBracket(r []rune, index int) bool

	IsLeftBracketI() int

	IsRightBracket(r []rune, index int) bool

	IsRightBracketI() int

	IsAnd(r []rune, index int) bool

	IsAndI() int

	IsOr(r []rune, index int) bool

	IsOrI() int

	IsNot(r []rune, index int) bool

	IsNotI() int

	IsExpRune(r []rune, index int) bool

	IsAssociativeOp(r []rune, index int) bool

	IsKeyword(r []rune, index int) bool

	IsQuote(r []rune, index int) bool

	IsSingleInvertedComma(r []rune, index int) bool

	IsDoubleInvertedComma(r []rune, index int) bool

	TokToString(t types.TokenType) string
}

type stateFn func(Definitions, []rune, int) (int, types.Token, stateFn, token_error.PosError)

// BooleanAlgebraLexer looks for boolean operations in the search text, and transforms the operations into a token list.
//
// If an unexpected token type comes up, outside the expected order, false defs.Is returned. In this situation, a message
// shows where in the line the exception occurred, with a descriptive message included that states what
// was expected.
func BooleanAlgebraLexer(defs Definitions, raw []rune) ([]types.Token, token_error.PosError) {
	var tokens []types.Token
	var token types.Token
	var err token_error.PosError

	if len(raw) == 0 {
		return tokens, token_error.New("Empty search text", 0)
	}

	index := skipWhitespace(raw, 0)
	for state := determineIfExpressionOrLeftBracketOrNot; state != nil; {
		index, token, state, err = state(defs, raw, index)
		if err != nil || token.Typ == types.UNKNOWN {
			return tokens, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

// lexLeftBracket Lex a left bracket
//
// A left bracket can be proceeded by an expression, a left bracket or a not operator
func lexLeftBracket(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token
	tok.Typ = types.LBR // we don't need to track that it's a left bracket anymore
	tok.Exp = string(rawData[index : index+defs.IsLeftBracketI()])
	i := skipWhitespace(rawData, index+1)

	if i < len(rawData) {
		if defs.IsExpRune(rawData, i) {
			return i, tok, lexExpression, nil
		} else if defs.IsLeftBracket(rawData, i) {
			return i, tok, lexLeftBracket, nil
		} else if defs.IsNot(rawData, i) {
			return i, tok, lexNotSymbolAndCreateTrueToken, nil
		}
	}

	return prepareAndReturnError(defs, rawData, i, types.EXP)
}

// lexRightBracket Lex a right bracket.
//
// A right bracket can be proceeded by an operator or a right bracket
func lexRightBracket(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token
	tok.Typ = types.RBR // we don't need to track that it's a right bracket anymore
	tok.Exp = string(rawData[index : index+defs.IsRightBracketI()])
	i := skipWhitespace(rawData, index+1)

	if i < len(rawData) {
		if defs.IsAssociativeOp(rawData, i) {
			return i, tok, lexOperator, nil
		} else if defs.IsRightBracket(rawData, i) {
			return i, tok, lexRightBracket, nil
		}
	} else {
		return i, tok, nil, nil
	}

	return prepareAndReturnError(defs, rawData, i, types.RBR)
}

// lexExpressionWithQuotes lex expression with surrounding quotes
//
// We determine nextIndex in return because lexExpressionWithoutQuotes inherently determines this, so although we don't,
// we do
func lexExpressionWithQuotes(defs Definitions, rawData []rune, index int, doubleQuote bool) (int, int, bool) {
	var isQ func([]rune, int) bool
	// if an expression starts with a double quote, it must end with a double quote, and same for a single quote
	if doubleQuote {
		isQ = defs.IsDoubleInvertedComma
	} else {
		isQ = defs.IsSingleInvertedComma
	}

	i := index
	for ; i < len(rawData) && !isQ(rawData, i); i++ {
	}

	endExp := i

	nextIndex := i + 1 // move over quote symbol
	// now skip whitespace
	nextIndex = skipWhitespace(rawData, nextIndex)
	if nextIndex == len(rawData) || defs.IsRightBracket(rawData, nextIndex) || defs.IsAssociativeOp(rawData, nextIndex) {
		return nextIndex, endExp, true
	}

	return i, i, false
}

// lexExpressionWithoutQuotes lex expression without quotes
// Since we do not have quotes to bind the expression, we must wait until we find a keyword
// Should only be called by lexExpression, as it check for quotes first.
func lexExpressionWithoutQuotes(defs Definitions, rawData []rune, index int) (int, int, bool) {
	i := index
	trailingWhitespace := 0

	foundNextToken := func(raw []rune, cur int) bool {
		return defs.IsRightBracket(raw, cur) || defs.IsAssociativeOp(raw, cur)
	}

	for ; i < len(rawData) && !(i != index && foundNextToken(rawData, i)); i++ {
		if defs.IsExpRune(rawData, i) {
			trailingWhitespace = 0
		} else if types.IsWhitespace(rawData, i) {
			trailingWhitespace++
		} else {
			return i, i, false
		}
	}

	return i, i - trailingWhitespace, true
}

// lexExpression lex an expression with or without quotes
// Determines if quotes are used and then hands off to lexExpressionWithQuotes or lexExpressionWithoutQuotes
//
// Expressions are followed by either operators or right brackets
func lexExpression(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token
	tok.Typ = types.EXP

	success := false
	// we need to track the start and end of the expression
	startExp := index
	endExp := index
	// and also when the next token starts
	nextI := index

	if defs.IsSingleInvertedComma(rawData, index) {
		startExp++
		nextI, endExp, success = lexExpressionWithQuotes(defs, rawData, startExp, false)
	} else if defs.IsDoubleInvertedComma(rawData, index) {
		startExp++
		nextI, endExp, success = lexExpressionWithQuotes(defs, rawData, startExp, true)
	} else {
		nextI, endExp, success = lexExpressionWithoutQuotes(defs, rawData, startExp)
	}

	if !success {
		return prepareAndReturnError(defs, rawData, nextI, types.EXP)
	}

	tok.Exp = string(rawData[startExp:endExp])

	nextI = skipWhitespace(rawData, nextI)

	if nextI < len(rawData) {
		if defs.IsAssociativeOp(rawData, nextI) {
			return nextI, tok, lexOperator, nil
		} else if defs.IsRightBracket(rawData, nextI) {
			return nextI, tok, lexRightBracket, nil
		}
	}

	return nextI, tok, nil, nil
}

// getNextFuncAfterOperator finds what function to call based on the last operator processed by the lexer.
//
// Operators are followed by expressions or left brackets
func getNextFuncAfterOperator(defs Definitions, rawData []rune, index int, tok types.Token) (int, types.Token, stateFn, token_error.PosError) {
	index = skipWhitespace(rawData, index)
	if index < len(rawData) {
		if defs.IsExpRune(rawData, index) {
			return index, tok, lexExpression, nil
		} else if defs.IsLeftBracket(rawData, index) {
			return index, tok, lexLeftBracket, nil
		}
	}

	return prepareAndReturnError(defs, rawData, index+1, types.EXP, types.LBR)
}

// lexNotSymbolAndCreateTrueToken for not symbols at the front of a condition or after a left bracket, we have to push
// a TRUE token first to form a valid binary operation. This function passes out that types.TRUE token, before calling
// the more common not-operator-lexing function lexNotSymbolAndCreateAndNotToken
func lexNotSymbolAndCreateTrueToken(_ Definitions, _ []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token

	tok.Typ = types.TRUE
	tok.Exp = "true"

	return index, tok, lexNotSymbolAndCreateAndNotToken, nil
}

// lexNotSymbolAndCreateAndNotToken Takes a types.NOT and makes it into a types.ANDNOT token type
// by combining the preceding types.AND token using this function
func lexNotSymbolAndCreateAndNotToken(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token

	tok.Typ = types.ANDNOT
	tok.Exp = string(rawData[index : index+defs.IsNotI()])

	i := skipWhitespace(rawData, index+defs.IsNotI())

	return getNextFuncAfterOperator(defs, rawData, i, tok)
}

// lexOperator Creates an operator token, and then calls getNextFuncAfterOperator to determine next step
func lexOperator(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	var tok types.Token
	indexEnd := 0

	tok.Typ, indexEnd = determineTokenType(defs, rawData, index)
	tok.Exp = string(rawData[index : index+indexEnd])

	return getNextFuncAfterOperator(defs, rawData, index+indexEnd, tok)
}

// determineTokenType determines the token type for rawData rune array at numeric index 'index'
func determineTokenType(defs Definitions, rawData []rune, index int) (types.TokenType, int) {
	typ := types.UNKNOWN
	typeSize := 0

	// first see if it is an operator
	opType, opLen := determineTokenTypeForOperator(defs, rawData, index)

	if opType != types.UNKNOWN {
		typ = opType
		typeSize = opLen
	} else if defs.IsRightBracket(rawData, index) {
		typ = types.RBR
		typeSize = defs.IsRightBracketI()
	} else if defs.IsLeftBracket(rawData, index) {
		typ = types.LBR
		typeSize = defs.IsLeftBracketI()
	} else if defs.IsSingleInvertedComma(rawData, index) {
		typ = types.SQUOTE
	} else if defs.IsDoubleInvertedComma(rawData, index) {
		typ = types.DQUOTE
	}

	return typ, typeSize
}

// determineTokenType determines the token type for rawData rune array at numeric index 'index'
func determineTokenTypeForOperator(defs Definitions, rawData []rune, index int) (types.TokenType, int) {
	typ := types.UNKNOWN
	typeSize := 0
	couldBeComplex := false

	// Let's try and make it a known type
	if defs.IsAnd(rawData, index) {
		typ = types.AND
		typeSize = defs.IsAndI()
		couldBeComplex = true
	} else if defs.IsOr(rawData, index) {
		typ = types.OR
		typeSize = defs.IsOrI()
		couldBeComplex = true
	}

	// if we made it a known type
	// Although not UNKNOWN is equivalent to couldBeComplex, it won't be in future improvements
	if typ != types.UNKNOWN && couldBeComplex {
		skipForward := skipWhitespace(rawData, index+typeSize)

		// if a not follows, then it's a complex operator
		if defs.IsNot(rawData, skipForward) {
			if typ == types.AND {
				typ = types.ANDNOT
			} else if typ == types.OR {
				typ = types.ORNOT
			}

			typeSize = (skipForward - index) + defs.IsNotI()
		}
	}

	return typ, typeSize
}

// prepareAndReturnError Prints an error when an error occurs in lexing.
//
// To be deprecated and replaced with an error with the same level of detail.
func prepareAndReturnError(defs Definitions, rawData []rune, index int, expected ...types.TokenType) (int, types.Token, stateFn, token_error.PosError) {
	var found types.TokenType
	index = skipWhitespace(rawData, index)
	if (index + 1) < len(rawData) {
		found, _ = determineTokenType(defs, rawData, index)
	} else if index < len(rawData) {
		found, _ = determineTokenType(defs, rawData, index)
	} else {
		found = types.EOL
	}
	message := tokensToList(defs, expected) + " expected, instead found " + defs.TokToString(found)
	return returnErrorMessage(defs, rawData, index, message)
}

// returnErrorMessage Builds an error message based on the current context of the lexer
func returnErrorMessage(_ Definitions, _ []rune, index int, message string) (int, types.Token, stateFn, token_error.PosError) {
	return index, getErrorToken(), nil, token_error.New(message, index)
}

// skipWhitespace skips whitespace in rawData array at index
func skipWhitespace(rawData []rune, index int) int {
	i := index

	for ; i < len(rawData); i++ {
		if !types.IsWhitespace(rawData, i) {
			return i
		}
	}

	return i
}

// determineIfExpressionOrLeftBracketOrNot Determines if next is a left bracket, an expression, or a not.
//
// Can be considered the entrypoint of the lexer. The types it looks for a valid starting values
func determineIfExpressionOrLeftBracketOrNot(defs Definitions, rawData []rune, index int) (int, types.Token, stateFn, token_error.PosError) {
	if defs.IsLeftBracket(rawData, index) {
		return lexLeftBracket(defs, rawData, index)
	} else if defs.IsExpRune(rawData, index) {
		return lexExpression(defs, rawData, index)
	} else if defs.IsNot(rawData, index) {
		return lexNotSymbolAndCreateTrueToken(defs, rawData, index)
	}

	return prepareAndReturnError(defs, rawData, index, types.LBR, types.EXP)
}

func getErrorToken() types.Token {
	return types.Token{Typ: types.UNKNOWN}
}

// tokensToList lists input tokens in a string, based on Definitions object
func tokensToList(defs Definitions, expected []types.TokenType) string {
	var res strings.Builder
	for index, exp := range expected {
		if len(expected) > 1 && index == len(expected)-1 {
			res.WriteString(" or ")
		}
		res.WriteString(defs.TokToString(exp))
		if index < len(expected)-2 {
			res.WriteString(", ")
		}

	}
	return res.String()
}
