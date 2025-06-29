package main

import (
	"context"
	"fmt"
	"time"
	"github.com/segmentio/kafka-go"
	"log"
)

func main () {
	conn, _ := kafka.Dial("tcp", "localhost:9092")
	defer conn.Close()

	err := conn.CreateTopics(kafka.TopicConfig{
		Topic:             "logs",
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	if err != nil {
		log.Fatal("Failed initializing Kafka")
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "logs",
        Balancer: &kafka.LeastBytes{},
    })

	messages := []string{
        "User logged in",
        "File uploaded",
        "Email sent",
        "User logged out",
    }

    for i := 0; i < len(messages); i++ {
        err := writer.WriteMessages(context.Background(),
            kafka.Message{
                Key:   []byte("key"),
                Value: []byte(messages[i]),
            },
        )
        if err != nil {
            fmt.Println("Failed to write message:", err)
        } else {
            fmt.Println("Sent:", messages[i])
        }
        time.Sleep(1 * time.Second)
        
        if i == len(messages) - 1 {
            i = -1
        }
    }

    writer.Close()
}