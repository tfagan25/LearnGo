package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)

func main() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"localhost:9092"},
        Topic:     "logs",
        Partition: 0,
        GroupID:   "logger-group",
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })
    defer r.Close()

    fmt.Println("Consumer started. Listening for messages...")
    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            fmt.Println("Error reading message:", err)
            continue
        }
        fmt.Printf("Received message: %s\n", string(m.Value))
    }
}