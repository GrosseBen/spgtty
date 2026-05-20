# Configuration

spgtty uses [Viper](https://github.com/spf13/viper) for hierarchical configuration management.

## Configuration Hierarchy

Settings are loaded in this order (later overrides earlier):

1. **Defaults** - Built into spgtty
2. **Global config** - `~/.config/spgtty/config.yaml`
3. **Local config** - `.spgtty.yaml` in current directory
4. **Environment variables** - `SPGTTY_*`
5. **CLI flags** - `--host`, `--script-id`, etc.

## Config File Format

### `.spgtty.yaml` (local project config)

```yaml
# Shelly device settings
shelly:
  host: "192.168.1.100"    # IP or hostname of Shelly device
  script_id: 1             # Script slot (1-10)

# Build settings
build:
  entry: "main.js"         # Entry point file
  output: "dist/main.js"   # Output file path
  minify: true             # Enable minification
```

### `~/.config/spgtty/config.yaml` (global config)

```yaml
# Default Shelly device (used if no local config)
shelly:
  host: "shelly-living-room.local"
  script_id: 1

# Default build settings
build:
  minify: true
```

## Environment Variables

All config keys can be set via environment variables with `SPGTTY_` prefix:

| Variable | Config Key | Example |
|----------|------------|---------|
| `SPGTTY_SHELLY_HOST` | `shelly.host` | `192.168.1.100` |
| `SPGTTY_SHELLY_SCRIPT_ID` | `shelly.script_id` | `1` |
| `SPGTTY_BUILD_ENTRY` | `build.entry` | `src/main.js` |
| `SPGTTY_BUILD_OUTPUT` | `build.output` | `dist/app.js` |
| `SPGTTY_BUILD_MINIFY` | `build.minify` | `true` |

### Example Usage

```bash
# Set host via environment
export SPGTTY_SHELLY_HOST="192.168.1.100"
spgtty upload

# Or inline
SPGTTY_SHELLY_HOST="192.168.1.100" spgtty upload
```

## CLI Flags

CLI flags have the highest priority and override all other settings.

### Global Flags (available on all commands)

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--version` | `-v` | Show version info | - |
| `--help` | `-h` | Show help | - |

### Build Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--notMinimize` | `-m` | Disable minification | `false` |
| `--output` | `-o` | Output file path | `dist/main.js` |

### Upload Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--host` | `-H` | Shelly IP/hostname | (from config) |
| `--script-id` | | Script slot (1-10) | `1` |

## Default Values

```go
// Defined in pkg/config/config.go
viper.SetDefault("build.entry", "main.js")
viper.SetDefault("build.output", "dist/main.js")
viper.SetDefault("build.minify", true)
viper.SetDefault("shelly.script_id", 1)
```

## Precedence Example

Given these configurations:

**Global config** (`~/.config/spgtty/config.yaml`):
```yaml
shelly:
  host: "192.168.1.50"
  script_id: 2
```

**Local config** (`.spgtty.yaml`):
```yaml
shelly:
  host: "192.168.1.100"
```

**Environment**:
```bash
export SPGTTY_SHELLY_SCRIPT_ID=3
```

**CLI**:
```bash
spgtty upload --host 192.168.1.200
```

**Result**:
- `host` = `192.168.1.200` (from CLI flag - highest priority)
- `script_id` = `3` (from environment - overrides config files)

## Creating Config with `spgtty init`

The `init` command creates a `.spgtty.yaml` template:

```bash
spgtty init
# Creates .spgtty.yaml with placeholder values

spgtty init my-project
# Creates my-project/ directory with .spgtty.yaml inside
```
