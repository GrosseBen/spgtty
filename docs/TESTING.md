# Testing

## Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package tests
go test -v ./pkg/builder/...

# Run with coverage
go test -cover ./...
```

## Test Structure

```
spgtty/
├── pkg/
│   └── builder/
│       ├── builder.go
│       └── builder_test.go    # Unit tests
└── test/
    ├── main.js                # Simple integration test
    └── fixtures/              # Test inputs per feature
        ├── simple.js
        ├── variables.js
        ├── functions.js
        ├── arrow.js
        ├── objects.js
        ├── arrays.js
        ├── imports.js
        ├── helper.js          # Module for imports.js
        └── shelly_api.js
```

## Test Fixtures

Each fixture tests a specific JavaScript language feature. This makes it easy to identify what works and what doesn't with Shelly's JavaScript engine.

### `simple.js` - Basic Output

```javascript
print("hello Shelly");
```
**Tests:** Basic `print()` function works.

### `variables.js` - Variable Declarations

```javascript
let name = "Shelly";
const VERSION = "1.0";
var legacy = true;
print(name, VERSION, legacy);
```
**Tests:** `let`, `const`, `var` declarations.

### `functions.js` - Function Declarations

```javascript
function greet(name) {
    return "Hello " + name;
}
print(greet("World"));
```
**Tests:** Function declaration, parameters, return values.

### `arrow.js` - Arrow Functions

```javascript
const add = (a, b) => a + b;
const square = x => x * x;
const multiLine = (x) => {
    return x * 2;
};
print(add(2, 3), square(4), multiLine(5));
```
**Tests:** Arrow function syntax variations.

### `objects.js` - Object Literals

```javascript
const config = {
    name: "test",
    value: 42,
    nested: {
        a: 1,
        b: 2,
    },
};
print(config.name, config.nested.a);
```
**Tests:** Object literals, nested objects, trailing commas (must be removed!).

### `arrays.js` - Arrays

```javascript
const items = [1, 2, 3,];  // trailing comma
const doubled = items.map(x => x * 2);
print(JSON.stringify(doubled));
```
**Tests:** Array literals, trailing commas, array methods.

### `imports.js` + `helper.js` - ES Modules

```javascript
// imports.js
import { helper, CONSTANT } from "./helper.js";
print(helper(), CONSTANT);

// helper.js
export function helper() {
    return "I'm a helper";
}
export const CONSTANT = 42;
```
**Tests:** ES module bundling - imports should be resolved and inlined.

### `shelly_api.js` - Shelly-Specific APIs

```javascript
// These functions exist on Shelly devices
Shelly.call("Switch.GetStatus", {id: 0}, function(result) {
    print(JSON.stringify(result));
});

Timer.set(1000, true, function() {
    print("Timer fired!");
});
```
**Tests:** Shelly API calls compile without errors (runtime testing requires actual device).

## Test Cases

### `pkg/builder/builder_test.go`

```go
package builder

import (
    "strings"
    "testing"
)

func TestBuildSimpleScript(t *testing.T) {
    result, err := BuildShellyScript("../../test/fixtures/simple.js", true)
    if err != nil {
        t.Fatalf("Build failed: %v", err)
    }
    if len(result) == 0 {
        t.Fatal("Empty output")
    }
    if !strings.Contains(string(result), "print") {
        t.Error("Output should contain print statement")
    }
}

func TestBuildWithMinify(t *testing.T) {
    minified, _ := BuildShellyScript("../../test/fixtures/simple.js", true)
    unminified, _ := BuildShellyScript("../../test/fixtures/simple.js", false)
    
    if len(minified) >= len(unminified) {
        t.Errorf("Minified (%d bytes) should be smaller than unminified (%d bytes)", 
            len(minified), len(unminified))
    }
}

func TestTrailingCommaRemoval(t *testing.T) {
    result, _ := BuildShellyScript("../../test/fixtures/objects.js", false)
    output := string(result)
    
    if strings.Contains(output, ",}") {
        t.Error("Trailing commas before } should be removed")
    }
    if strings.Contains(output, ",]") {
        t.Error("Trailing commas before ] should be removed")
    }
}

func TestBuildWithImports(t *testing.T) {
    result, err := BuildShellyScript("../../test/fixtures/imports.js", true)
    if err != nil {
        t.Fatalf("Build with imports failed: %v", err)
    }
    
    output := string(result)
    if strings.Contains(output, "import ") {
        t.Error("Imports should be bundled, not left as import statements")
    }
    if strings.Contains(output, "export ") {
        t.Error("Exports should be removed after bundling")
    }
}

func TestBuildNonExistentFile(t *testing.T) {
    _, err := BuildShellyScript("nonexistent.js", true)
    if err == nil {
        t.Error("Expected error for non-existent file")
    }
}

func TestBuildAllFixtures(t *testing.T) {
    fixtures := []string{
        "simple.js",
        "variables.js", 
        "functions.js",
        "arrow.js",
        "objects.js",
        "arrays.js",
        "imports.js",
        "shelly_api.js",
    }
    
    for _, fixture := range fixtures {
        t.Run(fixture, func(t *testing.T) {
            path := "../../test/fixtures/" + fixture
            result, err := BuildShellyScript(path, true)
            if err != nil {
                t.Errorf("Failed to build %s: %v", fixture, err)
            }
            if len(result) == 0 {
                t.Errorf("Empty output for %s", fixture)
            }
        })
    }
}
```

## Adding New Tests

### When to add a test

1. **New feature** - Add fixture demonstrating the feature
2. **Bug found** - Add fixture that reproduces the bug
3. **Shelly limitation discovered** - Document in fixture comments

### How to add a test

1. Create fixture file in `test/fixtures/`
2. Add test case in `builder_test.go`
3. Run `go test -v ./pkg/builder/...`
4. If it fails on Shelly but passes in tests, document the limitation

### Example: Adding a test for template literals

```javascript
// test/fixtures/template_literals.js
const name = "World";
const greeting = `Hello ${name}!`;  // Template literal
print(greeting);
```

```go
// Add to builder_test.go
func TestTemplateLiterals(t *testing.T) {
    result, err := BuildShellyScript("../../test/fixtures/template_literals.js", true)
    if err != nil {
        t.Fatalf("Template literals failed: %v", err)
    }
    // Template literals should be transpiled to string concatenation
    // because Shelly might not support them
}
```

## Known Limitations

Document any discovered Shelly JavaScript limitations here:

| Feature | Status | Notes |
|---------|--------|-------|
| `let`, `const` | ✅ Works | - |
| Arrow functions | ✅ Works | Transpiled to ES5 |
| Template literals | ⚠️ Test | May need transpilation |
| async/await | ❌ No | Use callbacks instead |
| Trailing commas | ❌ No | Removed by post-processing |
| ES modules | ✅ Works | Bundled by esbuild |
| Classes | ⚠️ Test | May need transpilation |

## CI Integration

The GitHub Actions workflow runs tests on every push:

```yaml
# .github/workflows/go.yml
- run: go test -v ./...
```
