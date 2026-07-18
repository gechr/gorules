# gorules

🛡️ Shared [ruleguard](https://github.com/quasilyte/go-ruleguard) lint rules for Go, consumed across repos through [go-critic](https://github.com/go-critic/go-critic)'s `ruleguard` checker under [golangci-lint](https://golangci-lint.run).

## Rules

| Rule             | Flags                                                                                                                                                                                              |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `stringEquality` | `assert`/`require` `Contains`/`NotContains` on a **string** (slice/map membership is left alone), and any `ErrorContains` - use `assert.Equal` on the full string, or `assert.ErrorIs` for errors. |

## Usage

Add a loader in the consuming repo. A dot-directory keeps it out of the normal build (the Go toolchain ignores `.`-prefixed directories), so no build tag is needed:

```go
// .ruleguard/rules.go
package gorules

import (
	bundle "github.com/gechr/gorules"
	"github.com/quasilyte/go-ruleguard/dsl"
)

func init() { dsl.ImportRules("gechr", bundle.Bundle) }
```

Wire it into `.golangci.yml` (the `rules:` path is required - golangci-lint does not auto-discover it):

```yaml
linters:
  settings:
    gocritic:
      enabled-checks:
        - ruleguard
      settings:
        ruleguard:
          rules: .ruleguard/rules.go
```

The bundle must be a `require` in the consumer's `go.mod` so ruleguard can resolve it from the module cache. Because the loader lives in an ignored directory, pin it deliberately (a `tools`-tagged import or a manual `require`) - `go mod tidy` will not add it on its own.
