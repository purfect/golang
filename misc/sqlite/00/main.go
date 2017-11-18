package main

import (
    "database/sql"
    "fmt"
    "strconv"
    "os"
    _ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

func main() {
    var id int
    var username string
    var alias string
	var randomint int

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	database, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("DB konnte nicht erstellt oder geöffnet werden, Fehler:", err)
		os.Exit(1)
	}
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, username TEXT, alias TEXT, randomint INT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (username, alias, randomint) VALUES (?, ?, ?)")
	statement.Exec("User"+strconv.Itoa(r1.Intn(100)), "=^_^=", r1.Intn(100) )
	rows, err := database.Query("SELECT id, username, alias, randomint FROM people")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
   	for rows.Next() {
        rows.Scan(&id, &username, &alias, &randomint)
        fmt.Println(strconv.Itoa(id) + ": " + username + " " + alias + " Alter:" + strconv.Itoa(randomint))
    }
}
