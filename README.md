# Bitwarden CLI

Official Bitwarden CLI tool built with Go and distributed via Homebrew. This is a standalone implementation compatible with Bitwarden's official CLI interface.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap yourusername/homebrew-tap
brew install bitwarden-cli
```

### Manual Installation

Download the appropriate binary for your platform from the [releases page](https://github.com/yourusername/bitwarden-cli/releases) and add it to your PATH.

## Usage

```bash
bw --help
```

## Building from Source

```bash
git clone https://github.com/yourusername/bitwarden-cli.git
cd bitwarden-cli
go build -o bw main.go
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