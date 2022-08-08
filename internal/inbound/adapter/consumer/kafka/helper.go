package kafka

import (
	"encoding/json"
	"ftgo-order/pkg/message"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func toMessage(kafkaMsg *kafka.Message) message.Message {
	resMsg := message.Message{}

	header := map[string][]byte{}
	header["key"] = kafkaMsg.Key
	for _, msgHeader := range kafkaMsg.Headers {
		header[msgHeader.Key] = msgHeader.Value
	}

	resMsg.Header, _ = json.Marshal(header)
	resMsg.Payload = kafkaMsg.Value
	return resMsg
}

func mapToSlice[T comparable](in map[T]struct{}) []T {
	var out []T
	for key, _ := range in {
		out = append(out, key)
	}
	return out
}
