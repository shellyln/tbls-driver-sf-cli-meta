# tbls driver for Salesforce CLI's metadata

[tbls](https://github.com/k1LoW/tbls?tab=readme-ov-file#external-database-driver) external database driver for Salesforce CLI's metadata.  
Driver for reading "Source Format" metadata of Salesforce CLI [sf command](https://developer.salesforce.com/docs/atlas.en-us.sfdx_setup.meta/sfdx_setup/sfdx_setup_install_cli.htm) on local files.

[![Test](https://github.com/shellyln/tbls-driver-sf-cli-meta/actions/workflows/test.yml/badge.svg)](https://github.com/shellyln/tbls-driver-sf-cli-meta/actions/workflows/test.yml)
[![release](https://img.shields.io/github/v/release/shellyln/tbls-driver-sf-cli-meta)](https://github.com/shellyln/tbls-driver-sf-cli-meta/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/shellyln/tbls-driver-sf-cli-meta)](https://github.com/shellyln/tbls-driver-sf-cli-meta)

## 🪄 Install

### Go install (recommended):
#### 🪟 Windows prerequirements:
```bash
choco install golang
# or
# scoop bucket add main
# scoop install main/go

go install github.com/k1LoW/tbls@latest
```
* https://community.chocolatey.org/
* https://scoop.sh/

#### 🍎 Mac prerequirements:
```bash
brew install go
go install github.com/k1LoW/tbls@latest
```
* https://brew.sh/

#### ⚡️ Install:
```bash
go install github.com/shellyln/tbls-driver-sf-cli-meta@latest
```

### Manually:
#### 🪟 Windows prerequirements:
```bash
choco install git

choco install golang
choco install make
# or
# scoop bucket add main
# scoop install main/go
# scoop install main/make

go install github.com/k1LoW/tbls@latest
```
* https://community.chocolatey.org/
* https://scoop.sh/

#### 🍎 Mac prerequirements:
```bash
brew install git
brew install go
brew install make
go install github.com/k1LoW/tbls@latest
```
* https://brew.sh/

#### ⚡️ Install:
```bash
git clone https://github.com/shellyln/tbls-driver-cli-meta.git
cd tbls-driver-cli-meta

# Build it and copy the binary to the $GOPATH/bin, which is probably in the PATH.
# tbls external driver executable SHOULD be in the PATH.
make && make install
```


## 🚀 Getting Started
Add `.tbls.yml` file to your repository.  
See [tbls documentation](https://github.com/k1LoW/tbls?tab=readme-ov-file#document-format) and [sample](./.tbls.yml).
```yaml
# .tbls.yml

# Relative path to the metadata.
dsn: sf-cli-meta:.
# Absolute path to the metadata.
# dsn: sf-cli-meta:///path/to/repository

docPath: doc/schema

format:
  # Adjust the column width of Markdown format table
  # Default is false
  adjust: false
  # Sort the order of table list and columns
  # Default is false
  sort: true
```

Reading `flows`, `globalValueSets`, `restrictionRules`,`triggers` and `objects` metadata.  
Metadata should be located in the repository as follows:
```
.
└── force-app/
    └── main/
        └── default/
```

Run tbls to generate document.
```bash
tbls doc
```
[Sample document](sample/README.md)

## ⚖️ License

MIT  
Copyright (c) 2025 Shellyl_N and Authors.
