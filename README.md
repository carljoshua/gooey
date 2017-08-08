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

	* Pure Go code
	* Supports many database drivers
		* MySQL
		* Sqlite


## Requirements

	* Go 1.8
	* Third-party driver for Go's [database/sql](https://golang.org/pkg/database/sql/) package

---------------------------------------

## Warning
This package is meant for development purposes only. Do not release your project with this in your code because it doesn't have authentication procedures. Another reason is this package is not yet fully developed.

## Installation
Make sure you have you set your GOPATH (https://github.com/golang/go/wiki/GOPATH).
Type the command below to install the package.

```bash
$ go get github.com/carljoshua/golangmyadmin
```

## Usage
Import the driver you like (I used Sqlite driver (https://github.com/mattn/go-sqlite3) in this example) and open a connection into the database. Pass the connection, the name of the driver and the port number where you want the GUI to be served into Run().

```go
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "github.com/carljoshua/golangmyadmin"

db, _ := sql.Open("sqlite3", "resources.db")
golangmyadmin.Run(db, "sqlite3", ":8002")
```

After running the code above, a folder named "UI" will be created in your project directory. This where the files of the website is stored. If you are planning to change the look of the user interface, you can edit or change the files in this folder.

All you have to do now is to go to your browser of your choice and go to "localhost:(port number)/golangmyadmin" (e.g localhost:8002/golangmyadmin).
