# spgtty Configuration

spgtty uses a JSON configuration file named `.spgtty` to store device settings and build options.

## Configuration File Format

The `.spgtty` file should be a valid JSON file with the following structure:

```json
{
  "device": "shellyplus1pm-123456",  // Required: Shelly device ID (used for mDNS)
  "ip": "192.168.33.1",             // Optional: Override device IP
  "minified": false,                 // Optional: Minify output (default: false)
  "scriptId": 1                      // Optional: Script ID (default: 1)
}
```

## Configuration File Locations

spgtty searches for configuration files in the following order (first found wins):

1. `./.spgtty` - Current directory
2. `~/.spgtty` - User home directory
3. `/etc/spgtty/.spgtty` - System-wide configuration

## Configuration Priority

Configuration values are merged with the following priority (highest to lowest):

1. **CLI flags** - `--device`, `--ip`, `--minified`, `--script-id`
2. **Config file** - Values from `.spgtty` file
3. **Defaults** - Hardcoded default values

## Field Reference

### `device` (required)
- **Type**: String
- **Description**: Shelly device ID used for mDNS resolution (e.g., `shellyplus1pm-123456.local`)
- **Example**: `"shellyplus1pm-123456"`
- **CLI override**: `--device` flag

### `ip` (optional)
- **Type**: String
- **Description**: Override the device IP address. If not specified, mDNS will be used.
- **Example**: `"192.168.33.1"`
- **CLI override**: `--ip` flag

### `minified` (optional)
- **Type**: Boolean
- **Description**: Whether to minify the output JavaScript code.
- **Default**: `false`
- **Example**: `true` or `false`
- **CLI override**: `--minified` flag

### `scriptId` (optional)
- **Type**: Integer
- **Description**: Script ID for the Shelly device.
- **Default**: `1`
- **Example**: `1`, `2`, `3`, etc.
- **CLI override**: `--script-id` flag

## Getting Started

To create a new project with a default configuration file:

```bash
# Initialize a new project (interactive)
spgtty init

# Initialize with specific device
spgtty init --device shellyplus1pm-abc123

# Initialize with all options
spgtty init --device shellyplus1pm-abc123 --ip 192.168.1.100 --minified true --script-id 2
```

## Example Usage

### Basic configuration
```json
{
  "device": "shellyplus1pm-abc123",
  "minified": true
}
```

### Full configuration with all options
```json
{
  "device": "shellyplus1pm-abc123",
  "ip": "192.168.1.100",
  "minified": true,
  "scriptId": 2
}
```

### CLI override example
```bash
# Use config file values but override device
spgtty build --device shellyplus1pm-xyz789

# Use config file values but override minification
spgtty build --minified false
```

## Error Handling

If no configuration file is found, spgtty will:
- Require the `--device` flag to be specified
- Use default values for optional fields
- Continue with CLI flag values if provided

If a configuration file is found but contains invalid JSON, spgtty will display an error message and exit.