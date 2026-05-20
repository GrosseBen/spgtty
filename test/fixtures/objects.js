// Test: Object literals with trailing commas
// Expected: Trailing commas MUST be removed (Shelly doesn't support them)

const config = {
    name: "test",
    value: 42,
    nested: {
        a: 1,
        b: 2,
    },
};

print("Name:", config.name);
print("Value:", config.value);
print("Nested a:", config.nested.a);
