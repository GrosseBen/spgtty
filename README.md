# **spgtty** – The Shelly Script Spaghetti Generator

> *"Because writing Shelly scripts by hand is fun... until it isn't. Let computers generate the spaghetti for you!"*

---

## **The Problem**
Shelly scripts are great for small automations, but as soon as things get complex, they turn into a **nightmare of unmaintainable code**.
No version control, no structure—just one giant file of chaos. **Painful.**

I tried to create a first *useful* program in ShellyScript so I wouldn't have to flash the entire device and could keep the original firmware features.
After a few hours of fiddling with the Shelly API, I found myself stuck in a copy-paste loop between my IDE, Git, and the Shelly web UI,
with a mental checklist of *"Okay, I'll always copy this snippet first"*—while accidentally building a glorious mess of spaghetti code.
One step back, a sip of coffee later, and the idea for **spgtty** was born: *Let's make a tool for this.*

**Solution:** `spgtty` (pronounced *"spaghetti"*)—your CLI tool that:
- **Automatically generates** Shelly scripts from JS/TS projects (because computers are better at this than we are).
- **Enforces separation of concerns** (by splitting your logic into clean files—*after* turning it into glorious spaghetti 🍝).
- **Plays nice with Git**—finally, version control for your scripts without the shame.

---

## **Installation**
Grab it with Go (because we don't want Node.js bloat, right?):
```bash
go install github.com/GrosseBen/spgtty
```

(Prerequisite: Go installed. If not, go.dev/dl.)

or without Go but with Homebrew:
```bash
brew install GrosseBen/spgtty/spgtty
```

or

download []()


## Features
Currently, spgtty can:

Generate Shelly scripts from actual JS (JavaScript) projects containing folders/imports. 
It generates only Shelly-supported subset of JS.

**Feels like compile to Shelly**

This enables you to create projects in Git, libraries for reuse, and follow DRY (Don't Repeat Yourself).

Minify output (or not, with -no-minify for debugging).
Write to dist/main.js by default (or wherever you want: -out path/to/your/script.js).

## howToUse

```bash
sh-3.2$ cat main.js
function main() {
  print("hallo welt");
}
main();
sh-3.2$ ./spgtty
2026/01/09 22:57:40 ✅ Code nach dist/main.js geschrieben (38 Bytes)
```
... and copy ```dist/main.js``` to shelly and run it.

## Why This?

- For the home automation community: So no one has to reinvent the wheel (or produce real spaghetti code).
- One binary: A single binary that does it all—no npm install with 500 dependencies.
- Hobby project: Yes, this is a for-fun thing. Merge requests are welcome, but please bring humor and patience.

## Contributing
Want to help? Awesome!

- Fork the repo.
- Do your thing (but please, no real spaghetti code in Go, okay?).
- Open a merge request—no pressure, this is all about fun.

## ⚠️ Important:

- This is a hobby project. If it crashes, restart your Shelly and file a bug.
- **Target: Shelly Gen2+** (because older ones don't support scripting).
