package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gabiSmachado/intents/datamodel"
	"github.com/segmentio/kafka-go"
)

const (
	DefaultHostPMS         = "nonrtricgateway.nonrtric.svc.cluster.local"
	BasePathPMS            = "/a1-policy/v2"
	DefaultHostRAppCatalog = "rappcatalogueservice.nonrtric.svc.cluster.local"
	BasePathRAppCatalog    = "/services"
)

var (
	baseURLRAppCatalogue  = DefaultHostRAppCatalog + BasePathRAppCatalog
	baseURLPMS            = DefaultHostPMS + BasePathPMS
)

func registerServiceRAppCatalogue(intent datamodel.Intent){
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
		fmt.Errorf("failed to marshal policy (%s)", err)
	}

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))
    if err != nil {
        fmt.Errorf("failed to create request (%s)", err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Errorf("failed to send policy (%s)", err)
    }
    defer resp.Body.Close()
}



func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka-service.smo.svc.cluster.local"},
		Topic:    "intent",
		GroupID:  "mygroup",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	var msg datamodel.Intent
	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Fatalf("Error reading message: %v", err)
	}

	err = json.Unmarshal(m.Value, &msg)
	if err != nil {
		log.Fatalf("Error parsing message: %v", err)
	}

	log.Printf("Processed message: %v\n",msg) 

}