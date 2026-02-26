// Shelly Gen2+ JavaScript example
// This is a basic example to get you started

print("Hello from Shelly device: " + Shelly.getDeviceInfo().id);

// Example: Toggle relay 0 every 5 seconds
// Timer.set(5000, true, function() {
//   let state = Shelly.getComponentStatus("switch:0").output;
//   Shelly.call("switch.set", {id: 0, on: !state});
// });
