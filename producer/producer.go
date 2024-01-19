package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gabiSmachado/intents/datamodel"
	"github.com/segmentio/kafka-go"
)

func WriteMsg(intent datamodel.Intent) error {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"kafka-teste.smo.svc.cluster.local"},
        Topic:    "intent",
        Balancer: &kafka.LeastBytes{},
    })
	
	marshal, err := json.Marshal(intent)
	if err != nil {
		fmt.Printf("failed to marshal intent (%s)\n", err)
        return err
	}
	message := kafka.Message{
		Value:     marshal,}

    err = writer.WriteMessages(context.Background(), message)
    
    if err != nil {
        fmt.Printf("Error writing message: (%s)\n", err)
        return err
    }
    
    fmt.Printf("Message written successfully %s\n",intent.Name)
    writer.Close()
    return nil
}