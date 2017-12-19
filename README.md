# Gooey

Gooey is a command-line tool which creates web-based GUI that can be used to manage different types of database.

This package also have a API for managing the database. The API can be found at http://localhost:8000/db if you use the default port number.

Note: If you don't like the GUI, you can edit it in `$GOPATH/src/github.com/carljoshua/gooey/core/UI`. I am not the best designer so feel free to change it.

---------------------------------------
  * [Features](#features)
  * [Installation](#installation)
  * [Usage](#usage)
  * [Contributing](#contibuting)
---------------------------------------

## Features
    * Pure Go code (Except for the GUI)
    * Light-weight
    * GUI is editable
    * Supports different kinds of database
    	* MySQL
        * Sqlite3

---------------------------------------

## Installation
Make sure you have you set your GOPATH (https://github.com/golang/go/wiki/GOPATH). Type the command below to install the package.

```bash
$ go get github.com/carljoshua/gooey
```

## Usage
In mysql:

```bash
$ gooey -v mysql -d mydb -u root -p mypasswd
```

In sqlite3:

```bash
$ gooey -v sqlite3 -d ./mydatabase.db
```

The GUI can now be found at http://localhost:8000/gooey.

## Contributing
Fork it and create a new branch. If you have a working feature, create a pull request.

I'll appreciate any contribution, suggestion or criticism.
