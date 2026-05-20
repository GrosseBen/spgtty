# Shelly RPC API Reference

This document covers the Shelly Gen2+ RPC API used by spgtty for script deployment.

## Overview

Shelly Gen2+ devices expose an HTTP-based RPC API at:

```
http://<device-ip>/rpc/<method>
```

Or via POST to:

```
http://<device-ip>/rpc/
```

## Authentication

If the device has authentication enabled:

```
http://admin:<password>@<device-ip>/rpc/
```

**Note:** spgtty currently doesn't support authentication. Add device to network without auth for development.

## Script Methods

### Script.List

List all scripts on the device.

**Request:**
```bash
curl http://192.168.1.100/rpc/Script.List
```

**Response:**
```json
{
  "scripts": [
    {"id": 1, "name": "script1", "enable": true, "running": false},
    {"id": 2, "name": "script2", "enable": false, "running": false}
  ]
}
```

### Script.Create

Create a new script slot.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.Create",
  "params": {
    "name": "my-script"
  }
}
```

**Response:**
```json
{
  "id": 1,
  "result": {
    "id": 3
  }
}
```

### Script.PutCode

Upload code to a script. This is the main method used by spgtty.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.PutCode",
  "params": {
    "id": 1,
    "code": "print('hello');",
    "append": false
  }
}
```

**Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| `id` | int | Script slot (1-10) |
| `code` | string | JavaScript code (max ~1024 bytes per request) |
| `append` | bool | `false` = overwrite, `true` = append to existing |

**Response:**
```json
{
  "id": 1,
  "result": {
    "len": 15
  }
}
```

### Chunked Upload

Shelly has a limit of approximately 1024 bytes per request. For larger scripts, use chunked upload:

```go
// First chunk - overwrite existing
{"id": 1, "method": "Script.PutCode", "params": {"id": 1, "code": "...", "append": false}}

// Subsequent chunks - append
{"id": 1, "method": "Script.PutCode", "params": {"id": 1, "code": "...", "append": true}}
{"id": 1, "method": "Script.PutCode", "params": {"id": 1, "code": "...", "append": true}}
// ... continue until all code is uploaded
```

### Script.GetCode

Retrieve the code from a script.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.GetCode",
  "params": {
    "id": 1
  }
}
```

**Response:**
```json
{
  "id": 1,
  "result": {
    "data": "print('hello');",
    "left": 0
  }
}
```

### Script.Start

Start a script.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.Start",
  "params": {
    "id": 1
  }
}
```

### Script.Stop

Stop a running script.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.Stop",
  "params": {
    "id": 1
  }
}
```

### Script.GetStatus

Get script status.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.GetStatus",
  "params": {
    "id": 1
  }
}
```

**Response:**
```json
{
  "id": 1,
  "result": {
    "id": 1,
    "running": true,
    "mem_used": 1234,
    "mem_peak": 2345,
    "mem_free": 5678
  }
}
```

### Script.SetConfig

Configure script settings (enable on boot, etc.).

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.SetConfig",
  "params": {
    "id": 1,
    "config": {
      "enable": true
    }
  }
}
```

### Script.Delete

Delete a script.

**Request:**
```json
POST /rpc/
{
  "id": 1,
  "method": "Script.Delete",
  "params": {
    "id": 1
  }
}
```

## Shelly JavaScript API

These are the functions available within Shelly scripts (not RPC methods).

### print()

Output to console/log.

```javascript
print("Hello World");
print("Value:", 42, "Name:", "test");
```

### Shelly.call()

Call RPC methods from within a script.

```javascript
Shelly.call("Switch.Set", {id: 0, on: true}, function(result, error) {
    if (error) {
        print("Error:", error);
    } else {
        print("Success:", JSON.stringify(result));
    }
});
```

### Shelly.addEventHandler()

Handle device events.

```javascript
Shelly.addEventHandler(function(event) {
    print("Event:", JSON.stringify(event));
});
```

### Timer.set()

Create timers.

```javascript
// One-shot timer (5 seconds)
Timer.set(5000, false, function() {
    print("Timer fired!");
});

// Repeating timer (every 10 seconds)
Timer.set(10000, true, function() {
    print("Tick!");
});
```

### Timer.clear()

Cancel a timer.

```javascript
let timerId = Timer.set(5000, false, function() {});
Timer.clear(timerId);
```

### HTTP Requests

```javascript
Shelly.call("HTTP.GET", {url: "http://example.com/api"}, function(result) {
    print("Response:", result.body);
});
```

## Limitations

### Script Size

- Maximum script size varies by device (typically 16KB-32KB)
- Use minification to reduce size

### JavaScript Engine

Shelly uses a limited JavaScript engine (mJS or similar):

| Feature | Supported |
|---------|-----------|
| `var`, `let`, `const` | ✅ |
| Functions | ✅ |
| Arrow functions | ✅ (may be transpiled) |
| Objects, Arrays | ✅ |
| JSON | ✅ |
| Trailing commas | ❌ |
| async/await | ❌ |
| Promises | ❌ |
| Classes | ⚠️ Limited |
| Template literals | ⚠️ May work |
| Modules (import/export) | ❌ (bundled by spgtty) |

### Memory

- Scripts have limited memory
- Monitor with `Script.GetStatus` → `mem_used`, `mem_free`

### Execution Time

- Long-running synchronous code may cause watchdog resets
- Use timers for periodic tasks

## Error Handling

**RPC Error Response:**
```json
{
  "id": 1,
  "error": {
    "code": -103,
    "message": "Script not found"
  }
}
```

**Common Error Codes:**
| Code | Meaning |
|------|---------|
| -103 | Script not found |
| -104 | Script already running |
| -105 | Script not running |
| -106 | Invalid script ID |

## Useful Links

- [Shelly Script Documentation](https://shelly-api-docs.shelly.cloud/gen2/Scripts/)
- [Shelly Gen2 API Reference](https://shelly-api-docs.shelly.cloud/gen2/)
- [Shelly Community Forum](https://www.shelly-support.eu/)
