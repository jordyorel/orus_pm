# ka - The Orus Package Manager

`ka` is a next-generation package manager for the Orus programming language, designed to be fast, reliable, and easy to use.

## Features

* **Dependency Management:** Easily add, remove, and upgrade packages.
* **Project Initialization:** Quickly scaffold a new Orus project.
* **Build System Integration:** Seamlessly build, run, and test your Orus code.
* **Registry and Publishing:** Publish your own packages to the Orus registry.
* **Offline and Local Caching:** Work offline and install packages from local paths or Git URLs.
* **Scripts and Automation:** Define and run custom scripts for your project.
* **Workspaces and Monorepo Support:** Manage multiple related packages in a single repository.
* **Package Locking:** Ensure deterministic and reproducible builds.

## Installation

To install `ka`, you need to have Go installed. Then, run the following command:

```bash
go get -u github.com/orus-lang/ka
```

## Usage

### Initialize a new project

```bash
ka init
```

This will create a new Orus project with the following structure:

```
.
├── ka.toml
├── src/
│   └── main.orus
└── README.md
```

### Add a dependency

```bash
ka add <package>
```

### Build your project

```bash
ka build
```

### Run your project

```bash
ka run
```

### Test your project

```bash
ka test
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

`ka` is licensed under the MIT License.
