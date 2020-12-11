package kafkautils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaUrl string, topic string, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaUrl, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 1,        // 1B
		MaxBytes: 57671680, // 10MB
		// MaxBytes: 10e6, // 10MB
	})
}

func RunConsumer(topic string, groupID string, resultChannel chan Message) {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("KAFKA_URL")
	responseMessage := &Message{}

	reader := getKafkaReader(kafkaURL, topic, groupID)
	defer reader.Close()

	log.Println("Run Consumer...")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		unMErr := json.Unmarshal([]byte(string(m.Value)), responseMessage)
		if unMErr != nil {
			log.Fatalln(unMErr)
		}

		// send result to main thread
		resultChannel <- *responseMessage
		log.Printf("message at topic:%v partition:%v offset:%v	%s = %s | %s | %+v\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value), responseMessage.Action, responseMessage.Context)
	}
}
