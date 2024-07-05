package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	//kafka configurationn
	brocker := []string{"localhost:29092"}
	topic := "TEST_TOPIC"
	goupID := "consumerGroup"

	ctx := context.Background()

	reder := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brocker,
		Topic:    topic,
		GroupID:  goupID,
		MinBytes: 10e3,
		MaxBytes: 10e3,
	})

	defer reder.Close()

	fmt.Println("Start Consuming...")

	for{
		msg,err:=reder.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Message at offset - %d: %s= %s\n",msg.Offset,string(msg.Key),string(msg.Value))
		//simulate processing time

		time.Sleep(1 *time.Second)
	}

}


