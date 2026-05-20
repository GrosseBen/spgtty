// Test: Arrow functions
// Expected: Should be transpiled to regular functions for Shelly compatibility

const add = (a, b) => a + b;
const square = x => x * x;
const multiLine = (x) => {
    let result = x * 2;
    return result;
};

print("add(2, 3) =", add(2, 3));
print("square(4) =", square(4));
print("multiLine(5) =", multiLine(5));
