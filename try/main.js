// Shelly Gen2+ JavaScript example
// This is a basic example to get you started

const AGrateGreeting = "Hello from Shelly device: ";
function sayHello() {
  print(AGrateGreeting + Shelly.getDeviceInfo().id);
}

sayHello();

// Example: Toggle relay 0 every 5 seconds
// Timer.set(5000, true, function() {
//   let state = Shelly.getComponentStatus("switch:0").output;
//   Shelly.call("switch.set", {id: 0, on: !state});
// });
