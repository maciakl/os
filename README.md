# os

Cross platform cli os detection tool.

## Usage

    Usage: os [options]
    Options:
      -v, -version    Print version information and exit
      -h, -help       Print this message and exit
      -i, -icon       Show the icon of the operating system
      -n, -name       Hide the name of the platform
      -d, -detail     Show the distribution name and version

## What is this tool?

This tool aims to answer one simple question: "what operating system am I running?"

It seems like an easy question to answer, but it is not.

If you're on Windows you would probably run one or two `wmic` queries to get the name and version number, but who can remember all of those. Or maybe you run `systeminfo` and parse the output. On Linux do you run `uname`, `lsb_release` or `cat /etc/os-release`. On MacOS you typically run `sw_vers` a non-POSIX command that no one ever remembers.

There isn't a single cross platform command to get a simple platform, OS name and version number regardless on what computer you are running on.

Until now.

## What does it do?

By default `os` prints out the name of the platform you are running on such as `linux`, `windows`, `darwin`, `freebsd`, etc..

If you use the `-i` flag it will show a cute little icon symbolizing the operating system (such as the apple logo for macOS). If you happen to be running Linux it will try to show your distribution's logo instead.

If you use the `-n` flag it will hide the name of the platform. You can use this flag in combination with `-i` to only show the icon.

Finally the `-d` flag will show the detailed distribution name and version number. For example: `Microsoft Windows 10 Pro 22H2` or `Ubuntu 20.04.3 LTS`.

## Installing

There are few different ways:

### Platform Independent

 Install via `go`:
 
    go install github.com/maciakl/os@latest

### Linux

On Linux (requires `wget` & `unzip`, installs to `/usr/local/bin`):

    p="os" && wget -qN "https://github.com/maciakl/${p}/releases/latest/download/${p}_lin.zip" && unzip -oq ${p}_lin.zip && rm -f ${p}_lin.zip && chmod +x ${p} && sudo mv ${p} /usr/local/bin

To uninstall, simply delete it:

    rm -f /usr/local/bin/jjmp

### Windows

On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

 First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install os

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
