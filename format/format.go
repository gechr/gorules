// Package format holds shared ruleguard *formatter* rules: ones that use
// Suggest (a rewrite) rather than Report (a diagnostic). They are separated
// from the diagnostics in the parent package because a Suggest rule matches
// on AST shape, not on whether the code is already formatted the way the
// suggestion would rewrite it - so it re-reports on already-formatted code
// and must run only as a `--fix` formatting pass, never in the lint gate.
// Consumers import this bundle from a dedicated ruleguard config (see the
// repo README), keeping it out of the config that gates lint.
package format

//go:generate go run gen/clog.go

import "github.com/quasilyte/go-ruleguard/dsl"

// Bundle exposes every formatter rule in this package for import via
// dsl.ImportRules.
var Bundle = dsl.Bundle{}
