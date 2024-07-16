# RandGen
[![Go Reference](https://pkg.go.dev/badge/dev.shib.me/randgen.svg)](https://pkg.go.dev/dev.shib.me/randgen)
[![Go Report Card](https://goreportcard.com/badge/dev.shib.me/randgen)](https://goreportcard.com/report/dev.shib.me/randgen)
[![Release Status](https://github.com/shibme/randgen/actions/workflows/release.yaml/badge.svg)](https://github.com/shibme/randgen/actions/workflows/release.yaml)
[![License](https://img.shields.io/github/license/shibme/randgen)](https://github.com/shibme/randgen/blob/main/LICENSE)

RandGen is a simple CLI tool to generate random files of a given size.

## CLI Usage
Download the latest binary from the [releases](https://github.com/shibme/randgen/releases/latest) page and add it to your path.

### Homebrew
RandGen can be installed with brew using the following command on macOS
```zsh
brew install shibme/lab/randgen
```

### Install Script

#### Install Latest Version
**With Shell (MacOs/Linux):**
```sh
curl -fsSL https://dev.shib.me/randgen/install.sh | sh
```
**With PowerShell (Windows):**
```powershell
irm https://dev.shib.me/randgen/install.ps1 | iex
```

#### Install Specific Version
**With Shell (MacOs/Linux):**
```sh
curl -fsSL https://dev.shib.me/randgen/install.sh | sh -s v1.1.0
```
**With PowerShell (Windows):**
```powershell
$v="1.1.0"; irm https://dev.shib.me/randgen/install.ps1 | iex
```

### Docker
You can also run RandGen without installing, using Docker:
```zsh
docker run --rm -v $PWD:/workspace -it ghcr.io/shibme/randgen help
```
