package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func main() {
	sqlQuery(30)
	fmt.Println("==================")
	sqlQueyRow("B001")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery(age int) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}

}

func sqlQueyRow(id string) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = student{}
	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name : %s \n grade: %d \n", result.name, result.grade)
}
