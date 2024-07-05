package main

import (
	"fmt"
	"net"

	"github.com/segmentio/kafka-go"
)

func createTopic(brockers []string, topic string) error {
	conn, err := kafka.Dial("tcp", brockers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		fmt.Println("err :", err)
		return err
	}

	controllerConn,err:=kafka.Dial("tcp",net.JoinHostPort(controller.Host,fmt.Sprint(controller.Port)))
	
	if err!=nil {
		return err
	}

	defer controllerConn.Close()


	topicConfig:=kafka.TopicConfig{
		Topic: topic,
		NumPartitions: 1,
		ReplicationFactor: 1,
	}

	return controllerConn.CreateTopics(topicConfig)
}
