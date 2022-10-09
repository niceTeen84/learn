package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

const (
	TOPIC       = "renbw"
	CHAN        = "exp"
	ADDR        = "127.0.0.1:4150"
	ADDR_HOST   = "127.0.0.1:4150"
	LOOKUP_ADDR = "127.0.0.1:4161"
	DELAY       = 10 * time.Second
)

type MyHandler struct {
}

func (h *MyHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	fmt.Println(string(m.Body))
	time.Sleep(DELAY)
	fmt.Println("done!")
	return nil
}

func ProduceMsg(cnf *nsq.Config) {
	producer, err := nsq.NewProducer(ADDR_HOST, cnf)
	defer producer.Stop()

	if err != nil {
		log.Fatal("init nsq failed ", err.Error())
	}

	messageBody := []byte("{\"name\": \"jerry\", \"age\": 15}")

	err = producer.Publish(TOPIC, messageBody)
	if err != nil {
		log.Fatal("sending msg failed ", err.Error())
	}
}

func Consumer(cnf *nsq.Config, idx int) {
	fmt.Println("worker ", idx, " started")
	consumer, err := nsq.NewConsumer(TOPIC, CHAN, cnf)
	// defer consumer.Stop()
	if err != nil {
		log.Fatal("init consumer failed ", err.Error())
	}
	consumer.AddHandler(&MyHandler{})
	err = consumer.ConnectToNSQLookupd(LOOKUP_ADDR)
	if err != nil {
		log.Fatal("connect to lookup failed ", err.Error())
	}

}

func main() {
	cnf := nsq.NewConfig()
	// ProduceMsg(cnf)
	fmt.Println("init...wait for msg")

	Consumer(cnf, 1)
	Consumer(cnf, 2)
	Consumer(cnf, 3)

	// wait signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM|syscall.SIGINT)
	<-sigChan
}
