# Tokenized Protocol Specification

Reference for the Tokenized Protocol structure. Definitions are stored in YAML to be parsed and compiled to other languages. Definitions are grouped by their purpose and are versioned within.

## Supported Languages

- Go
- Python
- Markdown
- _your language here_ (contributions welcome)

## Getting Started

To build the project from source, clone the GitHub repo and in a command line terminal.

    mkdir -p $GOPATH/src/github.com/tokenized
    cd $GOPATH/src/github.com/tokenized
    git clone git@github.com:tokenized/specification
    cd specification

Navigate to the root directory and run:

    make

## Project Structure

- `cmd/tokenized` - Command line interface
- `src` - Source files for protocol messages (see Source Definitions)
- `dist` - Latest library for each language
- `internal` - Compiler logic and templates

## Source Definitions

The Tokenized Protocol must retain the entire definition history in order to parse legacy messages from prior versions. The working version for each definition is `develop`. Previous versions are stored in a directory that specifies that release version, for example, `v0001`.

### Protocol

All actions related to the Tokenized Protocol.

### Assets

All asset payloads for the Asset related actions.

### Resources

All resources and message payloads for message related actions.

# License

Copyright 2018-2019 nChain Holdings Limited.
