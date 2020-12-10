package kafkautils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaUrl string, topic string, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaUrl, ",")
	fmt.Println("getKafkaReader")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 1,        // 1B
		MaxBytes: 57671680, // 10MB
		// MaxBytes: 10e6, // 10MB
	})
}

func RunConsumer() {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	responseMessage := &Message{}
	_ = responseMessage

	// kafkaURL := "kafka-service:9092"
	// topic := "admintome-test"
	// groupID := "consumer-group-id"

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	log.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		unMErr := json.Unmarshal([]byte(string(m.Value)), responseMessage)
		if unMErr != nil {
			panic(unMErr)
		}

		log.Printf("message at topic:%v partition:%v offset:%v	%s = %s | %s | %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value), responseMessage.Action, responseMessage.Context.Field)
	}
}
