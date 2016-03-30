package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"log"
)

type Message struct {
	key       string
	value     string
	offset    int64
	partition int32
	topic     string
}

func (m *Message) String() string {
	return fmt.Sprintf("topic:[%s] parition:[%d] offset:[%d] key:[%s] value:%s", m.topic, m.partition, m.offset, m.key, m.value)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "<broker1;broker2;...> <topic>")
		os.Exit(0)
	}
	brokers := os.Args[1]
	topic := os.Args[2]

	consumer, err := sarama.NewConsumer(strings.Split(brokers, ";"), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	partitions, err := consumer.Partitions(topic)
	log.Println("partitions:", partitions)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for _, partition := range partitions {
		pcon, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		go func() {
			mc := pcon.Messages()
			for {
				msg := <-mc
				m := Message{key:string(msg.Key), value:string(msg.Value), topic:msg.Topic, partition:msg.Partition, offset:msg.Offset}
				fmt.Println(m.String())
			}
		}()
	}

	killChan := make(chan os.Signal)
	signal.Notify(killChan, os.Kill, os.Interrupt)

	kill := <-killChan
	fmt.Println(kill)
}