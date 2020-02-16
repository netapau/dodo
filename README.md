# dodo

Like a list of tasks of a rare extinct species... [dodo](<https://fr.wikipedia.org/wiki/Dodo_(oiseau)>)

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

* [cobra](https://github.com/spf13/cobra) : go get -u github.com/spf13/cobra/cobra
* [viper](https://github.com/spf13/viper) : go get -u github.com/spf13/viper
* [go-sqlite](github.com/mattn/go-sqlite3): go get github.com/mattn/go-sqlite3

__go-sqlite3 and go modules__
See for environment var GOPRIVATE here : [go-sqlite & GOPRIVATE](https://github.com/mattn/go-sqlite3/issues/755#issuecomment-555419067)

[Go Documentation](https://golang.org/doc/go1.13)

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

![Alt gplv3-or-later](./gplv3-or-later.png?raw=true "gplv3")
