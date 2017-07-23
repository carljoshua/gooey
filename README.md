# GolangMyAdmin

Web-based GUI for managing databases (similar to PHPMyAdmin)

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  *	[Warning](#warning)	
  * [Installation](#installation)
  * [Usage](#usage)
---------------------------------------

## Features

	* Pure Go (I don't know C bindings YET. Don't judge me)
	* Supports many database drivers
		* MySQL
		* Sqlite


## Requirements

	* Go 1.8
	* Third-party driver for Go's [database/sql](https://golang.org/pkg/database/sql/) package

---------------------------------------

## Warning
This package is meant for development purposes only. Do not release your project with this in your code because it doesn't have authentication procedures.

## Installation
Make sure you have you set your GOPATH (https://github.com/golang/go/wiki/GOPATH "GOPATH").
Type the command below to install the package.

```bash
$ go get [nothing yet]
```

## Usage
Import the driver you like (I used Sqlite driver (https://github.com/mattn/go-sqlite3) in this example) and open a connection into the database. Pass the connection and the port where you want the GUI to be served into Sync().

```go
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import [nothing yet]

db, err := sql.Open("sqlite3", "resources.db")
[nothing yet].Sync(db, ":8002")
```
