# Architecture

## Project Overview

spgtty is a CLI tool that bundles JavaScript projects for Shelly Gen2+ smart home devices.
It takes multi-file JS projects with imports and outputs a single, minified, Shelly-compatible script.

## Project Structure

```
spgtty/
├── main.go              # Entry point (calls cmd.Execute())
├── cmd/                 # Cobra CLI commands
│   ├── root.go          # Root command, global flags, Viper init
│   ├── build.go         # Build command (default action)
│   ├── upload.go        # Upload to Shelly device
│   ├── init.go          # Create new project
│   └── version.go       # Version display logic
├── pkg/                 # Reusable packages
│   ├── builder/         # JS bundling with esbuild
│   │   ├── builder.go   # BuildShellyScript() function
│   │   └── builder_test.go  # Tests (planned)
│   ├── config/          # Viper configuration (planned)
│   │   └── config.go
│   ├── deployer/        # Shelly RPC communication
│   │   ├── deployer.go      # High-level deploy functions
│   │   ├── upload.go        # Chunked upload logic
│   │   ├── callRPC.go       # Generic RPC caller
│   │   ├── callRPCWithResult.go  # RPC with response parsing
│   │   ├── helper.go        # Struct definitions
│   │   ├── ensureScriptExists.go
│   │   └── abortIfRunning.go
│   └── utils/           # Utilities
│       └── utils.go     # Version info
├── test/                # Test files and fixtures
│   ├── main.js          # Simple test script
│   ├── fixtures/        # Test inputs per feature (planned)
│   └── dist/            # Test build output
├── docs/                # Documentation
└── .opencode/           # AI agent context
```

## Dependencies

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/spf13/cobra` | CLI framework | v1.10.2 |
| `github.com/spf13/viper` | Configuration (planned) | v1.18+ |
| `github.com/evanw/esbuild` | JavaScript bundler | v0.27.2 |

## Data Flow

### Build Command

```
User Input                Processing                     Output
───────────────────────────────────────────────────────────────────
main.js          →   esbuild (bundle + minify)    →   dist/main.js
  ↓                        ↓
imports          →   resolve & inline             →   single file
  ↓                        ↓
ES6+ syntax      →   transpile to ES2015          →   Shelly-compatible
  ↓                        ↓
trailing commas  →   regex removal                →   clean JS
```

### Upload Command (planned)

```
dist/main.js  →  Read file  →  Chunk (1024 bytes)  →  HTTP POST  →  Shelly
                                     ↓
                              append=false (first)
                              append=true (rest)
```

## Key Design Decisions

1. **esbuild for bundling** - Fast, supports ES modules, configurable minification
2. **Chunked uploads** - Shelly has 1024 byte limit per RPC request
3. **Post-processing regex** - Remove trailing commas (Shelly doesn't support them)
4. **Cobra + Viper** - Standard Go CLI stack with hierarchical config

## Configuration Hierarchy

See [CONFIG.md](CONFIG.md) for details.

```
Priority (lowest to highest):
1. Defaults (hardcoded)
2. Global config (~/.config/spgtty/config.yaml)
3. Local config (.spgtty.yaml)
4. Environment variables (SPGTTY_*)
5. CLI flags (--host, --script-id, etc.)
```

## Error Handling

- Builder errors: Wrapped with context, returned to caller
- Deployer errors: Return error, let CLI decide how to display
- CLI: Use `log.Fatalf()` for user-facing errors with exit code 1

## Future Considerations

- Watch mode for development (`spgtty watch`)
- Multiple script support (deploy multiple scripts at once)
- Device profiles (save multiple Shelly configurations)
- TypeScript support (via esbuild)
