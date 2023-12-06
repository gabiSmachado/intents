package main
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBconnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/intents")
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		fmt.Println("Connection successful!")
		return db, nil
	}

}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS intents (
		id int primary key AUTO_INCREMENT,
		name varchar(255),
		label varchar(255),
		day_of_the_week varchar(255),
		start_tiime varchar(255),
		end_time varchar(255),
		minimum_cell_offset int,
		maximum_cell_offset int)`)

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}

	log.Printf("Creating table successfully!\n")
	return nil
}

func main(){
	db, err := DBconnect()
	if err != nil {
		println(err)
	}
	defer db.Close()

	createTable(db)
}