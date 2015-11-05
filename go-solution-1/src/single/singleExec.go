/**
 * Database schema
 * users(id number(10), username varchar(50), password varchar(50), description varchar(50), timestmp timestamp)
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "gopkg.in/rana/ora.v3"
)



// Stolen function from havoc
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Same as the above function
func checkErrCallback(err error, fn func()) {
	if err != nil {
		if fn != nil {
			fn()
		}
		log.Fatal(err)
	}
}

func createUsers(db *sql.DB) {
	var stmt *sql.Stmt
	var res sql.Result
	
	_ = res
	
	// User 1
	// Propably, this is not the proper syntax for oracle (it's actually for postgres)
	stmt, err := db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(1, "john", "1234", "The janitor")
	checkError(err)

	// User 2
	stmt, err = db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(2, "maria", "4321", "The secretary")
	checkError(err)
	
	// User 3
	stmt, err = db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(3, "wiston", "whiskey", "The president")
	checkError(err)
	
	// User 4
	stmt, err = db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(4, "jeff", "qw@##lmn45", "The code-monkey")
	checkError(err)
}

func displayUsers(db *sql.DB) {
	var (
		id int
		username string
		password string
		description string
		timestmp string
	)
	
	rows, err := db.Query("SELECT id, username, password, description, timestmp FROM users")
	checkError(err)
	
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &description, &timestmp)
		checkError(err)
		
		fmt.Println(id, username, password, description, timestmp)
	}
}

func updateUsers(db *sql.DB) {
	var stmt *sql.Stmt
	var res sql.Result
	
	_ = res
	
	// User update with id=1
	stmt, err := db.Prepare("UPDATE users SET description=$1 WHERE id=$2")
	checkError(err)
	
	res, err = stmt.Exec("The millionaire", 1)
	checkError(err)
	
	// User update with id=2
	stmt, err = db.Prepare("UPDATE users SET description=$1 WHERE id=$2")
	checkError(err)
	
	res, err = stmt.Exec("The queen of Norway", 2)
	checkError(err)
}

func deleteUsers(db *sql.DB) {
	// Delete user with id=3
	_, err := db.Exec("DELETE FROM users WHERE id=3")
	checkError(err)
	
	// Delete user with id=4
	_, err = db.Exec("DELETE FROM users WHERE id=4")
	checkError(err)
}

// It doesn't run because of bad configuration. It throws an exit status 3221225785.
func main() {
	db, err := sql.Open("ora.v3", "lalakis/v1de0@localhot:1521/xe")
	
	checkErrCallback(err, func() {
		fmt.Println(err)
		log.Fatal("here")
	})
	
	err = db.Ping()
	
	createUsers(db)
	displayUsers(db)
	updateUsers(db)
	deleteUsers(db)
	displayUsers(db)
}
