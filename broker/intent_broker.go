package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gabiSmachado/intents/datamodel"
	"github.com/segmentio/kafka-go"
)
type PutBody struct {
	RicID         string `json:"ric_id"`
	PolicyId  	  int    `json:"policy_id"`
	ServiceID     string `json:"service_id"`
	PolicyData	  PolicyData`json:"policy_data"`
	PolicyTypeId  string    `json:"policytype_id"`
}

type PolicyData struct {
	Threshold int   `json:"threshold"`
}

type rApp struct {
	Version       string `json:"version"`
	DisplayName         string `json:"display_name"`
	Description 	string    `json:"description"`
}


func registerServiceRAppCatalogue(){
	body := rApp{
		Version:  "0.0.1",
		DisplayName: "Hello Word rApp",
		Description: "Hello Word rApp for testing Non-RT RIC guide development of future rApps and demo purposes",
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("failed to marshal rapp body (%s)", err)
	}

	req, err := http.NewRequest("PUT",
	"http://rappcatalogueservice.nonrtric.svc.cluster.local:9085/services/IntentBrokerApp",
		 bytes.NewBuffer(marshal))

    if err != nil {
        fmt.Printf("failed to create rApp request (%s)", err)
    }
    req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Printf("failed to register rApp(%s)", err)
    }else{
		fmt.Print("rApp registered" )
	}
    defer resp.Body.Close()
}


func PutPolicy(intent *datamodel.Intent){
	threshold := rand.Intn(10^15 - 1)
	typeid := strconv.Itoa(intent.PolicyTypeId)

	policy:= PutBody{
		intent.RicID,
		intent.PolicyId,
		intent.ServiceID,
		PolicyData{
			threshold,
		},
		typeid,
	}

	marshal, err := json.Marshal(policy)
	if err != nil {
		fmt.Printf("failed to marshal policy (%s)", err)
	}

    req, err := http.NewRequest("PUT",
		"http://nonrtricgateway.nonrtric.svc.cluster.local:9090/a1-policy/v2/policies",
		 bytes.NewBuffer(marshal))
    if err != nil {
        fmt.Printf("failed to create request (%s)", err)
    }
    req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Printf("failed to send policy (%s)", err)
    }else{
		fmt.Print("policy registered" )
	}
    defer resp.Body.Close()
}



func main() {
	registerServiceRAppCatalogue()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka-teste.smo.svc.cluster.local"},
		Topic:    "intent",
		GroupID:  "mygroup",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	var msg *datamodel.Intent
	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		fmt.Printf("Error reading message: %v", err)
	}

	err = json.Unmarshal(m.Value, &msg)
	if err != nil {
		fmt.Printf("Error parsing message: %v", err)
	}

	fmt.Printf("Processed message: %v\n",msg) 

	PutPolicy(msg)
}