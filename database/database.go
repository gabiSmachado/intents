package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gabiSmachado/intents/datamodel"
	_ "github.com/go-sql-driver/mysql"
)

func CurrentId(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM intents").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}


func Insert(db *sql.DB, intent datamodel.Intent) (int, error) {
	_, err := db.Exec(`INSERT INTO intents (name,description,ric_id,policy_id,service_id,policy_type_id) VALUES (?,?,?,?,?,?)`, 
			intent.Name, intent.Description, intent.RicID, intent.PolicyId,intent.ServiceID,intent.PolicyTypeId)
	if err != nil {
		log.Printf("Error %s when inserting in table", err)
		return 0, err
	}
	id, _ := CurrentId(db)
	return id, nil
}

func ListIntents(db *sql.DB) ([]datamodel.Intent, error) {
	rows, err := db.Query("SELECT id,name FROM intents")
	if err != nil {
		log.Printf("Error %s when listing intents", err)
		return nil, err
	}
	defer rows.Close()

	var intents []datamodel.Intent
	var name string
	var id int
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Printf("Error %s when listing intents", err)
		}
		intent := datamodel.Intent{Idx: id, Name: name}
		intents = append(intents, intent)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error %s when listing intents", err)
		return nil, err
	}

	return intents, nil

}

func DeleteIntent(db *sql.DB, id int) error {
	_, err := db.Query("DELETE FROM intents WHERE id = ?", id)
	return err
}

func DBconnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:teste@tcp(mariadb-service.smo.svc.cluster.local)/intent")

	if err != nil {
		fmt.Println("Connection error")
		//panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection db ping")
		log.Fatalln(err)
		return nil, err
	} else {
		fmt.Println("Connection successful!")
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
		log.Printf("Error %s when selecting intent", err)
		return intent, err
	}
	return intent, nil
}