# spgtty Improvement Plan

## Overview

This document outlines the planned improvements for spgtty - the Shelly Script Spagetty Generator.

## Goals

1. **Viper Configuration** - Hierarchical config (global + local + CLI flags + env vars)
2. **Refactor Deployer** - Remove standalone `main()`, export `Upload()` function
3. **Implement Upload Command** - Connect `cmd/upload.go` with deployer logic
4. **Implement Init Command** - Create project scaffolding
5. **Add Tests** - Builder tests with fixtures per language feature
6. **Fix Typos** - Correct spelling errors (but keep "Spagetty" wordplay!)

## Phases

### Phase 1: Foundation (Merge + Viper)

- [ ] Merge v0.2.5 into dev branch (brings Cobra structure)
- [ ] Add Viper dependency to go.mod
- [ ] Create `pkg/config/config.go` for configuration handling
- [ ] Update `cmd/root.go` to initialize Viper
- [ ] Bind CLI flags to Viper config keys

### Phase 2: Deployer Refactoring

- [ ] Remove `main()` function from `pkg/deployer/upload.go`
- [ ] Export `Upload(host, scriptID, filePath)` function
- [ ] Clean up unused `flag` import
- [ ] Update error handling to return errors instead of `os.Exit()`

### Phase 3: Upload Command

- [ ] Implement `cmd/upload.go` with proper Viper integration
- [ ] Add `--host` and `--script-id` flags
- [ ] Support reading file path from args or config
- [ ] Add progress output for chunked uploads

### Phase 4: Init Command

- [ ] Implement `cmd/init.go`
- [ ] Create `.spgtty.yaml` template
- [ ] Create `main.js` template with Shelly boilerplate
- [ ] Create directory structure (`src/`, `dist/`)
- [ ] Create `.gitignore`
- [ ] Print helpful next-steps message

### Phase 5: Testing

- [ ] Create `test/fixtures/` directory with JS test files
- [ ] Add `pkg/builder/builder_test.go`
- [ ] Test cases: simple, variables, functions, arrow functions, objects, imports
- [ ] Verify trailing comma removal works

### Phase 6: Cleanup

- [ ] Fix typos in all files (see docs/TYPOS.md or decisions.md)
- [ ] Update README.md
- [ ] Final review and testing

## Status Legend

- [ ] Not started
- [x] Completed
- [~] In progress

## Current Status

**Branch:** dev  
**Last Updated:** 2025-05-20  
**Current Phase:** All phases completed!

### Completed Tasks

- [x] Phase 1: Merged v0.2.5 into dev, added Viper configuration
- [x] Phase 2: Refactored deployer, removed standalone main()
- [x] Phase 3: Implemented upload command with Viper integration
- [x] Phase 4: Implemented init command with project scaffolding
- [x] Phase 5: Added tests and fixtures for all JS language features
- [x] Phase 6: Fixed typos (kept "Spagetty" wordplay)
