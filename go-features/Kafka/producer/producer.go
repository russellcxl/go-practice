package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:   "my-topic",
		Balancer: &kafka.LeastBytes{},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}
		if input == "stop" {
			break
		}
		fmt.Println("writing message:", input)
		if err := w.WriteMessages(context.Background(), kafka.Message{
			Key: []byte("Random key"),
			Value: []byte(input),
		}); err != nil {
			log.Fatal("failed to write messages:", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
