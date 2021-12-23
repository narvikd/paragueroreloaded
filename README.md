[![Go Report Card](https://goreportcard.com/badge/github.com/narvikd/paragueroreloaded)](https://goreportcard.com/report/github.com/narvikd/paragueroreloaded)  [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

<h1 align="center">Deprecated</h1>
<p align="center">
Paraguas Bot Reloaded has been deprecated. Important bugs may not be fixed, but PRs are still welcomed.
</p>


# Paraguas Bot Reloaded

## Usage
You need an env file ("env.env") on the root directory which contains a "TOKEN" key/value pair.
Example:
```shell
TOKEN=YOURSECRETKEYGOESHERE
```
If you don't have one, the bot will create it for you.

Use "/help" after you start the bot to see the list of commands available

## Supported Platforms
Alpha version tested under these platforms:
* MacOS Intel
* Linux 64 bits
* Windows 64 bits LTSC 2019

## Docker instructions:
```shell
git clone $repo
cd $repo
docker-compose up --build --detach
```

## Compilation instructions:
```shell
git clone $repo
cd $repo
go mod tidy -v
go build
```

## License
[GNU GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)