# dodo

Like a list of tasks of a rare extinct species...

__A CLI app__ in Go with SQLite.

## Overview

dodo allows the user the following functions to manage their tasks.

* Tasks list
* Addition
* Check as done
* Deleting task
* ...

just do it ;)

## Install

Check dependencies

* cobra : go get -u github.com/spf13/cobra/cobra
* viper : go get -u github.com/spf13/viper
* go-sqlite: go get github.com/mattn/go-sqlite3

You need a $GOBIN at go install !!! otherwise you have problems with database connection.

see:

```bash
export GOBIN=$HOME/go/bin
source .bash_profile
```

## Licenses

* [Go](https://golang.org/LICENSE)
* [SQLite](https://www.sqlite.org) : [Public Domain](https://www.sqlite.org/copyright.html)
* [Cobra](https://github.com/spf13/cobra) : Cobra is released under the Apache 2.0 license. See LICENSE
* __dodo__ is released under the GNU General Public License. See [gpl-3.0.txt](./gpl-3.0.txt)

<img src="gplv3-or-later.svg" />

