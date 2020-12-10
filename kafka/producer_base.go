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
	Context MessageContext
}

type MessageContext struct {
	Id       string
	Field    string
	Previous string
	New      string
}

func producerHandler(kafkaWriter *kafka.Writer) {
	value := &Message{
		Action: "user_updated",
		Context: MessageContext{
			Id:       "1343",
			Field:    "first_name",
			Previous: "Robert",
			New:      "Bob",
		},
	}

	jsonOut, jsonErr := json.Marshal(value)
	if jsonErr != nil {
		panic(jsonErr)
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", 1)),
		Value: []byte(fmt.Sprint(string(jsonOut))),
	}
	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 10 * time.Millisecond,
	}
}

func RunProducer() {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	// groupID := os.Getenv("KAFKA_GROUP_ID")
	// kafkaURL := "kafka-service:9092"
	// topic := "admintome-test"
	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer kafkaWriter.Close()

	// Add handle func for producer.
	producerHandler(kafkaWriter)

	// Run the web server.
	log.Println("start producer-api ... !!")
}
