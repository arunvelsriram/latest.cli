# latest.cli

[![Build Status](https://travis-ci.org/arunvelsriram/latest.cli.svg?branch=master)](https://travis-ci.org/arunvelsriram/latest.cli)

A CLI to find the latest version of a Ruby Gem, Node module, Java JAR etc.

## Usage

### Prerequisites

On *Linux* `curl` dependencies are required.

Debian based Linux:

`apt-get install curl`

RPM based Linux:

`yum install curl`

### Install

#### MacOS and Linux

1. Download appropriate tar file from [releases](https://github.com/arunvelsriram/latest.cli/releases) page
2. Extract the tar file
3. Move the binary to `/usr/local/bin`
5. Run `latest`

#### Windows

Not tested yet! :P

## TODO
- [ ] Make latest installable through brew
- [ ] Findest latest version of a JAR  
- [ ] Pass multiple names (eg. `latest rails puma rspec`)
