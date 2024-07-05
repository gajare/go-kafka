package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	brocker := []string{"localhost:29092"}
	topic := "TEST_TOPIC"
	fmt.Println("BROCKER AND TOPIC :", brocker, topic)

	err := createTopic(brocker, topic)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Topic Created Successfully")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brocker,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	ctx := context.Background()
	err = writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("key-A"),
			Value: []byte("Hello Amot"),
		})

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Successfully Published the Message  to Topic")
	fmt.Println("brocker , topic :", brocker, topic)
}

// write one function
