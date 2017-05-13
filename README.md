# gitbook-gen

[![GoDoc][I1]][L1] [![License][I2]][L2] [![Build Status][I3]][L3] [![Coverage Status][I4]][L4]

[I1]: http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
[L1]: http://godoc.org/github.com/yulefox/gitbook-gen
[I2]: http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[L2]: LICENSE
[I3]: https://img.shields.io/travis/yulefox/gitbook-gen.svg?style=flat-square
[L3]: https://travis-ci.org/yulefox/gitbook-gen
[I4]: https://img.shields.io/codecov/c/github/yulefox/gitbook-gen.svg?style=flat-square
[L4]: https://codecov.io/gh/yulefox/gitbook-gen

Generate `SUMMARY.md` for [Gitbook](https://github.com/GitbookIO/gitbook).

## Features

- Recursively scan the specified directory(current for default) with depth(2 for default) to generate `SUMMARY.md`
- The posts with the prefix `_` are private, can be skipped when generates `SUMMARY.md`
- Exclude specified directories

## Install

```sh
glide update
go install
```

## Test

```sh
./test.sh
```

## Usage

```sh
gitbook-gen help

USAGE:
   gitbook-gen [global options] command [command options] [arguments...]

VERSION:
   1.0.0

DESCRIPTION:
   gitbook directory, current directory for default

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --depth DEPTH, -d DEPTH                 DEPTH of TOC (default: 2)
   --extensions EXTENSIONS, -e EXTENSIONS  post EXTENSIONS (separated by commas, NO spaces) (default: ".md,.markdown")
   --excludes DIRECTORIES                  exclude DIRECTORIES (separated by commas, NO spaces) (default: "_book")
   --show-all                              show all posts(include private, posts )
   --help, -h                              show help
   --version, -v                           print the version
```

## TODO

- Embbedded in [Gitbook](https://github.com/GitbookIO/gitbook)

## License

[MIT](LICENSE)
