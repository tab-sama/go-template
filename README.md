# 🚀 Go Template

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/tab-sama/go-template)](https://goreportcard.com/report/github.com/tab-sama/go-template)
[![moonrepo](https://img.shields.io/badge/moonrepo-managed-blue.svg)](https://moonrepo.dev)
[![lefthook](https://img.shields.io/badge/lefthook-enabled-green.svg)](https://github.com/evilmartians/lefthook)
[![golangci-lint](https://img.shields.io/badge/golangci--lint-enabled-brightgreen.svg)](https://golangci-lint.run/)

This is a modern GitHub template repository for Go projects. Use this template to create a new Go application with a standardized structure, automated tooling, and best practices built-in.

## 🎯 Why This Template?

- **📁 Standard Layout**: Adopts the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for a clean and organized codebase
- **🔧 Modern Tooling**: Pre-configured with golangci-lint, lefthook, and moonrepo for seamless development
- **📋 Best Practices**: Follows [Effective Go](https://go.dev/doc/effective_go) guidelines with automated formatting and linting
- **🚀 Ready to Deploy**: Optimized build configuration with cross-compilation support
- **⚡ Fast Development**: Moonrepo provides consistent task running and dependency management across the project


## ✨ Features

This template comes pre-configured with:

- 🐹 **Go 1.24.5**: Latest stable Go version with module support
- 🌙 **Moonrepo**: Modern build system for consistent task running and project management
- 🔍 **golangci-lint**: Comprehensive linting with 40+ linters for code quality
- 🪝 **Lefthook**: Fast Git hooks for automated quality checks
- 📁 **Standard Go Layout**: Organized project structure following community standards
- 🚀 **Optimized Builds**: Cross-compilation ready with size optimization flags
- 📝 **GitHub Templates**: CODE_OF_CONDUCT.md, SECURITY.md, and LICENSE included
- ⚡ **Proto Tool Manager**: Automated tool management with [moonrepo proto](https://moonrepo.dev/proto)
- 🔧 **Makefile**: Convenient command shortcuts for common development tasks

## 🚀 Quick Start

### Prerequisites

- [proto](https://moonrepo.dev/proto) - A multi-language version manager that will manage all required tools
- Alternatively, you can install tools separately:
    - [golangci-lint](https://golangci-lint.run/usage/install/)
    - [lefthook](https://github.com/evilmartians/lefthook)
    - [Go](https://golang.org/dl/)

### Installation

1. **Use this template** by clicking the "Use this template" button on GitHub
2. **Clone your new repository**:
   ```bash
   git clone https://github.com/yourusername/your-project-name.git
   cd your-project-name
   ```
3. **Install tools and dependencies**:
   ```bash
   # Install all required tools (golangci-lint, lefthook) using proto
   proto install
   
   # Set up the project (download dependencies and install git hooks)
   moon :setup
   ```

### 🏃‍♂️ Running the Project

```bash
# Run the application directly (development)
moon :run
# or
go run ./cmd/app

# Build and run the production binary
moon :build
./bin/app
```

## 🛠️ Development
### 🛠️ Available Commands

#### Moonrepo Commands

The project defines several tasks in the `moon.yml` file:

- `moon :setup` - Set up the project dependencies
- `moon :clean` - Clean build artifacts
- `moon :format` - Format the code
- `moon :lint` - Lint the code
- `moon :test` - Run tests
- `moon :build` - Build the project
- `moon :run` - Run the application

#### Make Commands

The Makefile provides the following commands (all of which call the corresponding Moonrepo commands):

```shell
make help     # Show available commands
make install  # Install tools and dependencies
make clean    # Clean build artifacts
make format   # Format the code
make lint     # Lint the code
make test     # Run tests
make build    # Build the project
make run      # Run the application
make          # Equivalent to 'make clean build'
```

### 🧹 Code Quality

This template uses **golangci-lint** for comprehensive code quality checks with 40+ linters:

```bash
# Check for linting issues
moon :lint
# or
make lint

# Auto-fix linting issues and format code
moon :format
# or
make format
```

### 🪝 Git Hooks & Conventional Commits

This template includes **Lefthook** for automated Git hooks to ensure code quality:

#### Automatic Quality Checks

Git hooks will automatically run on:

- **Pre-commit**: Format code and run golangci-lint
- **Commit-msg**: Validate commit message format
- **Pre-push**: Final lint and test checks

#### Conventional Commits

All commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```bash
# ✅ Valid commit messages
git commit -m "feat: add user authentication"
git commit -m "fix: resolve memory leak in data processing"
git commit -m "docs: update API documentation"
git commit -m "refactor: simplify error handling logic"

# ❌ Invalid commit messages
git commit -m "add feature"           # Missing type
git commit -m "Fix bug"              # Wrong case
git commit -m "feat!: breaking change" # Use BREAKING CHANGE footer instead
```

**Available commit types:**
- `feat` - New features
- `fix` - Bug fixes
- `docs` - Documentation changes
- `style` - Code style changes (formatting, etc.)
- `refactor` - Code refactoring
- `perf` - Performance improvements
- `test` - Adding or updating tests
- `build` - Build system changes
- `ci` - CI configuration changes
- `chore` - Other changes (maintenance, etc.)
- `revert` - Reverting previous commits

#### Managing Git Hooks

```bash
# Install hooks (automatically runs during `moon :setup`)
lefthook install

# Skip hooks for a single commit (use sparingly)
git commit -m "feat: add feature" --no-verify

# Temporarily disable hooks
lefthook uninstall

# Re-enable hooks
lefthook install
```

### 📁 Project Structure

```
go-template/
├── cmd/
│   └── app/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   └── foo/
│       └── foo.go               # Example business logic package
├── configs/
│   └── golangci.yml             # golangci-lint configuration
├── scripts/
│   └── install.sh               # Installation scripts
├── bin/                         # Built binaries (created by build)
├── .moon/
│   └── plugins/
│       └── golangci-lint.yml    # Moon plugin configuration
├── go.mod                       # Go module definition
├── moon.yml                     # Moonrepo project configuration
├── Makefile                     # Make commands for convenience
├── .lefthook.yml                # Git hooks configuration
├── .prototools                  # Proto tool versions
├── CODE_OF_CONDUCT.md           # Code of conduct
├── SECURITY.md                  # Security policy
├── LICENSE                      # Apache 2.0 license
└── README.md                    # This file
```

## 🔧 Configuration

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linter: `moon :test && moon :lint`
5. Format your code: `moon :format`
6. Commit your changes (`git commit -m 'feat: add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## 📋 Customization

To customize this template for your project:

1. **Update go.mod** with your module path (replace `github.com/tab-sama/go-template`)
2. **Modify the application** in `cmd/app/main.go` to fit your needs
3. **Add your business logic** in the `internal/` directory
4. **Adjust golangci-lint config** in `configs/golangci.yml` as needed
5. **Update moon.yml** with your project name and description
6. **Update this README** with your project-specific information

## 🔒 Security

Please see [SECURITY.md](SECURITY.md) for our security policy and how to report security vulnerabilities.

## 📄 License

This project is licensed under the Apache License—see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Go](https://golang.org/) for the amazing programming language
- [golangci-lint](https://golangci-lint.run/) for comprehensive code quality checks
- [Moonrepo](https://moonrepo.dev/) for consistent build tooling and task management
- [Lefthook](https://github.com/evilmartians/lefthook) for fast and reliable Git hooks

---

**Happy coding! 🎉** If you find this template useful, please give it a ⭐️
