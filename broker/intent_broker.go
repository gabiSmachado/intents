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

const (
	DefaultHostPMS         = "http://nonrtricgateway.nonrtric.svc.cluster.local:9090"
	BasePathPMS            = "/a1-policy/v2"
	DefaultHostRAppCatalog = "http://rappcatalogueservice.nonrtric.svc.cluster.local:9085"
	BasePathRAppCatalog    = "/services"
)

var (
	baseURLRAppCatalogue  = DefaultHostRAppCatalog + BasePathRAppCatalog
	baseURLPMS            = DefaultHostPMS + BasePathPMS
)

func registerServiceRAppCatalogue(intent *datamodel.Intent){
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

	url :=	baseURLPMS + "/policies"

	marshal, err := json.Marshal(policy)
	if err != nil {
		fmt.Printf("failed to marshal policy (%s)", err)
	}

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))
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

	registerServiceRAppCatalogue(msg)
}