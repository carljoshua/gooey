// Gooey is a command-line tool which creates web-based GUI that can be used to manage different types of database.
package main

import (
    "flag"
    "os"
    "fmt"
    "database/sql"
    "github.com/carljoshua/gooey/core"

    _ "github.com/mattn/go-sqlite3"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    driver := flag.String("v", "", "driver to use.\n\nSupported drivers: \n\tmysql\n\tsqlite3")
    uname := flag.String("u", "", "username")
    passwd := flag.String("p", "", "password")
    dbname := flag.String("d", "", "database name/file path")
	flag.Parse()

    var data string
    switch *driver {
    case "mysql":
        data = *uname + ":" + *passwd + "@/" + *dbname
    case "sqlite3":
        data = *dbname
    default:
        fmt.Println("Missing or Invalid Driver")
        os.Exit(1)
    }

    db, err := sql.Open(*driver, data)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    test(db, *driver, *dbname)
    core.Start(db, *driver)
}

//Test if the the system has connected to the database  successfully
func test(db *sql.DB, driver string, dbname string) {
    var err error
    switch driver {
	case "mysql":
        _, err = db.Query("SELECT '1' FROM dual")
	case "sqlite3":
        _, err = os.Stat(dbname)
	}

	if err != nil {
		fmt.Println("Error connecting into the database")
		os.Exit(1)
	}
}
