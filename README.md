## Hexlet path size

[![Actions Status](https://github.com/Lirikman/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Lirikman/go-project-242/actions)

### Description

Print size of a file or directory

USAGE: hexlet-path-size [global options]

GLOBAL OPTIONS:
   
   --recursive, -r  recursive size of directories (default: false)
   
   --human, -H      human-readable sizes (auto-select unit) (default: false)
   
   --all, -a        include hidden files and directories (default: false)
   
   --help, -h       show help

### Requirements

* Go 1.25
* Make
* urfave/cli v3

### Setup

```bash
git clone git@github.com:Lirikman/go-project-242.git
cd go-project-242
```

### Run build

```bash
make build
```

### Run golangci-lint 

```bash
make lint
```

### Run golangci-lint and automatic correction

```bash
make lint-fix
```

### Run tests

```bash
make test
```

### Example

```bash
./bin/hexlet-path-size project/ -H -a
27.0MB  project/

bin/hexlet-path-size project/ -H -a -r
31.0MB  project/
```

### Asciinema

https://asciinema.org/connect/f5b636e0-6e25-4aae-acce-8975c6ce0b50
