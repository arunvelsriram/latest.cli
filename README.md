# latest.cli

[![Build Status](https://travis-ci.org/arunvelsriram/latest.cli.svg?branch=master)](https://travis-ci.org/arunvelsriram/latest.cli)

A CLI to find the latest version of a Ruby Gem, Node module, Java JAR etc.

## Usage

### Prerequisites

On **Linux** `curl` dependencies are required.

##### Debian based Linux:

`apt-get install curl`

##### RPM based Linux:

`yum install curl`

### Install

#### Linux

Based on your Linux variant download `deb` or `rpm` file from the [latest release](https://github.com/arunvelsriram/latest.cli/releases/latest).

##### Debian based Linux:

`dpkg -i latest.cli.deb`

##### RPM based Linux:

`rpm -i latest.cli.rpm`

#### MacOS and Linux (using binary)

1. Download tar file from the [latest release](https://github.com/arunvelsriram/latest.cli/releases/latest)
2. Extract the tar file
3. Move the binary to `/usr/local/bin`
4. Run `latest`

#### Windows

Not tested yet! :P

## TODO
- [ ] Make latest installable through brew
- [ ] Support Python packages
- [ ] Findest latest version of a JAR  
- [ ] Pass multiple names (eg. `latest rails puma rspec`)
