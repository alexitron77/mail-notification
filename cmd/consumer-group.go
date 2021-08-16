package cmd

import (
	"context"
	"fmt"
	"log"
	"sync"

	kafka "github.com/segmentio/kafka-go"
	env "mail.notification.com/pkg"
)

func ConsumeGroup() {
	var wg sync.WaitGroup

	env := env.GetEnv()

	wg.Add(2)
	go consumer1([]string{env.Kafka.Url}, env.Kafka.Topic)
	go consumer2([]string{env.Kafka.Url}, env.Kafka.Topic)

	wg.Wait()
	defer wg.Done()

}

func consumer1(broker []string, topic string) {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     broker,
		Topic:       topic,
		GroupID:     "3",
		StartOffset: kafka.LastOffset,
	})

	for {
		msg, err := k.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Kafka reader has failed, err: %v", err)
		}
		fmt.Printf("Consumer 1, %v\n", string(msg.Value))
	}
}

func consumer2(broker []string, topic string) {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     broker,
		Topic:       topic,
		GroupID:     "3",
		StartOffset: kafka.LastOffset,
	})

	for {
		msg, err := k.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Kafka reader has failed, err: %v", err)
		}
		fmt.Printf("consumer 2, %v\n", string(msg.Value))
	}
}
