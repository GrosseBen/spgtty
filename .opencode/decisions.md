# Architecture Decision Records

This document tracks important decisions made during spgtty development.

---

## ADR-001: Keep "Spagetty" Spelling

**Date:** 2025-05-20  
**Status:** Accepted

### Context
The project name "spgtty" stands for "spaghetti" - a wordplay on messy code that gets bundled into one file.

### Decision
The intentional misspelling "Spagetty" in user-facing text (CLI descriptions, README branding) should be preserved.

### Consequences
- Only fix actual typos like "Shally" → "Shelly", "inmplemented" → "implemented"
- Keep "Spagetty" in: CLI help text, README title, branding
- Code identifiers should use correct spelling where practical

---

## ADR-002: Use Viper for Configuration

**Date:** 2025-05-20  
**Status:** Planned

### Context
The tool needs configuration for Shelly device IP, script ID, build settings. Users want flexibility: config files, environment variables, and CLI flags.

### Decision
Use `spf13/viper` for hierarchical configuration management.

### Configuration Hierarchy (lowest to highest priority)
1. Defaults (hardcoded in Go)
2. Global config: `~/.config/spgtty/config.yaml`
3. Local config: `.spgtty.yaml` in project directory
4. Environment variables: `SPGTTY_*`
5. CLI flags: `--host`, `--script-id`, etc.

### Consequences
- Add viper dependency to go.mod
- Create `pkg/config/config.go` for initialization
- Bind CLI flags to viper in `cmd/root.go`
- Document config options in `docs/CONFIG.md`

---

## ADR-003: Chunked Upload for Shelly

**Date:** 2025-05-20  
**Status:** Implemented (in pkg/deployer)

### Context
Shelly Gen2+ devices have a limit of approximately 1024 bytes per RPC request body.

### Decision
Implement chunked upload:
1. First chunk: `append: false` (overwrite existing script)
2. Subsequent chunks: `append: true` (append to script)

### Implementation
```go
const symbolsInChunk = 1024

for pos := 0; pos < len(code); {
    end := min(pos + symbolsInChunk, len(code))
    chunk := code[pos:end]
    putChunk(host, id, chunk, pos > 0)  // append=true after first
    pos = end
}
```

### Consequences
- Large scripts work reliably
- Multiple HTTP requests needed (slight latency)
- Progress reporting possible (X of Y chunks)

---

## ADR-004: Remove Trailing Commas in Post-Processing

**Date:** 2025-05-20  
**Status:** Implemented

### Context
Shelly's JavaScript engine (based on mJS) doesn't support trailing commas in object literals or arrays, which are valid in modern JavaScript.

### Decision
Use regex post-processing after esbuild to remove trailing commas.

### Implementation
```go
re := regexp.MustCompile(`,[\s]*([}\]])`)
cleanedCode := re.ReplaceAll(jsCode, []byte("$1"))
```

### Pattern Matched
```javascript
// Before
{a: 1, b: 2,}  // trailing comma
[1, 2, 3,]     // trailing comma

// After
{a: 1, b: 2}
[1, 2, 3]
```

### Consequences
- Shelly compatibility ensured
- Slight processing overhead (negligible)
- May need adjustment if false positives occur (commas in strings)

---

## ADR-005: Test Fixtures Per Language Feature

**Date:** 2025-05-20  
**Status:** Planned

### Context
Need to test JavaScript bundling for various language features. Also need to document what works and what doesn't with Shelly.

### Decision
Create one test fixture file per JavaScript language feature in `test/fixtures/`.

### Fixture Files
| File | Feature |
|------|---------|
| `simple.js` | Basic print() |
| `variables.js` | let, const, var |
| `functions.js` | Function declarations |
| `arrow.js` | Arrow functions |
| `objects.js` | Object literals |
| `arrays.js` | Arrays |
| `imports.js` | ES modules |
| `shelly_api.js` | Shelly-specific APIs |

### Consequences
- Easy to identify what works/fails
- Tests serve as documentation
- New issues → new fixtures
- Can track Shelly compatibility matrix

---

## ADR-006: esbuild for Bundling

**Date:** 2025-01-XX (original decision)  
**Status:** Implemented

### Context
Need to bundle multiple JavaScript files into a single file for Shelly. Need to transpile modern JS to ES2015.

### Decision
Use `evanw/esbuild` as the bundler/transpiler.

### Configuration
```go
api.Build(api.BuildOptions{
    EntryPoints:       []string{entryPath},
    Bundle:            true,
    Target:            api.ES2015,
    Format:            api.FormatCommonJS,
    MinifyIdentifiers: minify,
    MinifySyntax:      minify,
    MinifyWhitespace:  minify,
    Write:             false,
})
```

### Consequences
- Very fast builds
- Handles ES modules
- Transpiles to ES2015 (Shelly-compatible)
- Small dependency footprint

---

## ADR-007: Cobra for CLI

**Date:** 2025-01-XX (original decision)  
**Status:** Implemented

### Context
Need a CLI framework for subcommands (build, upload, init) and flag handling.

### Decision
Use `spf13/cobra` as the CLI framework.

### Structure
```
spgtty          → runs build (default)
spgtty build    → explicit build
spgtty upload   → upload to Shelly
spgtty init     → create project
```

### Consequences
- Standard Go CLI pattern
- Subcommand support
- Built-in help generation
- Integrates well with Viper

---

## ADR-008: Library Functions, Not Standalone Binaries

**Date:** 2025-05-20  
**Status:** Planned

### Context
`pkg/deployer/upload.go` contains a `main()` function, suggesting it was originally a standalone tool. This causes issues when importing as a library.

### Decision
Remove standalone `main()` functions from packages. All functionality should be exported functions callable from `cmd/` commands.

### Changes Required
- Remove `main()` from `pkg/deployer/upload.go`
- Export `Upload()` function
- Remove `flag` package usage (use Viper instead)
- Return `error` instead of calling `os.Exit()`

### Consequences
- Clean package structure
- Testable functions
- Single entry point in `main.go`

---

## Template for New Decisions

```markdown
## ADR-XXX: Title

**Date:** YYYY-MM-DD  
**Status:** Proposed | Accepted | Implemented | Deprecated

### Context
What is the issue that we're seeing that is motivating this decision?

### Decision
What is the change that we're proposing and/or doing?

### Consequences
What becomes easier or more difficult to do because of this change?
```
