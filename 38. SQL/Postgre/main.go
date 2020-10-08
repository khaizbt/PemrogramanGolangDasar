package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sayangkamu"
	dbname   = "belajar"
)

func main() {
	sqlQuery(28)
	fmt.Println("==================")
	sqlQueyRow("B002")
	fmt.Println("==================")
	sqlPrepare()
	fmt.Println("==================")
	sqlExec()

}

func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
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

	rows, err := db.Query("select id, name, grade from tb_students where age = $1", age)

	if err != nil {
		fmt.Println("database", err.Error())
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
		fmt.Println(each.name, string('|'), each.grade)
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
		QueryRow("select name, grade from tb_students where id = $1", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name : %s \n grade: %d \n", result.name, result.grade)
}

func sqlPrepare() {
	//Teknik Penulisan Query bisa reuse/eksekusi berbeda2
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_students where id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\n grade : %d \n", result1.name, result1.grade)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\n grade : %d \n", result2.name, result2.grade)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\n grade : %d \n", result3.name, result3.grade)
}

//CRUD menggunakan Exec
func sqlExec() {
	db, err := connect()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer db.Close()

	//Create
	_, err = db.Exec("insert into tb_students values ($1, $2, $3, $4)",
		"K001", "Khaiz Badaru Tammam", 23, 2)

	if err != nil {
		fmt.Println("Insert Data Failed", err.Error())
		return
	}
	fmt.Println("Insert Success")

	//Update
	_, err = db.Exec("update tb_students set age = $1 where id = $2",
		22, "K002")
	if err != nil {
		fmt.Println("Update Data Failed", err.Error())
		return
	}
	fmt.Println("Update Success!")

	//Delete
	_, err = db.Exec("delete from tb_students where id = $1",
		"B001")
	if err != nil {
		fmt.Println("Delete Data Failed", err.Error())
		return
	}
	fmt.Println("Delete Data Success")
}
