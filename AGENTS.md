# AGENTS.md - GoFrame (GF) Project

This file provides guidance to AI agents working on the GoFrame project.

## Project Overview

GoFrame (GF) is a modular, powerful, high-performance, and enterprise-class Go application development framework. It provides:

- **Language**: Go (Golang) 1.23+
- **Framework**: Modular application framework
- **Key Features**: Web server, ORM, caching, logging, configuration, etc.
- **Architecture**: Modular and extensible design

## Key Configuration Files

- `go.mod` - Go module definition and dependencies
- `go.sum` - Dependency checksums
- `Makefile` - Build automation scripts
- `.golangci.yml` - GolangCI-Lint configuration
- `README.MD` - Project documentation

## Build and Test Commands

### Installation
```bash
# Install dependencies
go mod download

# Install development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Development
```bash
# Build the project
go build ./...

# Run specific packages
go run cmd/main.go

# Make targets (see Makefile)
make tidy    # Clean and format code
make test    # Run tests
make build   # Build project
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./path/to/package

# Linting
golangci-lint run
```

### Makefile Commands
```bash
# Common make commands
make tidy      # Format and clean code
make test      # Run test suite
make build     # Build the project
make version   # Update version information
```

## Code Style Guidelines

- **Go Standards**: Follow official Go code review comments
- **Formatting**: Use `gofmt` for code formatting
- **Linting**: GolangCI-Lint configuration in `.golangci.yml`
- **Naming**: Use camelCase for variables, PascalCase for exported types
- **Error Handling**: Explicit error handling (no panic for expected errors)
- **Documentation**: Add godoc comments for exported functions/types

## Project Structure

```
cmd/                    # Command-line applications
container/              # Container/dependency injection
crypto/                 # Cryptography utilities
database/               # Database drivers and ORM
debug/                  # Debugging utilities
encoding/               # Data encoding/decoding
errors/                 # Error handling
frame/                  # Framework core
internal/               # Internal packages
net/                    # Network utilities
os/                     # Operating system utilities
text/                   # Text processing
util/                   # Utilities
test/                   # Test files
```

## Testing Instructions

- Tests are located in `test/` directory and alongside source files
- Use Go's built-in testing package
- Follow table-driven test pattern for complex logic
- Mock external dependencies when possible
- Test both happy paths and error cases

## Security Considerations

- **Dependencies**: Regularly update dependencies with `go get -u`
- **Input Validation**: Validate all external inputs
- **Error Messages**: Don't expose sensitive information in errors
- **Configuration**: Keep secrets out of code (use environment variables)
- **SQL Injection**: Use parameterized queries with the ORM

## Deployment

- The framework is designed as a library, not a standalone application
- Applications using GF should have their own deployment processes
- Follow Go deployment best practices
- Use static linking for production builds

## Git Conventions

- Follow standard Git practices
- Use meaningful commit messages
- Keep commits focused and atomic
- Use branches for features/bugfixes
- Regularly update from main branch

## CI/CD

The project likely uses GitHub Actions or similar:
- Automated testing on push/PR
- Linting checks
- Build verification
- Release automation

## Performance Considerations

- GF is designed for high performance
- Avoid unnecessary allocations in hot paths
- Use sync.Pool for object reuse when appropriate
- Profile before optimizing
- Follow Go performance best practices

## Internationalization

- The framework supports internationalization (i18n)
- Translation files in `i18n/` directory
- Use framework's i18n utilities for localized applications

## AI Agent Workflow

### 1. Requirements Discovery
- **Primary Source**: `PRD.md` (Always prioritize this if present).
- **Secondary**: `requirements.txt`, `README.md`, or specific task files.
- **Goal**: Understand the full scope before writing code.

### 2. Implementation Protocol
- **Branching**: Work on a dedicated feature branch (e.g., `feat/implementation-details`).
- **Development**:
  - Analyze code structure.
  - Implement changes in `src/` or relevant directories.
  - Adhere to existing code style.
- **Verification**:
  - Run build commands (see above).
  - Run test suite (see above).
  - Ensure no regressions.

### 3. Delivery
- **Commit**: Use conventional commits (e.g., `feat: ...`, `fix: ...`).
- **PR Creation**:
  - Push branch: `git push -u origin <branch-name>`
  - Create a Pull Request against the main branch.
  - Summary: Link to `PRD.md` requirements solved.

## Task Implementation
1. **Analyze Requirements**: Refer to `README.md` for detailed feature specifications and system design.
2. **Implementation**: Modify source code in the respective directories (e.g., `src/`, `internal/`).
3. **Verification**: Run provided build and test commands (see above) to ensure correctness.
4. **Push Changes**:
   - Commit changes: `git commit -m "feat: implement <feature>"`
   - Push to remote: `git push origin <branch-name>`
