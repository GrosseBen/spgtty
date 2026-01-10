# **spgtty** – The Shelly Script Spaghetti Generator

> *"Because writing Shelly scripts by hand is fun... until it isn't. Let computers generate the spaghetti for you!"*

---

## **The Problem**
Shelly scripts are great for small automations, but as soon as things get complex,
it turns into a **nightmare of unmaintainable code**. 
No version control, no structure—just one giant file of chaos. **Painful.**

I trie'd to crate a first use full Programm in ShellyScript, so that I don't need
to flash the hole device. So that I keep the is Part of orginal Firmware features.
After a veiw houres of learing ShellyApis via playing arraound, i found me in a
coppy and past action with IDE, Git and a mose driven way of combinig parts of code
into shallys web UI. With a set of "OK this I will coppy always in front" of my code...
By creating a m(a/e)ss of spagetty code. One Stepp back and a cooffe later
the idea of **spgtty** was born, lets do this a tool.

**Solution:** `spgtty` (pronounced *"spaghetti"*) – your CLI tool that:
- **Automatically generates** Shelly scripts from projects of Js or ts structures
(because computers are better at this than we are).
- **Enforces separation of concerns** (by splitting your logic into clean
files—*after* turning it into glorious spaghetti 🍝).
- **Plays nice with Git**—finally, version control for your scripts without the shame.

---

## **Installation**
Just grab it with go (because we don’t want Node.js bloat, Benjamin):
```bash
go get github.com/GrosseBen/spgtty
```

(Prerequisite: Go installed. If not, go.dev/dl.)

## Features
Currently, spgtty can:

Generate Shelly scripts from actual JS (JavaScript) projects containing folders / imports a shelly script that is reduced to
the supported language features. It Generates only shelly supported supset of js.

**Feels like compile to shelly**

This enables, to crate projects in git. Liberjs for reuse and follow a DRY (don't Repad Yoursleve).

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

- This is a hobby project. If it crashes, resart  shally and file a Bug.
- **Target: Shelly Gen 2+** (because older ones don’t support scripting).
