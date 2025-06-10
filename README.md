# Orus CLI (`orus`)

The official CLI tool and package manager for the Orus programming language.

---

## 🚀 Features

- Project scaffolding (`orus init`)
- Dependency management (`orus add`, `orus install`)
- Modular project structure
- Run any `.orus` file
- Integrated testing and publishing

---

## 🧱 Architecture

- `orus`: Manages the project and uses `orusc` internally to execute code.
- `orusc`: Bare interpreter, useful for raw `.orus` execution.

---

## 🛠️ Usage Examples

```bash
# Create new project
orus init

# Add a dependency
orus add math

# Install all dependencies
orus install

# Run the main project file
orus run

# Run a specific file with project context
orus run tests/math_test.orus

# Run file directly with raw interpreter (no dependencies)
orusc tests/math_test.orus
```

---

## 📁 Project Layout

```
my_project/
├── orus.toml
├── orus.lock
├── src/
│   └── main.orus
├── .orus/
│   └── math/
```

---

## 📌 Coming Soon

- `orus test` command
- `orus build` compiler support
- Official package registry
