# BATT Boolean Algebra To Tokens

[![Build](https://github.com/kranzuft/stoc/actions/workflows/goBuild.yml/badge.svg)](https://github.com/kranzuft/stoc/actions/workflows/goBuild.yml)
[![Test](https://github.com/kranzuft/stoc/actions/workflows/goTest.yml/badge.svg)](https://github.com/kranzuft/stoc/actions/workflows/goTest.yml)

## Overview

Bridges the gap between simple searching on a single string, and more powerful search using regex.
Plain english and simple boolean algebra is used.

## Dependencies

Only dependencies are for unit testing using ginkgo and gomega

## Current features

- and, or, not conditions
- quotes around strings
- custom types (see tokens_definition.go ```prepareDefaultTokensDefinition()``` function and 'Custom Examples' section
  below)

## Basic Examples

- should not contain foo or baa, and must have baz
  ```! ( "foo" | 'baa' ) & baz```
    - brackets work with unary and binary conditionals
- The word foo or the phrase "baa or baz"
  ```foo | "baa | baz"```
    - keywords can be in expressions using quotes
- The phrase foo and the phrase foo
  ```foo & "foo"```
    - this is clearly redundant and is equivalent to the next example
- The phrase foo
  ```foo```
    - simple search is always an option

## Custom Examples

You can use custom types with stoc library. For instance using words instead of symbols:

- ```not { "foo" or 'baa' } and baz```
- ```foo or "baa or baz"```
- ```foo and "foo"```
- ```foo```

this can be configured with the following code snippet:

```go
package main

import (
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc"
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc/types"
)

func main() {
	def := types.TokensDefinition{}
	customTypes := def.DefineTokenInfo(types.AND, "and", "and").
		DefineTokenInfo(types.OR, "or", "or").
		DefineTokenInfo(types.NOT, "not", "not").
		DefineTokenInfo(types.ANDNOT, "and not", "and not").
		DefineTokenInfo(types.ORNOT, "or not", "or not").
		DefineTokenInfo(types.TRUE, "True", "true").
		DefineTokenInfo(types.LBR, "{", "left bracket").
		DefineTokenInfo(types.RBR, "}", "right bracket").
		DefineTokenInfo(types.EOL, "\n", "end of line").
		DefineTokenInfo(types.EXP, "", "expression").
		DefineTokenInfo(types.DQUOTE, "\"", "double inverted comma").
		DefineTokenInfo(types.SQUOTE, "'", "single inverted comma").
		Finalise()
	success, err := stoc.SearchStringCustom(customTypes, "Hello or hi", "Hello world")
}
```

## RoadMap

- escaping quotes in expressions
- some way to make strings case-sensitive and case-insensitive, likely using a character in front of the expression, similar to python 3's f string
- the reference app (command line) (WIP, see Frontends)
- match support (return a list of points where a match was found, for highlighting)
- context support (include lines before and after a matched line)

## Potential additions (still considering)

- overriding keywords e.g. \& or \and to override keywords when not using quotes (might be overkill)
- might be good to include basic regex concepts like \w\d\s word, digit, whitespace \W\D\S not word, digit,
  whitespace, and potentially ^ and $

## Frontends

https://github.com/kranzuft/stoc-cli (reference implementation)
