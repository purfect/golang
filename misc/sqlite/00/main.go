package main

import (
    "database/sql"
    "fmt"
    "strconv"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    var id int
    var username string
    var alias string

    database, _ := sql.Open("sqlite3", "./foo.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, username TEXT, alias TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO people (username, alias) VALUES (?, ?)")
    statement.Exec("rasputin", "=^_^=")
    rows, _ := database.Query("SELECT id, username, alias FROM people")
   for rows.Next() {
        rows.Scan(&id, &username, &alias)
        fmt.Println(strconv.Itoa(id) + ": " + username + " " + alias)
    }
}
