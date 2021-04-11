package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	// RoutePackage "github.com/neocite/simulator/application/route"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	KafkaApplicationPackage "github.com/neocite/simulator/application/kafka"
	KafkaPackage "github.com/neocite/simulator/infra/kafka"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("erro loading .env file")
	}

}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := KafkaPackage.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		go KafkaApplicationPackage.Produce(msg)
		fmt.Println(string(msg.Value))
	}

}
