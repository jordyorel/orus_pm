# 📦 Orus Package Manager Roadmap

This roadmap outlines the design and development stages of the official package manager and CLI for the Orus programming language. The goal is to enable code sharing, dependency management, and modular application development for `.orus` projects, while clearly distinguishing the CLI tool (`orus`) from the interpreter (`orusc`).

---

## 🔧 Naming Convention

- `orus`: Official CLI and package manager for managing projects, dependencies, testing, and running code.
- `orusc`: The raw Orus interpreter binary, used internally by `orus` or directly by users who want low-level access.

---

## ✅ Stage 1: Foundation & Specification

### 1.1 CLI Interface Design
```
orus init
orus add <package>
orus install
orus run                # Runs main.orus
orus run path/file.orus
orus publish
```

- Advanced:
  ```
  orusc path/to/file.orus  # Run standalone file
  ```

### 1.2 Package Manifest Design (`orus.toml`)
```toml
[package]
name = "example"
version = "0.1.0"
authors = ["Author Name"]
description = "My Orus library"

[dependencies]
math = "0.2.1"
logging = { git = "https://github.com/user/logging" }
```

---

## 🧰 Stage 2: Core Functionality

### 2.1 `init` Command
- Initializes `orus.toml`, `src/main.orus`
- Creates standard folders: `src/`, `.orus/`

### 2.2 `add` and `install`
- Adds to `orus.toml`
- Resolves Git/local/registry deps
- Installs under `.orus/`

---

## 🌐 Stage 3: Dependency Management

- Semantic versioning support
- Lock file generation (`orus.lock`)
- Deduplication and reproducibility

---

## 📚 Stage 4: Module Linking

- Support `use` resolution from `.orus/`
- Detect circular imports
- Flattened namespace index for `use`

---

## 🚀 Stage 5: Package Registry & Publishing

- `orus publish` uploads validated packages
- Optional central registry
- `orus install` reads from registry

---

## 🔒 Stage 6: Security

- SHA-256 verification
- Signed lockfile
- Future: sandboxed execution

---

## 🧪 Stage 7: Testing & CI

- `orus test`: runs test files in `tests/`
- `--frozen` and `--offline` flags

---

## 💡 Future Enhancements

- Binary builds with `orus build`
- VS Code plugin integration
- Native binary deps (WASM?)

---

## 📌 Example Directory Layout

```
my_project/
├── orus.toml
├── orus.lock
├── src/
│   └── main.orus
├── tests/
│   └── math_test.orus
```

---

## 🛠️ CLI Command Summary

| Command                  | Description                          |
|--------------------------|--------------------------------------|
| `orus init`              | Initialize a new Orus project        |
| `orus add <pkg>`         | Add dependency to manifest           |
| `orus install`           | Install dependencies                 |
| `orus run`               | Run main project file                |
| `orus run path.orus`     | Run specific file                    |
| `orusc path.orus`        | Run file with raw interpreter        |
| `orus publish`           | Publish package                      |
| `orus test`              | Run tests                            |

---

## 🗓️ Suggested Timeline

| Milestone                  | Target Date     |
|---------------------------|-----------------|
| CLI Scaffolding           | Week 1          |
| Manifest and `init`       | Week 2          |
| `add`, `install`, lockfile| Week 3–4        |
| `run` and `orusc` split   | Week 5          |
| Registry mockup           | Week 6–7        |
| Public alpha              | Week 8          |

---

Happy hacking with Orus!
