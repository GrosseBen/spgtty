// Test: ES module imports
// Expected: Imports should be bundled into a single file (no import statements in output)

import { helper, CONSTANT, multiply } from "./helper.js";

print("Helper says:", helper());
print("Constant value:", CONSTANT);
print("3 * 4 =", multiply(3, 4));
