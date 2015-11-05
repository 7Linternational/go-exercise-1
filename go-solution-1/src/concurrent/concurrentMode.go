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


func createGroupOne(db *sql.DB, done chan<- bool) {
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
	
	done <- true
}

func createGroupTwo(db *sql.DB, done chan<- bool) {
	var stmt *sql.Stmt
	var res sql.Result
	
	_ = res
	
	// User 3
	// Propably, this is not the proper syntax for oracle (it's actually for postgres)
	stmt, err := db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(3, "wiston", "whiskey", "The president")
	checkError(err)

	// User 4
	stmt, err = db.Prepare("INSERT INTO users(id, username, password, description, timestmp) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)")
	checkError(err)
	
	res, err = stmt.Exec(4, "jeff", "qw@##lmn45", "The code-monkey")
	checkError(err)
	
	done <- true
}

func updateUsers(db *sql.DB, id int, description string, checkSign chan<- bool) {
	var stmt *sql.Stmt
	var res sql.Result
	
	_ = res
	
	stmt, err := db.Prepare("UPDATE users SET description=$1 WHERE id=$2")
	checkError(err)
	
	res, err = stmt.Exec(description, id)
	checkError(err)
	
	checkSign <- true
}

func deleteUsers(db *sql.DB, checkSign chan<- bool) {
	// Delete user with id=3
	_, err := db.Exec("DELETE FROM users WHERE id=3")
	checkError(err)
	
	// Delete user with id=4
	_, err = db.Exec("DELETE FROM users WHERE id=4")
	checkError(err)
	
	checkSign <- true
}

func displayUsers(db *sql.DB, doneOne <-chan bool, doneTwo <-chan bool) {
	var (
		id int
		username string
		password string
		description string
		timestmp string
	)
	
	for {
		if <-doneOne && <-doneTwo {
			break
		}
	}
	
	rows, err := db.Query("SELECT id, username, password, description, timestmp FROM users")
	checkError(err)
	
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &description, &timestmp)
		checkError(err)
		
		fmt.Println(id, username, password, description, timestmp)
	}
}

func displayAllUsers(db *sql.DB, checkOne <-chan bool, checkTwo <-chan bool, checkThree <-chan bool) {
	var (
		id int
		username string
		password string
		description string
		timestmp string
	)
	
	for {
		if <-checkOne && <-checkTwo  && <-checkThree {
			break
		}
	}
	
	rows, err := db.Query("SELECT id, username, password, description, timestmp FROM users")
	checkError(err)
	
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &description, &timestmp)
		checkError(err)
		
		fmt.Println(id, username, password, description, timestmp)
	}
}

func main() {
	db, err := sql.Open("ora.v3", "lalakis/v1de0@localhot:1521/xe")
	
	checkErrCallback(err, func() {
		fmt.Println(err)
		log.Fatal("here")
	})
	
	err = db.Ping()
	
	// a. and b.
	doneOne := make(chan bool, 1)
	doneTwo := make(chan bool, 1)
	
	doneOne <- false
	doneTwo <- false
	
	go createGroupOne(db, doneOne)
	go createGroupTwo(db, doneTwo)
	
	displayUsers(db, doneOne, doneTwo)
	
	// c.
	checkOne := make(chan bool, 1)
	checkTwo := make(chan bool, 1)
	checkThree := make(chan bool, 1)
	
	checkOne <- false
	checkTwo <- false
	checkThree <- false
	
	go updateUsers(db, 1, "Master of puppets", checkOne)
	go updateUsers(db, 2, "Lord commander", checkTwo)
	go deleteUsers(db, checkThree)
	
	displayAllUsers(db,  checkOne, checkTwo, checkThree)
}