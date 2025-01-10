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

If you use the `-i` flag it will show a cute little icon symbolizing the operating system such as:  or
. If you happen to be running Linux it will try to show your distribution's logo instead.

If you use the `-n` flag it will hide the name of the platform. You can use this flag in combination with `-i` to only show the icon.

Finally the `-d` flag will show the detailed distribution name and version number. For example: `Microsoft Windows 10 Pro 22H2` or `Ubuntu 20.04.3 LTS`.


    
