package main

import (
	"testing"

	"github.com/nsqio/go-nsq"
)

func TestMain(t *testing.T) {
	ProduceMsg(nsq.NewConfig())
}
