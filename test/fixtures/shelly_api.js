// Test: Shelly-specific API calls
// Expected: Should compile without errors (runtime behavior depends on actual device)
// Note: These APIs only work on actual Shelly devices, not in Node.js

// Get device info
Shelly.call("Shelly.GetDeviceInfo", {}, function(result, error) {
    if (error) {
        print("Error:", error);
    } else {
        print("Device:", result.name || result.id);
    }
});

// Set up a timer
let timerId = Timer.set(5000, false, function() {
    print("Timer fired after 5 seconds!");
});

// Example: Toggle a switch (commented out for safety)
// Shelly.call("Switch.Toggle", {id: 0}, function(result, error) {
//     print("Switch toggled");
// });

// Add an event handler
Shelly.addEventHandler(function(event) {
    print("Event received:", JSON.stringify(event));
});

print("Shelly script initialized");
