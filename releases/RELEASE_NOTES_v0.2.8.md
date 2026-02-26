# spgtty v0.2.8 Release Notes

## 🎉 New Features

### Safe Project Initialization
- **Non-destructive init command**: `spgtty init` now preserves existing `main.js` files
- **Smart file handling**: Checks if `main.js` exists before creating a new one
- **Clear user feedback**: Shows "ℹ️ main.js already exists - preserving your code" when file exists

## 🔧 Improvements

### User Experience
- **Safer workflow**: No risk of accidentally overwriting existing code
- **Better messaging**: Clear indication of what files were created vs. preserved
- **Flexible initialization**: Can be run multiple times without losing work

### Code Quality
- **Better file handling**: Uses `os.Stat()` to check file existence
- **Error handling**: Proper handling of file system errors
- **Clean code**: Well-structured conditional logic

## 📝 Documentation

- **Organized release notes**: Moved to `releases/` directory for better project structure
- **Clear change logging**: Detailed description of new safe initialization feature

## 🚀 Usage Examples

```bash
# Initialize new project (creates main.js)
spgtty init --device shellyplus1pm-abc123

# Re-run init on existing project (preserves main.js)
spgtty init --device shellyplus1pm-xyz789

# Output when main.js exists:
# ℹ️  main.js already exists - preserving your code
# ✅ Project initialized successfully!
# 📄 Created .spgtty config file
# 📁 Created dist/ directory
```

## 🎯 Upgrade Notes

No breaking changes. The `init` command is now safer and more user-friendly.

## 📊 Impact

- **User safety**: Eliminates risk of accidental code loss
- **Workflow improvement**: Can safely re-run init command
- **Professional quality**: Better file handling practices

## 🙏 Contributors

Thanks to all contributors who helped make this release possible!

---
*Released on 2026-02-26*
*Part of the v0.2.x improvement series*