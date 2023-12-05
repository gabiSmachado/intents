package producer

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
	"github.com/gabiSmachado/lbapp/datamodel"
	"encoding/json"
)

func WriteMsg(intent datamodel.Intent) error {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "test",
        Balancer: &kafka.LeastBytes{},
    })
	
	marshal, err := json.Marshal(intent)
	if err != nil {
		log.Fatalf("failed to marshal intent (%s)\n", err)
        return err
	}

	message := kafka.Message{
		Value:     marshal,}

    err = writer.WriteMessages(context.Background(), message)
    if err != nil {
        log.Fatalf("Error writing message: (%s)\n", err)
        return err
    }
    
    log.Printf("Message written successfully %s\n",intent.Name)
    writer.Close()
    return nil
}