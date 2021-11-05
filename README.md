# Paraguas Bot Reloaded

## Usage
You need an env file ("env.env") on the root directory which contains a "TOKEN" key/value pair.

Example:
```shell
TOKEN=YOURSECRETKEYGOESHERE
```

## Supported Platforms
Alpha version tested under these platforms:
* MacOS Intel
* Linux 64 bits

## Compilation instructions:
```shell
git clone $repo
cd $repo/app
go mod tidy -v
go build
```

## License
[GNU GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)