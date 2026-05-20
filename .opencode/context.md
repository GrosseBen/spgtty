# Quick Context for spgtty

Use this file to quickly understand the project when starting with fresh context.

## One-Line Summary

CLI tool that bundles JavaScript for Shelly smart home devices (like webpack for Shelly).

## Name Origin

"spgtty" = "spaghetti" - wordplay on messy code that gets bundled into one file.
**Keep this spelling intentional!**

## Tech Stack

- **Language:** Go 1.25
- **CLI:** Cobra (spf13/cobra)
- **Config:** Viper (planned, spf13/viper)
- **JS Bundler:** esbuild

## Project State

| Component | Status |
|-----------|--------|
| Build command | ✅ Working |
| Upload command | ❌ Stub |
| Init command | ❌ Stub |
| Viper config | ❌ Not implemented |
| Tests | ❌ None yet |

## Branch Info

- **dev** - Development branch (currently checked out, but outdated)
- **main** - Stable branch
- **v0.2.5** - Latest tag with Cobra structure

**Action needed:** Merge v0.2.5 into dev

## Key Commands

```bash
# Build
spgtty                    # Build main.js → dist/main.js
spgtty build src/app.js   # Build specific file
spgtty -m                 # No minification

# Planned
spgtty upload             # Upload to Shelly
spgtty init               # Create project
```

## File Structure

```
spgtty/
├── main.go           # Entry → cmd.Execute()
├── cmd/              # CLI commands (Cobra)
├── pkg/
│   ├── builder/      # esbuild wrapper ← Core logic here
│   ├── deployer/     # Shelly RPC ← Needs refactoring
│   └── utils/        # Version info
├── test/             # Test files
├── docs/             # Documentation
└── .opencode/        # AI context (you are here)
```

## Files to Read First

1. `docs/PLAN.md` - What needs to be done
2. `docs/ARCHITECTURE.md` - How it works
3. `cmd/root.go` - CLI entry (on v0.2.5)
4. `pkg/builder/builder.go` - Core bundling logic

## Known Issues

1. **Orphan main()** in `pkg/deployer/upload.go` - needs to be a library function
2. **Typos** throughout codebase - "Shally", "inmplemented", etc.
3. **No tests** - need to add builder tests with fixtures
4. **dev branch outdated** - needs merge from v0.2.5

## Shelly Constraints

- No trailing commas (removed by post-processing)
- No async/await (use callbacks)
- Max 1024 bytes per upload request (chunking needed)
- ES modules bundled into single file

## Quick Build & Test

```bash
go build
./spgtty test/main.js
cat dist/main.js
```

## Links

- Repo: https://github.com/GrosseBen/spgtty
- Shelly Docs: https://shelly-api-docs.shelly.cloud/gen2/Scripts/
