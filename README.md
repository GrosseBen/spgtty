# **spgtty** – The Shelly Script Spaghetti Generator

> *"Because writing Shelly scripts by hand is fun... until it isn't. Let computers generate the spaghetti for you!"*

---

## **The Problem**
Shelly scripts are great for small automations, but as soon as things get complex, they turn into a **nightmare of unmaintainable code**.
No version control, no structure—just one giant file of chaos. **Painful.**

I tried to create a first *useful* program in ShellyScript so I wouldn’t have to flash the entire device and could keep the original firmware features.
After a few hours of fiddling with the Shelly API, I found myself stuck in a copy-paste loop between my IDE, Git, and the Shelly web UI,
with a mental checklist of *"Okay, I’ll always copy this snippet first"*—while accidentally building a glorious mess of spaghetti code.
One step back, a sip of coffee later, and the idea for **spgtty** was born: *Let’s make a tool for this.*

**Solution:** `spgtty` (pronounced *"spaghetti"*)—your CLI tool that:
- **Automatically generates** Shelly scripts from JS/TS projects (because computers are better at this than we are).
- **Enforces separation of concerns** (by splitting your logic into clean files—*after* turning it into glorious spaghetti 🍝).
- **Plays nice with Git**—finally, version control for your scripts without the shame.

---

## **Installation**
Grab it with Go (because we don’t want Node.js bloat, right?):
```bash
go install github.com/GrosseBen/spgtty
