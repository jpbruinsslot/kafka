package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic  = "first_topic"
	broker = "localhost:32792"
)

func main() {
	i := 0

	l := log.New(os.Stdout, "kafka producer: ", 0)

	ctx := context.Background()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
		Logger:  l,
	})

	for {
		err := w.WriteMessages(
			ctx,
			kafka.Message{
				Key:   []byte(strconv.Itoa(i)),
				Value: []byte(fmt.Sprintf("message: %d", i)),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		i++
		time.Sleep(time.Second)
	}
}
