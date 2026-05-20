# Agent Instructions for spgtty

## Project Summary

**spgtty** (pronounced "spaghetti") is a CLI tool that bundles JavaScript for Shelly Gen2+ smart home devices. Think of it as "webpack for Shelly" - it takes multi-file JS projects and outputs a single Shelly-compatible script.

## Quick Start for Agents

```bash
# Current branch
git branch --show-current  # Should be: dev

# Build the tool
go build

# Test it works
./spgtty test/main.js
cat dist/main.js

# Run tests
go test ./...
```

## Code Style & Conventions

### Go Code
- Standard `gofmt` formatting
- Comments in English
- Error messages: German or English (currently mixed, can be unified later)
- Use `log.Fatalf()` for user-facing errors
- Return `error` from library functions, don't `os.Exit()` in packages

### Naming
- **IMPORTANT:** Keep "Spagetty" spelling in branding - it's an intentional wordplay!
- Fix actual typos: "Shally" → "Shelly", "inmplemented" → "implemented"

### Imports
```go
import (
    // Standard library first
    "fmt"
    "os"
    
    // External packages
    "github.com/spf13/cobra"
    
    // Internal packages
    "github.com/GrosseBen/spgtty/pkg/builder"
)
```

## Key Files to Understand

| File | Purpose | Status |
|------|---------|--------|
| `main.go` | Entry point | ✅ Working |
| `cmd/root.go` | CLI setup, flags | ✅ Working (on v0.2.5) |
| `cmd/build.go` | Build command | ✅ Working |
| `cmd/upload.go` | Upload command | ❌ Stub (panic) |
| `cmd/init.go` | Init command | ❌ Stub (panic) |
| `pkg/builder/builder.go` | esbuild integration | ✅ Working |
| `pkg/deployer/upload.go` | Shelly upload | ⚠️ Has orphan main() |

## Current Branch Status

**dev branch** is older and missing Cobra structure.  
**v0.2.5 tag** has the complete Cobra CLI.

**TODO:** Merge v0.2.5 into dev to get Cobra structure.

## Architecture Overview

```
User runs: spgtty build main.js
                ↓
        cmd/build.go
                ↓
    pkg/builder/builder.go
                ↓
          esbuild
                ↓
    Post-process (remove trailing commas)
                ↓
        dist/main.js
```

## Important Constraints

### Shelly JavaScript Limitations
- No trailing commas in objects/arrays
- No async/await (use callbacks)
- No ES modules at runtime (bundled by esbuild)
- Limited memory (~16-32KB scripts)

### Upload Chunking
Shelly has 1024 byte limit per RPC request. Large scripts must be uploaded in chunks:
```go
// First chunk
append: false  // overwrite

// Subsequent chunks  
append: true   // append to existing
```

## Planned Changes

See `docs/PLAN.md` for the full roadmap. Summary:

1. **Viper config** - Hierarchical configuration (global/local/env/flags)
2. **Deployer refactor** - Remove `main()` from `pkg/deployer/upload.go`
3. **Upload command** - Implement `cmd/upload.go`
4. **Init command** - Implement `cmd/init.go`
5. **Tests** - Add `pkg/builder/builder_test.go` with fixtures
6. **Typo fixes** - Correct spelling errors

## Testing

```bash
# Run all tests
go test ./...

# Test the build manually
./spgtty test/main.js
cat dist/main.js

# Test fixtures are in test/fixtures/ (to be created)
```

## Common Tasks

### Add a new CLI flag
1. Define in `cmd/root.go` or specific command file
2. Bind to Viper config key (after Viper is added)
3. Access via `viper.GetString("key")` or flag variable

### Add a new command
1. Create `cmd/newcommand.go`
2. Define `var newCmd = &cobra.Command{...}`
3. Add to root in `init()`: `rootCmd.AddCommand(newCmd)`

### Fix a Shelly compatibility issue
1. Add test fixture in `test/fixtures/`
2. Add test case in `pkg/builder/builder_test.go`
3. Fix in `pkg/builder/builder.go` (often post-processing regex)

## Dependencies

| Package | Purpose |
|---------|---------|
| `spf13/cobra` | CLI framework |
| `spf13/viper` | Config (planned) |
| `evanw/esbuild` | JS bundler |

## Contact / Resources

- GitHub: https://github.com/GrosseBen/spgtty
- Shelly Docs: https://shelly-api-docs.shelly.cloud/gen2/Scripts/
