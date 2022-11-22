# TODO CLI

A command line TODO tool written in Go.

## Prerequisites

go >= 1.11

## Set up

Bulid executable binary to a folder that you register as $PATH.

```sh
go build -o ~/path/to/bin/todo
```

It uses sqlite3 for storage so you have to specify the path to storage like this.

```sh
export TODO_CLI_SQLITE_STORAGE_PATH="/path/to/sqlite/file"
```

Then you can use TODO CLI!

```sh
$ todo add -t "Fix a bug" -d "It's hard work."
# Todo was added successfully.
# ID        Title         Description
# 2caef4    Fix a bug     It's hard work.
$ todo list
# ID        Title         Description
# 2caef4    Fix a bug     It's hard work.
$ todo mod 2caef4 -d "It's easy!"
# Todo was updated successfully.
# ID        Title         Description
# 2caef4    Fix a bug     It's easy!
$ todo del 2caef4
# Todo was deleted successfully.
# ID        Title         Description
$ todo clear
# All todos were deleted successfully.
```
