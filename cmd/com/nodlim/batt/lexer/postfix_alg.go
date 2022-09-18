package lexer

import (
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/commons"
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/pos_error"
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/types"
)

// TokenShuntingAlgorithm Shunting algorithm, based on the mathematical implementation available on the shunting algorithm wiki page
//
// A snapshot of that algorithm is below (this implementation is a simplified version with some new additions):
//
// while there are unparsed to be read:
//
//	read a token.
//	if the token is an expression, then:
//	    push it to the output queue.
//	else if the token is an operator then:
//	    while ((there is an operator at the top of the operator stack)
//	          and ((the operator at the top of the operator stack has greater precedence)
//	              or (the operator at the top of the operator stack has equal precedence and the token is left associative))
//	          and (the operator at the top of the operator stack is not a left parenthesis)):
//	        pop operators from the operator stack onto the output queue.
//	    push it onto the operator stack.
//	else if the token is a left parenthesis (i.e. "("), then:
//	    push it onto the operator stack.
//	else if the token is a right parenthesis (i.e. ")"), then:
//	    while the operator at the top of the operator stack is not a left parenthesis:
//	        pop the operator from the operator stack onto the output queue.
//	    -> If the stack runs out without finding a left parenthesis, then there are mismatched parentheses.
//	    if there is a left parenthesis at the top of the operator stack, then:
//	        pop the operator from the operator stack and discard it
//	    if there is a function token at the top of the operator stack, then:
//	        pop the function from the operator stack onto the output queue.
//
// -> After while loop, if operator stack not null, pop everything to output queue
// if there are no more parsed to read then:
//
//	while there are still operator parsed on the stack:
//	    -> If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses.
//	    pop the operator from the operator stack onto the output queue.
//
// exit.
func TokenShuntingAlgorithm(toks []types.Token) ([]types.Token, pos_error.PosError) {
	var parsed []types.Token
	var operator []types.Token
	for i, tok := range toks {
		// This implementation does not implement composite functions, functions with variable number of arguments,
		// and unary operators. (Note: unary operators are represented using and-not, or-not instead)
		if tok.Typ == types.EXP || tok.Typ == types.TRUE {
			parsed = append(parsed, tok)
		} else if types.IsOp(tok.Typ) {
			for len(operator) > 0 && (types.IsLeftAssociative(end(operator).Typ) || types.IsComplexOp(end(operator).Typ)) {
				parsed, operator = moveEnd(parsed, operator)
			}
			operator = append(operator, tok)
		} else if tok.Typ == types.LBR {
			operator = append(operator, tok)
		} else if tok.Typ == types.RBR {
			// move all the operators to the parsed until we find a left bracket
			for len(operator) > 0 && operator[len(operator)-1].Typ != types.LBR {
				parsed, operator = moveEnd(parsed, operator)
			}

			// If the stack runs out without finding a left parenthesis, then there are mismatched parentheses.
			if len(operator) == 0 {
				return parsed, pos_error.New("shunting error, missing right bracket", i)
			} else if operator[len(operator)-1].Typ == types.LBR {
				operator = operator[:endIndex(operator)]
			}
		}
	}

	// After while loop, add every operator to output queue
	// (didn't bother popping from stack, just read array from back)
	for i := range operator {
		opI := operator[len(operator)-1-i] // reverse because this is a stack
		// If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses.
		if opI.Typ == types.LBR || opI.Typ == types.RBR {
			lastBR := commons.LastIndexOf(toks, func(t types.Token) bool {
				return t.Typ == types.LBR || t.Typ == types.RBR
			})
			return parsed, pos_error.New("shunting error, some brackets were not matched together", lastBR)
		}
		parsed = append(parsed, opI)
	}

	return parsed, nil
}

/**
 * Utility methods only used by shunting algorithm
 */

func endIndex(tokens []types.Token) int {
	return len(tokens) - 1
}

func end(tokens []types.Token) types.Token {
	return tokens[endIndex(tokens)]
}

func pop(tokens []types.Token) (types.Token, []types.Token) {
	tok := end(tokens)
	return tok, tokens[:endIndex(tokens)]
}

func moveEnd(as []types.Token, bs []types.Token) ([]types.Token, []types.Token) {
	b, bs := pop(bs)
	as = append(as, b)
	return as, bs
}
