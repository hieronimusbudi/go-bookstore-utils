package kafkautils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// {
// 	"action": "user_updated",
// 	"context": {
// 	  "id": 1343,
// 	  "field": "first_name",
// 	  "previous": "Robert",
// 	  "new": "Bob"
// 	}
// }

type Message struct {
	Action  string
	Context interface{}
}

// type MessageContext struct {
// 	Id       string
// 	Field    string
// 	Previous string
// 	New      string
// }

type MessageContext struct{}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 10 * time.Millisecond,
	}
}

func producerHandler(kafkaWriter *kafka.Writer, message *Message) {
	jsonOut, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", 1)),
		Value: []byte(fmt.Sprint(string(jsonOut))),
	}
	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatalln(err)
	}
}

func RunProducer(message *Message, topic string) {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaWriter := getKafkaWriter(kafkaURL, topic)
	defer kafkaWriter.Close()

	// Add handle func for producer.
	producerHandler(kafkaWriter, message)

	// Run the web server.
	log.Println("Run Producer...")
}
