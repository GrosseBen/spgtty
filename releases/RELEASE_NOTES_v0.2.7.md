# spgtty v0.2.7 Release Notes

## 🎉 New Features

### Configuration System Overhaul
- **Multi-location config support**: `.spgtty` files are now searched in multiple locations:
  - `./.spgtty` (current directory)
  - `~/.spgtty` (user home directory)
  - `/etc/spgtty/.spgtty` (system-wide)

- **Config merging strategy**: Proper priority order implemented:
  - CLI flags > config file > defaults

- **`spgtty init` command**: New command to initialize projects with default configuration:
  ```bash
  spgtty init --device shellyplus1pm-abc123
  ```
  Creates `.spgtty` config file, example `main.js`, and `dist/` directory

### Improved Developer Experience

- **Enhanced `justfile`**: Cleaner recipes with emoji icons and better output:
  - `just build` - Build the binary
  - `just test` - Run tests (now with actual test coverage!)
  - `just try` - Try spgtty in isolated workspace (now shows version info)
  - `just clean` - Clean up try directory

- **Version display**: `just try` now shows version information

### Testing

- **Comprehensive test suite**: Added tests for core functionality:
  - Configuration loading and validation
  - Config merging with CLI flags
  - JavaScript building and minification

## 🐛 Bug Fixes

- **JSON parsing**: Fixed config file parsing errors by removing comments from `.spgtty` files
- **Error handling**: Fixed error wrapping in version.go
- **Path resolution**: Fixed issues with justfile path resolution

## 📝 Documentation

- **Complete config documentation**: Added `doc/config.md` with:
  - Configuration file format
  - File locations and priority
  - Field reference with examples
  - Usage examples

## 🔧 Technical Improvements

- **Clean separation**: `try/` directory is now completely isolated from project root
- **Better error messages**: More descriptive error messages throughout
- **Code quality**: Improved validation and error handling

## 📊 Statistics

- **Test coverage**: 7 tests covering core functionality
- **Code quality**: Improved validation, error handling, and user experience
- **Documentation**: Complete config documentation added

## 🚀 Usage Examples

```bash
# Initialize a new project
spgtty init --device shellyplus1pm-abc123

# Build using config file
spgtty build

# Override config with CLI flags
spgtty build --device shellyplus1pm-xyz789 --minified true

# Try it out
just try

# Run tests
just test
```

## 🎯 Upgrade Notes

No breaking changes. Existing configurations will continue to work.

## 🙏 Contributors

Thanks to all contributors who helped make this release possible!

---
*Released on 2026-02-26*