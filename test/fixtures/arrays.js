// Test: Arrays with trailing commas and array methods
// Expected: Trailing commas removed, basic array methods should work

const items = [1, 2, 3,];  // trailing comma
const doubled = [];

for (let i = 0; i < items.length; i++) {
    doubled.push(items[i] * 2);
}

print("Items:", JSON.stringify(items));
print("Doubled:", JSON.stringify(doubled));
