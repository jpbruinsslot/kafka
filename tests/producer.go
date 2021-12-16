package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	VERSION = "0.1.0"
	USAGE   = `NAME: producer - test script to add messages to a kafka topic

VERSION:
    %s
`
)

var (
	flgTopic  string
	flgBroker string
)

func init() {
	flag.StringVar(
		&flgTopic,
		"t",
		"",
		"topic name",
	)

	flag.StringVar(
		&flgTopic,
		"topic",
		"",
		"topic name",
	)

	flag.StringVar(
		&flgBroker,
		"b",
		"",
		"broker host",
	)

	flag.StringVar(
		&flgBroker,
		"broker",
		"",
		"broker host",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}
}

func main() {
	flag.Parse()
	if flgTopic == "" || flgBroker == "" {
		flag.Usage()
		os.Exit(1)
	}

	l := log.New(os.Stdout, "kafka producer: ", 0)

	ctx := context.Background()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{flgBroker},
		Topic:   flgTopic,
		Logger:  l,
	})

	i := 0
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
