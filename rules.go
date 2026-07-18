// Package gorules holds shared gocritic ruleguard lint rules, consumed across
// repos via a `.ruleguard/rules.go` loader that calls dsl.ImportRules.
package gorules

//go:generate go run gen/clog.go

import "github.com/quasilyte/go-ruleguard/dsl"

// Bundle exposes every rule in this package for import via dsl.ImportRules.
var Bundle = dsl.Bundle{}

// stringEquality forbids substring-style testify assertions on strings and
// errors: exact equality catches malformed output (wrong separators, stray
// fields, leading/trailing junk) that Contains silently passes. Slice/map
// membership checks (a non-string container) are left untouched.
func StringEquality(m dsl.Matcher) {
	m.Match(
		`assert.Contains($t, $s, $sub, $*_)`,
		`assert.NotContains($t, $s, $sub, $*_)`,
		`require.Contains($t, $s, $sub, $*_)`,
		`require.NotContains($t, $s, $sub, $*_)`,
	).
		Where(m["s"].Type.Is("string")).
		Report(`use assert.Equal with the full expected string, not Contains/NotContains`)

	m.Match(
		`assert.ErrorContains($t, $err, $sub, $*_)`,
		`require.ErrorContains($t, $err, $sub, $*_)`,
	).
		Report(`use assert.ErrorIs or assert.Equal on the full error string, not ErrorContains`)
}
