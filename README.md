# 🚀 Go Project Template

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> **Note**: This is a GitHub template repository. Use this template to create a new Go project with a standardized structure and tooling.

## 📋 How to Use This Template

1. Click the "Use this template" button at the top of this repository page
2. Name your project and provide a description
3. Clone your new repository
4. Customize the project by:
   - Updating this README with your project name and description
   - Modifying the code in `cmd/main.go` to implement your application
   - Updating the `go.mod` file with your module path

## 📚 About This Template

- This template adopts the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for a clean
  and organized codebase.

- We recommend following the guidelines outlined in [Effective Go](https://go.dev/doc/effective_go). Most formatting can
  be automatically handled using the `go fmt` tool.

- For projects involving API design, we highly recommend referring to the
  [Google API Design Guide](https://cloud.google.com/apis/design) for best practices.

- This template uses [Moonrepo](https://moonrepo.dev/) as its primary build system. Moonrepo provides consistent tooling,
  task running, and dependency management across the project.

## 🔧 Setup

### Using Moonrepo (Recommended)

To set up the project with Moonrepo:

```shell
# Install Moonrepo's proto tool if you don't have it
curl -fsSL https://moonrepo.dev/install/proto.sh | bash

# Install required tools and dependencies
proto use
moon :setup
```

### Using Make (Compatibility)

A Makefile is provided for compatibility with traditional workflows, but it's just a wrapper around Moonrepo commands:

```shell
make install
```

### Manual Setup

You can also manually install dependencies with:

```shell
go mod download
```

### Running the Code

To run the code:

```shell
# Using Moonrepo
moon :run

# Using Make
make run

# Using Go directly
go run ./cmd
```

## 🔨 Build

### Using Moonrepo (Recommended)

To build the project with Moonrepo:

```shell
moon :build
```

The binary file will be created in the `bin` directory.

### Using Make (Compatibility)

You can also use Make, which will call the Moonrepo build command:

```shell
make
# or
make build
```

## 🛠️ Available Commands

### Moonrepo Commands

The project defines several tasks in the `moon.yml` file:

- `moon :setup` - Set up the project dependencies
- `moon :clean` - Clean build artifacts
- `moon :format` - Format the code
- `moon :lint` - Lint the code
- `moon :test` - Run tests
- `moon :build` - Build the project
- `moon :run` - Run the application

### Make Commands

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
