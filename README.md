# Bitwarden CLI

Official Bitwarden CLI tool built with Go and distributed via Homebrew. This is a standalone implementation compatible with Bitwarden's official CLI interface.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap codefuturist/homebrew-tap
brew install bitwarden-cli
```

### Manual Installation

Download the appropriate binary for your platform from the [releases page](https://github.com/codefuturist/bitwarden-cli/releases) and add it to your PATH.

### Version

**Current Version**: 2.29.0 (matches official Bitwarden CLI)

## Usage

```bash
bw --help
```

## Building from Source

```bash
git clone https://github.com/codefuturist/bitwarden-cli.git
cd bitwarden-cli
go build -o bw main.go
```

## GitHub Actions

This repository includes automated workflows:

- **Test**: Runs on every push/PR to verify builds and configuration
- **Release**: Automatically builds and publishes when tags are pushed
- **Homebrew Update**: Weekly formula updates

### Creating a Release

```bash
# Tag a new version
git tag v2.29.0
git push origin v2.29.0

# GitHub Actions will automatically:
# 1. Build binaries for all platforms
# 2. Create GitHub release with assets
# 3. Update Homebrew tap
# 4. Publish to Homebrew
```

## Features

- **Standalone binary** - Minimal external dependencies
- Cross-platform support (Linux, Windows, macOS)
- Multiple architecture support (amd64, arm64, arm, i386)
- Statically compiled with `CGO_ENABLED=0`
- Homebrew tap for easy installation
- Single binary deployment
- No complex runtime requirements

## License

MIT