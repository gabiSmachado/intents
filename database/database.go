package database

import (
	"database/sql"
	"fmt"

	"github.com/gabiSmachado/intents/datamodel"
)

func CurrentId(db *sql.DB) (int, error) {
	fmt.Println("GETTING IDs")
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM intents").Scan(&id)
	if err != nil {
		fmt.Printf("ERRO GETTING IDs",err)
		return 0, err
	}
	return id, nil
}


func Insert(db *sql.DB, intent datamodel.Intent) (int, error) {
	fmt.Println("INSERT DB")
	_, err := db.Exec(`INSERT INTO intents (name,description,ric_id,policy_id,service_id,policy_type_id) VALUES (?,?,?,?,?,?)`, 
			intent.Name, intent.Description, intent.RicID, intent.PolicyId,intent.ServiceID,intent.PolicyTypeId)
	fmt.Println("INSERT DB ending")
	if err != nil {
		fmt.Printf("Error %s when inserting in table", err)
		return 0, err
	}
	id, _ := CurrentId(db)
	return id, nil
}

func ListIntents(db *sql.DB) ([]datamodel.Intent, error) {
	rows, err := db.Query("SELECT id,name FROM intents")
	if err != nil {
		fmt.Printf("Error %s when listing intents", err)
		return nil, err
	}
	defer rows.Close()

	var intents []datamodel.Intent
	var name string
	var id int
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Printf("Error %s when listing intents", err)
		}
		intent := datamodel.Intent{Idx: id, Name: name}
		intents = append(intents, intent)
	}
	if err = rows.Err(); err != nil {
		fmt.Printf("Error %s when listing intents", err)
		return nil, err
	}

	return intents, nil

}

func DeleteIntent(db *sql.DB, id int) error {
	_, err := db.Query("DELETE FROM intents WHERE id = ?", id)
	return err
}

func DBconnect() (*sql.DB, error) {
	fmt.Println("/nDB/n")
	db, err := sql.Open("mysql", "root:mudemeja@tcp(mariadb-service.smo.svc.cluster.local)/intent")
	fmt.Println("Starting DB")
	if err != nil {
		fmt.Println("Connection error")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection to database error")
		return nil, err
	} else {
		fmt.Println("Connection to database successful!")
		return db, nil
	}

}

func IntentShow(db *sql.DB, id int) (datamodel.Intent, error) {
	var name, ricId, serviceID, description string
var policyId, policyTypeId int
	err := db.QueryRow("SELECT * FROM intents WHERE id = ?", id).Scan(&id, &name,
		&description,&ricId,&policyId,&serviceID,&policyTypeId)
	intent := datamodel.Intent{
		Name: name,
		Description: description,
		RicID: ricId,
		PolicyId: policyId,
		ServiceID: serviceID,
		PolicyTypeId: policyTypeId,
	} 
		
	if err != nil {
		fmt.Printf("Error %s when selecting intent", err)
		return intent, err
	}
	return intent, nil
}