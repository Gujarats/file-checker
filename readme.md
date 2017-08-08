# File checker 
This cli applicaiton will print all the files that has zero byte in specific directory

## how to install 
Install GO in your os. then
```shell 
$ go build
$ ./file-zero-checker

```
copy the  binary file your home PATH. like `/user/bin/local`

## how to use

```shell
$ ./file-zero-checker --help

NAME:
   MyCli - To check csv files

USAGE:
   file-zero-checker [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     find, f  Show file result
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --location value  specify files location (default: "files/")
   --help, -h        show help
   --version, -v     print the version
```
